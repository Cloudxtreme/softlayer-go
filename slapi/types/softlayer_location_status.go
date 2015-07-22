package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Location_Status - SoftLayer_Location_Status models the state of any location. SoftLayer
// uses the following status codes: The location is currently active and available for public usage.
// Used when a location is planned but not yet active. Used when a location has been retired and no
// longer active. Locations in use should stay in the state. If a locations status ever reads anything
// else and contains active hardware then please contact SoftLayer support.
type SoftLayer_Location_Status struct {

	// Id - no documentation
	Id int `json:"id"`

	// Status - A Location's status code. See the SoftLayer_Locaiton_Status Overview for ''status'''
	// possible values.
	Status string `json:"status"`
}
