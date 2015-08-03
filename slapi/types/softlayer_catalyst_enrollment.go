package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
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

// GetAffiliates - <nil>
func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) GetAffiliates(ctx *slapi.RequestContext) ([]*SoftLayer_Catalyst_Affiliate, error) {
	var returnValue []*SoftLayer_Catalyst_Affiliate
	return returnValue, nil
}

// GetCompanyTypes - <nil>
func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) GetCompanyTypes(ctx *slapi.RequestContext) ([]*SoftLayer_Catalyst_Company_Type, error) {
	var returnValue []*SoftLayer_Catalyst_Company_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Catalyst_Enrollment, error) {
	var returnValue *SoftLayer_Catalyst_Enrollment
	return returnValue, nil
}

// SelfEnrollNewAccount - <nil>
func (softlayer_catalyst_enrollment *SoftLayer_Catalyst_Enrollment) SelfEnrollNewAccount(ctx *slapi.RequestContext, accountTemplate SoftLayer_Account) (*SoftLayer_Account, error) {
	var returnValue *SoftLayer_Account
	return returnValue, nil
}
