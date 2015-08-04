package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Partner_Referral_Prospect - <nil>
type SoftLayer_Account_Partner_Referral_Prospect struct {

	// EmailAddress - <nil>
	EmailAddress string `json:"emailAddress,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`
}

func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect) String() string {
	return "SoftLayer_Account_Partner_Referral_Prospect"
}

// SoftLayer_Account_Partner_Referral_Prospect_Extended is SoftLayer_Account_Partner_Referral_Prospect with all maskable types.
type SoftLayer_Account_Partner_Referral_Prospect_Extended struct {
	SoftLayer_Account_Partner_Referral_Prospect

	// AssignedEmployeeCount - no documentation
	AssignedEmployeeCount uint64 `json:"assignedEmployeeCount,omitempty"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_Prospect_Type `json:"type,omitempty"`

	// AssignedEmployees - <nil>
	AssignedEmployees []*SoftLayer_User_Employee `json:"assignedEmployees,omitempty"`

	// Quotes - <nil>
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes,omitempty"`
}

func (softlayer_account_partner_referral_prospect *SoftLayer_Account_Partner_Referral_Prospect_Extended) String() string {
	return "SoftLayer_Account_Partner_Referral_Prospect"
}
