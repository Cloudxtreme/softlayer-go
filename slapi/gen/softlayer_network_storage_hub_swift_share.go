package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Storage_Hub_Swift_Share - <nil>
type SoftLayer_Network_Storage_Hub_Swift_Share struct {
}

// GetContainerList - This method returns a collection of container objects.
func (softlayer_network_storage_hub_swift_share *SoftLayer_Network_Storage_Hub_Swift_Share) GetContainerList() ([]*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Folder, error) {
	var returnValue []*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_Folder
	return returnValue, nil
}

// GetFile - This method returns a file object given the file's full name.
func (softlayer_network_storage_hub_swift_share *SoftLayer_Network_Storage_Hub_Swift_Share) GetFile(fileName string, container string) (*SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File, error) {
	var returnValue *SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File
	return returnValue, nil
}

// GetFileList - This method returns a collection of the file objects within a container and the given
// path.
func (softlayer_network_storage_hub_swift_share *SoftLayer_Network_Storage_Hub_Swift_Share) GetFileList(container string, path string) ([]*SoftLayer_Container_Utility_File_Entity, error) {
	var returnValue []*SoftLayer_Container_Utility_File_Entity
	return returnValue, nil
}