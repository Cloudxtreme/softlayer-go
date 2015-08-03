package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor - <nil>
type SoftLayer_Billing_Payment_Processor struct {

	// BrandAssignmentCount - no documentation
	BrandAssignmentCount uint64 `json:"brandAssignmentCount"`

	// BrandAssignments - <nil>
	BrandAssignments []*SoftLayer_Brand_Payment_Processor `json:"brandAssignments"`

	// Description - <nil>
	Description string `json:"description"`

	// Name - <nil>
	Name string `json:"name"`

	// OwnerAccount - <nil>
	OwnerAccount *SoftLayer_Account `json:"ownerAccount"`

	// PaymentMethodCount - no documentation
	PaymentMethodCount uint64 `json:"paymentMethodCount"`

	// PaymentMethods - <nil>
	PaymentMethods []*SoftLayer_Billing_Payment_Processor_Method `json:"paymentMethods"`

	// Type - <nil>
	Type *SoftLayer_Billing_Payment_Processor_Type `json:"type"`
}

func (softlayer_billing_payment_processor *SoftLayer_Billing_Payment_Processor) String() string {
	return "SoftLayer_Billing_Payment_Processor"
}
