package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}

// SoftLayer_Catalyst_Enrollment_Extended is SoftLayer_Catalyst_Enrollment with all maskable types.
type SoftLayer_Catalyst_Enrollment_Extended struct {
	SoftLayer_Catalyst_Enrollment

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType"`

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment_Extended) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
