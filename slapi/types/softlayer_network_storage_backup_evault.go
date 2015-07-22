package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Storage_Backup_Evault - The SoftLayer_Network_Storage_Backup_Evault contains
// general information regarding an EVault Storage service such as account id, username, maximum
// capacity, password, Storage's product type and the server id.
type SoftLayer_Network_Storage_Backup_Evault struct {
}

// DeleteTasks - This method can be used to help maintain the storage space on a vault. When a job is
// removed from the Webcc, the task and stored usage still exists on the vault. This method can be used
// to delete the associated task and its usage. All that is required for the use of the method is to
// pass in an integer array of task(s).
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) DeleteTasks(commonOptions *slapi.CommonOptions, tasks []int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetHardwareWithEvaultFirst - Retrieve a list of hardware associated with a SoftLayer customer
// account, placing all hardware with associated EVault storage accounts at the beginning of the list.
// The return type is SoftLayer_Hardware_Server[] contains the results; the number of items returned in
// the result will be returned in the soap header (totalItems). ''getHardwareWithEvaultFirst'' is
// useful in situations where you wish to search for hardware and provide paginated output. Results are
// only returned for hardware belonging to the account of the user making the API call. This method
// drives the backup page of the SoftLayer customer portal. It serves a very specific function, but we
// have exposed it as it may prove useful for API developers too.
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) GetHardwareWithEvaultFirst(commonOptions *slapi.CommonOptions, option string, exactMatch bool, criteria string, mode string) ([]*SoftLayer_Hardware, error) {
	var returnValue []*SoftLayer_Hardware
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Storage_Backup_Evault object whose ID
// corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Storage_Backup_Evault service.
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Storage_Backup_Evault, error) {
	var returnValue *SoftLayer_Network_Storage_Backup_Evault
	return returnValue, nil
}

// GetWebCCAuthenticationDetails - <nil>
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) GetWebCCAuthenticationDetails(commonOptions *slapi.CommonOptions) (*SoftLayer_Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details, error) {
	var returnValue *SoftLayer_Container_Network_Storage_Backup_Evault_WebCc_Authentication_Details
	return returnValue, nil
}

// InitiateBareMetalRestore - Evault Bare Metal Restore is a special version of Rescue Kernel designed
// specifically for making full system restores made with Evault's BMR backup. This process works very
// similar to Rescue Kernel, except only the Evault restore program is available. The process takes
// approximately 10 minutes. Once completed you will be able to access your server to do a restore
// through VNC or your servers KVM-over-IP. IP information and credentials can be found on the hardware
// page of the customer portal. The Evault Application will be running automatically upon startup, and
// will walk you through the restore process.
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) InitiateBareMetalRestore(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// InitiateBareMetalRestoreForServer - This method operates the same as the initiateBareMetalRestore()
// method. However, using this method, the Bare Metal Restore can be initiated on any Windows server
// under the account.
func (softlayer_network_storage_backup_evault *SoftLayer_Network_Storage_Backup_Evault) InitiateBareMetalRestoreForServer(commonOptions *slapi.CommonOptions, hardwareId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
