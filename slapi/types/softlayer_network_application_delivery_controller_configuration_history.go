package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Application_Delivery_Controller_Configuration_History - The
// SoftLayer_Network_Application_Delivery_Controller_Configuration_History data type models a single
// instance of a configuration history entry for an application delivery controller. The configuration
// history entries are used to support creating backups of an application delivery controller's
// configuration state in order to restore them later if needed.
type SoftLayer_Network_Application_Delivery_Controller_Configuration_History struct {

	// Controller - The application delivery controller that a configuration history record belongs to.
	Controller *SoftLayer_Network_Application_Delivery_Controller `json:"controller"`

	// CreateDate - The date a configuration history record was created.
	CreateDate *time.Time `json:"createDate"`

	// Id - An configuration history record's unique identifier
	Id int `json:"id"`

	// Notes - Editable notes used to describe a configuration history record
	Notes string `json:"notes"`
}

// DeleteObject - deleteObject permanently removes a configuration history record
func (softlayer_network_application_delivery_controller_configuration_history *SoftLayer_Network_Application_Delivery_Controller_Configuration_History) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_application_delivery_controller_configuration_history *SoftLayer_Network_Application_Delivery_Controller_Configuration_History) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Application_Delivery_Controller_Configuration_History, error) {
	var returnValue *SoftLayer_Network_Application_Delivery_Controller_Configuration_History
	return returnValue, nil
}
