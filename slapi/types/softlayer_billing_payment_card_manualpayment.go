package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Card_ManualPayment - The SoftLayer_Billing_Payment_Card_ManualPayment data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ManualPayment struct {

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice"`

	// CardAccountHash - no documentation
	CardAccountHash string `json:"cardAccountHash"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber"`

	// Id - The unique identifier for a single manual payment request.
	Id int `json:"id"`

	// CancelUrl - The cancel URL is the page to which PayPal redirects if payment is not approved.
	CancelUrl string `json:"cancelUrl"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName"`

	// Type - no documentation
	Type string `json:"type"`

	// CardType - The method key of the type payment issued (Visa - 001, Mastercard - 002, American Express
	// - 003, Discover - 004, PayPal - paypal).
	CardType string `json:"cardType"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber"`

	// Notes - Notes generated as a result of the payment request.
	Notes string `json:"notes"`

	// AuthorizedPayPalTransactionId - The unique identifier of an attempted PayPal transaction.
	AuthorizedPayPalTransactionId int `json:"authorizedPayPalTransactionId"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount"`

	// FromIpAddress - The IP address from which the transaction originates.
	FromIpAddress string `json:"fromIpAddress"`

	// ReturnUrl - The return URL is the page to which PayPal redirects after payment is approved.
	ReturnUrl string `json:"returnUrl"`

	// AuthorizedCreditCardTransactionId - The unique identifier of an attempted credit card transaction.
	AuthorizedCreditCardTransactionId int `json:"authorizedCreditCardTransactionId"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst"`

	// BillingState - no documentation
	BillingState string `json:"billingState"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear"`
}

// SoftLayer_Billing_Payment_Card_ManualPayment_Extended is SoftLayer_Billing_Payment_Card_ManualPayment with all maskable types.
type SoftLayer_Billing_Payment_Card_ManualPayment_Extended struct {
	SoftLayer_Billing_Payment_Card_ManualPayment

	// AuthorizedCreditCardTransaction - This is the credit card transaction data tied to a credit card
	// manual payment.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card manual payment.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount"`

	// CapturePayPalTransaction - The SoftLayer_Billing_Payment_PayPal_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// via PayPal.
	CapturePayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction"`

	// TicketAttachmentReferences - These are tickets tied to a credit card manual payment.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences"`

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// with a credit card.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AuthorizedPayPalTransaction - This is the PayPal transaction data tied to a PayPal manual payment.
	AuthorizedPayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction"`
}

func (softlayer_billing_payment_card_manualpayment *SoftLayer_Billing_Payment_Card_ManualPayment) String() string {
	return "SoftLayer_Billing_Payment_Card_ManualPayment"
}
