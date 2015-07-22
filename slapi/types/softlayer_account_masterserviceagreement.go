package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_MasterServiceAgreement - <nil>
type SoftLayer_Account_MasterServiceAgreement struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Guid - <nil>
	Guid string `json:"guid"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetObject - <nil>
func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_MasterServiceAgreement, error) {
	var returnValue *SoftLayer_Account_MasterServiceAgreement
	return returnValue, nil
}

// GetPdf - <nil>
func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement) GetPdf(commonOptions *slapi.CommonOptions, accountId int, guid string) (string, error) {
	var returnValue string
	return returnValue, nil
}