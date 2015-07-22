package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Software_Component_OperatingSystem - SoftLayer_Software_Component_OperatingSystem extends
// the [[SoftLayer_Software_Component]] data type to include operating system specific properties.
type SoftLayer_Software_Component_OperatingSystem struct {

	// LicenseExpirationDate - The date in which the license for this software expires.
	LicenseExpirationDate *time.Time `json:"licenseExpirationDate"`

	// PartitionTemplateCount - A count of an operating system's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]] that can be used to
	// configure a hardware drive.
	PartitionTemplateCount uint64 `json:"partitionTemplateCount"`

	// PartitionTemplates - An operating system's associated
	// [[SoftLayer_Hardware_Component_Partition_Template|Partition Templates]] that can be used to
	// configure a hardware drive.
	PartitionTemplates []*SoftLayer_Hardware_Component_Partition_Template `json:"partitionTemplates"`

	// ReloadTransactionGroup - An operating systems associated
	// [[SoftLayer_Provisioning_Version1_Transaction_Group|Transaction Group]]. A transaction group is a
	// list of operations that will occur during the installment of an operating system.
	ReloadTransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup"`
}
