package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Storage_Evault_Vault_Task - SoftLayer's StorageLayer Evault services
// provides details regarding the the purchased vault. When a job is created using the Webcc Console,
// the job created is identified as a task on the vault. Using this service, information regarding the
// task can be retrieved.
type SoftLayer_Container_Network_Storage_Evault_Vault_Task struct {

	// Id - no documentation
	Id uint `json:"id"`

	// Name - The hostname provided at time of agent registration.
	Name string `json:"name"`

	// UsedPoolsize - no documentation
	UsedPoolsize uint64 `json:"usedPoolsize"`
}

func (softlayer_container_network_storage_evault_vault_task *SoftLayer_Container_Network_Storage_Evault_Vault_Task) String() string {
	return "SoftLayer_Container_Network_Storage_Evault_Vault_Task"
}
