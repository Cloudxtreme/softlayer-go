package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Utility_File_Entity - SoftLayer_Container_Utility_File_Entity data type models a
// single entity on a storage resource. Entities can include anything within a storage volume including
// files, folders, directories, and CloudLayer storage projects.
type SoftLayer_Container_Utility_File_Entity struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DeleteDate - The date a CloudLayer storage file entity was moved into the recycle bin. This field
	// applies to files that are pending deletion in the recycle bin.
	DeleteDate *time.Time `json:"deleteDate"`

	// Id - Unique identifier for the file. This can be either a number or guid.
	Id string `json:"id"`

	// IsShared - Whether a CloudLayer storage file entity is shared with another CloudLayer user.
	IsShared int `json:"isShared"`

	// Owner - The owner is usually the account who first upload or created the file on the resource or the
	// account who is responsible for the file at the moment.
	Owner string `json:"owner"`

	// Size - no documentation
	Size uint64 `json:"size"`

	// Content - no documentation
	Content string `json:"content"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// Type - A CloudLayer storage file entity's type. Types can include "file", "folder", "dir", and
	// "project".
	Type string `json:"type"`

	// Version - The latest revision of a file on a CloudLayer storage volume. This number increments each
	// time a new revision of the file is uploaded.
	Version int `json:"version"`

	// ContentType - no documentation
	ContentType string `json:"contentType"`
}

func (softlayer_container_utility_file_entity *SoftLayer_Container_Utility_File_Entity) String() string {
	return "SoftLayer_Container_Utility_File_Entity"
}
