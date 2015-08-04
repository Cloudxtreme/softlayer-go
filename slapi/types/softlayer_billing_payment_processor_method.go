package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor_Method - <nil>
type SoftLayer_Billing_Payment_Processor_Method struct {

	// MethodKey - <nil>
	MethodKey string `json:"methodKey,omitempty"`

	// MultipleCurrencyFlag - <nil>
	MultipleCurrencyFlag bool `json:"multipleCurrencyFlag,omitempty"`

	// PaymentProcessor - <nil>
	PaymentProcessor *SoftLayer_Billing_Payment_Processor `json:"paymentProcessor,omitempty"`

	// PaymentType - <nil>
	PaymentType *SoftLayer_Billing_Payment_Type `json:"paymentType,omitempty"`
}

func (softlayer_billing_payment_processor_method *SoftLayer_Billing_Payment_Processor_Method) String() string {
	return "SoftLayer_Billing_Payment_Processor_Method"
}
