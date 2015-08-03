package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand_Payment_Processor - <nil>
type SoftLayer_Brand_Payment_Processor struct {
}

// SoftLayer_Brand_Payment_Processor_Extended is SoftLayer_Brand_Payment_Processor with all maskable types.
type SoftLayer_Brand_Payment_Processor_Extended struct {
	SoftLayer_Brand_Payment_Processor

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// PaymentProcessor - <nil>
	PaymentProcessor *SoftLayer_Billing_Payment_Processor `json:"paymentProcessor"`
}

func (softlayer_brand_payment_processor *SoftLayer_Brand_Payment_Processor) String() string {
	return "SoftLayer_Brand_Payment_Processor"
}
