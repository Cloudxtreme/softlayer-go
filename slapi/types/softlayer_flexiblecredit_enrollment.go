package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_FlexibleCredit_Enrollment - <nil>
type SoftLayer_FlexibleCredit_Enrollment struct {

	// CompanyDescription - no documentation
	CompanyDescription string `json:"companyDescription"`

	// MonthlyCreditAmount - no documentation
	MonthlyCreditAmount float64 `json:"monthlyCreditAmount"`

	// AffiliateId - ID of the corresponding Flexible Credit Program Affiliate
	AffiliateId int `json:"affiliateId"`

	// AgreementCompleteFlag - Indicates signing of Flexible Credit agreement (independent from
	AgreementCompleteFlag int `json:"agreementCompleteFlag"`

	// RepresentativeEmployeeId - no documentation
	RepresentativeEmployeeId int `json:"representativeEmployeeId"`

	// CompanyTypeId - ID of the Flexible Credit Program Company classification for this enrollment
	CompanyTypeId int `json:"companyTypeId"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// EnrollmentDate - Date when participation in the Flexible Credit program began
	EnrollmentDate *time.Time `json:"enrollmentDate"`

	// GraduationDate - no documentation
	GraduationDate *time.Time `json:"graduationDate"`
}

func (softlayer_flexiblecredit_enrollment *SoftLayer_FlexibleCredit_Enrollment) String() string {
	return "SoftLayer_FlexibleCredit_Enrollment"
}

// SoftLayer_FlexibleCredit_Enrollment_Extended is SoftLayer_FlexibleCredit_Enrollment with all maskable types.
type SoftLayer_FlexibleCredit_Enrollment_Extended struct {
	SoftLayer_FlexibleCredit_Enrollment

	// CompanyType - no documentation
	CompanyType *SoftLayer_FlexibleCredit_Company_Type `json:"companyType"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Affiliate - no documentation
	Affiliate *SoftLayer_FlexibleCredit_Affiliate `json:"affiliate"`

	// FlexibleCreditProgram - no documentation
	FlexibleCreditProgram *SoftLayer_FlexibleCredit_Program `json:"flexibleCreditProgram"`

	// IsActiveFlag - Flag indicating whether an enrollment is active (true) or inactive (false)
	IsActiveFlag bool `json:"isActiveFlag"`

	// Representative - no documentation
	Representative *SoftLayer_User_Employee `json:"representative"`
}

func (softlayer_flexiblecredit_enrollment *SoftLayer_FlexibleCredit_Enrollment_Extended) String() string {
	return "SoftLayer_FlexibleCredit_Enrollment"
}
