package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail - <nil>
type SoftLayer_Container_Dns_Domain_Registration_Registrant_Verification_StatusDetail struct {

	// Status - no documentation
	Status *SoftLayer_Dns_Domain_Registration_Registrant_Verification_Status `json:"status"`

	// VerificationDeadlineDate - no documentation
	VerificationDeadlineDate *time.Time `json:"verificationDeadlineDate"`
}
