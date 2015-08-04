package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Utility_File_Attachment - At times,such as when attaching files to tickets, it
// is necessary to send files to SoftLayer API methods. The SoftLayer_Container_Utility_File_Attachment
// data type models a single file to upload to the
type SoftLayer_Container_Utility_File_Attachment struct {

	// Data - The contents of a file that is uploaded to the SoftLayer
	Data string `json:"data,omitempty"`

	// Filename - The name of a file that is uploaded to the SoftLayer
	Filename string `json:"filename,omitempty"`
}

func (softlayer_container_utility_file_attachment *SoftLayer_Container_Utility_File_Attachment) String() string {
	return "SoftLayer_Container_Utility_File_Attachment"
}
