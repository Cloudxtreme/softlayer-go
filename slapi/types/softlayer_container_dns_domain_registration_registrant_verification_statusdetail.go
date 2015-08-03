package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail - <nil>
type SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail struct {

	// VerificationDeadlineDate - no documentation
	VerificationDeadlineDate *time.Time `json:"verificationDeadlineDate"`

	// Status - no documentation
	Status *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status `json:"status"`
}

func (softlayer_container_dns_domain_registration_registrant_verification_statusdetail *SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail) String() string {
	return "SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail"
}
