package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor_Type - <nil>
type SoftLayer_Billing_Payment_Processor_Type struct {

	// Description - <nil>
	Description string `json:"description"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// SoftLayer_Billing_Payment_Processor_Type_Extended is SoftLayer_Billing_Payment_Processor_Type with all maskable types.
type SoftLayer_Billing_Payment_Processor_Type_Extended struct {
	SoftLayer_Billing_Payment_Processor_Type

	// PaymentProcessorCount - no documentation
	PaymentProcessorCount uint64 `json:"paymentProcessorCount"`

	// PaymentProcessors - <nil>
	PaymentProcessors []*SoftLayer_Billing_Payment_Processor `json:"paymentProcessors"`
}

func (softlayer_billing_payment_processor_type *SoftLayer_Billing_Payment_Processor_Type) String() string {
	return "SoftLayer_Billing_Payment_Processor_Type"
}
