package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Provisioning_Maintenance_Classification - The
// SoftLayer_Provisioning_Maintenance_Classification represent a maintenance type for the specific
// hardware maintenance desired.
type SoftLayer_Provisioning_Maintenance_Classification struct {

	// Id - no documentation
	Id int `json:"id"`

	// ItemCategories - <nil>
	ItemCategories []*SoftLayer_Provisioning_Maintenance_Classification_Item_Category `json:"itemCategories"`

	// ItemCategoryCount - no documentation
	ItemCategoryCount uint64 `json:"itemCategoryCount"`

	// Slots - The number of slots required for the maintenance classification.
	Slots int `json:"slots"`

	// Type - The type or name of the maintenance classification.
	Type string `json:"type"`
}

func (softlayer_provisioning_maintenance_classification *SoftLayer_Provisioning_Maintenance_Classification) String() string {
	return "SoftLayer_Provisioning_Maintenance_Classification"
}

// GetMaintenanceClassification - Retrieve an array of
// SoftLayer_Provisioning_Maintenance_Classification data types, which contain all maintenance
// classifications.
func (softlayer_provisioning_maintenance_classification *SoftLayer_Provisioning_Maintenance_Classification) GetMaintenanceClassification(ctx *slapi.RequestContext, maintenanceClassificationId int) ([]*SoftLayer_Provisioning_Maintenance_Classification, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Classification
	return returnValue, nil
}

// GetMaintenanceClassificationsByItemCategory - Retrieve an array of
// SoftLayer_Provisioning_Maintenance_Classification data types, which contain all maintenance
// classifications.
func (softlayer_provisioning_maintenance_classification *SoftLayer_Provisioning_Maintenance_Classification) GetMaintenanceClassificationsByItemCategory(ctx *slapi.RequestContext) ([]*SoftLayer_Provisioning_Maintenance_Classification_Item_Category, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Classification_Item_Category
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_provisioning_maintenance_classification *SoftLayer_Provisioning_Maintenance_Classification) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Provisioning_Maintenance_Classification, error) {
	var returnValue *SoftLayer_Provisioning_Maintenance_Classification
	return returnValue, nil
}
