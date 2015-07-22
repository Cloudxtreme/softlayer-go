package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage - The SoftLayer_Network_Storage data type contains general information
// regarding a Storage product such as account id, access username and password, the Storage product
// type, and the server the Storage service is associated with. Currently, only EVault backup storage
// has an associated server.
type SoftLayer_Network_Storage struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the SoftLayer customer account that a Storage account belongs
	// to.
	AccountId int `json:"accountId"`

	// AccountPassword - Other usernames and passwords associated with a Storage volume.
	AccountPassword *SoftLayer_Account_Password `json:"accountPassword"`

	// ActiveTransactionCount - no documentation
	ActiveTransactionCount uint64 `json:"activeTransactionCount"`

	// ActiveTransactions - no documentation
	ActiveTransactions []*SoftLayer_Provisioning_Version1_Transaction `json:"activeTransactions"`

	// AllowableHardware - The list of Hardware that can be attached to this storage.
	AllowableHardware []*SoftLayer_Hardware `json:"allowableHardware"`

	// AllowableHardwareCount - A count of the list of Hardware that can be attached to this storage.
	AllowableHardwareCount uint64 `json:"allowableHardwareCount"`

	// AllowableVirtualGuestCount - A count of the list of Virtual Guest that can be attached to this
	// storage.
	AllowableVirtualGuestCount uint64 `json:"allowableVirtualGuestCount"`

	// AllowableVirtualGuests - The list of Virtual Guest that can be attached to this storage.
	AllowableVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowableVirtualGuests"`

	// AllowedHardware - The SoftLayer_Hardware objects which are allowed access to this storage volume.
	AllowedHardware []*SoftLayer_Hardware `json:"allowedHardware"`

	// AllowedHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access to this
	// storage volume.
	AllowedHardwareCount uint64 `json:"allowedHardwareCount"`

	// AllowedIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume.
	AllowedIpAddressCount uint64 `json:"allowedIpAddressCount"`

	// AllowedIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed access to this
	// storage volume.
	AllowedIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedIpAddresses"`

	// AllowedReplicationHardware - The SoftLayer_Hardware objects which are allowed access to this storage
	// volume's Replicant.
	AllowedReplicationHardware []*SoftLayer_Hardware `json:"allowedReplicationHardware"`

	// AllowedReplicationHardwareCount - A count of the SoftLayer_Hardware objects which are allowed access
	// to this storage volume's Replicant.
	AllowedReplicationHardwareCount uint64 `json:"allowedReplicationHardwareCount"`

	// AllowedReplicationIpAddressCount - A count of the SoftLayer_Network_Subnet_IpAddress objects which
	// are allowed access to this storage volume's Replicant.
	AllowedReplicationIpAddressCount uint64 `json:"allowedReplicationIpAddressCount"`

	// AllowedReplicationIpAddresses - The SoftLayer_Network_Subnet_IpAddress objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"allowedReplicationIpAddresses"`

	// AllowedReplicationSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationSubnetCount uint64 `json:"allowedReplicationSubnetCount"`

	// AllowedReplicationSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationSubnets []*SoftLayer_Network_Subnet `json:"allowedReplicationSubnets"`

	// AllowedReplicationVirtualGuestCount - A count of the SoftLayer_Hardware objects which are allowed
	// access to this storage volume's Replicant.
	AllowedReplicationVirtualGuestCount uint64 `json:"allowedReplicationVirtualGuestCount"`

	// AllowedReplicationVirtualGuests - The SoftLayer_Hardware objects which are allowed access to this
	// storage volume's Replicant.
	AllowedReplicationVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedReplicationVirtualGuests"`

	// AllowedSubnetCount - A count of the SoftLayer_Network_Subnet objects which are allowed access to
	// this storage volume.
	AllowedSubnetCount uint64 `json:"allowedSubnetCount"`

	// AllowedSubnets - The SoftLayer_Network_Subnet objects which are allowed access to this storage
	// volume.
	AllowedSubnets []*SoftLayer_Network_Subnet `json:"allowedSubnets"`

	// AllowedVirtualGuestCount - A count of the SoftLayer_Virtual_Guest objects which are allowed access
	// to this storage volume.
	AllowedVirtualGuestCount uint64 `json:"allowedVirtualGuestCount"`

	// AllowedVirtualGuests - The SoftLayer_Virtual_Guest objects which are allowed access to this storage
	// volume.
	AllowedVirtualGuests []*SoftLayer_Virtual_Guest `json:"allowedVirtualGuests"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// BillingItemCategory - <nil>
	BillingItemCategory *SoftLayer_Product_Item_Category `json:"billingItemCategory"`

	// BytesUsed - no documentation
	BytesUsed string `json:"bytesUsed"`

	// CapacityGb - A Storage account's capacity, measured in gigabytes.
	CapacityGb int `json:"capacityGb"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// CreationScheduleId - The schedule id which was executed to create a snapshot.
	CreationScheduleId string `json:"creationScheduleId"`

	// CredentialCount - no documentation
	CredentialCount uint64 `json:"credentialCount"`

	// Credentials - <nil>
	Credentials []*SoftLayer_Network_Storage_Credential `json:"credentials"`

	// DailySchedule - The Daily Schedule which is associated with this network storage volume.
	DailySchedule *SoftLayer_Network_Storage_Schedule `json:"dailySchedule"`

	// EventCount - A count of the events which have taken place on a network storage volume.
	EventCount uint64 `json:"eventCount"`

	// Events - The events which have taken place on a network storage volume.
	Events []*SoftLayer_Network_Storage_Event `json:"events"`

	// GuestId - The unique identification number of the guest associated with a Storage volume.
	GuestId int `json:"guestId"`

	// Hardware - When applicable, the hardware associated with a Storage service.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - The server that is associated with a Storage service.
	HardwareId int `json:"hardwareId"`

	// HostId - The unique identification number of the host associated with a Storage volume.
	HostId int `json:"hostId"`

	// HourlySchedule - The Hourly Schedule which is associated with this network storage volume.
	HourlySchedule *SoftLayer_Network_Storage_Schedule `json:"hourlySchedule"`

	// Id - no documentation
	Id int `json:"id"`

	// Iops - The maximum number of IOPs selected for this volume.
	Iops string `json:"iops"`

	// IscsiLunCount - A count of relationship between a container volume and iSCSI LUNs.
	IscsiLunCount uint64 `json:"iscsiLunCount"`

	// IscsiLuns - Relationship between a container volume and iSCSI LUNs.
	IscsiLuns []*SoftLayer_Network_Storage `json:"iscsiLuns"`

	// MetricTrackingObject - A network storage volume's metric tracking object. This object records all
	// periodic polled data available to this volume.
	MetricTrackingObject *SoftLayer_Metric_Tracking_Object `json:"metricTrackingObject"`

	// MountableFlag - Whether or not a network storage volume may be mounted.
	MountableFlag string `json:"mountableFlag"`

	// NasType - A Storage account's type. Valid examples are and
	NasType string `json:"nasType"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// NotificationSubscriberCount - A count of the subscribers that will be notified for usage amount
	// warnings and overages.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount"`

	// NotificationSubscribers - The subscribers that will be notified for usage amount warnings and
	// overages.
	NotificationSubscribers []*SoftLayer_Notification_User_Subscriber `json:"notificationSubscribers"`

	// OsType - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type.
	OsType *SoftLayer_Network_Storage_Iscsi_OS_Type `json:"osType"`

	// OsTypeId - A volume's configured SoftLayer_Network_Storage_Iscsi_OS_Type
	OsTypeId string `json:"osTypeId"`

	// ParentPartnershipCount - A count of the volumes or snapshots partnered with a network storage volume
	// in a parental role.
	ParentPartnershipCount uint64 `json:"parentPartnershipCount"`

	// ParentPartnerships - The volumes or snapshots partnered with a network storage volume in a parental
	// role.
	ParentPartnerships []*SoftLayer_Network_Storage_Partnership `json:"parentPartnerships"`

	// ParentVolume - The parent volume of a volume in a complex storage relationship.
	ParentVolume *SoftLayer_Network_Storage `json:"parentVolume"`

	// PartnershipCount - A count of the volumes or snapshots partnered with a network storage volume.
	PartnershipCount uint64 `json:"partnershipCount"`

	// Partnerships - The volumes or snapshots partnered with a network storage volume.
	Partnerships []*SoftLayer_Network_Storage_Partnership `json:"partnerships"`

	// Password - The password used to access a non-EVault Storage volume. This password is used to
	// register the EVault server agent with the vault backup system.
	Password string `json:"password"`

	// PermissionsGroupCount - A count of all permissions group(s) this volume is in.
	PermissionsGroupCount uint64 `json:"permissionsGroupCount"`

	// PermissionsGroups - no documentation
	PermissionsGroups []*SoftLayer_Network_Storage_Group `json:"permissionsGroups"`

	// Properties - The properties used to provide additional details about a network storage volume.
	Properties []*SoftLayer_Network_Storage_Property `json:"properties"`

	// PropertyCount - A count of the properties used to provide additional details about a network storage
	// volume.
	PropertyCount uint64 `json:"propertyCount"`

	// ReplicatingLunCount - A count of the iSCSI LUN volumes being replicated by a volume.
	ReplicatingLunCount uint64 `json:"replicatingLunCount"`

	// ReplicatingLuns - The iSCSI LUN volumes being replicated by a volume.
	ReplicatingLuns []*SoftLayer_Network_Storage `json:"replicatingLuns"`

	// ReplicatingVolume - The network storage volume being replicated by a volume.
	ReplicatingVolume *SoftLayer_Network_Storage `json:"replicatingVolume"`

	// ReplicationEventCount - no documentation
	ReplicationEventCount uint64 `json:"replicationEventCount"`

	// ReplicationEvents - no documentation
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents"`

	// ReplicationPartnerCount - A count of the network storage volumes configured to be replicants of a
	// volume.
	ReplicationPartnerCount uint64 `json:"replicationPartnerCount"`

	// ReplicationPartners - The network storage volumes configured to be replicants of a volume.
	ReplicationPartners []*SoftLayer_Network_Storage `json:"replicationPartners"`

	// ReplicationStatus - no documentation
	ReplicationStatus string `json:"replicationStatus"`

	// ScheduleCount - A count of the schedules which are associated with a network storage volume.
	ScheduleCount uint64 `json:"scheduleCount"`

	// Schedules - The schedules which are associated with a network storage volume.
	Schedules []*SoftLayer_Network_Storage_Schedule `json:"schedules"`

	// ServiceProviderId - no documentation
	ServiceProviderId int `json:"serviceProviderId"`

	// ServiceResource - The network resource a Storage service is connected to.
	ServiceResource *SoftLayer_Network_Service_Resource `json:"serviceResource"`

	// ServiceResourceBackendIpAddress - no documentation
	ServiceResourceBackendIpAddress string `json:"serviceResourceBackendIpAddress"`

	// ServiceResourceName - no documentation
	ServiceResourceName string `json:"serviceResourceName"`

	// SnapshotCapacityGb - no documentation
	SnapshotCapacityGb string `json:"snapshotCapacityGb"`

	// SnapshotCount - A count of the snapshots associated with this SoftLayer_Network_Storage volume.
	SnapshotCount uint64 `json:"snapshotCount"`

	// SnapshotCreationTimestamp - The creation timestamp of the snapshot on the storage platform.
	SnapshotCreationTimestamp string `json:"snapshotCreationTimestamp"`

	// SnapshotDeletionThresholdPercentage - The percentage of used snapshot space after which to delete
	// automated snapshots.
	SnapshotDeletionThresholdPercentage string `json:"snapshotDeletionThresholdPercentage"`

	// SnapshotSizeBytes - no documentation
	SnapshotSizeBytes string `json:"snapshotSizeBytes"`

	// SnapshotSpaceAvailable - no documentation
	SnapshotSpaceAvailable string `json:"snapshotSpaceAvailable"`

	// Snapshots - The snapshots associated with this SoftLayer_Network_Storage volume.
	Snapshots []*SoftLayer_Network_Storage `json:"snapshots"`

	// StorageGroupCount - A count of the network storage groups this volume is attached to.
	StorageGroupCount uint64 `json:"storageGroupCount"`

	// StorageGroups - The network storage groups this volume is attached to.
	StorageGroups []*SoftLayer_Network_Storage_Group `json:"storageGroups"`

	// TotalBytesUsed - no documentation
	TotalBytesUsed string `json:"totalBytesUsed"`

	// TotalScheduleSnapshotRetentionCount - <nil>
	TotalScheduleSnapshotRetentionCount uint `json:"totalScheduleSnapshotRetentionCount"`

	// UpgradableFlag - This flag indicates whether this storage type is upgradable or not.
	UpgradableFlag bool `json:"upgradableFlag"`

	// UsageNotification - no documentation
	UsageNotification *SoftLayer_Notification `json:"usageNotification"`

	// Username - The username used to access a non-EVault Storage volume. This username is used to
	// register the EVault server agent with the vault backup system.
	Username string `json:"username"`

	// VendorName - no documentation
	VendorName string `json:"vendorName"`

	// VirtualGuest - When applicable, the virtual guest associated with a Storage service.
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VolumeHistory - The username and password history for a Storage service.
	VolumeHistory []*SoftLayer_Network_Storage_History `json:"volumeHistory"`

	// VolumeHistoryCount - A count of the username and password history for a Storage service.
	VolumeHistoryCount uint64 `json:"volumeHistoryCount"`

	// VolumeStatus - no documentation
	VolumeStatus string `json:"volumeStatus"`

	// WebccAccount - The account username and password for the EVault webCC interface.
	WebccAccount *SoftLayer_Account_Password `json:"webccAccount"`

	// WeeklySchedule - The Weekly Schedule which is associated with this network storage volume.
	WeeklySchedule *SoftLayer_Network_Storage_Schedule `json:"weeklySchedule"`
}

// AllowAccessFromHardware - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Hardware objects which have been allowed access to this storage will be listed
// in the allowedHardware property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromHardware(ctx *slapi.RequestContext, hardwareObjectTemplate SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromHardwareList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromHardwareList(ctx *slapi.RequestContext, hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromIpAddress - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromIpAddress(ctx *slapi.RequestContext, ipAddressObjectTemplate SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromIpAddressList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromIpAddressList(ctx *slapi.RequestContext, ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromSubnet - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Hardware objects which have been allowed access to this storage will be listed
// in the allowedHardware property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromSubnet(ctx *slapi.RequestContext, subnetObjectTemplate SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromSubnetList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromSubnetList(ctx *slapi.RequestContext, subnetObjectTemplates []SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromVirtualGuest - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage will be
// listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromVirtualGuest(ctx *slapi.RequestContext, virtualGuestObjectTemplate SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessFromVirtualGuestList - This method is used to modify the access control list for this
// Storage volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage
// will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessFromVirtualGuestList(ctx *slapi.RequestContext, virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromHardware - This method is used to modify the access control list for this
// Storage replicant volume. The SoftLayer_Hardware objects which have been allowed access to this
// storage will be listed in the allowedHardware property of this storage replicant volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromHardware(ctx *slapi.RequestContext, hardwareObjectTemplate SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromHardwareList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Hardware objects which have been allowed access to this
// storage volume's replica will be listed in the allowedReplicationHardware property of this storage
// volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromHardwareList(ctx *slapi.RequestContext, hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromIpAddress - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromIpAddress(ctx *slapi.RequestContext, ipAddressObjectTemplate SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromIpAddressList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Network_Subnet_IpAddress objects which have been
// allowed access to this storage volume's replica will be listed in the allowedReplicationIpAddresses
// property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromIpAddressList(ctx *slapi.RequestContext, ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromSubnet - This method is used to modify the access control list for this
// Storage replicant volume. The SoftLayer_Hardware objects which have been allowed access to this
// storage will be listed in the allowedHardware property of this storage replicant volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromSubnet(ctx *slapi.RequestContext, subnetObjectTemplate SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromSubnetList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Network_Subnet objects which have been allowed access
// to this storage volume's replica will be listed in the allowedReplicationSubnets property of this
// storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromSubnetList(ctx *slapi.RequestContext, subnetObjectTemplates []SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromVirtualGuest - This method is used to modify the access control list for
// this Storage replicant volume. The SoftLayer_Virtual_Guest objects which have been allowed access to
// this storage will be listed in the allowedVirtualGuests property of this storage replicant volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromVirtualGuest(ctx *slapi.RequestContext, virtualGuestObjectTemplate SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AllowAccessToReplicantFromVirtualGuestList - This method is used to modify the access control list
// for this Storage volume's replica. The SoftLayer_Virtual_Guest objects which have been allowed
// access to this storage volume's replica will be listed in the allowedReplicationVirtualGuests
// property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) AllowAccessToReplicantFromVirtualGuestList(ctx *slapi.RequestContext, virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AssignCredential - This method will assign an existing credential to the current volume. The
// credential must have been created using the 'addNewCredential' method. The volume type must support
// an additional credential.
func (softlayer_network_storage *SoftLayer_Network_Storage) AssignCredential(ctx *slapi.RequestContext, username string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AssignNewCredential - This method will set up a new credential for the remote storage volume. The
// storage volume must support an additional credential. Once created, the credential will be
// automatically assigned to the current volume. If there are no volumes assigned to the credential it
// will be automatically deleted.
func (softlayer_network_storage *SoftLayer_Network_Storage) AssignNewCredential(ctx *slapi.RequestContext, type_ string) (*SoftLayer_Network_Storage_Credential, error) {
	var returnValue *SoftLayer_Network_Storage_Credential
	return returnValue, nil
}

// ChangePassword - The method will change the password for the given Storage/Virtual Server Storage
// account.
func (softlayer_network_storage *SoftLayer_Network_Storage) ChangePassword(ctx *slapi.RequestContext, username string, currentPassword string, newPassword string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CollectBandwidth - {{CloudLayerOnlyMethod}} collectBandwidth() Retrieve the bandwidth usage for the
// current billing cycle.
func (softlayer_network_storage *SoftLayer_Network_Storage) CollectBandwidth(ctx *slapi.RequestContext, type_ string, startDate time.Time, endDate time.Time) (uint64, error) {
	var returnValue uint64
	return returnValue, nil
}

// CollectBytesUsed - {{CloudLayerOnlyMethod}} collectBytesUsed() retrieves the number of bytes
// capacity currently in use on a Storage account.
func (softlayer_network_storage *SoftLayer_Network_Storage) CollectBytesUsed(ctx *slapi.RequestContext) (uint64, error) {
	var returnValue uint64
	return returnValue, nil
}

// CreateFolder - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) CreateFolder(ctx *slapi.RequestContext, folder string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateSnapshot - Create a new snapshot of the data on the network storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) CreateSnapshot(ctx *slapi.RequestContext, notes string) (*SoftLayer_Network_Storage, error) {
	var returnValue *SoftLayer_Network_Storage
	return returnValue, nil
}

// DeleteAllFiles - {{CloudLayerOnlyMethod}} Delete all files within a Storage account. Depending on
// the type of Storage account, Deleting either deletes files permanently or sends files to your
// account's recycle bin. Currently, Virtual Server storage is the only type of Storage account that
// sends files to a recycle bin when deleted. When called against a Virtual Server storage account ,
// this method also determines if the files are in the account's recycle bin. If the files exist in the
// recycle bin, then they are permanently deleted. Please note, files can not be restored once they are
// permanently deleted.
func (softlayer_network_storage *SoftLayer_Network_Storage) DeleteAllFiles(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteFile - {{CloudLayerOnlyMethod}} Delete an individual file within a Storage account. Depending
// on the type of Storage account, Deleting a file either deletes the file permanently or sends the
// file to your account's recycle bin. Currently, Virtual Server storage is the only type of Storage
// account that sends files to a recycle bin when deleted. When called against a Virtual Server storage
// account , this method also determines if the file is in the account's recycle bin. If the file exist
// in the recycle bin, then it is permanently deleted. Please note, a file can not be restored once it
// is permanently deleted.
func (softlayer_network_storage *SoftLayer_Network_Storage) DeleteFile(ctx *slapi.RequestContext, fileId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteFiles - {{CloudLayerOnlyMethod}} Delete multiple files within a Storage account. Depending on
// the type of Storage account, Deleting either deletes files permanently or sends files to your
// account's recycle bin. Currently, Virtual Server storage is the only type of Storage account that
// sends files to a recycle bin when deleted. When called against a Virtual Server storage account ,
// this method also determines if the files are in the account's recycle bin. If the files exist in the
// recycle bin, then they are permanently deleted. Please note, files can not be restored once they are
// permanently deleted.
func (softlayer_network_storage *SoftLayer_Network_Storage) DeleteFiles(ctx *slapi.RequestContext, fileIds []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteFolder - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) DeleteFolder(ctx *slapi.RequestContext, folder string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObject - Delete a network storage volume. '''This cannot be undone.''' At this time only
// network storage snapshots may be deleted with this method. ''deleteObject'' returns Boolean ''true''
// on successful deletion or ''false'' if it was unable to remove a volume;
func (softlayer_network_storage *SoftLayer_Network_Storage) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DisableSnapshots - This method is not valid for Legacy iSCSI Storage Volumes. Disable scheduled
// snapshots of this storage volume. Scheduling options include and schedules.
func (softlayer_network_storage *SoftLayer_Network_Storage) DisableSnapshots(ctx *slapi.RequestContext, scheduleType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DownloadFile - {{CloudLayerOnlyMethod}} Download a file from a Storage account. This method returns
// a file's details including the file's raw content.
func (softlayer_network_storage *SoftLayer_Network_Storage) DownloadFile(ctx *slapi.RequestContext, fileId string) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// EditCredential - This method will change the password of a credential created using the
// 'addNewCredential' method. If the credential exists on multiple storage volumes it will change for
// those volumes as well.
func (softlayer_network_storage *SoftLayer_Network_Storage) EditCredential(ctx *slapi.RequestContext, username string, newPassword string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - The password and/or notes may be modified for the Storage service except evault
// passwords and notes.
func (softlayer_network_storage *SoftLayer_Network_Storage) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Network_Storage) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EnableSnapshots - This method is not valid for Legacy iSCSI Storage Volumes. Enable scheduled
// snapshots of this storage volume. Scheduling options include and schedules. For schedules, provide
// relevant data for $scheduleType, $retentionCount and $minute. For schedules, provide relevant data
// for $scheduleType, $retentionCount, $minute, and $hour. For schedules, provide relevant data for all
// parameters of this method.
func (softlayer_network_storage *SoftLayer_Network_Storage) EnableSnapshots(ctx *slapi.RequestContext, scheduleType string, retentionCount int, minute int, hour int, dayOfWeek string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FailbackFromReplicant - Failback from a volume replicant. In order to failback the volume must have
// already been failed over to a replicant.
func (softlayer_network_storage *SoftLayer_Network_Storage) FailbackFromReplicant(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FailoverToReplicant - Failover to a volume replicant. During the time which the replicant is in use
// the local nas volume will not be available.
func (softlayer_network_storage *SoftLayer_Network_Storage) FailoverToReplicant(ctx *slapi.RequestContext, replicantId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllFiles - {{CloudLayerOnlyMethod}} Retrieve details such as id, name, size, create date for all
// files in a Storage account's root directory. This does not download file content.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetAllFiles(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue []*SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetAllFilesByFilter - {{CloudLayerOnlyMethod}} Retrieve details such as id, name, size, create date
// for all files matching the filter's criteria in a Storage account's root directory. This does not
// download file content.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetAllFilesByFilter(ctx *slapi.RequestContext, filter SoftLayer_Container_Utility_File_Entity) ([]*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue []*SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetAllowedHostsLimit - Retrieves the total number of allowed hosts limit per volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetAllowedHostsLimit(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetByUsername - Retrieve network storage accounts by username and storage account type. Use this
// method if you wish to retrieve a storage record by username rather than by id. The ''type''
// parameter must correspond to one of the available ''nasType'' values in the
// SoftLayer_Network_Storage data type.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetByUsername(ctx *slapi.RequestContext, username string, type_ string) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetCdnUrls - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetCdnUrls(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl, error) {
	var returnValue []*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_ContentDeliveryUrl
	return returnValue, nil
}

// GetClusterResource - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetClusterResource(ctx *slapi.RequestContext) (*SoftLayer_Network_Service_Resource, error) {
	var returnValue *SoftLayer_Network_Service_Resource
	return returnValue, nil
}

// GetFileByIdentifier - {{CloudLayerOnlyMethod}} Retrieve details such as id, name, size, create date
// of a file within a Storage account. This does not download file content.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFileByIdentifier(ctx *slapi.RequestContext, identifier string) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetFileCount - {{CloudLayerOnlyMethod}} Retrieve the file number of files in a Virtual Server
// Storage account's root directory. This does not include the files stored in the recycle bin.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFileCount(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetFileList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFileList(ctx *slapi.RequestContext, folder string, path string) ([]*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue []*SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetFilePendingDeleteCount - {{CloudLayerOnlyMethod}} Retrieve the number of files pending deletion
// in a Storage account's recycle bin. Files in an account's recycle bin may either be restored to the
// account's root directory or permanently deleted.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFilePendingDeleteCount(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetFilesPendingDelete - {{CloudLayerOnlyMethod}} Retrieve a list of files that are pending deletion
// in a Storage account's recycle bin. Files in an account's recycle bin may either be restored to the
// account's root directory or permanently deleted. This method does not download file content.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFilesPendingDelete(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue []*SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetFolderList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetFolderList(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Folder, error) {
	var returnValue []*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Folder
	return returnValue, nil
}

// GetGraph - {{CloudLayerOnlyMethod}} getGraph() retrieves a Storage account's usage and returns a PNG
// graph image, title, and the minimum and maximum dates included in the graphed date range. Virtual
// Server storage accounts can also graph upload and download bandwidth usage.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time, type_ string) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetNetworkConnectionDetails - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetNetworkConnectionDetails(ctx *slapi.RequestContext) (*SoftLayer_Container_Network_Storage_NetworkConnectionInformation, error) {
	var returnValue *SoftLayer_Container_Network_Storage_NetworkConnectionInformation
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Storage object whose ID corresponds to the ID
// number of the init parameter passed to the SoftLayer_Network_Storage service. Please use the
// associated methods in the [[SoftLayer_Network_Storage]] service to retrieve a Storage account's id.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Network_Storage, error) {
	var returnValue *SoftLayer_Network_Storage
	return returnValue, nil
}

// GetObjectsByCredential - Retrieve network storage accounts by SoftLayer_Network_Storage_Credential
// object. Use this method if you wish to retrieve a storage record by a credential rather than by id.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetObjectsByCredential(ctx *slapi.RequestContext, credentialObject SoftLayer_Network_Storage_Credential) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetRecycleBinFileByIdentifier - {{CloudLayerOnlyMethod}} Retrieve the details of a file that is
// pending deletion in a Storage account's a recycle bin.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetRecycleBinFileByIdentifier(ctx *slapi.RequestContext, fileId string) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// GetRemainingAllowedHosts - Retrieves the remaining number of allowed hosts per volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetRemainingAllowedHosts(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetSnapshotsForVolume - Retrieves a list of snapshots for this SoftLayer_Network_Storage volume.
// This method works with the result limits and offset to support pagination.
func (softlayer_network_storage *SoftLayer_Network_Storage) GetSnapshotsForVolume(ctx *slapi.RequestContext) ([]*SoftLayer_Network_Storage, error) {
	var returnValue []*SoftLayer_Network_Storage
	return returnValue, nil
}

// GetStorageGroupsNetworkConnectionDetails - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetStorageGroupsNetworkConnectionDetails(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Network_Storage_NetworkConnectionInformation, error) {
	var returnValue []*SoftLayer_Container_Network_Storage_NetworkConnectionInformation
	return returnValue, nil
}

// GetValidReplicationTargetDatacenterLocations - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) GetValidReplicationTargetDatacenterLocations(ctx *slapi.RequestContext) ([]*SoftLayer_Location, error) {
	var returnValue []*SoftLayer_Location
	return returnValue, nil
}

// ImmediateFailoverToReplicant - Immediate Failover to a volume replicant. During the time which the
// replicant is in use the local nas volume will not be available.
func (softlayer_network_storage *SoftLayer_Network_Storage) ImmediateFailoverToReplicant(ctx *slapi.RequestContext, replicantId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsBlockingOperationInProgress - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) IsBlockingOperationInProgress(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromHardware - This method is used to modify the access control list for this Storage
// volume. The SoftLayer_Hardware objects which have been allowed access to this storage will be listed
// in the allowedHardware property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromHardware(ctx *slapi.RequestContext, hardwareObjectTemplate SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromHardwareList - This method is used to modify the access control list for this
// Storage volume. The SoftLayer_Hardware objects which have been allowed access to this storage will
// be listed in the allowedHardware property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromHardwareList(ctx *slapi.RequestContext, hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromIpAddress - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromIpAddress(ctx *slapi.RequestContext, ipAddressObjectTemplate SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromIpAddressList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromIpAddressList(ctx *slapi.RequestContext, ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromSubnet - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromSubnet(ctx *slapi.RequestContext, subnetObjectTemplate SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromSubnetList - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromSubnetList(ctx *slapi.RequestContext, subnetObjectTemplates []SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromVirtualGuest - This method is used to modify the access control list for this
// Storage volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage
// will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromVirtualGuest(ctx *slapi.RequestContext, virtualGuestObjectTemplate SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessFromVirtualGuestList - This method is used to modify the access control list for this
// Storage volume. The SoftLayer_Virtual_Guest objects which have been allowed access to this storage
// will be listed in the allowedVirtualGuests property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessFromVirtualGuestList(ctx *slapi.RequestContext, virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromHardwareList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Hardware objects which have been allowed access to this
// storage volume's replica will be listed in the allowedReplicationHardware property of this storage
// volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessToReplicantFromHardwareList(ctx *slapi.RequestContext, hardwareObjectTemplates []SoftLayer_Hardware) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromIpAddressList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Network_Subnet_IpAddress objects which have been
// allowed access to this storage volume's replica will be listed in the allowedReplicationIpAddresses
// property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessToReplicantFromIpAddressList(ctx *slapi.RequestContext, ipAddressObjectTemplates []SoftLayer_Network_Subnet_IpAddress) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromSubnet - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessToReplicantFromSubnet(ctx *slapi.RequestContext, subnetObjectTemplate SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromSubnetList - This method is used to modify the access control list for
// this Storage volume's replica. The SoftLayer_Network_Subnet objects which have been allowed access
// to this storage volume's replica will be listed in the allowedReplicationSubnets property of this
// storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessToReplicantFromSubnetList(ctx *slapi.RequestContext, subnetObjectTemplates []SoftLayer_Network_Subnet) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAccessToReplicantFromVirtualGuestList - This method is used to modify the access control list
// for this Storage volume's replica. The SoftLayer_Virtual_Guest objects which have been allowed
// access to this storage volume's replica will be listed in the allowedReplicationVirtualGuests
// property of this storage volume.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveAccessToReplicantFromVirtualGuestList(ctx *slapi.RequestContext, virtualGuestObjectTemplates []SoftLayer_Virtual_Guest) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveCredential - This method will remove a credential from the current volume. The credential must
// have been created using the 'addNewCredential' method.
func (softlayer_network_storage *SoftLayer_Network_Storage) RemoveCredential(ctx *slapi.RequestContext, username string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RestoreFile - {{CloudLayerOnlyMethod}} Restore an individual file so that it may be used as it was
// before it was deleted. If a file is deleted from a Virtual Server Storage account, the file is
// placed into the account's recycle bin and not permanently deleted. Therefore, restoreFile can be
// used to place the file back into your Virtual Server account's root directory.
func (softlayer_network_storage *SoftLayer_Network_Storage) RestoreFile(ctx *slapi.RequestContext, fileId string) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}

// RestoreFromSnapshot - Restore the volume from a snapshot that was previously taken.
func (softlayer_network_storage *SoftLayer_Network_Storage) RestoreFromSnapshot(ctx *slapi.RequestContext, snapshotId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SendPasswordReminderEmail - The method will retrieve the password for the StorageLayer or Virtual
// Server Storage Account and email the password. The Storage Account passwords will be emailed to the
// master user. For Virtual Server Storage, the password will be sent to the email address used as the
// username.
func (softlayer_network_storage *SoftLayer_Network_Storage) SendPasswordReminderEmail(ctx *slapi.RequestContext, username string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetMountable - Enable or disable the mounting of a Storage volume. When mounting is enabled the
// Storage volume will be mountable or available for use. For Virtual Server volumes, disabling
// mounting will deny access to the Virtual Server Account, remove published material and deny all file
// interaction including uploads and downloads. Enabling or disabling mounting for Storage volumes is
// not possible if mounting has been disabled by SoftLayer or a parent account.
func (softlayer_network_storage *SoftLayer_Network_Storage) SetMountable(ctx *slapi.RequestContext, mountable bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetSnapshotAllocation - <nil>
func (softlayer_network_storage *SoftLayer_Network_Storage) SetSnapshotAllocation(ctx *slapi.RequestContext, capacityGb int) error {
	return nil
}

// UpgradeVolumeCapacity - Upgrade the Storage volume to one of the upgradable packages (for example
// from 10 Gigs of EVault storage to 100 Gigs of EVault storage).
func (softlayer_network_storage *SoftLayer_Network_Storage) UpgradeVolumeCapacity(ctx *slapi.RequestContext, itemId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UploadFile - {{CloudLayerOnlyMethod}} Upload a file to a Storage account's root directory. Once
// uploaded, this method returns new file entity identifier for the upload file. The following
// properties are required in the ''file'' parameter. *'''name''': The name of the file you wish to
// upload *'''content''': The raw contents of the file you wish to upload. *'''contentType''': The
// MIME-type of content that you wish to upload.
func (softlayer_network_storage *SoftLayer_Network_Storage) UploadFile(ctx *slapi.RequestContext, file SoftLayer_Container_Utility_File_Entity) (*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue *SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}
