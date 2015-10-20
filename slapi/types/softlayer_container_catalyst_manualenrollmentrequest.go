package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Catalyst_ManualEnrollmentRequest - Contains user information used to request a
// manual Catalyst enrollment.
type SoftLayer_Container_Catalyst_ManualEnrollmentRequest struct {

	// CustomerEmail - no documentation
	CustomerEmail string `json:"customerEmail,omitempty"`

	// CustomerName - no documentation
	CustomerName string `json:"customerName,omitempty"`

	// StartupName - no documentation
	StartupName string `json:"startupName,omitempty"`

	// VentureAffiliationFlag - Flag indicating whether (true) or not (false) and applicant is
	VentureAffiliationFlag bool `json:"ventureAffiliationFlag,omitempty"`

	// VentureFundName - Name of the venture capital fund, if any, applicant is affiliated with
	VentureFundName string `json:"ventureFundName,omitempty"`
}

func (softlayer_container_catalyst_manualenrollmentrequest *SoftLayer_Container_Catalyst_ManualEnrollmentRequest) String() string {
	return "SoftLayer_Container_Catalyst_ManualEnrollmentRequest"
}
