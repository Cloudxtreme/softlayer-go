package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_Prospect - <nil>
type SoftLayer_User_Customer_Prospect struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AssignedEmployeeCount - no documentation
	AssignedEmployeeCount uint64 `json:"assignedEmployeeCount"`

	// AssignedEmployees - <nil>
	AssignedEmployees []*SoftLayer_User_Employee `json:"assignedEmployees"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount"`

	// Quotes - <nil>
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes"`

	// Type - <nil>
	Type *SoftLayer_User_Customer_Prospect_Type `json:"type"`
}
