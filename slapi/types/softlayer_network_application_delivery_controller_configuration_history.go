package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Application_Delivery_Controller_Configuration_History - The
// SoftLayer_Network_Application_Delivery_Controller_Configuration_History data type models a single
// instance of a configuration history entry for an application delivery controller. The configuration
// history entries are used to support creating backups of an application delivery controller's
// configuration state in order to restore them later if needed.
type SoftLayer_Network_Application_Delivery_Controller_Configuration_History struct {

	// Notes - Editable notes used to describe a configuration history record
	Notes string `json:"notes,omitempty"`

	// CreateDate - The date a configuration history record was created.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - An configuration history record's unique identifier
	Id int `json:"id,omitempty"`
}

func (softlayer_network_application_delivery_controller_configuration_history *SoftLayer_Network_Application_Delivery_Controller_Configuration_History) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_Configuration_History"
}

// SoftLayer_Network_Application_Delivery_Controller_Configuration_History_Extended is SoftLayer_Network_Application_Delivery_Controller_Configuration_History with all maskable types.
type SoftLayer_Network_Application_Delivery_Controller_Configuration_History_Extended struct {
	SoftLayer_Network_Application_Delivery_Controller_Configuration_History

	// Controller - The application delivery controller that a configuration history record belongs to.
	Controller *SoftLayer_Network_Application_Delivery_Controller `json:"controller,omitempty"`
}

func (softlayer_network_application_delivery_controller_configuration_history *SoftLayer_Network_Application_Delivery_Controller_Configuration_History_Extended) String() string {
	return "SoftLayer_Network_Application_Delivery_Controller_Configuration_History"
}
