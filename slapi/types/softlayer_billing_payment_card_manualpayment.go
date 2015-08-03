package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Card_ManualPayment - The SoftLayer_Billing_Payment_Card_ManualPayment data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ManualPayment struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount"`

	// AuthorizedCreditCardTransaction - This is the credit card transaction data tied to a credit card
	// manual payment.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction"`

	// AuthorizedCreditCardTransactionId - The unique identifier of an attempted credit card transaction.
	AuthorizedCreditCardTransactionId int `json:"authorizedCreditCardTransactionId"`

	// AuthorizedPayPalTransaction - This is the PayPal transaction data tied to a PayPal manual payment.
	AuthorizedPayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction"`

	// AuthorizedPayPalTransactionId - The unique identifier of an attempted PayPal transaction.
	AuthorizedPayPalTransactionId int `json:"authorizedPayPalTransactionId"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode"`

	// BillingState - no documentation
	BillingState string `json:"billingState"`

	// CancelUrl - The cancel URL is the page to which PayPal redirects if payment is not approved.
	CancelUrl string `json:"cancelUrl"`

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// with a credit card.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction"`

	// CapturePayPalTransaction - The SoftLayer_Billing_Payment_PayPal_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// via PayPal.
	CapturePayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction"`

	// CardAccountHash - no documentation
	CardAccountHash string `json:"cardAccountHash"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear"`

	// CardType - The method key of the type payment issued (Visa - 001, Mastercard - 002, American Express
	// - 003, Discover - 004, PayPal - paypal).
	CardType string `json:"cardType"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId"`

	// FromIpAddress - The IP address from which the transaction originates.
	FromIpAddress string `json:"fromIpAddress"`

	// Id - The unique identifier for a single manual payment request.
	Id int `json:"id"`

	// Notes - Notes generated as a result of the payment request.
	Notes string `json:"notes"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType"`

	// ReturnUrl - The return URL is the page to which PayPal redirects after payment is approved.
	ReturnUrl string `json:"returnUrl"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card manual payment.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount"`

	// TicketAttachmentReferences - These are tickets tied to a credit card manual payment.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences"`

	// Type - no documentation
	Type string `json:"type"`
}

func (softlayer_billing_payment_card_manualpayment *SoftLayer_Billing_Payment_Card_ManualPayment) String() string {
	return "SoftLayer_Billing_Payment_Card_ManualPayment"
}
