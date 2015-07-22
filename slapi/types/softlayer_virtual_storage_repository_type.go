package sl

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

	// StorageRepositories - The storage repositories on a SoftLayer customer account that belong to this
	// type.
	StorageRepositories []*SoftLayer_Virtual_Storage_Repository `json:"storageRepositories"`

	// StorageRepositoryCount - A count of the storage repositories on a SoftLayer customer account that
	// belong to this type.
	StorageRepositoryCount uint64 `json:"storageRepositoryCount"`
}
