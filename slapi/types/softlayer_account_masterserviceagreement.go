package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_MasterServiceAgreement - <nil>
type SoftLayer_Account_MasterServiceAgreement struct {

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// Guid - <nil>
	Guid string `json:"guid,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`
}

func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement) String() string {
	return "SoftLayer_Account_MasterServiceAgreement"
}

// SoftLayer_Account_MasterServiceAgreement_Extended is SoftLayer_Account_MasterServiceAgreement with all maskable types.
type SoftLayer_Account_MasterServiceAgreement_Extended struct {
	SoftLayer_Account_MasterServiceAgreement

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_masterserviceagreement *SoftLayer_Account_MasterServiceAgreement_Extended) String() string {
	return "SoftLayer_Account_MasterServiceAgreement"
}
