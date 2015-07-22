package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Provisioning_Maintenance_Classification_Item_Category - <nil>
type SoftLayer_Provisioning_Maintenance_Classification_Item_Category struct {

	// ItemCategoryId - <nil>
	ItemCategoryId int `json:"itemCategoryId"`

	// MaintenanceClassification - <nil>
	MaintenanceClassification *SoftLayer_Provisioning_Maintenance_Classification `json:"maintenanceClassification"`

	// MaintenanceClassificationId - <nil>
	MaintenanceClassificationId int `json:"maintenanceClassificationId"`
}

// GetObject - <nil>
func (softlayer_provisioning_maintenance_classification_item_category *SoftLayer_Provisioning_Maintenance_Classification_Item_Category) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Provisioning_Maintenance_Classification_Item_Category, error) {
	var returnValue *SoftLayer_Provisioning_Maintenance_Classification_Item_Category
	return returnValue, nil
}
