package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Prospect - <nil>
type SoftLayer_User_Customer_Prospect struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AssignedEmployees - <nil>
	AssignedEmployees []*SoftLayer_User_Employee `json:"assignedEmployees,omitempty"`

	// Quotes - <nil>
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes,omitempty"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_Prospect_Type `json:"type,omitempty"`

	// AssignedEmployeeCount - no documentation
	AssignedEmployeeCount uint64 `json:"assignedEmployeeCount,omitempty"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount,omitempty"`
}

func (softlayer_user_customer_prospect *SoftLayer_User_Customer_Prospect) String() string {
	return "SoftLayer_User_Customer_Prospect"
}
