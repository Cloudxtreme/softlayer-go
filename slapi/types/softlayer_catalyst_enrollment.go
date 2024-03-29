package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId,omitempty"`

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag,omitempty"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate,omitempty"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate,omitempty"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount slapi.Float64 `json:"monthlyCreditAmount,omitempty"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId,omitempty"`

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType,omitempty"`

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag,omitempty"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative,omitempty"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
