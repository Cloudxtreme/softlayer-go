package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_MasterServiceAgreement - <nil>
type SoftLayer_Account_MasterServiceAgreement struct {

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Guid - <nil>
	Guid string `json:"guid"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement) String() string {
	return "SoftLayer_Account_MasterServiceAgreement"
}

// SoftLayer_Account_MasterServiceAgreement_Extended is SoftLayer_Account_MasterServiceAgreement with all maskable types.
type SoftLayer_Account_MasterServiceAgreement_Extended struct {
	SoftLayer_Account_MasterServiceAgreement

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`
}

func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement_Extended) String() string {
	return "SoftLayer_Account_MasterServiceAgreement"
}
