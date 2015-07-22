package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Dns_Domain_Registration_List - <nil>
type SoftLayer_Container_Dns_Domain_Registration_List struct {

	// DomainName - no documentation
	DomainName string `json:"domainName"`

	// EncodingType - Three-character language tag for the IDN domain that you're trying to register. This
	// is only required for IDN domains.
	EncodingType string `json:"encodingType"`

	// RegistrationPeriod - The length of the registration period in years. Valid values are 1 – 10.
	RegistrationPeriod int `json:"registrationPeriod"`
}
