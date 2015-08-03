package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Processor_Method - <nil>
type SoftLayer_Billing_Payment_Processor_Method struct {

	// MethodKey - <nil>
	MethodKey string `json:"methodKey"`

	// MultipleCurrencyFlag - <nil>
	MultipleCurrencyFlag bool `json:"multipleCurrencyFlag"`
}

func (softlayer_billing_payment_processor_method *SoftLayer_Billing_Payment_Processor_Method) String() string {
	return "SoftLayer_Billing_Payment_Processor_Method"
}

// SoftLayer_Billing_Payment_Processor_Method_Extended is SoftLayer_Billing_Payment_Processor_Method with all maskable types.
type SoftLayer_Billing_Payment_Processor_Method_Extended struct {
	SoftLayer_Billing_Payment_Processor_Method

	// PaymentProcessor - <nil>
	PaymentProcessor *SoftLayer_Billing_Payment_Processor `json:"paymentProcessor"`

	// PaymentType - <nil>
	PaymentType *SoftLayer_Billing_Payment_Type `json:"paymentType"`
}

func (softlayer_billing_payment_processor_method *SoftLayer_Billing_Payment_Processor_Method_Extended) String() string {
	return "SoftLayer_Billing_Payment_Processor_Method"
}
