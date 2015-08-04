package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Catalyst_Enrollment - <nil>
type SoftLayer_Catalyst_Enrollment struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// AgreementCompleteFlag - <nil>
	AgreementCompleteFlag int `json:"agreementCompleteFlag,omitempty"`

	// MonthlyCreditAmount - <nil>
	MonthlyCreditAmount slapi.Float64 `json:"monthlyCreditAmount,omitempty"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// GraduationDate - <nil>
	GraduationDate *time.Time `json:"graduationDate,omitempty"`

	// CompanyTypeId - <nil>
	CompanyTypeId int `json:"companyTypeId,omitempty"`

	// EnrollmentDate - <nil>
	EnrollmentDate *time.Time `json:"enrollmentDate,omitempty"`

	// AffiliateId - <nil>
	AffiliateId int `json:"affiliateId,omitempty"`

	// RepresentativeEmployeeId - <nil>
	RepresentativeEmployeeId int `json:"representativeEmployeeId,omitempty"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}

// SoftLayer_Catalyst_Enrollment_Extended is SoftLayer_Catalyst_Enrollment with all maskable types.
type SoftLayer_Catalyst_Enrollment_Extended struct {
	SoftLayer_Catalyst_Enrollment

	// IsActiveFlag - <nil>
	IsActiveFlag bool `json:"isActiveFlag,omitempty"`

	// Representative - <nil>
	Representative *SoftLayer_User_Employee `json:"representative,omitempty"`

	// Affiliate - <nil>
	Affiliate *SoftLayer_Catalyst_Affiliate `json:"affiliate,omitempty"`

	// CompanyType - <nil>
	CompanyType *SoftLayer_Catalyst_Company_Type `json:"companyType,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment_Extended) String() string {
	return "SoftLayer_Catalyst_Enrollment"
}
