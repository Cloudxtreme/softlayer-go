package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_FlexibleCredit_Program - <nil>
type SoftLayer_FlexibleCredit_Program struct {

	// Id - no documentation
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// Name - no documentation
	Name string `json:"name"`
}

// GetAffiliatesAvailableForSelfEnrollmentByVerificationType - <nil>
func (softlayer_flexiblecredit_program *SoftLayer_FlexibleCredit_Program) GetAffiliatesAvailableForSelfEnrollmentByVerificationType(ctx *slapi.RequestContext, verificationTypeKeyName string) ([]*SoftLayer_FlexibleCredit_Affiliate, error) {
	var returnValue []*SoftLayer_FlexibleCredit_Affiliate
	return returnValue, nil
}

// GetCompanyTypes - <nil>
func (softlayer_flexiblecredit_program *SoftLayer_FlexibleCredit_Program) GetCompanyTypes(ctx *slapi.RequestContext) ([]*SoftLayer_FlexibleCredit_Company_Type, error) {
	var returnValue []*SoftLayer_FlexibleCredit_Company_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_flexiblecredit_program *SoftLayer_FlexibleCredit_Program) GetObject(ctx *slapi.RequestContext) (*SoftLayer_FlexibleCredit_Program, error) {
	var returnValue *SoftLayer_FlexibleCredit_Program
	return returnValue, nil
}

// SelfEnrollNewAccount - <nil>
func (softlayer_flexiblecredit_program *SoftLayer_FlexibleCredit_Program) SelfEnrollNewAccount(ctx *slapi.RequestContext, accountTemplate SoftLayer_Account) (*SoftLayer_Account, error) {
	var returnValue *SoftLayer_Account
	return returnValue, nil
}
