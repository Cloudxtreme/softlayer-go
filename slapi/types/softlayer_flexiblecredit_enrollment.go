package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_FlexibleCredit_Enrollment - <nil>
type SoftLayer_FlexibleCredit_Enrollment struct {

	// CompanyTypeId - ID of the Flexible Credit Program Company classification for this enrollment
	CompanyTypeId int `json:"companyTypeId"`

	// EnrollmentDate - Date when participation in the Flexible Credit program began
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// RepresentativeEmployeeId - no documentation
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`

	// AgreementCompleteFlag - Indicates signing of Flexible Credit agreement (independent from
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// GraduationDate - no documentation
	GraduationDate *time.Time `json:"graduationDate"`

	// AffiliateId - ID of the corresponding Flexible Credit Program Affiliate
	AffiliateId int `json:"affiliateId"`

	// MonthlyCreditAmount - no documentation
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// CompanyDescription - no documentation
	CompanyDescription string `json:"companyDescription"`
}

// SoftLayer_FlexibleCredit_Enrollment_Extended is SoftLayer_FlexibleCredit_Enrollment with all maskable types.
type SoftLayer_FlexibleCredit_Enrollment_Extended struct {
	SoftLayer_FlexibleCredit_Enrollment

	// Affiliate - no documentation
	Affiliate *SoftLayer_FlexibleCredit_Affiliate `json:"affiliate"`

	// FlexibleCreditProgram - no documentation
	FlexibleCreditProgram *SoftLayer_FlexibleCredit_Program `json:"flexibleCreditProgram"`

	// IsActiveFlag - Flag indicating whether an enrollment is active (true) or inactive (false)
	IsActiveFlag bool `json:"isActiveFlag"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Representative - no documentation
	Representative *SoftLayer_User_Employee `json:"representative"`

	// CompanyType - no documentation
	CompanyType *SoftLayer_FlexibleCredit_Company_Type `json:"companyType"`
}

func (softlayer_flexiblecredit_enrollment *SoftLayer_FlexibleCredit_Enrollment) String() string {
	return "SoftLayer_FlexibleCredit_Enrollment"
}
