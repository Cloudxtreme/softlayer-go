package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_FlexibleCredit_Enrollment - <nil>
type SoftLayer_FlexibleCredit_Enrollment struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// Affiliate - no documentation
	Affiliate *SoftLayer_FlexibleCredit_Affiliate `json:"affiliate"`

	// AffiliateId - ID of the corresponding Flexible Credit Program Affiliate
	AffiliateId int `json:"affiliateId"`

	// AgreementCompleteFlag - Indicates signing of Flexible Credit agreement (independent from
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// CompanyDescription - no documentation
	CompanyDescription string `json:"companyDescription"`

	// CompanyType - no documentation
	CompanyType *SoftLayer_FlexibleCredit_Company_Type `json:"companyType"`

	// CompanyTypeId - ID of the Flexible Credit Program Company classification for this enrollment
	CompanyTypeId int `json:"companyTypeId"`

	// EnrollmentDate - Date when participation in the Flexible Credit program began
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// FlexibleCreditProgram - no documentation
	FlexibleCreditProgram *SoftLayer_FlexibleCredit_Program `json:"flexibleCreditProgram"`

	// GraduationDate - no documentation
	GraduationDate *time.Time `json:"graduationDate"`

	// IsActiveFlag - Flag indicating whether an enrollment is active (true) or inactive (false)
	IsActiveFlag bool `json:"isActiveFlag"`

	// MonthlyCreditAmount - no documentation
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// Representative - no documentation
	Representative *SoftLayer_User_Employee `json:"representative"`

	// RepresentativeEmployeeId - no documentation
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`
}

func (softlayer_flexiblecredit_enrollment *SoftLayer_FlexibleCredit_Enrollment) String() string {
	return "SoftLayer_FlexibleCredit_Enrollment"
}
