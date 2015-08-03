package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Storage_Repository_Type - SoftLayer employs many different types of repositories
// that computing instances use as their storage volume. SoftLayer_Virtual_Storage_Repository_Type
// models a single storage type. Common types of storage repositories include networked file systems,
// logical volume management, and local disk volumes for swap and page file management.
type SoftLayer_Virtual_Storage_Repository_Type struct {

	// Description - no documentation
	Description string `json:"description"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_virtual_storage_repository_type *SoftLayer_Virtual_Storage_Repository_Type) String() string {
	return "SoftLayer_Virtual_Storage_Repository_Type"
}

// SoftLayer_Virtual_Storage_Repository_Type_Extended is SoftLayer_Virtual_Storage_Repository_Type with all maskable types.
type SoftLayer_Virtual_Storage_Repository_Type_Extended struct {
	SoftLayer_Virtual_Storage_Repository_Type

	// StorageRepositoryCount - A count of the storage repositories on a SoftLayer customer account that
	// belong to this type.
	StorageRepositoryCount uint64 `json:"storageRepositoryCount"`

	// StorageRepositories - The storage repositories on a SoftLayer customer account that belong to this
	// type.
	StorageRepositories []*SoftLayer_Virtual_Storage_Repository `json:"storageRepositories"`
}

func (softlayer_virtual_storage_repository_type *SoftLayer_Virtual_Storage_Repository_Type_Extended) String() string {
	return "SoftLayer_Virtual_Storage_Repository_Type"
}
