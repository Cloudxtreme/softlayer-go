package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor_Type - <nil>
type SoftLayer_Billing_Payment_Processor_Type struct {

	// KeyName - <nil>
	KeyName string `json:"keyName,omitempty"`

	// Name - <nil>
	Name string `json:"name,omitempty"`

	// Description - <nil>
	Description string `json:"description,omitempty"`
}

func (softlayer_billing_payment_processor_type *SoftLayer_Billing_Payment_Processor_Type) String() string {
	return "SoftLayer_Billing_Payment_Processor_Type"
}

// SoftLayer_Billing_Payment_Processor_Type_Extended is SoftLayer_Billing_Payment_Processor_Type with all maskable types.
type SoftLayer_Billing_Payment_Processor_Type_Extended struct {
	SoftLayer_Billing_Payment_Processor_Type

	// PaymentProcessors - <nil>
	PaymentProcessors []*SoftLayer_Billing_Payment_Processor `json:"paymentProcessors,omitempty"`

	// PaymentProcessorCount - no documentation
	PaymentProcessorCount uint64 `json:"paymentProcessorCount,omitempty"`
}

func (softlayer_billing_payment_processor_type *SoftLayer_Billing_Payment_Processor_Type_Extended) String() string {
	return "SoftLayer_Billing_Payment_Processor_Type"
}
