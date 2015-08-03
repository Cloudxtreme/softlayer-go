package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Partner_Referral_Prospect - <nil>
type SoftLayer_Account_Partner_Referral_Prospect struct {

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// EmailAddress - <nil>
	EmailAddress string `json:"emailAddress"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// Id - <nil>
	Id int `json:"id"`

	// LastName - <nil>
	LastName string `json:"lastName"`
}

func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect) String() string {
	return "SoftLayer_Account_Partner_Referral_Prospect"
}

// CreateProspect - no documentation
func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect) CreateProspect(ctx *slapi.RequestContext, templateObject SoftLayer_Container_Referral_Partner_Prospect, commit bool) (*SoftLayer_Account_Partner_Referral_Prospect, error) {
	var returnValue *SoftLayer_Account_Partner_Referral_Prospect
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Partner_Referral_Prospect, error) {
	var returnValue *SoftLayer_Account_Partner_Referral_Prospect
	return returnValue, nil
}

// GetSurveyQuestions - no documentation
func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect) GetSurveyQuestions(ctx *slapi.RequestContext) ([]*SoftLayer_Survey_Question, error) {
	var returnValue []*SoftLayer_Survey_Question
	return returnValue, nil
}
