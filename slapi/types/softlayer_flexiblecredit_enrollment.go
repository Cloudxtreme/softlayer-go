package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_FlexibleCredit_Enrollment - <nil>
type SoftLayer_FlexibleCredit_Enrollment struct {

	// AffiliateId - ID of the corresponding Flexible Credit Program Affiliate
	AffiliateId int `json:"affiliateId,omitempty"`

	// AgreementCompleteFlag - Indicates signing of Flexible Credit agreement (independent from
	AgreementCompleteFlag int `json:"agreementCompleteFlag,omitempty"`

	// GraduationDate - no documentation
	GraduationDate *time.Time `json:"graduationDate,omitempty"`

	// RepresentativeEmployeeId - no documentation
	RepresentativeEmployeeId int `json:"representativeEmployeeId,omitempty"`

	// AccountId - no documentation
	AccountId int `json:"accountId,omitempty"`

	// CompanyDescription - no documentation
	CompanyDescription string `json:"companyDescription,omitempty"`

	// EnrollmentDate - Date when participation in the Flexible Credit program began
	EnrollmentDate *time.Time `json:"enrollmentDate,omitempty"`

	// CompanyTypeId - ID of the Flexible Credit Program Company classification for this enrollment
	CompanyTypeId int `json:"companyTypeId,omitempty"`

	// MonthlyCreditAmount - no documentation
	MonthlyCreditAmount slapi.Float64 `json:"monthlyCreditAmount,omitempty"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account,omitempty"`

	// FlexibleCreditProgram - no documentation
	FlexibleCreditProgram *SoftLayer_FlexibleCredit_Program `json:"flexibleCreditProgram,omitempty"`

	// Representative - no documentation
	Representative *SoftLayer_User_Employee `json:"representative,omitempty"`

	// Affiliate - no documentation
	Affiliate *SoftLayer_FlexibleCredit_Affiliate `json:"affiliate,omitempty"`

	// CompanyType - no documentation
	CompanyType *SoftLayer_FlexibleCredit_Company_Type `json:"companyType,omitempty"`

	// IsActiveFlag - Flag indicating whether an enrollment is active (true) or inactive (false)
	IsActiveFlag bool `json:"isActiveFlag,omitempty"`
}

func (softlayer_flexiblecredit_enrollment *SoftLayer_FlexibleCredit_Enrollment) String() string {
	return "SoftLayer_FlexibleCredit_Enrollment"
}
