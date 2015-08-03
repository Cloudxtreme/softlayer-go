package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId"`

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate"`

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative"`

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
