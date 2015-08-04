package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Status - SoftLayer_Hardware_Status models the inventory state of any piece of
// hardware in SoftLayer's inventory. Most of these statuses are used by SoftLayer while a server is
// not provisioned or undergoing provisioning. SoftLayer uses the following status codes: This server
// is active and in use. Used during server provisioning. *'''DEPLOY2''': Used during server
// provisioning. Used during server provisioning. This server has been reclaimed by SoftLayer and is
// awaiting de-provisioning. Servers in production and in use should stay in the state. If a server's
// status ever reads anything else then please contact SoftLayer support.
type SoftLayer_Hardware_Status struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Status - A hardware's status code. See the SoftLayer_Hardware_Status Overview for ''status'''
	// possible values.
	Status string `json:"status,omitempty"`
}

func (softlayer_hardware_status *SoftLayer_Hardware_Status) String() string {
	return "SoftLayer_Hardware_Status"
}
