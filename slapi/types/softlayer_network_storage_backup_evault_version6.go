package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Backup_Evault_Version6 - The
// SoftLayer_Network_Storage_Backup_Evault_Version6 contains the same properties as the
// SoftLayer_Network_Storage_Backup_Evault. Additional properties available for the EVault Storage
// type: softwareComponent, totalBytesUsed, backupJobDetails, restoreJobDetails and agentStatuses
type SoftLayer_Network_Storage_Backup_Evault_Version6 struct {

	// NasType - A Storage account's type. Valid examples are and
	NasType string `json:"nasType,omitempty"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// HardwareId - The server that is associated with a Storage service.
	HardwareId int `json:"hardwareId,omitempty"`

	// CapacityGb - A Storage account's capacity, measured in gigabytes.
	CapacityGb int `json:"capacityGb,omitempty"`

	// UpgradableFlag - This flag indicates whether this storage type is upgradable or not.
	UpgradableFlag bool `json:"upgradableFlag,omitempty"`

	// HostId - The unique identification number of the host associated with a Storage volume.
	HostId int `json:"hostId,omitempty"`

	// AccountId - The internal identifier of the SoftLayer customer account that a Storage account belongs
	// to.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// StorageTypeId - no documentation
	StorageTypeId string `json:"storageTypeId,omitempty"`

	// Username - The username used to access a non-EVault Storage volume. This username is used to
	// register the EVault server agent with the vault backup system.
	Username string `json:"username,omitempty"`

	// GuestId - The unique identification number of the guest associated with a Storage volume.
	GuestId int `json:"guestId,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// Password - The password used to access a non-EVault Storage volume. This password is used to
	// register the EVault server agent with the vault backup system.
	Password string `json:"password,omitempty"`

	// RestoreJobDetails - All the of the restore jobs for the EVault Storage account.
	RestoreJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"restoreJobDetails,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// ReplicationStatus - The current replication status of a network storage volume. Indicates Failover
	// or Failback status.
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// UsageNotification - no documentation
	UsageNotification *SoftLayer_Notification `json:"usageNotification,omitempty"`

	// AllowedIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed access to this
	// storage volume.
	AllowedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedIpAddresses,omitempty"`

	// SnapshotSizeBytes - no documentation
	SnapshotSizeBytes string `json:"snapshotSizeBytes,omitempty"`

	// AllowedSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed access to
	// this storage volume.
	AllowedSubnetCount uint64 `json:"allowedSubnetCount,omitempty"`

	// EventCount - A count of the events which have taken place on a network storage volume.
	EventCount uint64 `json:"eventCount,omitempty"`

	// ActiveTransactionCount - A count of the currently active transactions on a network storage volume.
	ActiveTransactionCount uint64 `json:"activeTransactionCount,omitempty"`

	// ScheduleCount - A count of the schedules which are associated with a network storage volume.
	ScheduleCount uint64 `json:"scheduleCount,omitempty"`

	// ParentPartnerships - The volumes or snapshots partnered with a network storage volume in a parental
	// role.
	ParentPartnerships []*SoftLayer_Network_Storage_Partnership `json:"parentPartnerships,omitempty"`

	// AllowedVirtualGuests - The SoftLayer_Virtual_Guest objects which are allowed access to this storage
	// volume.
	AllowedVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedVirtualGuests,omitempty"`

	// BillingItemCategory - <nil>
	BillingItemCategory *SoftLayer_Product_Item_Category `json:"billingItemCategory,omitempty"`

	// ReplicatingLunCount - A count of the iSCSI LUN volumes being replicated by this network storage
	// volume.
	ReplicatingLunCount uint64 `json:"replicatingLunCount,omitempty"`

	// MetricTrackingObject - A network storage volume's metric tracking object. This object records all
	// periodic polled data available to this volume.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// AllowedHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access to this
	// storage volume.
	AllowedHardwareCount uint64 `json:"allowedHardwareCount,omitempty"`

	// AllowedIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume.
	AllowedIpAddressCount uint64 `json:"allowedIpAddressCount,omitempty"`

	// Snapshots - The snapshots associated with this SoftLayer_Network_Storage volume.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots,omitempty"`

	// IscsiLuns - Relationship between a container volume and iSCSI LUNs.
	IscsiLuns []*SoftLayer_Network_Storage `json:"iscsiLuns,omitempty"`

	// VolumeHistory - The username and password history for a Storage service.
	VolumeHistory []*SoftLayer_Network_Storage_History `json:"volumeHistory,omitempty"`

	// AllowedHardware - The SoftLayer_Hardware objects which are allowed access to this storage volume.
	AllowedHardware []*SoftLayer_Hardware `json:"allowedHardware,omitempty"`

	// AllowedReplicationSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationSubnets []*SoftLayer_Network_Subnet `json:"allowedReplicationSubnets,omitempty"`

	// AgentStatusCount - A count of statuses (most of the time will be one status) for the agent tied to
	// the EVault Storage services.
	AgentStatusCount uint64 `json:"agentStatusCount,omitempty"`

	// ServiceResourceName - no documentation
	ServiceResourceName string `json:"serviceResourceName,omitempty"`

	// ReplicationEvents - no documentation
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents,omitempty"`

	// AllowedReplicationHardware - The SoftLayer_Hardware objects which are allowed access to this storage
	// volume's Replicant.
	AllowedReplicationHardware []*SoftLayer_Hardware `json:"allowedReplicationHardware,omitempty"`

	// BackupJobDetails - All the of the backup jobs for the EVault Storage account.
	BackupJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"backupJobDetails,omitempty"`

	// AllowedReplicationIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which
	// are allowed access to this storage volume's Replicant.
	AllowedReplicationIpAddressCount uint64 `json:"allowedReplicationIpAddressCount,omitempty"`

	// PartnershipCount - A count of the volumes or snapshots partnered with a network storage volume.
	PartnershipCount uint64 `json:"partnershipCount,omitempty"`

	// StorageGroups - The network storage groups this volume is attached to.
	StorageGroups []*SoftLayer_Network_Storage_Group `json:"storageGroups,omitempty"`

	// ReplicationEventCount - no documentation
	ReplicationEventCount uint64 `json:"replicationEventCount,omitempty"`

	// Credentials - <nil>
	Credentials []*SoftLayer_Network_Storage_Credential `json:"credentials,omitempty"`

	// NotificationSubscribers - The subscribers that will be notified for usage amount warnings and
	// overages.
	NotificationSubscribers []*SoftLayer_Notification_User_Subscriber `json:"notificationSubscribers,omitempty"`

	// VolumeStatus - no documentation
	VolumeStatus string `json:"volumeStatus,omitempty"`

	// PermissionsGroups - no documentation
	PermissionsGroups []*SoftLayer_Network_Storage_Group `json:"permissionsGroups,omitempty"`

	// OsType - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type.
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType,omitempty"`

	// VendorName - no documentation
	VendorName string `json:"vendorName,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// OsTypeId - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type
	OsTypeId string `json:"osTypeId,omitempty"`

	// TotalScheduleSnapshotRetentionCount - The total snapshot retention count of all schedules on this
	// network storage volume.
	TotalScheduleSnapshotRetentionCount uint `json:"totalScheduleSnapshotRetentionCount,omitempty"`

	// AllowedReplicationVirtualGuests - The SoftLayer_Hardware objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedReplicationVirtualGuests,omitempty"`

	// DailySchedule - The Daily Schedule which is associated with this network storage volume.
	DailySchedule *SoftLayer_Network_Storage_Schedule `json:"dailySchedule,omitempty"`

	// AgentStatuses - Statuses (most of the time will be one status) for the agent tied to the EVault
	// Storage services.
	AgentStatuses []*SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus `json:"agentStatuses,omitempty"`

	// Hardware - When applicable, the hardware associated with a Storage service.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// Iops - The maximum number of IOPs selected for this volume.
	Iops string `json:"iops,omitempty"`

	// AllowedSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this storage
	// volume.
	AllowedSubnets []*SoftLayer_Network_Subnet `json:"allowedSubnets,omitempty"`

	// ReplicatingLuns - The iSCSI LUN volumes being replicated by this network storage volume.
	ReplicatingLuns []*SoftLayer_Network_Storage `json:"replicatingLuns,omitempty"`

	// BytesUsed - no documentation
	BytesUsed string `json:"bytesUsed,omitempty"`

	// AllowedReplicationIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedReplicationIpAddresses,omitempty"`

	// ReplicationPartnerCount - A count of the network storage volumes configured to be replicants of a
	// volume.
	ReplicationPartnerCount uint64 `json:"replicationPartnerCount,omitempty"`

	// SnapshotCount - A count of the snapshots associated with this SoftLayer_Network_Storage volume.
	SnapshotCount uint64 `json:"snapshotCount,omitempty"`

	// Events - The events which have taken place on a network storage volume.
	Events []*SoftLayer_Network_Storage_Event `json:"events,omitempty"`

	// StorageType - no documentation
	StorageType *SoftLayer_Network_Storage_Type `json:"storageType,omitempty"`

	// AccountPassword - Other usernames and passwords associated with a Storage volume.
	AccountPassword *SoftLayer_Account_Password `json:"accountPassword,omitempty"`

	// HourlySchedule - The Hourly Schedule which is associated with this network storage volume.
	HourlySchedule *SoftLayer_Network_Storage_Schedule `json:"hourlySchedule,omitempty"`

	// CurrentCyclePeakUsage - Peak number of bytes used in the vault for the current billing cycle.
	CurrentCyclePeakUsage uint `json:"currentCyclePeakUsage,omitempty"`

	// PluginBillingItemCount - A count of the billing items for plugins tied to the EVault Storage
	// service.
	PluginBillingItemCount uint64 `json:"pluginBillingItemCount,omitempty"`

	// ReplicationPartners - The network storage volumes configured to be replicants of a volume.
	ReplicationPartners []*SoftLayer_Network_Storage `json:"replicationPartners,omitempty"`

	// VirtualGuest - When applicable, the virtual guest associated with a Storage service.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// AllowedReplicationVirtualGuestCount - A count of the SoftLayer_Hardware objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationVirtualGuestCount uint64 `json:"allowedReplicationVirtualGuestCount,omitempty"`

	// SnapshotSpaceAvailable - no documentation
	SnapshotSpaceAvailable string `json:"snapshotSpaceAvailable,omitempty"`

	// CreationScheduleId - The schedule id which was executed to create a snapshot.
	CreationScheduleId string `json:"creationScheduleId,omitempty"`

	// StorageTierLevel - <nil>
	StorageTierLevel string `json:"storageTierLevel,omitempty"`

	// ServiceResource - The network resource a Storage service is connected to.
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource,omitempty"`

	// SnapshotCreationTimestamp - The creation timestamp of the snapshot on the storage platform.
	SnapshotCreationTimestamp string `json:"snapshotCreationTimestamp,omitempty"`

	// Tasks - Retrieve the task information for the EVault Storage service.
	Tasks []*SoftLayer_Container_Network_Storage_Evault_Vault_Task `json:"tasks,omitempty"`

	// RestoreJobDetailCount - A count of all the of the restore jobs for the EVault Storage account.
	RestoreJobDetailCount uint64 `json:"restoreJobDetailCount,omitempty"`

	// CredentialCount - no documentation
	CredentialCount uint64 `json:"credentialCount,omitempty"`

	// SnapshotDeletionThresholdPercentage - The percentage of used snapshot space after which to delete
	// automated snapshots.
	SnapshotDeletionThresholdPercentage string `json:"snapshotDeletionThresholdPercentage,omitempty"`

	// StorageGroupCount - A count of the network storage groups this volume is attached to.
	StorageGroupCount uint64 `json:"storageGroupCount,omitempty"`

	// AllowedReplicationSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationSubnetCount uint64 `json:"allowedReplicationSubnetCount,omitempty"`

	// MountableFlag - Whether or not a network storage volume may be mounted.
	MountableFlag string `json:"mountableFlag,omitempty"`

	// TotalBytesUsed - no documentation
	TotalBytesUsed string `json:"totalBytesUsed,omitempty"`

	// ActiveTransactions - The currently active transactions on a network storage volume.
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions,omitempty"`

	// SoftwareComponent - no documentation
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent,omitempty"`

	// PropertyCount - A count of the properties used to provide additional details about a network storage
	// volume.
	PropertyCount uint64 `json:"propertyCount,omitempty"`

	// SnapshotCapacityGb - no documentation
	SnapshotCapacityGb string `json:"snapshotCapacityGb,omitempty"`

	// ParentVolume - The parent volume of a volume in a complex storage relationship.
	ParentVolume *SoftLayer_Network_Storage `json:"parentVolume,omitempty"`

	// ServiceResourceBackendIpAddress - no documentation
	ServiceResourceBackendIpAddress string `json:"serviceResourceBackendIpAddress,omitempty"`

	// PluginBillingItems - The billing items for plugins tied to the EVault Storage service.
	PluginBillingItems []*SoftLayer_Billing_Item `json:"pluginBillingItems,omitempty"`

	// AllowedVirtualGuestCount - A count of the SoftLayer_Virtual_Guest objects which are allowed access
	// to this storage volume.
	AllowedVirtualGuestCount uint64 `json:"allowedVirtualGuestCount,omitempty"`

	// Schedules - The schedules which are associated with a network storage volume.
	Schedules []*SoftLayer_Network_Storage_Schedule `json:"schedules,omitempty"`

	// PreviousCyclePeakUsage - Peak number of bytes used in the vault for the previous billing cycle.
	PreviousCyclePeakUsage uint `json:"previousCyclePeakUsage,omitempty"`

	// BackupJobDetailCount - A count of all the of the backup jobs for the EVault Storage account.
	BackupJobDetailCount uint64 `json:"backupJobDetailCount,omitempty"`

	// AllowedReplicationHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access
	// to this storage volume's Replicant.
	AllowedReplicationHardwareCount uint64 `json:"allowedReplicationHardwareCount,omitempty"`

	// IscsiLunCount - A count of relationship between a container volume and iSCSI LUNs.
	IscsiLunCount uint64 `json:"iscsiLunCount,omitempty"`

	// WeeklySchedule - The Weekly Schedule which is associated with this network storage volume.
	WeeklySchedule *SoftLayer_Network_Storage_Schedule `json:"weeklySchedule,omitempty"`

	// Properties - The properties used to provide additional details about a network storage volume.
	Properties []*SoftLayer_Network_Storage_Property `json:"properties,omitempty"`

	// ParentPartnershipCount - A count of the volumes or snapshots partnered with a network storage volume
	// in a parental role.
	ParentPartnershipCount uint64 `json:"parentPartnershipCount,omitempty"`

	// LunId - no documentation
	LunId string `json:"lunId,omitempty"`

	// Partnerships - The volumes or snapshots partnered with a network storage volume.
	Partnerships []*SoftLayer_Network_Storage_Partnership `json:"partnerships,omitempty"`

	// WebccAccount - The account username and password for the EVault webCC interface.
	WebccAccount *SoftLayer_Account_Password `json:"webccAccount,omitempty"`

	// ReplicatingVolume - The network storage volume being replicated by a volume.
	ReplicatingVolume *SoftLayer_Network_Storage `json:"replicatingVolume,omitempty"`

	// PermissionsGroupCount - A count of all permissions group(s) this volume is in.
	PermissionsGroupCount uint64 `json:"permissionsGroupCount,omitempty"`

	// TaskCount - A count of retrieve the task information for the EVault Storage service.
	TaskCount uint64 `json:"taskCount,omitempty"`

	// VolumeHistoryCount - A count of the username and password history for a Storage service.
	VolumeHistoryCount uint64 `json:"volumeHistoryCount,omitempty"`

	// NotificationSubscriberCount - A count of the subscribers that will be notified for usage amount
	// warnings and overages.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount,omitempty"`

	// ReplicationSchedule - The Replication Schedule associated with a network storage volume.
	ReplicationSchedule *SoftLayer_Network_Storage_Schedule `json:"replicationSchedule,omitempty"`
}

func (softlayer_network_storage_backup_evault_version6 *SoftLayer_Network_Storage_Backup_Evault_Version6) String() string {
	return "SoftLayer_Network_Storage_Backup_Evault_Version6"
}
