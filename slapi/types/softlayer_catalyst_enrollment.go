package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag,omitempty"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount,omitempty"`

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId,omitempty"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate,omitempty"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId,omitempty"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate,omitempty"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}

// SoftLayer_Catalyst_Enrollment_Extended is SoftLayer_Catalyst_Enrollment with all maskable types.
type SoftLayer_Catalyst_Enrollment_Extended struct {
	SoftLayer_Catalyst_Enrollment

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag,omitempty"`

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate,omitempty"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType,omitempty"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative,omitempty"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment_Extended) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
