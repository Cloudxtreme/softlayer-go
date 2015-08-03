package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Billing_Information - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place an order with SoftLayer.
type SoftLayer_Container_Product_Order_Billing_Information struct {

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode"`

	// CardExpirationMonth - no documentation
	CardExpirationMonth int `json:"cardExpirationMonth"`

	// TaxExempt - Tax exempt status. 1 = exempt (not taxable), 0 = not exempt (taxable)
	TaxExempt int `json:"taxExempt"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax"`

	// CreditCardVerificationNumber - no documentation
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany"`

	// CardExpirationYear - no documentation
	CardExpirationYear int `json:"cardExpirationYear"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber"`

	// VatId - no documentation
	VatId string `json:"vatId"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast"`

	// BillingState - no documentation
	BillingState string `json:"billingState"`
}

func (softlayer_container_product_order_billing_information *SoftLayer_Container_Product_Order_Billing_Information) String() string {
	return "SoftLayer_Container_Product_Order_Billing_Information"
}
