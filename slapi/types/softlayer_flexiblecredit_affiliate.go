package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_FlexibleCredit_Affiliate - <nil>
type SoftLayer_FlexibleCredit_Affiliate struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// FlexibleCreditProgram - no documentation
	FlexibleCreditProgram *SoftLayer_FlexibleCredit_Program `json:"flexibleCreditProgram,omitempty"`
}

func (softlayer_flexiblecredit_affiliate *SoftLayer_FlexibleCredit_Affiliate) String() string {
	return "SoftLayer_FlexibleCredit_Affiliate"
}
