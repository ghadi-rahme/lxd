package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/lxd/cluster"
	"github.com/lxc/lxd/lxd/db"
	"github.com/lxc/lxd/lxd/state"
	"github.com/lxc/lxd/lxd/util"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
	"github.com/lxc/lxd/shared/logger"
	"github.com/lxc/lxd/shared/version"

	log "github.com/lxc/lxd/shared/log15"
)

var profilesCmd = Command{
	name: "profiles",
	get:  profilesGet,
	post: profilesPost,
}

var profileCmd = Command{
	name:   "profiles/{name}",
	get:    profileGet,
	put:    profilePut,
	delete: profileDelete,
	post:   profilePost,
	patch:  profilePatch,
}

/* This is used for both profiles post and profile put */
func profilesGet(d *Daemon, r *http.Request) Response {
	results, err := d.cluster.Profiles("default")
	if err != nil {
		return SmartError(err)
	}

	recursion := util.IsRecursionRequest(r)

	resultString := make([]string, len(results))
	resultMap := make([]*api.Profile, len(results))
	i := 0
	for _, name := range results {
		if !recursion {
			url := fmt.Sprintf("/%s/profiles/%s", version.APIVersion, name)
			resultString[i] = url
		} else {
			profile, err := doProfileGet(d.State(), name)
			if err != nil {
				logger.Error("Failed to get profile", log.Ctx{"profile": name})
				continue
			}
			resultMap[i] = profile
		}
		i++
	}

	if !recursion {
		return SyncResponse(true, resultString)
	}

	return SyncResponse(true, resultMap)
}

func profilesPost(d *Daemon, r *http.Request) Response {
	project := projectParam(r)
	req := api.ProfilesPost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return BadRequest(err)
	}

	// Sanity checks
	if req.Name == "" {
		return BadRequest(fmt.Errorf("No name provided"))
	}

	_, profile, _ := d.cluster.ProfileGet(req.Name)
	if profile != nil {
		return BadRequest(fmt.Errorf("The profile already exists"))
	}

	if strings.Contains(req.Name, "/") {
		return BadRequest(fmt.Errorf("Profile names may not contain slashes"))
	}

	if shared.StringInSlice(req.Name, []string{".", ".."}) {
		return BadRequest(fmt.Errorf("Invalid profile name '%s'", req.Name))
	}

	err := containerValidConfig(d.os, req.Config, true, false)
	if err != nil {
		return BadRequest(err)
	}

	err = containerValidDevices(d.cluster, req.Devices, true, false)
	if err != nil {
		return BadRequest(err)
	}

	// Update DB entry
	err = d.cluster.Transaction(func(tx *db.ClusterTx) error {
		profile := db.Profile{
			Project:     project,
			Name:        req.Name,
			Description: req.Description,
			Config:      req.Config,
			Devices:     req.Devices,
		}
		_, err := tx.ProfileCreate(profile)
		return err
	})
	if err != nil {
		return SmartError(
			fmt.Errorf("Error inserting %s into database: %s", req.Name, err))
	}

	return SyncResponseLocation(true, nil, fmt.Sprintf("/%s/profiles/%s", version.APIVersion, req.Name))
}

func doProfileGet(s *state.State, name string) (*api.Profile, error) {
	_, profile, err := s.Cluster.ProfileGet(name)
	if err != nil {
		return nil, err
	}

	cts, err := s.Cluster.ProfileContainersGet(name)
	if err != nil {
		return nil, err
	}

	usedBy := []string{}
	for _, ct := range cts {
		usedBy = append(usedBy, fmt.Sprintf("/%s/containers/%s", version.APIVersion, ct))
	}
	profile.UsedBy = usedBy

	return profile, nil
}

func profileGet(d *Daemon, r *http.Request) Response {
	name := mux.Vars(r)["name"]

	resp, err := doProfileGet(d.State(), name)
	if err != nil {
		return SmartError(err)
	}

	etag := []interface{}{resp.Config, resp.Description, resp.Devices}
	return SyncResponseETag(true, resp, etag)
}

func getContainersWithProfile(s *state.State, project, profile string) []container {
	results := []container{}

	output, err := s.Cluster.ProfileContainersGet(profile)
	if err != nil {
		return results
	}

	for _, name := range output {
		c, err := containerLoadByProjectAndName(s, project, name)
		if err != nil {
			logger.Error("Failed opening container", log.Ctx{"container": name})
			continue
		}
		results = append(results, c)
	}

	return results
}

func profilePut(d *Daemon, r *http.Request) Response {
	// Get the project
	project := projectParam(r)

	// Get the profile
	name := mux.Vars(r)["name"]

	if isClusterNotification(r) {
		// In this case the ProfilePut request payload contains
		// information about the old profile, since the new one has
		// already been saved in the database.
		old := api.ProfilePut{}
		err := json.NewDecoder(r.Body).Decode(&old)
		if err != nil {
			return BadRequest(err)
		}
		err = doProfileUpdateCluster(d, project, name, old)
		return SmartError(err)

	}

	id, profile, err := d.cluster.ProfileGet(name)
	if err != nil {
		return SmartError(fmt.Errorf("Failed to retrieve profile='%s'", name))
	}

	// Validate the ETag
	etag := []interface{}{profile.Config, profile.Description, profile.Devices}
	err = util.EtagCheck(r, etag)
	if err != nil {
		return PreconditionFailed(err)
	}

	req := api.ProfilePut{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return BadRequest(err)
	}

	err = doProfileUpdate(d, project, name, id, profile, req)

	if err == nil && !isClusterNotification(r) {
		// Notify all other nodes. If a node is down, it will be ignored.
		notifier, err := cluster.NewNotifier(d.State(), d.endpoints.NetworkCert(), cluster.NotifyAlive)
		if err != nil {
			return SmartError(err)
		}
		err = notifier(func(client lxd.ContainerServer) error {
			return client.UpdateProfile(name, profile.ProfilePut, "")
		})
		if err != nil {
			return SmartError(err)
		}
	}

	return SmartError(err)
}

func profilePatch(d *Daemon, r *http.Request) Response {
	// Get the project
	project := projectParam(r)

	// Get the profile
	name := mux.Vars(r)["name"]
	id, profile, err := d.cluster.ProfileGet(name)
	if err != nil {
		return SmartError(fmt.Errorf("Failed to retrieve profile='%s'", name))
	}

	// Validate the ETag
	etag := []interface{}{profile.Config, profile.Description, profile.Devices}
	err = util.EtagCheck(r, etag)
	if err != nil {
		return PreconditionFailed(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return InternalError(err)
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(body))

	reqRaw := shared.Jmap{}
	if err := json.NewDecoder(rdr1).Decode(&reqRaw); err != nil {
		return BadRequest(err)
	}

	req := api.ProfilePut{}
	if err := json.NewDecoder(rdr2).Decode(&req); err != nil {
		return BadRequest(err)
	}

	// Get Description
	_, err = reqRaw.GetString("description")
	if err != nil {
		req.Description = profile.Description
	}

	// Get Config
	if req.Config == nil {
		req.Config = profile.Config
	} else {
		for k, v := range profile.Config {
			_, ok := req.Config[k]
			if !ok {
				req.Config[k] = v
			}
		}
	}

	// Get Devices
	if req.Devices == nil {
		req.Devices = profile.Devices
	} else {
		for k, v := range profile.Devices {
			_, ok := req.Devices[k]
			if !ok {
				req.Devices[k] = v
			}
		}
	}

	return SmartError(doProfileUpdate(d, project, name, id, profile, req))
}

// The handler for the post operation.
func profilePost(d *Daemon, r *http.Request) Response {
	name := mux.Vars(r)["name"]

	if name == "default" {
		return Forbidden(errors.New("The 'default' profile cannot be renamed"))
	}

	req := api.ProfilePost{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return BadRequest(err)
	}

	// Sanity checks
	if req.Name == "" {
		return BadRequest(fmt.Errorf("No name provided"))
	}

	// Check that the name isn't already in use
	id, _, _ := d.cluster.ProfileGet(req.Name)
	if id > 0 {
		return Conflict(fmt.Errorf("Name '%s' already in use", req.Name))
	}

	if strings.Contains(req.Name, "/") {
		return BadRequest(fmt.Errorf("Profile names may not contain slashes"))
	}

	if shared.StringInSlice(req.Name, []string{".", ".."}) {
		return BadRequest(fmt.Errorf("Invalid profile name '%s'", req.Name))
	}

	err := d.cluster.ProfileUpdate(name, req.Name)
	if err != nil {
		return SmartError(err)
	}

	return SyncResponseLocation(true, nil, fmt.Sprintf("/%s/profiles/%s", version.APIVersion, req.Name))
}

// The handler for the delete operation.
func profileDelete(d *Daemon, r *http.Request) Response {
	project := projectParam(r)
	name := mux.Vars(r)["name"]

	if name == "default" {
		return Forbidden(errors.New("The 'default' profile cannot be deleted"))
	}

	_, err := doProfileGet(d.State(), name)
	if err != nil {
		return SmartError(err)
	}

	clist := getContainersWithProfile(d.State(), project, name)
	if len(clist) != 0 {
		return BadRequest(fmt.Errorf("Profile is currently in use"))
	}

	err = d.cluster.ProfileDelete(name)
	if err != nil {
		return SmartError(err)
	}

	return EmptySyncResponse
}
