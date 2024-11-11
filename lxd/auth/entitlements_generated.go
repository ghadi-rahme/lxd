// Code generated by `make update-auth`. DO NOT EDIT.

package auth

import (
	"github.com/canonical/lxd/shared/entity"
)

// Entitlement is a representation of the relations that group members can have with entity types.
type Entitlement string

const (
	// EntitlementCanView is the "can_view" entitlement. It applies to the following entities: entity.TypeCertificate, entity.TypeAuthGroup, entity.TypeIdentity, entity.TypeIdentityProviderGroup, entity.TypeImage, entity.TypeImageAlias, entity.TypeInstance, entity.TypeNetwork, entity.TypeNetworkACL, entity.TypeNetworkZone, entity.TypeProfile, entity.TypeProject, entity.TypeStorageBucket, entity.TypeStorageVolume.
	EntitlementCanView Entitlement = "can_view"

	// EntitlementCanEdit is the "can_edit" entitlement. It applies to the following entities: entity.TypeCertificate, entity.TypeAuthGroup, entity.TypeIdentity, entity.TypeIdentityProviderGroup, entity.TypeImage, entity.TypeImageAlias, entity.TypeInstance, entity.TypeNetwork, entity.TypeNetworkACL, entity.TypeNetworkZone, entity.TypeProfile, entity.TypeProject, entity.TypeServer, entity.TypeStorageBucket, entity.TypeStoragePool, entity.TypeStorageVolume.
	EntitlementCanEdit Entitlement = "can_edit"

	// EntitlementCanDelete is the "can_delete" entitlement. It applies to the following entities: entity.TypeCertificate, entity.TypeAuthGroup, entity.TypeIdentity, entity.TypeIdentityProviderGroup, entity.TypeImage, entity.TypeImageAlias, entity.TypeInstance, entity.TypeNetwork, entity.TypeNetworkACL, entity.TypeNetworkZone, entity.TypeProfile, entity.TypeProject, entity.TypeStorageBucket, entity.TypeStoragePool, entity.TypeStorageVolume.
	EntitlementCanDelete Entitlement = "can_delete"

	// EntitlementAdmin is the "admin" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementAdmin Entitlement = "admin"

	// EntitlementViewer is the "viewer" entitlement. It applies to the following entities: entity.TypeProject, entity.TypeServer.
	EntitlementViewer Entitlement = "viewer"

	// EntitlementPermissionManager is the "permission_manager" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementPermissionManager Entitlement = "permission_manager"

	// EntitlementCanViewPermissions is the "can_view_permissions" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewPermissions Entitlement = "can_view_permissions"

	// EntitlementCanCreateIdentities is the "can_create_identities" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanCreateIdentities Entitlement = "can_create_identities"

	// EntitlementCanViewIdentities is the "can_view_identities" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewIdentities Entitlement = "can_view_identities"

	// EntitlementCanEditIdentities is the "can_edit_identities" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanEditIdentities Entitlement = "can_edit_identities"

	// EntitlementCanDeleteIdentities is the "can_delete_identities" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanDeleteIdentities Entitlement = "can_delete_identities"

	// EntitlementCanCreateGroups is the "can_create_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanCreateGroups Entitlement = "can_create_groups"

	// EntitlementCanViewGroups is the "can_view_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewGroups Entitlement = "can_view_groups"

	// EntitlementCanEditGroups is the "can_edit_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanEditGroups Entitlement = "can_edit_groups"

	// EntitlementCanDeleteGroups is the "can_delete_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanDeleteGroups Entitlement = "can_delete_groups"

	// EntitlementCanCreateIdentityProviderGroups is the "can_create_identity_provider_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanCreateIdentityProviderGroups Entitlement = "can_create_identity_provider_groups"

	// EntitlementCanViewIdentityProviderGroups is the "can_view_identity_provider_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewIdentityProviderGroups Entitlement = "can_view_identity_provider_groups"

	// EntitlementCanEditIdentityProviderGroups is the "can_edit_identity_provider_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanEditIdentityProviderGroups Entitlement = "can_edit_identity_provider_groups"

	// EntitlementCanDeleteIdentityProviderGroups is the "can_delete_identity_provider_groups" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanDeleteIdentityProviderGroups Entitlement = "can_delete_identity_provider_groups"

	// EntitlementStoragePoolManager is the "storage_pool_manager" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementStoragePoolManager Entitlement = "storage_pool_manager"

	// EntitlementCanCreateStoragePools is the "can_create_storage_pools" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanCreateStoragePools Entitlement = "can_create_storage_pools"

	// EntitlementCanEditStoragePools is the "can_edit_storage_pools" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanEditStoragePools Entitlement = "can_edit_storage_pools"

	// EntitlementCanDeleteStoragePools is the "can_delete_storage_pools" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanDeleteStoragePools Entitlement = "can_delete_storage_pools"

	// EntitlementProjectManager is the "project_manager" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementProjectManager Entitlement = "project_manager"

	// EntitlementCanCreateProjects is the "can_create_projects" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanCreateProjects Entitlement = "can_create_projects"

	// EntitlementCanViewProjects is the "can_view_projects" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewProjects Entitlement = "can_view_projects"

	// EntitlementCanEditProjects is the "can_edit_projects" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanEditProjects Entitlement = "can_edit_projects"

	// EntitlementCanDeleteProjects is the "can_delete_projects" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanDeleteProjects Entitlement = "can_delete_projects"

	// EntitlementCanOverrideClusterTargetRestriction is the "can_override_cluster_target_restriction" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanOverrideClusterTargetRestriction Entitlement = "can_override_cluster_target_restriction"

	// EntitlementCanViewPrivilegedEvents is the "can_view_privileged_events" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewPrivilegedEvents Entitlement = "can_view_privileged_events"

	// EntitlementCanViewResources is the "can_view_resources" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewResources Entitlement = "can_view_resources"

	// EntitlementCanViewMetrics is the "can_view_metrics" entitlement. It applies to the following entities: entity.TypeProject, entity.TypeServer.
	EntitlementCanViewMetrics Entitlement = "can_view_metrics"

	// EntitlementCanViewWarnings is the "can_view_warnings" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewWarnings Entitlement = "can_view_warnings"

	// EntitlementCanViewUnmanagedNetworks is the "can_view_unmanaged_networks" entitlement. It applies to the following entities: entity.TypeServer.
	EntitlementCanViewUnmanagedNetworks Entitlement = "can_view_unmanaged_networks"

	// EntitlementOperator is the "operator" entitlement. It applies to the following entities: entity.TypeInstance, entity.TypeProject.
	EntitlementOperator Entitlement = "operator"

	// EntitlementImageManager is the "image_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementImageManager Entitlement = "image_manager"

	// EntitlementCanCreateImages is the "can_create_images" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateImages Entitlement = "can_create_images"

	// EntitlementCanViewImages is the "can_view_images" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewImages Entitlement = "can_view_images"

	// EntitlementCanEditImages is the "can_edit_images" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditImages Entitlement = "can_edit_images"

	// EntitlementCanDeleteImages is the "can_delete_images" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteImages Entitlement = "can_delete_images"

	// EntitlementImageAliasManager is the "image_alias_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementImageAliasManager Entitlement = "image_alias_manager"

	// EntitlementCanCreateImageAliases is the "can_create_image_aliases" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateImageAliases Entitlement = "can_create_image_aliases"

	// EntitlementCanViewImageAliases is the "can_view_image_aliases" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewImageAliases Entitlement = "can_view_image_aliases"

	// EntitlementCanEditImageAliases is the "can_edit_image_aliases" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditImageAliases Entitlement = "can_edit_image_aliases"

	// EntitlementCanDeleteImageAliases is the "can_delete_image_aliases" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteImageAliases Entitlement = "can_delete_image_aliases"

	// EntitlementInstanceManager is the "instance_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementInstanceManager Entitlement = "instance_manager"

	// EntitlementCanCreateInstances is the "can_create_instances" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateInstances Entitlement = "can_create_instances"

	// EntitlementCanViewInstances is the "can_view_instances" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewInstances Entitlement = "can_view_instances"

	// EntitlementCanEditInstances is the "can_edit_instances" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditInstances Entitlement = "can_edit_instances"

	// EntitlementCanDeleteInstances is the "can_delete_instances" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteInstances Entitlement = "can_delete_instances"

	// EntitlementCanOperateInstances is the "can_operate_instances" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanOperateInstances Entitlement = "can_operate_instances"

	// EntitlementNetworkManager is the "network_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementNetworkManager Entitlement = "network_manager"

	// EntitlementCanCreateNetworks is the "can_create_networks" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateNetworks Entitlement = "can_create_networks"

	// EntitlementCanViewNetworks is the "can_view_networks" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewNetworks Entitlement = "can_view_networks"

	// EntitlementCanEditNetworks is the "can_edit_networks" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditNetworks Entitlement = "can_edit_networks"

	// EntitlementCanDeleteNetworks is the "can_delete_networks" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteNetworks Entitlement = "can_delete_networks"

	// EntitlementNetworkACLManager is the "network_acl_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementNetworkACLManager Entitlement = "network_acl_manager"

	// EntitlementCanCreateNetworkACLs is the "can_create_network_acls" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateNetworkACLs Entitlement = "can_create_network_acls"

	// EntitlementCanViewNetworkACLs is the "can_view_network_acls" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewNetworkACLs Entitlement = "can_view_network_acls"

	// EntitlementCanEditNetworkACLs is the "can_edit_network_acls" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditNetworkACLs Entitlement = "can_edit_network_acls"

	// EntitlementCanDeleteNetworkACLs is the "can_delete_network_acls" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteNetworkACLs Entitlement = "can_delete_network_acls"

	// EntitlementNetworkZoneManager is the "network_zone_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementNetworkZoneManager Entitlement = "network_zone_manager"

	// EntitlementCanCreateNetworkZones is the "can_create_network_zones" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateNetworkZones Entitlement = "can_create_network_zones"

	// EntitlementCanViewNetworkZones is the "can_view_network_zones" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewNetworkZones Entitlement = "can_view_network_zones"

	// EntitlementCanEditNetworkZones is the "can_edit_network_zones" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditNetworkZones Entitlement = "can_edit_network_zones"

	// EntitlementCanDeleteNetworkZones is the "can_delete_network_zones" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteNetworkZones Entitlement = "can_delete_network_zones"

	// EntitlementProfileManager is the "profile_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementProfileManager Entitlement = "profile_manager"

	// EntitlementCanCreateProfiles is the "can_create_profiles" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateProfiles Entitlement = "can_create_profiles"

	// EntitlementCanViewProfiles is the "can_view_profiles" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewProfiles Entitlement = "can_view_profiles"

	// EntitlementCanEditProfiles is the "can_edit_profiles" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditProfiles Entitlement = "can_edit_profiles"

	// EntitlementCanDeleteProfiles is the "can_delete_profiles" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteProfiles Entitlement = "can_delete_profiles"

	// EntitlementStorageVolumeManager is the "storage_volume_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementStorageVolumeManager Entitlement = "storage_volume_manager"

	// EntitlementCanCreateStorageVolumes is the "can_create_storage_volumes" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateStorageVolumes Entitlement = "can_create_storage_volumes"

	// EntitlementCanViewStorageVolumes is the "can_view_storage_volumes" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewStorageVolumes Entitlement = "can_view_storage_volumes"

	// EntitlementCanEditStorageVolumes is the "can_edit_storage_volumes" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditStorageVolumes Entitlement = "can_edit_storage_volumes"

	// EntitlementCanDeleteStorageVolumes is the "can_delete_storage_volumes" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteStorageVolumes Entitlement = "can_delete_storage_volumes"

	// EntitlementStorageBucketManager is the "storage_bucket_manager" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementStorageBucketManager Entitlement = "storage_bucket_manager"

	// EntitlementCanCreateStorageBuckets is the "can_create_storage_buckets" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanCreateStorageBuckets Entitlement = "can_create_storage_buckets"

	// EntitlementCanViewStorageBuckets is the "can_view_storage_buckets" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewStorageBuckets Entitlement = "can_view_storage_buckets"

	// EntitlementCanEditStorageBuckets is the "can_edit_storage_buckets" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanEditStorageBuckets Entitlement = "can_edit_storage_buckets"

	// EntitlementCanDeleteStorageBuckets is the "can_delete_storage_buckets" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanDeleteStorageBuckets Entitlement = "can_delete_storage_buckets"

	// EntitlementCanViewOperations is the "can_view_operations" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewOperations Entitlement = "can_view_operations"

	// EntitlementCanViewEvents is the "can_view_events" entitlement. It applies to the following entities: entity.TypeProject.
	EntitlementCanViewEvents Entitlement = "can_view_events"

	// EntitlementUser is the "user" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementUser Entitlement = "user"

	// EntitlementCanUpdateState is the "can_update_state" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementCanUpdateState Entitlement = "can_update_state"

	// EntitlementCanManageSnapshots is the "can_manage_snapshots" entitlement. It applies to the following entities: entity.TypeInstance, entity.TypeStorageVolume.
	EntitlementCanManageSnapshots Entitlement = "can_manage_snapshots"

	// EntitlementCanManageBackups is the "can_manage_backups" entitlement. It applies to the following entities: entity.TypeInstance, entity.TypeStorageVolume.
	EntitlementCanManageBackups Entitlement = "can_manage_backups"

	// EntitlementCanConnectSFTP is the "can_connect_sftp" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementCanConnectSFTP Entitlement = "can_connect_sftp"

	// EntitlementCanAccessFiles is the "can_access_files" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementCanAccessFiles Entitlement = "can_access_files"

	// EntitlementCanAccessConsole is the "can_access_console" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementCanAccessConsole Entitlement = "can_access_console"

	// EntitlementCanExec is the "can_exec" entitlement. It applies to the following entities: entity.TypeInstance.
	EntitlementCanExec Entitlement = "can_exec"
)

var EntityTypeToEntitlements = map[entity.Type][]Entitlement{
	entity.TypeCertificate: {
		// Grants permission to view the certificate.
		EntitlementCanView,
		// Grants permission to edit the certificate.
		EntitlementCanEdit,
		// Grants permission to delete the certificate.
		EntitlementCanDelete,
	},
	entity.TypeAuthGroup: {
		// Grants permission to view the group. Identities can always view groups that they are a member of.
		EntitlementCanView,
		// Grants permission to edit the group.
		EntitlementCanEdit,
		// Grants permission to delete the group.
		EntitlementCanDelete,
	},
	entity.TypeIdentity: {
		// Grants permission to view the identity.
		EntitlementCanView,
		// Grants permission to edit the identity.
		EntitlementCanEdit,
		// Grants permission to delete the identity.
		EntitlementCanDelete,
	},
	entity.TypeIdentityProviderGroup: {
		// Grants permission to view the identity provider group.
		EntitlementCanView,
		// Grants permission to edit the identity provider group.
		EntitlementCanEdit,
		// Grants permission to delete the identity provider group.
		EntitlementCanDelete,
	},
	entity.TypeImage: {
		// Grants permission to edit the image.
		EntitlementCanEdit,
		// Grants permission to delete the image.
		EntitlementCanDelete,
		// Grants permission to view the image.
		EntitlementCanView,
	},
	entity.TypeImageAlias: {
		// Grants permission to edit the image alias.
		EntitlementCanEdit,
		// Grants permission to delete the image alias.
		EntitlementCanDelete,
		// Grants permission to view the image alias.
		EntitlementCanView,
	},
	entity.TypeInstance: {
		// Grants permission to view the instance, to access files, and to start a terminal or console session.
		EntitlementUser,
		// Grants permission to view the instance, to access files, start a terminal or console session, and to manage snapshots and backups.
		EntitlementOperator,
		// Grants permission to edit the instance.
		EntitlementCanEdit,
		// Grants permission to delete the instance.
		EntitlementCanDelete,
		// Grants permission to view the instance.
		EntitlementCanView,
		// Grants permission to change the instance state.
		EntitlementCanUpdateState,
		// Grants permission to create and delete snapshots of the instance.
		EntitlementCanManageSnapshots,
		// Grants permission to create and delete backups of the instance.
		EntitlementCanManageBackups,
		// Grants permission to get an SFTP client for the instance.
		EntitlementCanConnectSFTP,
		// Grants permission to push or pull files into or out of the instance.
		EntitlementCanAccessFiles,
		// Grants permission to start a console session.
		EntitlementCanAccessConsole,
		// Grants permission to start a terminal session.
		EntitlementCanExec,
	},
	entity.TypeNetwork: {
		// Grants permission to edit the network.
		EntitlementCanEdit,
		// Grants permission to delete the network.
		EntitlementCanDelete,
		// Grants permission to view the network.
		EntitlementCanView,
	},
	entity.TypeNetworkACL: {
		// Grants permission to edit the network ACL.
		EntitlementCanEdit,
		// Grants permission to delete the network ACL.
		EntitlementCanDelete,
		// Grants permission to view the network ACL.
		EntitlementCanView,
	},
	entity.TypeNetworkZone: {
		// Grants permission to edit the network zone.
		EntitlementCanEdit,
		// Grants permission to delete the network zone.
		EntitlementCanDelete,
		// Grants permission to view the network zone.
		EntitlementCanView,
	},
	entity.TypeProfile: {
		// Grants permission to edit the profile.
		EntitlementCanEdit,
		// Grants permission to delete the profile.
		EntitlementCanDelete,
		// Grants permission to view the profile.
		EntitlementCanView,
	},
	entity.TypeProject: {
		// Grants permission to create, view, edit, and delete all resources belonging to the project, but does not grant permission to edit the project configuration itself.
		EntitlementOperator,
		// Grants permission to view all resources belonging to the project.
		EntitlementViewer,
		// Grants permission to view the project.
		EntitlementCanView,
		// Grants permission to edit the project.
		EntitlementCanEdit,
		// Grants permission to delete the project.
		EntitlementCanDelete,
		// Grants permission to create, view, edit, and delete all images belonging to the project.
		EntitlementImageManager,
		// Grants permission to create images.
		EntitlementCanCreateImages,
		// Grants permission to view images.
		EntitlementCanViewImages,
		// Grants permission to edit images.
		EntitlementCanEditImages,
		// Grants permission to delete images.
		EntitlementCanDeleteImages,
		// Grants permission to create, view, edit, and delete all image aliases belonging to the project.
		EntitlementImageAliasManager,
		// Grants permission to create image aliases.
		EntitlementCanCreateImageAliases,
		// Grants permission to view image aliases.
		EntitlementCanViewImageAliases,
		// Grants permission to edit image aliases.
		EntitlementCanEditImageAliases,
		// Grants permission to delete image aliases.
		EntitlementCanDeleteImageAliases,
		// Grants permission to create, view, edit, and delete all instances belonging to the project.
		EntitlementInstanceManager,
		// Grants permission to create instances.
		EntitlementCanCreateInstances,
		// Grants permission to view instances.
		EntitlementCanViewInstances,
		// Grants permission to edit instances.
		EntitlementCanEditInstances,
		// Grants permission to delete instances.
		EntitlementCanDeleteInstances,
		// Grants permission to view instances, manage their state, manage their snapshots and backups, start terminal or console sessions, and access their files.
		EntitlementCanOperateInstances,
		// Grants permission to create, view, edit, and delete all networks belonging to the project.
		EntitlementNetworkManager,
		// Grants permission to create networks.
		EntitlementCanCreateNetworks,
		// Grants permission to view networks.
		EntitlementCanViewNetworks,
		// Grants permission to edit networks.
		EntitlementCanEditNetworks,
		// Grants permission to delete networks.
		EntitlementCanDeleteNetworks,
		// Grants permission to create, view, edit, and delete all network ACLs belonging to the project.
		EntitlementNetworkACLManager,
		// Grants permission to create network ACLs.
		EntitlementCanCreateNetworkACLs,
		// Grants permission to view network ACLs.
		EntitlementCanViewNetworkACLs,
		// Grants permission to edit network ACLs.
		EntitlementCanEditNetworkACLs,
		// Grants permission to delete network ACLs.
		EntitlementCanDeleteNetworkACLs,
		// Grants permission to create, view, edit, and delete all network zones belonging to the project.
		EntitlementNetworkZoneManager,
		// Grants permission to create network zones.
		EntitlementCanCreateNetworkZones,
		// Grants permission to view network zones.
		EntitlementCanViewNetworkZones,
		// Grants permission to edit network zones.
		EntitlementCanEditNetworkZones,
		// Grants permission to delete network zones.
		EntitlementCanDeleteNetworkZones,
		// Grants permission to create, view, edit, and delete all profiles belonging to the project.
		EntitlementProfileManager,
		// Grants permission to create profiles.
		EntitlementCanCreateProfiles,
		// Grants permission to view profiles.
		EntitlementCanViewProfiles,
		// Grants permission to edit profiles.
		EntitlementCanEditProfiles,
		// Grants permission to delete profiles.
		EntitlementCanDeleteProfiles,
		// Grants permission to create, view, edit, and delete all storage volumes belonging to the project.
		EntitlementStorageVolumeManager,
		// Grants permission to create storage volumes.
		EntitlementCanCreateStorageVolumes,
		// Grants permission to view storage volumes.
		EntitlementCanViewStorageVolumes,
		// Grants permission to edit storage volumes.
		EntitlementCanEditStorageVolumes,
		// Grants permission to delete storage volumes.
		EntitlementCanDeleteStorageVolumes,
		// Grants permission to create, view, edit, and delete all storage buckets belonging to the project.
		EntitlementStorageBucketManager,
		// Grants permission to create storage buckets.
		EntitlementCanCreateStorageBuckets,
		// Grants permission to view storage buckets.
		EntitlementCanViewStorageBuckets,
		// Grants permission to edit storage buckets.
		EntitlementCanEditStorageBuckets,
		// Grants permission to delete storage buckets.
		EntitlementCanDeleteStorageBuckets,
		// Grants permission to view operations relating to the project.
		EntitlementCanViewOperations,
		// Grants permission to view events relating to the project.
		EntitlementCanViewEvents,
		// Grants permission to view project level metrics.
		EntitlementCanViewMetrics,
	},
	entity.TypeServer: {
		// Grants full access to LXD as if via Unix socket.
		EntitlementAdmin,
		// Grants access to view all resources in the LXD server.
		EntitlementViewer,
		// Grants permission to edit server configuration, to edit cluster member configuration, to update the state of a cluster member, to create, edit, and delete cluster groups, to update cluster member certificates, and to edit or delete warnings.
		EntitlementCanEdit,
		// Grants permission to view permissions, to create, edit, and delete identities, to view, create, edit, and delete authorization groups, and to view, create, edit, and delete identity provider groups. Note that clients with this permission are able to elevate their own privileges.
		EntitlementPermissionManager,
		// Grants permission to view permissions.
		EntitlementCanViewPermissions,
		// Grants permission to create identities.
		EntitlementCanCreateIdentities,
		// Grants permission to view identities.
		EntitlementCanViewIdentities,
		// Grants permission to edit identities.
		EntitlementCanEditIdentities,
		// Grants permission to delete identities.
		EntitlementCanDeleteIdentities,
		// Grants permission to create authorization groups.
		EntitlementCanCreateGroups,
		// Grants permission to view authorization groups.
		EntitlementCanViewGroups,
		// Grants permission to edit authorization groups.
		EntitlementCanEditGroups,
		// Grants permission to delete authorization groups.
		EntitlementCanDeleteGroups,
		// Grants permission to create identity provider groups.
		EntitlementCanCreateIdentityProviderGroups,
		// Grants permission to view identity provider groups.
		EntitlementCanViewIdentityProviderGroups,
		// Grants permission to edit identity provider groups.
		EntitlementCanEditIdentityProviderGroups,
		// Grants permission to delete identity provider groups.
		EntitlementCanDeleteIdentityProviderGroups,
		// Grants permission to create, edit, and delete storage pools.
		EntitlementStoragePoolManager,
		// Grants permission to create storage pools.
		EntitlementCanCreateStoragePools,
		// Grants permission to edit storage pools.
		EntitlementCanEditStoragePools,
		// Grants permission to delete storage pools.
		EntitlementCanDeleteStoragePools,
		// Grants permission to view, create, edit, and delete projects, and to create, edit, and delete any resources that are owned by those projects.
		EntitlementProjectManager,
		// Grants permission to create projects.
		EntitlementCanCreateProjects,
		// Grants permission to view projects, and all resources within those projects.
		EntitlementCanViewProjects,
		// Grants permission to edit projects, and all resources within those projects.
		EntitlementCanEditProjects,
		// Grants permission to delete projects.
		EntitlementCanDeleteProjects,
		// If a project is configured with `restricted.cluster.target`, clients with this permission can override the restriction.
		EntitlementCanOverrideClusterTargetRestriction,
		// Grants permission to view privileged event types, such as logging events.
		EntitlementCanViewPrivilegedEvents,
		// Grants permission to view server and storage pool resource usage information.
		EntitlementCanViewResources,
		// Grants permission to view all server and project level metrics.
		EntitlementCanViewMetrics,
		// Grants permission to view warnings.
		EntitlementCanViewWarnings,
		// Grants permission to view unmanaged networks on the LXD host machines.
		EntitlementCanViewUnmanagedNetworks,
	},
	entity.TypeStorageBucket: {
		// Grants permission to edit the storage bucket.
		EntitlementCanEdit,
		// Grants permission to delete the storage bucket.
		EntitlementCanDelete,
		// Grants permission to view the storage bucket.
		EntitlementCanView,
	},
	entity.TypeStoragePool: {
		// Grants permission to edit the storage pool.
		EntitlementCanEdit,
		// Grants permission to delete the storage pool.
		EntitlementCanDelete,
	},
	entity.TypeStorageVolume: {
		// Grants permission to edit the storage volume.
		EntitlementCanEdit,
		// Grants permission to delete the storage volume.
		EntitlementCanDelete,
		// Grants permission to view the storage volume.
		EntitlementCanView,
		// Grants permission to create and delete snapshots of the storage volume.
		EntitlementCanManageSnapshots,
		// Grants permission to create and delete backups of the storage volume.
		EntitlementCanManageBackups,
	},
}
