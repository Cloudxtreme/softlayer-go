package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Backup_Evault_Version6 - The
// SoftLayer_Network_Storage_Backup_Evault_Version6 contains the same properties as the
// SoftLayer_Network_Storage_Backup_Evault. Additional properties available for the EVault Storage
// type: softwareComponent, totalBytesUsed, backupJobDetails, restoreJobDetails and agentStatuses
type SoftLayer_Network_Storage_Backup_Evault_Version6 struct {

	// AgentStatusCount - A count of statuses (most of the time will be one status) for the agent tied to
	// the EVault Storage services.
	AgentStatusCount uint64 `json:"agentStatusCount"`

	// AgentStatuses - Statuses (most of the time will be one status) for the agent tied to the EVault
	// Storage services.
	AgentStatuses []*SoftLayer_Container_Network_Storage_Evault_WebCc_AgentStatus `json:"agentStatuses"`

	// BackupJobDetailCount - A count of all the of the backup jobs for the EVault Storage account.
	BackupJobDetailCount uint64 `json:"backupJobDetailCount"`

	// BackupJobDetails - All the of the backup jobs for the EVault Storage account.
	BackupJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"backupJobDetails"`

	// PluginBillingItemCount - A count of the billing items for plugins tied to the EVault Storage
	// service.
	PluginBillingItemCount uint64 `json:"pluginBillingItemCount"`

	// PluginBillingItems - The billing items for plugins tied to the EVault Storage service.
	PluginBillingItems []*SoftLayer_Billing_Item `json:"pluginBillingItems"`

	// RestoreJobDetailCount - A count of all the of the restore jobs for the EVault Storage account.
	RestoreJobDetailCount uint64 `json:"restoreJobDetailCount"`

	// RestoreJobDetails - All the of the restore jobs for the EVault Storage account.
	RestoreJobDetails []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails `json:"restoreJobDetails"`

	// SoftwareComponent - no documentation
	SoftwareComponent *SoftLayer_Software_Component `json:"softwareComponent"`

	// TaskCount - A count of retrieve the task information for the EVault Storage service.
	TaskCount uint64 `json:"taskCount"`

	// Tasks - Retrieve the task information for the EVault Storage service.
	Tasks []*SoftLayer_Container_Network_Storage_Evault_Vault_Task `json:"tasks"`
}

func (softlayer_network_storage_backup_evault_version6 *SoftLayer_Network_Storage_Backup_Evault_Version6) String() string {
	return "SoftLayer_Network_Storage_Backup_Evault_Version6"
}
