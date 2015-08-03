package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId"`
}

// SoftLayer_Catalyst_Enrollment_Extended is SoftLayer_Catalyst_Enrollment with all maskable types.
type SoftLayer_Catalyst_Enrollment_Extended struct {
	SoftLayer_Catalyst_Enrollment

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate"`

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
