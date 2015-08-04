package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_File_Entity - SoftLayer_Container_Utility_File_Entity data type models a
// single entity on a storage resource. Entities can include anything within a storage volume including
// files, folders, directories, and CloudLayer storage projects.
type SoftLayer_Container_Utility_File_Entity struct {

	// ContentType - no documentation
	ContentType string `json:"contentType,omitempty"`

	// Id - Unique identifier for the file. This can be either a number or guid.
	Id string `json:"id,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Version - The latest revision of a file on a CloudLayer storage volume. This number increments each
	// time a new revision of the file is uploaded.
	Version int `json:"version,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Owner - The owner is usually the account who first upload or created the file on the resource or the
	// account who is responsible for the file at the moment.
	Owner string `json:"owner,omitempty"`

	// Size - no documentation
	Size uint64 `json:"size,omitempty"`

	// Type - A CloudLayer storage file entity's type. Types can include "file", "folder", "dir", and
	// "project".
	Type string `json:"type,omitempty"`

	// Content - no documentation
	Content string `json:"content,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// DeleteDate - The date a CloudLayer storage file entity was moved into the recycle bin. This field
	// applies to files that are pending deletion in the recycle bin.
	DeleteDate *time.Time `json:"deleteDate,omitempty"`

	// IsShared - Whether a CloudLayer storage file entity is shared with another CloudLayer user.
	IsShared int `json:"isShared,omitempty"`
}

func (softlayer_container_utility_file_entity *SoftLayer_Container_Utility_File_Entity) String() string {
	return "SoftLayer_Container_Utility_File_Entity"
}
