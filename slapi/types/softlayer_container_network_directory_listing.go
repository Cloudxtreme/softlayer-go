package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_Directory_Listing - SoftLayer_Container_Network_Directory_Listing
// represents a single entry in a listing of files within a remote directory. API methods that return
// remote directory listings typically return arrays of SoftLayer_Container_Network_Directory_Listing
// objects.
type SoftLayer_Container_Network_Directory_Listing struct {

	// Name - The name of a directory or a file within a directory listing.
	Name string `json:"name,omitempty"`

	// Type - The type of file in a directory listing. If a directory listing entry is a directory itself
	// then type is set to "directory". Otherwise, type is a blank string.
	Type string `json:"type,omitempty"`

	// FileCount - If the file in a directory listing is a directory itself then fileCount is the number of
	// files within the directory.
	FileCount int `json:"fileCount,omitempty"`
}

func (softlayer_container_network_directory_listing *SoftLayer_Container_Network_Directory_Listing) String() string {
	return "SoftLayer_Container_Network_Directory_Listing"
}
