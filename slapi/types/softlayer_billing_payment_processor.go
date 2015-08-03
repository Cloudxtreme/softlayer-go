package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor - <nil>
type SoftLayer_Billing_Payment_Processor struct {

	// Description - <nil>
	Description string `json:"description"`

	// Name - <nil>
	Name string `json:"name"`
}

func (softlayer_billing_payment_processor *SoftLayer_Billing_Payment_Processor) String() string {
	return "SoftLayer_Billing_Payment_Processor"
}

// SoftLayer_Billing_Payment_Processor_Extended is SoftLayer_Billing_Payment_Processor with all maskable types.
type SoftLayer_Billing_Payment_Processor_Extended struct {
	SoftLayer_Billing_Payment_Processor

	// BrandAssignments - <nil>
	BrandAssignments []*SoftLayer_Brand_Payment_Processor `json:"brandAssignments"`

	// OwnerAccount - <nil>
	OwnerAccount *SoftLayer_Account `json:"ownerAccount"`

	// PaymentMethods - <nil>
	PaymentMethods []*SoftLayer_Billing_Payment_Processor_Method `json:"paymentMethods"`

	// Type - <nil>
	Type *SoftLayer_Billing_Payment_Processor_Type `json:"type"`

	// BrandAssignmentCount - no documentation
	BrandAssignmentCount uint64 `json:"brandAssignmentCount"`

	// PaymentMethodCount - no documentation
	PaymentMethodCount uint64 `json:"paymentMethodCount"`
}

func (softlayer_billing_payment_processor *SoftLayer_Billing_Payment_Processor_Extended) String() string {
	return "SoftLayer_Billing_Payment_Processor"
}
