package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant - An iscsi replicant receives incoming
// data from an associated iscsi volume. While the replicant is not in failover mode it will not be
// mountable. Upon failover the replicant can be mounted and used as a normal volume. It is suggested
// to only do this as part of a disaster recovery plan.
type SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant struct {

	// Password - The password used to access a non-EVault Storage volume. This password is used to
	// register the EVault server agent with the vault backup system.
	Password string `json:"password,omitempty"`

	// NasType - A Storage account's type. Valid examples are and
	NasType string `json:"nasType,omitempty"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// HardwareId - The server that is associated with a Storage service.
	HardwareId int `json:"hardwareId,omitempty"`

	// UpgradableFlag - This flag indicates whether this storage type is upgradable or not.
	UpgradableFlag bool `json:"upgradableFlag,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CapacityGb - A Storage account's capacity, measured in gigabytes.
	CapacityGb int `json:"capacityGb,omitempty"`

	// AccountId - The internal identifier of the SoftLayer customer account that a Storage account belongs
	// to.
	AccountId int `json:"accountId,omitempty"`

	// GuestId - The unique identification number of the guest associated with a Storage volume.
	GuestId int `json:"guestId,omitempty"`

	// HostId - The unique identification number of the host associated with a Storage volume.
	HostId int `json:"hostId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Username - The username used to access a non-EVault Storage volume. This username is used to
	// register the EVault server agent with the vault backup system.
	Username string `json:"username,omitempty"`
}

func (softlayer_network_storage_iscsi_equallogic_version3_replicant *SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant) String() string {
	return "SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant"
}

// SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant_Extended is SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant with all maskable types.
type SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant_Extended struct {
	SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant

	// Hardware - When applicable, the hardware associated with a Storage service.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// BytesUsed - no documentation
	BytesUsed string `json:"bytesUsed,omitempty"`

	// SnapshotCreationTimestamp - The creation timestamp of the snapshot on the storage platform.
	SnapshotCreationTimestamp string `json:"snapshotCreationTimestamp,omitempty"`

	// Iops - The maximum number of IOPs selected for this volume.
	Iops string `json:"iops,omitempty"`

	// ReplicationPartnerCount - A count of the network storage volumes configured to be replicants of a
	// volume.
	ReplicationPartnerCount uint64 `json:"replicationPartnerCount,omitempty"`

	// HourlySchedule - The Hourly Schedule which is associated with this network storage volume.
	HourlySchedule *SoftLayer_Network_Storage_Schedule `json:"hourlySchedule,omitempty"`

	// AllowedReplicationVirtualGuests - The SoftLayer_Hardware objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedReplicationVirtualGuests,omitempty"`

	// ReplicationStatus - The current replication status of a network storage volume. Indicates Failover
	// or Failback status.
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// Schedules - The schedules which are associated with a network storage volume.
	Schedules []*SoftLayer_Network_Storage_Schedule `json:"schedules,omitempty"`

	// MetricTrackingObject - A network storage volume's metric tracking object. This object records all
	// periodic polled data available to this volume.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject,omitempty"`

	// CreationScheduleId - The schedule id which was executed to create a snapshot.
	CreationScheduleId string `json:"creationScheduleId,omitempty"`

	// ReplicatingLuns - The iSCSI LUN volumes being replicated by this network storage volume.
	ReplicatingLuns []*SoftLayer_Network_Storage `json:"replicatingLuns,omitempty"`

	// ServiceResourceBackendIpAddress - no documentation
	ServiceResourceBackendIpAddress string `json:"serviceResourceBackendIpAddress,omitempty"`

	// VolumeName - no documentation
	VolumeName string `json:"volumeName,omitempty"`

	// MountableFlag - Whether or not a network storage volume may be mounted.
	MountableFlag string `json:"mountableFlag,omitempty"`

	// CredentialCount - no documentation
	CredentialCount uint64 `json:"credentialCount,omitempty"`

	// PermissionsGroups - no documentation
	PermissionsGroups []*SoftLayer_Network_Storage_Group `json:"permissionsGroups,omitempty"`

	// WeeklySchedule - The Weekly Schedule which is associated with this network storage volume.
	WeeklySchedule *SoftLayer_Network_Storage_Schedule `json:"weeklySchedule,omitempty"`

	// SnapshotSizeBytes - no documentation
	SnapshotSizeBytes string `json:"snapshotSizeBytes,omitempty"`

	// AllowedReplicationSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationSubnetCount uint64 `json:"allowedReplicationSubnetCount,omitempty"`

	// PermissionsGroupCount - A count of all permissions group(s) this volume is in.
	PermissionsGroupCount uint64 `json:"permissionsGroupCount,omitempty"`

	// Credentials - <nil>
	Credentials []*SoftLayer_Network_Storage_Credential `json:"credentials,omitempty"`

	// ParentPartnershipCount - A count of the volumes or snapshots partnered with a network storage volume
	// in a parental role.
	ParentPartnershipCount uint64 `json:"parentPartnershipCount,omitempty"`

	// FailbackInProgressFlag - When a replicant is in the process of synchronizing with the parent volume
	// this flag will be true.
	FailbackInProgressFlag bool `json:"failbackInProgressFlag,omitempty"`

	// ServiceResource - The network resource a Storage service is connected to.
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource,omitempty"`

	// VolumeStatus - no documentation
	VolumeStatus string `json:"volumeStatus,omitempty"`

	// AllowedReplicationHardware - The SoftLayer_Hardware objects which are allowed access to this storage
	// volume's Replicant.
	AllowedReplicationHardware []*SoftLayer_Hardware `json:"allowedReplicationHardware,omitempty"`

	// WebccAccount - The account username and password for the EVault webCC interface.
	WebccAccount *SoftLayer_Account_Password `json:"webccAccount,omitempty"`

	// ReplicatingVolume - The network storage volume being replicated by a volume.
	ReplicatingVolume *SoftLayer_Network_Storage `json:"replicatingVolume,omitempty"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem,omitempty"`

	// ParentVolume - The parent volume of a volume in a complex storage relationship.
	ParentVolume *SoftLayer_Network_Storage `json:"parentVolume,omitempty"`

	// AllowableVirtualGuestCount - A count of the list of Virtual Guest that can be attached to this
	// storage.
	AllowableVirtualGuestCount uint64 `json:"allowableVirtualGuestCount,omitempty"`

	// AllowedHardware - The SoftLayer_Hardware objects which are allowed access to this storage volume.
	AllowedHardware []*SoftLayer_Hardware `json:"allowedHardware,omitempty"`

	// IscsiLuns - Relationship between a container volume and iSCSI LUNs.
	IscsiLuns []*SoftLayer_Network_Storage `json:"iscsiLuns,omitempty"`

	// ActiveTransactionCount - A count of the currently active transactions on a network storage volume.
	ActiveTransactionCount uint64 `json:"activeTransactionCount,omitempty"`

	// TotalBytesUsed - no documentation
	TotalBytesUsed string `json:"totalBytesUsed,omitempty"`

	// AllowedSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this storage
	// volume.
	AllowedSubnets []*SoftLayer_Network_Subnet `json:"allowedSubnets,omitempty"`

	// ReplicationEventCount - no documentation
	ReplicationEventCount uint64 `json:"replicationEventCount,omitempty"`

	// StorageGroupCount - A count of the network storage groups this volume is attached to.
	StorageGroupCount uint64 `json:"storageGroupCount,omitempty"`

	// AccountPassword - Other usernames and passwords associated with a Storage volume.
	AccountPassword *SoftLayer_Account_Password `json:"accountPassword,omitempty"`

	// Properties - The properties used to provide additional details about a network storage volume.
	Properties []*SoftLayer_Network_Storage_Property `json:"properties,omitempty"`

	// DailySchedule - The Daily Schedule which is associated with this network storage volume.
	DailySchedule *SoftLayer_Network_Storage_Schedule `json:"dailySchedule,omitempty"`

	// VendorName - no documentation
	VendorName string `json:"vendorName,omitempty"`

	// LunId - no documentation
	LunId string `json:"lunId,omitempty"`

	// ParentPartnerships - The volumes or snapshots partnered with a network storage volume in a parental
	// role.
	ParentPartnerships []*SoftLayer_Network_Storage_Partnership `json:"parentPartnerships,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// SnapshotSpaceAvailable - no documentation
	SnapshotSpaceAvailable string `json:"snapshotSpaceAvailable,omitempty"`

	// AllowedReplicationIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedReplicationIpAddresses,omitempty"`

	// AllowedReplicationSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationSubnets []*SoftLayer_Network_Subnet `json:"allowedReplicationSubnets,omitempty"`

	// TotalScheduleSnapshotRetentionCount - The total snapshot retention count of all schedules on this
	// network storage volume.
	TotalScheduleSnapshotRetentionCount uint `json:"totalScheduleSnapshotRetentionCount,omitempty"`

	// SnapshotDeletionThresholdPercentage - The percentage of used snapshot space after which to delete
	// automated snapshots.
	SnapshotDeletionThresholdPercentage string `json:"snapshotDeletionThresholdPercentage,omitempty"`

	// AllowedIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed access to this
	// storage volume.
	AllowedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedIpAddresses,omitempty"`

	// IscsiLunCount - A count of relationship between a container volume and iSCSI LUNs.
	IscsiLunCount uint64 `json:"iscsiLunCount,omitempty"`

	// ActiveTransactions - The currently active transactions on a network storage volume.
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions,omitempty"`

	// PartnershipCount - A count of the volumes or snapshots partnered with a network storage volume.
	PartnershipCount uint64 `json:"partnershipCount,omitempty"`

	// VolumeHistoryCount - A count of the username and password history for a Storage service.
	VolumeHistoryCount uint64 `json:"volumeHistoryCount,omitempty"`

	// ReplicationEvents - no documentation
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents,omitempty"`

	// SnapshotCount - A count of the snapshots associated with this SoftLayer_Network_Storage volume.
	SnapshotCount uint64 `json:"snapshotCount,omitempty"`

	// AllowableHardwareCount - A count of the list of Hardware that can be attached to this storage.
	AllowableHardwareCount uint64 `json:"allowableHardwareCount,omitempty"`

	// BillingItemCategory - <nil>
	BillingItemCategory *SoftLayer_Product_Item_Category `json:"billingItemCategory,omitempty"`

	// NotificationSubscriberCount - A count of the subscribers that will be notified for usage amount
	// warnings and overages.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount,omitempty"`

	// AllowedSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed access to
	// this storage volume.
	AllowedSubnetCount uint64 `json:"allowedSubnetCount,omitempty"`

	// VirtualGuest - When applicable, the virtual guest associated with a Storage service.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest,omitempty"`

	// Partnerships - The volumes or snapshots partnered with a network storage volume.
	Partnerships []*SoftLayer_Network_Storage_Partnership `json:"partnerships,omitempty"`

	// Events - The events which have taken place on a network storage volume.
	Events []*SoftLayer_Network_Storage_Event `json:"events,omitempty"`

	// ServiceResourceName - no documentation
	ServiceResourceName string `json:"serviceResourceName,omitempty"`

	// AllowedHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access to this
	// storage volume.
	AllowedHardwareCount uint64 `json:"allowedHardwareCount,omitempty"`

	// ReplicationPartners - The network storage volumes configured to be replicants of a volume.
	ReplicationPartners []*SoftLayer_Network_Storage `json:"replicationPartners,omitempty"`

	// Snapshots - The snapshots associated with this SoftLayer_Network_Storage volume.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots,omitempty"`

	// AllowedReplicationHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access
	// to this storage volume's Replicant.
	AllowedReplicationHardwareCount uint64 `json:"allowedReplicationHardwareCount,omitempty"`

	// SnapshotCapacityGb - no documentation
	SnapshotCapacityGb string `json:"snapshotCapacityGb,omitempty"`

	// PropertyCount - A count of the properties used to provide additional details about a network storage
	// volume.
	PropertyCount uint64 `json:"propertyCount,omitempty"`

	// NotificationSubscribers - The subscribers that will be notified for usage amount warnings and
	// overages.
	NotificationSubscribers []*SoftLayer_Notification_User_Subscriber `json:"notificationSubscribers,omitempty"`

	// AllowedIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume.
	AllowedIpAddressCount uint64 `json:"allowedIpAddressCount,omitempty"`

	// EventCount - A count of the events which have taken place on a network storage volume.
	EventCount uint64 `json:"eventCount,omitempty"`

	// AllowableHardware - The list of Hardware that can be attached to this storage.
	AllowableHardware []*SoftLayer_Hardware `json:"allowableHardware,omitempty"`

	// AllowableVirtualGuests - The list of Virtual Guest that can be attached to this storage.
	AllowableVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowableVirtualGuests,omitempty"`

	// StorageGroups - The network storage groups this volume is attached to.
	StorageGroups []*SoftLayer_Network_Storage_Group `json:"storageGroups,omitempty"`

	// AllowedVirtualGuestCount - A count of the SoftLayer_Virtual_Guest objects which are allowed access
	// to this storage volume.
	AllowedVirtualGuestCount uint64 `json:"allowedVirtualGuestCount,omitempty"`

	// AllowedVirtualGuests - The SoftLayer_Virtual_Guest objects which are allowed access to this storage
	// volume.
	AllowedVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedVirtualGuests,omitempty"`

	// UsageNotification - no documentation
	UsageNotification *SoftLayer_Notification `json:"usageNotification,omitempty"`

	// AllowedReplicationIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which
	// are allowed access to this storage volume's Replicant.
	AllowedReplicationIpAddressCount uint64 `json:"allowedReplicationIpAddressCount,omitempty"`

	// VolumeHistory - The username and password history for a Storage service.
	VolumeHistory []*SoftLayer_Network_Storage_History `json:"volumeHistory,omitempty"`

	// OsTypeId - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type
	OsTypeId string `json:"osTypeId,omitempty"`

	// ScheduleCount - A count of the schedules which are associated with a network storage volume.
	ScheduleCount uint64 `json:"scheduleCount,omitempty"`

	// ReplicatingLunCount - A count of the iSCSI LUN volumes being replicated by this network storage
	// volume.
	ReplicatingLunCount uint64 `json:"replicatingLunCount,omitempty"`

	// OsType - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type.
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType,omitempty"`

	// AllowedReplicationVirtualGuestCount - A count of the SoftLayer_Hardware objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationVirtualGuestCount uint64 `json:"allowedReplicationVirtualGuestCount,omitempty"`
}

func (softlayer_network_storage_iscsi_equallogic_version3_replicant *SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant_Extended) String() string {
	return "SoftLayer_Network_Storage_Iscsi_EqualLogic_Version3_Replicant"
}
