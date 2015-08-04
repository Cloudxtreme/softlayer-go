package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File -
// SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File provides specific details that only apply
// to files that are sent or received from CloudLayer storage resources.
type SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File struct {

	// Folder - <nil>
	Folder string `json:"folder,omitempty"`

	// Hash - <nil>
	Hash string `json:"hash,omitempty"`

	// IsShared - Whether a CloudLayer storage file entity is shared with another CloudLayer user.
	IsShared int `json:"isShared,omitempty"`

	// DeleteDate - The date a CloudLayer storage file entity was moved into the recycle bin. This field
	// applies to files that are pending deletion in the recycle bin.
	DeleteDate *time.Time `json:"deleteDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Size - no documentation
	Size uint64 `json:"size,omitempty"`

	// Owner - The owner is usually the account who first upload or created the file on the resource or the
	// account who is responsible for the file at the moment.
	Owner string `json:"owner,omitempty"`

	// Type - A CloudLayer storage file entity's type. Types can include "file", "folder", "dir", and
	// "project".
	Type string `json:"type,omitempty"`

	// Version - The latest revision of a file on a CloudLayer storage volume. This number increments each
	// time a new revision of the file is uploaded.
	Version int `json:"version,omitempty"`

	// Content - no documentation
	Content string `json:"content,omitempty"`

	// ContentType - no documentation
	ContentType string `json:"contentType,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - Unique identifier for the file. This can be either a number or guid.
	Id string `json:"id,omitempty"`
}

func (softlayer_container_network_storage_hub_objectstorage_file *SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File) String() string {
	return "SoftLayer_Container_Network_Storage_Hub_ObjectStorage_File"
}
