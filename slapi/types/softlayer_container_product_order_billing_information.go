package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Billing_Information - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place an order with SoftLayer.
type SoftLayer_Container_Product_Order_Billing_Information struct {

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2,omitempty"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	// BillingState - no documentation
	BillingState string `json:"billingState,omitempty"`

	// CreditCardVerificationNumber - no documentation
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber,omitempty"`

	// VatId - no documentation
	VatId string `json:"vatId,omitempty"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst,omitempty"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast,omitempty"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1,omitempty"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity,omitempty"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode,omitempty"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail,omitempty"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany,omitempty"`

	// TaxExempt - Tax exempt status. 1 = exempt (not taxable), 0 = not exempt (taxable)
	TaxExempt int `json:"taxExempt,omitempty"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax,omitempty"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice,omitempty"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber,omitempty"`

	// CardExpirationMonth - no documentation
	CardExpirationMonth int `json:"cardExpirationMonth,omitempty"`

	// CardExpirationYear - no documentation
	CardExpirationYear int `json:"cardExpirationYear,omitempty"`
}

func (softlayer_container_product_order_billing_information *SoftLayer_Container_Product_Order_Billing_Information) String() string {
	return "SoftLayer_Container_Product_Order_Billing_Information"
}
