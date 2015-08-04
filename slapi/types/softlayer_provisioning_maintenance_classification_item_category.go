package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Provisioning_Maintenance_Classification_Item_Category - <nil>
type SoftLayer_Provisioning_Maintenance_Classification_Item_Category struct {

	// ItemCategoryId - <nil>
	ItemCategoryId int `json:"itemCategoryId,omitempty"`

	// MaintenanceClassificationId - <nil>
	MaintenanceClassificationId int `json:"maintenanceClassificationId,omitempty"`
}

func (softlayer_provisioning_maintenance_classification_item_category *SoftLayer_Provisioning_Maintenance_Classification_Item_Category) String() string {
	return "SoftLayer_Provisioning_Maintenance_Classification_Item_Category"
}

// SoftLayer_Provisioning_Maintenance_Classification_Item_Category_Extended is SoftLayer_Provisioning_Maintenance_Classification_Item_Category with all maskable types.
type SoftLayer_Provisioning_Maintenance_Classification_Item_Category_Extended struct {
	SoftLayer_Provisioning_Maintenance_Classification_Item_Category

	// MaintenanceClassification - <nil>
	MaintenanceClassification *SoftLayer_Provisioning_Maintenance_Classification `json:"maintenanceClassification,omitempty"`
}

func (softlayer_provisioning_maintenance_classification_item_category *SoftLayer_Provisioning_Maintenance_Classification_Item_Category_Extended) String() string {
	return "SoftLayer_Provisioning_Maintenance_Classification_Item_Category"
}
