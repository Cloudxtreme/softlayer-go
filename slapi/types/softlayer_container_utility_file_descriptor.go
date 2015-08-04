package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_File_Descriptor - Used to describe a document in the file system on the
// file server
type SoftLayer_Container_Utility_File_Descriptor struct {

	// FileName - The name of a file as it exists on the file server.
	FileName string `json:"fileName,omitempty"`

	// FriendlyName - The friendly name of a file as it exists on the file server.
	FriendlyName string `json:"friendlyName,omitempty"`

	// ModifyDate - The date the file was last modified on the file server.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`
}

func (softlayer_container_utility_file_descriptor *SoftLayer_Container_Utility_File_Descriptor) String() string {
	return "SoftLayer_Container_Utility_File_Descriptor"
}
