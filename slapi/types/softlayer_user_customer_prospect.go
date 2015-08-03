package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Prospect - <nil>
type SoftLayer_User_Customer_Prospect struct {
}

// SoftLayer_User_Customer_Prospect_Extended is SoftLayer_User_Customer_Prospect with all maskable types.
type SoftLayer_User_Customer_Prospect_Extended struct {
	SoftLayer_User_Customer_Prospect

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AssignedEmployees - <nil>
	AssignedEmployees []*SoftLayer_User_Employee `json:"assignedEmployees"`

	// Quotes - <nil>
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_Prospect_Type `json:"type"`

	// AssignedEmployeeCount - no documentation
	AssignedEmployeeCount uint64 `json:"assignedEmployeeCount"`
}

func (softlayer_user_customer_prospect *SoftLayer_User_Customer_Prospect) String() string {
	return "SoftLayer_User_Customer_Prospect"
}
