package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Provisioning_Maintenance_Window - The SoftLayer_Provisioning_Maintenance_Window represent
// a time window that SoftLayer performs a hardware or software maintenance and upgrades.
type SoftLayer_Provisioning_Maintenance_Window struct {

	// BeginDate - The date and time a maintenance window is scheduled to begin.
	BeginDate *time.Time `json:"beginDate"`

	// DayOfWeek - An ISO-8601 numeric representation of the day of the week that a maintenance window is
	// performed. 1: Monday, 7: Sunday
	DayOfWeek int `json:"dayOfWeek"`

	// EndDate - The date and time a maintenance window is scheduled to end.
	EndDate *time.Time `json:"endDate"`

	// Id - no documentation
	Id int `json:"id"`

	// LocationId - An internal identifier of the location (data center) record that a maintenance window
	// takes place in.
	LocationId int `json:"locationId"`

	// PortalTzId - no documentation
	PortalTzId int `json:"portalTzId"`
}

// AddCustomerUpgradeWindow - no documentation
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) AddCustomerUpgradeWindow(commonOptions *slapi.CommonOptions, customerUpgradeWindow SoftLayer_Container_Provisioning_Maintenance_Window) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetMaintenanceClassifications - getMaintenanceClassifications() returns an object of maintenance
// classifications
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenanceClassifications(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Provisioning_Maintenance_Classification, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Classification
	return returnValue, nil
}

// GetMaintenanceStartEndTime - getMaintenanceStartEndTime() returns a specific maintenance window
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenanceStartEndTime(commonOptions *slapi.CommonOptions, ticketId int) (*SoftLayer_Provisioning_Maintenance_Window, error) {
	var returnValue *SoftLayer_Provisioning_Maintenance_Window
	return returnValue, nil
}

// GetMaintenanceWindowForTicket - getMaintenceWindowForTicket() returns a specific maintenance window
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenanceWindowForTicket(commonOptions *slapi.CommonOptions, maintenanceWindowId int) ([]*SoftLayer_Provisioning_Maintenance_Window, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Window
	return returnValue, nil
}

// GetMaintenanceWindowTicketsByTicketId - getMaintenanceWindowTicketsByTicketId() returns a list
// maintenance window ticket records by ticket id
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenanceWindowTicketsByTicketId(commonOptions *slapi.CommonOptions, ticketId int) ([]*SoftLayer_Provisioning_Maintenance_Ticket, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Ticket
	return returnValue, nil
}

// GetMaintenanceWindows - This method returns a list of available maintenance windows
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenanceWindows(commonOptions *slapi.CommonOptions, beginDate time.Time, endDate time.Time, locationId int, slotsNeeded int) ([]*SoftLayer_Provisioning_Maintenance_Window, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Window
	return returnValue, nil
}

// GetMaintenceWindows - Use
// [[SoftLayer_Provisioning_Maintenance_Window::getMaintenanceWindows|getMaintenanceWindows]] method.
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) GetMaintenceWindows(commonOptions *slapi.CommonOptions, beginDate time.Time, endDate time.Time, locationId int, slotsNeeded int) ([]*SoftLayer_Provisioning_Maintenance_Window, error) {
	var returnValue []*SoftLayer_Provisioning_Maintenance_Window
	return returnValue, nil
}

// UpdateCustomerUpgradeWindow - no documentation
func (softlayer_provisioning_maintenance_window *SoftLayer_Provisioning_Maintenance_Window) UpdateCustomerUpgradeWindow(commonOptions *slapi.CommonOptions, maintenanceStartTime time.Time, newMaintenanceWindowId int, ticketId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
