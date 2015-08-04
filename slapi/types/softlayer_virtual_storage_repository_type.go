package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Storage_Repository_Type - SoftLayer employs many different types of repositories
// that computing instances use as their storage volume. SoftLayer_Virtual_Storage_Repository_Type
// models a single storage type. Common types of storage repositories include networked file systems,
// logical volume management, and local disk volumes for swap and page file management.
type SoftLayer_Virtual_Storage_Repository_Type struct {

	// Description - no documentation
	Description string `json:"description,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// StorageRepositories - The storage repositories on a SoftLayer customer account that belong to this
	// type.
	StorageRepositories []*SoftLayer_Virtual_Storage_Repository `json:"storageRepositories,omitempty"`

	// StorageRepositoryCount - A count of the storage repositories on a SoftLayer customer account that
	// belong to this type.
	StorageRepositoryCount uint64 `json:"storageRepositoryCount,omitempty"`
}

func (softlayer_virtual_storage_repository_type *SoftLayer_Virtual_Storage_Repository_Type) String() string {
	return "SoftLayer_Virtual_Storage_Repository_Type"
}
