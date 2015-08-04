package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Provisioning_Maintenance_Classification - The
// SoftLayer_Provisioning_Maintenance_Classification represent a maintenance type for the specific
// hardware maintenance desired.
type SoftLayer_Provisioning_Maintenance_Classification struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Slots - The number of slots required for the maintenance classification.
	Slots int `json:"slots,omitempty"`

	// Type - The type or name of the maintenance classification.
	Type string `json:"type,omitempty"`

	// ItemCategories - <nil>
	ItemCategories []*SoftLayer_Provisioning_Maintenance_Classification_Item_Category `json:"itemCategories,omitempty"`

	// ItemCategoryCount - no documentation
	ItemCategoryCount uint64 `json:"itemCategoryCount,omitempty"`
}

func (softlayer_provisioning_maintenance_classification *SoftLayer_Provisioning_Maintenance_Classification) String() string {
	return "SoftLayer_Provisioning_Maintenance_Classification"
}
