package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Prospect - <nil>
type SoftLayer_User_Customer_Prospect struct {
}

func (softlayer_user_customer_prospect *SoftLayer_User_Customer_Prospect) String() string {
	return "SoftLayer_User_Customer_Prospect"
}

// SoftLayer_User_Customer_Prospect_Extended is SoftLayer_User_Customer_Prospect with all maskable types.
type SoftLayer_User_Customer_Prospect_Extended struct {
	SoftLayer_User_Customer_Prospect

	// AssignedEmployeeCount - no documentation
	AssignedEmployeeCount uint64 `json:"assignedEmployeeCount,omitempty"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AssignedEmployees - <nil>
	AssignedEmployees []*SoftLayer_User_Employee `json:"assignedEmployees,omitempty"`

	// Quotes - <nil>
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes,omitempty"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_Prospect_Type `json:"type,omitempty"`
}

func (softlayer_user_customer_prospect *SoftLayer_User_Customer_Prospect_Extended) String() string {
	return "SoftLayer_User_Customer_Prospect"
}
