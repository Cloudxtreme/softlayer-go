package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Security_Scanner_Request_Status - The
// SoftLayer_Network_Security_Scanner_Request_Status data type represents the current status of a
// vulnerability scan. The status messages are as follows: *Scan Pending *Scan Processing *Scan
// Complete *Scan Cancelled *Generating Report. The status of a vulnerability scan will change over the
// course of a scan's execution.
type SoftLayer_Network_Security_Scanner_Request_Status struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}