package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Provisioning_Maintenance_Window - This is the datatype that needs to be
// populated and sent to SoftLayer_Provisioning_Maintenance_Window::addCustomerUpgradeWindow. This
// datatype has everything required to place an order with SoftLayer.
type SoftLayer_Container_Provisioning_Maintenance_Window struct {

	// ClassificationIds - no documentation
	ClassificationIds []*SoftLayer_Provisioning_Maintenance_Classification `json:"classificationIds"`

	// ItemCategoryIds - no documentation
	ItemCategoryIds []*SoftLayer_Product_Item_Category `json:"itemCategoryIds"`

	// MaintenanceWindowId - no documentation
	MaintenanceWindowId int `json:"maintenanceWindowId"`

	// TicketId - no documentation
	TicketId int `json:"ticketId"`

	// WindowMaintenanceDate - no documentation
	WindowMaintenanceDate *time.Time `json:"windowMaintenanceDate"`
}

func (softlayer_container_provisioning_maintenance_window *SoftLayer_Container_Provisioning_Maintenance_Window) String() string {
	return "SoftLayer_Container_Provisioning_Maintenance_Window"
}
