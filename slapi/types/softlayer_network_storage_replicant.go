package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Replicant - <nil>
type SoftLayer_Network_Storage_Replicant struct {

	// UpgradableFlag - This flag indicates whether this storage type is upgradable or not.
	UpgradableFlag bool `json:"upgradableFlag,omitempty"`

	// CapacityGb - A Storage account's capacity, measured in gigabytes.
	CapacityGb int `json:"capacityGb,omitempty"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// StorageTypeId - no documentation
	StorageTypeId string `json:"storageTypeId,omitempty"`

	// NasType - A Storage account's type. Valid examples are and
	NasType string `json:"nasType,omitempty"`

	// AccountId - The internal identifier of the SoftLayer customer account that a Storage account belongs
	// to.
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// HostId - The unique identification number of the host associated with a Storage volume.
	HostId int `json:"hostId,omitempty"`

	// HardwareId - The server that is associated with a Storage service.
	HardwareId int `json:"hardwareId,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// GuestId - The unique identification number of the guest associated with a Storage volume.
	GuestId int `json:"guestId,omitempty"`

	// Password - The password used to access a non-EVault Storage volume. This password is used to
	// register the EVault server agent with the vault backup system.
	Password string `json:"password,omitempty"`

	// Username - The username used to access a non-EVault Storage volume. This username is used to
	// register the EVault server agent with the vault backup system.
	Username string `json:"username,omitempty"`

	// OsTypeId - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type
	OsTypeId string `json:"osTypeId,omitempty"`

	// AllowedReplicationVirtualGuestCount - A count of the SoftLayer_Hardware objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationVirtualGuestCount uint64 `json:"allowedReplicationVirtualGuestCount,omitempty"`

	// AllowedSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this storage
	// volume.
	AllowedSubnets []*SoftLayer_Network_Subnet `json:"allowedSubnets,omitempty"`

	// StorageTierLevel - <nil>
	StorageTierLevel string `json:"storageTierLevel,omitempty"`

	// BytesUsed - no documentation
	BytesUsed string `json:"bytesUsed,omitempty"`

	// Hardware - When applicable, the hardware associated with a Storage service.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// StorageType - no documentation
	StorageType *SoftLayer_Network_Storage_Type `json:"storageType,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// AllowedReplicationHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access
	// to this storage volume's Replicant.
	AllowedReplicationHardwareCount uint64 `json:"allowedReplicationHardwareCount,omitempty"`

	// TotalBytesUsed - no documentation
	TotalBytesUsed string `json:"totalBytesUsed,omitempty"`

	// IscsiLuns - Relationship between a container volume and iSCSI LUNs.
	IscsiLuns []*SoftLayer_Network_Storage `json:"iscsiLuns,omitempty"`

	// AllowedSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed access to
	// this storage volume.
	AllowedSubnetCount uint64 `json:"allowedSubnetCount,omitempty"`

	// NotificationSubscriberCount - A count of the subscribers that will be notified for usage amount
	// warnings and overages.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount,omitempty"`

	// StorageGroups - The network storage groups this volume is attached to.
	StorageGroups []*SoftLayer_Network_Storage_Group `json:"storageGroups,omitempty"`

	// VendorName - no documentation
	VendorName string `json:"vendorName,omitempty"`

	// SnapshotSizeBytes - no documentation
	SnapshotSizeBytes string `json:"snapshotSizeBytes,omitempty"`

	// Schedules - The schedules which are associated with a network storage volume.
	Schedules []*SoftLayer_Network_Storage_Schedule `json:"schedules,omitempty"`

	// VolumeStatus - no documentation
	VolumeStatus string `json:"volumeStatus,omitempty"`

	// ParentVolume - The parent volume of a volume in a complex storage relationship.
	ParentVolume *SoftLayer_Network_Storage `json:"parentVolume,omitempty"`

	// WeeklySchedule - The Weekly Schedule which is associated with this network storage volume.
	WeeklySchedule *SoftLayer_Network_Storage_Schedule `json:"weeklySchedule,omitempty"`

	// AllowedReplicationIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which
	// are allowed access to this storage volume's Replicant.
	AllowedReplicationIpAddressCount uint64 `json:"allowedReplicationIpAddressCount,omitempty"`

	// CredentialCount - no documentation
	CredentialCount uint64 `json:"credentialCount,omitempty"`

	// OsType - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type.
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType,omitempty"`

	// ServiceResourceName - no documentation
	ServiceResourceName string `json:"serviceResourceName,omitempty"`

	// StorageGroupCount - A count of the network storage groups this volume is attached to.
	StorageGroupCount uint64 `json:"storageGroupCount,omitempty"`

	// Snapshots - The snapshots associated with this SoftLayer_Network_Storage volume.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots,omitempty"`

	// ServiceResource - The network resource a Storage service is connected to.
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource,omitempty"`

	// EventCount - A count of the events which have taken place on a network storage volume.
	EventCount uint64 `json:"eventCount,omitempty"`

	// MetricTrackingObject - A network storage volume's metric tracking object. This object records all
	// periodic polled data available to this volume.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// VolumeHistoryCount - A count of the username and password history for a Storage service.
	VolumeHistoryCount uint64 `json:"volumeHistoryCount,omitempty"`

	// ActiveTransactionCount - A count of the currently active transactions on a network storage volume.
	ActiveTransactionCount uint64 `json:"activeTransactionCount,omitempty"`

	// FailbackInProgressFlag - When a replicant is in the process of synchronizing with the parent volume
	// this flag will be true.
	FailbackInProgressFlag string `json:"failbackInProgressFlag,omitempty"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// HourlySchedule - The Hourly Schedule which is associated with this network storage volume.
	HourlySchedule *SoftLayer_Network_Storage_Schedule `json:"hourlySchedule,omitempty"`

	// Properties - The properties used to provide additional details about a network storage volume.
	Properties []*SoftLayer_Network_Storage_Property `json:"properties,omitempty"`

	// SnapshotSpaceAvailable - no documentation
	SnapshotSpaceAvailable string `json:"snapshotSpaceAvailable,omitempty"`

	// ParentPartnershipCount - A count of the volumes or snapshots partnered with a network storage volume
	// in a parental role.
	ParentPartnershipCount uint64 `json:"parentPartnershipCount,omitempty"`

	// PermissionsGroups - no documentation
	PermissionsGroups []*SoftLayer_Network_Storage_Group `json:"permissionsGroups,omitempty"`

	// AllowedReplicationIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedReplicationIpAddresses,omitempty"`

	// PartnershipCount - A count of the volumes or snapshots partnered with a network storage volume.
	PartnershipCount uint64 `json:"partnershipCount,omitempty"`

	// AllowedHardware - The SoftLayer_Hardware objects which are allowed access to this storage volume.
	AllowedHardware []*SoftLayer_Hardware `json:"allowedHardware,omitempty"`

	// AllowedIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume.
	AllowedIpAddressCount uint64 `json:"allowedIpAddressCount,omitempty"`

	// AllowedHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access to this
	// storage volume.
	AllowedHardwareCount uint64 `json:"allowedHardwareCount,omitempty"`

	// PermissionsGroupCount - A count of all permissions group(s) this volume is in.
	PermissionsGroupCount uint64 `json:"permissionsGroupCount,omitempty"`

	// ReplicationEvents - no documentation
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents,omitempty"`

	// ReplicationStatus - The current replication status of a network storage volume. Indicates Failover
	// or Failback status.
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// SnapshotCapacityGb - no documentation
	SnapshotCapacityGb string `json:"snapshotCapacityGb,omitempty"`

	// ReplicationSchedule - The Replication Schedule associated with a network storage volume.
	ReplicationSchedule *SoftLayer_Network_Storage_Schedule `json:"replicationSchedule,omitempty"`

	// MountableFlag - Whether or not a network storage volume may be mounted.
	MountableFlag string `json:"mountableFlag,omitempty"`

	// Iops - The maximum number of IOPs selected for this volume.
	Iops string `json:"iops,omitempty"`

	// ReplicationEventCount - no documentation
	ReplicationEventCount uint64 `json:"replicationEventCount,omitempty"`

	// AllowedReplicationSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationSubnets []*SoftLayer_Network_Subnet `json:"allowedReplicationSubnets,omitempty"`

	// ReplicatingLunCount - A count of the iSCSI LUN volumes being replicated by this network storage
	// volume.
	ReplicatingLunCount uint64 `json:"replicatingLunCount,omitempty"`

	// PropertyCount - A count of the properties used to provide additional details about a network storage
	// volume.
	PropertyCount uint64 `json:"propertyCount,omitempty"`

	// TotalScheduleSnapshotRetentionCount - The total snapshot retention count of all schedules on this
	// network storage volume.
	TotalScheduleSnapshotRetentionCount uint `json:"totalScheduleSnapshotRetentionCount,omitempty"`

	// BillingItemCategory - <nil>
	BillingItemCategory *SoftLayer_Product_Item_Category `json:"billingItemCategory,omitempty"`

	// Credentials - <nil>
	Credentials []*SoftLayer_Network_Storage_Credential `json:"credentials,omitempty"`

	// AllowedReplicationSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationSubnetCount uint64 `json:"allowedReplicationSubnetCount,omitempty"`

	// ActiveTransactions - The currently active transactions on a network storage volume.
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions,omitempty"`

	// VolumeHistory - The username and password history for a Storage service.
	VolumeHistory []*SoftLayer_Network_Storage_History `json:"volumeHistory,omitempty"`

	// AllowedVirtualGuests - The SoftLayer_Virtual_Guest objects which are allowed access to this storage
	// volume.
	AllowedVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedVirtualGuests,omitempty"`

	// SnapshotCreationTimestamp - The creation timestamp of the snapshot on the storage platform.
	SnapshotCreationTimestamp string `json:"snapshotCreationTimestamp,omitempty"`

	// ReplicationPartners - The network storage volumes configured to be replicants of a volume.
	ReplicationPartners []*SoftLayer_Network_Storage `json:"replicationPartners,omitempty"`

	// LunId - no documentation
	LunId string `json:"lunId,omitempty"`

	// AllowedReplicationHardware - The SoftLayer_Hardware objects which are allowed access to this storage
	// volume's Replicant.
	AllowedReplicationHardware []*SoftLayer_Hardware `json:"allowedReplicationHardware,omitempty"`

	// UsageNotification - no documentation
	UsageNotification *SoftLayer_Notification `json:"usageNotification,omitempty"`

	// ReplicatingLuns - The iSCSI LUN volumes being replicated by this network storage volume.
	ReplicatingLuns []*SoftLayer_Network_Storage `json:"replicatingLuns,omitempty"`

	// DailySchedule - The Daily Schedule which is associated with this network storage volume.
	DailySchedule *SoftLayer_Network_Storage_Schedule `json:"dailySchedule,omitempty"`

	// ServiceResourceBackendIpAddress - no documentation
	ServiceResourceBackendIpAddress string `json:"serviceResourceBackendIpAddress,omitempty"`

	// SnapshotCount - A count of the snapshots associated with this SoftLayer_Network_Storage volume.
	SnapshotCount uint64 `json:"snapshotCount,omitempty"`

	// AllowedVirtualGuestCount - A count of the SoftLayer_Virtual_Guest objects which are allowed access
	// to this storage volume.
	AllowedVirtualGuestCount uint64 `json:"allowedVirtualGuestCount,omitempty"`

	// ReplicationPartnerCount - A count of the network storage volumes configured to be replicants of a
	// volume.
	ReplicationPartnerCount uint64 `json:"replicationPartnerCount,omitempty"`

	// ParentPartnerships - The volumes or snapshots partnered with a network storage volume in a parental
	// role.
	ParentPartnerships []*SoftLayer_Network_Storage_Partnership `json:"parentPartnerships,omitempty"`

	// Events - The events which have taken place on a network storage volume.
	Events []*SoftLayer_Network_Storage_Event `json:"events,omitempty"`

	// NotificationSubscribers - The subscribers that will be notified for usage amount warnings and
	// overages.
	NotificationSubscribers []*SoftLayer_Notification_User_Subscriber `json:"notificationSubscribers,omitempty"`

	// ScheduleCount - A count of the schedules which are associated with a network storage volume.
	ScheduleCount uint64 `json:"scheduleCount,omitempty"`

	// CreationScheduleId - The schedule id which was executed to create a snapshot.
	CreationScheduleId string `json:"creationScheduleId,omitempty"`

	// ReplicatingVolume - The network storage volume being replicated by a volume.
	ReplicatingVolume *SoftLayer_Network_Storage `json:"replicatingVolume,omitempty"`

	// WebccAccount - The account username and password for the EVault webCC interface.
	WebccAccount *SoftLayer_Account_Password `json:"webccAccount,omitempty"`

	// VirtualGuest - When applicable, the virtual guest associated with a Storage service.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// SnapshotDeletionThresholdPercentage - The percentage of used snapshot space after which to delete
	// automated snapshots.
	SnapshotDeletionThresholdPercentage string `json:"snapshotDeletionThresholdPercentage,omitempty"`

	// IscsiLunCount - A count of relationship between a container volume and iSCSI LUNs.
	IscsiLunCount uint64 `json:"iscsiLunCount,omitempty"`

	// Partnerships - The volumes or snapshots partnered with a network storage volume.
	Partnerships []*SoftLayer_Network_Storage_Partnership `json:"partnerships,omitempty"`

	// AccountPassword - Other usernames and passwords associated with a Storage volume.
	AccountPassword *SoftLayer_Account_Password `json:"accountPassword,omitempty"`

	// AllowedIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed access to this
	// storage volume.
	AllowedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedIpAddresses,omitempty"`

	// AllowedReplicationVirtualGuests - The SoftLayer_Hardware objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedReplicationVirtualGuests,omitempty"`
}

func (softlayer_network_storage_replicant *SoftLayer_Network_Storage_Replicant) String() string {
	return "SoftLayer_Network_Storage_Replicant"
}
