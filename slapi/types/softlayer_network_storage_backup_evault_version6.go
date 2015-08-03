package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Backup_Evault_Version6 - The
// SoftLayer_Network_Storage_Backup_Evault_Version6 contains the same properties as the
// SoftLayer_Network_Storage_Backup_Evault. Additional properties available for the EVault Storage
// type: softwareComponent, totalBytesUsed, backupJobDetails, restoreJobDetails and agentStatuses
type SoftLayer_Network_Storage_Backup_Evault_Version6 struct {
}

func (softlayer_network_storage_backup_evault_version6 *SoftLayer_Network_Storage_Backup_Evault_Version6) String() string {
	return "SoftLayer_Network_Storage_Backup_Evault_Version6"
}

// SoftLayer_Network_Storage_Backup_Evault_Version6_Extended is SoftLayer_Network_Storage_Backup_Evault_Version6 with all maskable types.
type SoftLayer_Network_Storage_Backup_Evault_Version6_Extended struct {
	SoftLayer_Network_Storage_Backup_Evault_Version6

	// SoftwareComponent - no documentation
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent"`

	// Tasks - Retrieve the task information for the EVault Storage service.
	Tasks []*SoftLayer_Container_Network_Storage_Evault_Vault_Task `json:"tasks"`

	// BackupJobDetailCount - A count of all the of the backup jobs for the EVault Storage account.
	BackupJobDetailCount uint64 `json:"backupJobDetailCount"`

	// PluginBillingItemCount - A count of the billing items for plugins tied to the EVault Storage
	// service.
	PluginBillingItemCount uint64 `json:"pluginBillingItemCount"`

	// AgentStatuses - Statuses (most of the time will be one status) for the agent tied to the EVault
	// Storage services.
	AgentStatuses []*SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus `json:"agentStatuses"`

	// BackupJobDetails - All the of the backup jobs for the EVault Storage account.
	BackupJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"backupJobDetails"`

	// AgentStatusCount - A count of statuses (most of the time will be one status) for the agent tied to
	// the EVault Storage services.
	AgentStatusCount uint64 `json:"agentStatusCount"`

	// RestoreJobDetailCount - A count of all the of the restore jobs for the EVault Storage account.
	RestoreJobDetailCount uint64 `json:"restoreJobDetailCount"`

	// TaskCount - A count of retrieve the task information for the EVault Storage service.
	TaskCount uint64 `json:"taskCount"`

	// PluginBillingItems - The billing items for plugins tied to the EVault Storage service.
	PluginBillingItems []*SoftLayer_Billing_Item `json:"pluginBillingItems"`

	// RestoreJobDetails - All the of the restore jobs for the EVault Storage account.
	RestoreJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"restoreJobDetails"`
}

func (softlayer_network_storage_backup_evault_version6 *SoftLayer_Network_Storage_Backup_Evault_Version6_Extended) String() string {
	return "SoftLayer_Network_Storage_Backup_Evault_Version6"
}
