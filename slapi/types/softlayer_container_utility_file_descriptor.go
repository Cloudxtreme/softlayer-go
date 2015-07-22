package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_File_Descriptor - Used to describe a document in the file system on the
// file server
type SoftLayer_Container_Utility_File_Descriptor struct {

	// FileName - The name of a file as it exists on the file server.
	FileName string `json:"fileName"`

	// FriendlyName - The friendly name of a file as it exists on the file server.
	FriendlyName string `json:"friendlyName"`

	// ModifyDate - The date the file was last modified on the file server.
	ModifyDate *time.Time `json:"modifyDate"`
}
