package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Card_ManualPayment - The SoftLayer_Billing_Payment_Card_ManualPayment data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ManualPayment struct {

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth,omitempty"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany,omitempty"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail,omitempty"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst,omitempty"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4,omitempty"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber,omitempty"`

	// Notes - Notes generated as a result of the payment request.
	Notes string `json:"notes,omitempty"`

	// Type - no documentation
	Type string `json:"type,omitempty"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode,omitempty"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount,omitempty"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2,omitempty"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice,omitempty"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity,omitempty"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear,omitempty"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType,omitempty"`

	// AuthorizedCreditCardTransactionId - The unique identifier of an attempted credit card transaction.
	AuthorizedCreditCardTransactionId int `json:"authorizedCreditCardTransactionId,omitempty"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast,omitempty"`

	// FromIpAddress - The IP address from which the transaction originates.
	FromIpAddress string `json:"fromIpAddress,omitempty"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1,omitempty"`

	// CardAccountHash - no documentation
	CardAccountHash string `json:"cardAccountHash,omitempty"`

	// CardType - The method key of the type payment issued (Visa - 001, Mastercard - 002, American Express
	// - 003, Discover - 004, PayPal - paypal).
	CardType string `json:"cardType,omitempty"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName,omitempty"`

	// ReturnUrl - The return URL is the page to which PayPal redirects after payment is approved.
	ReturnUrl string `json:"returnUrl,omitempty"`

	// AuthorizedPayPalTransactionId - The unique identifier of an attempted PayPal transaction.
	AuthorizedPayPalTransactionId int `json:"authorizedPayPalTransactionId,omitempty"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax,omitempty"`

	// BillingState - no documentation
	BillingState string `json:"billingState,omitempty"`

	// CancelUrl - The cancel URL is the page to which PayPal redirects if payment is not approved.
	CancelUrl string `json:"cancelUrl,omitempty"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber,omitempty"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId,omitempty"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId,omitempty"`

	// Id - The unique identifier for a single manual payment request.
	Id int `json:"id,omitempty"`
}

func (softlayer_billing_payment_card_manualpayment *SoftLayer_Billing_Payment_Card_ManualPayment) String() string {
	return "SoftLayer_Billing_Payment_Card_ManualPayment"
}

// SoftLayer_Billing_Payment_Card_ManualPayment_Extended is SoftLayer_Billing_Payment_Card_ManualPayment with all maskable types.
type SoftLayer_Billing_Payment_Card_ManualPayment_Extended struct {
	SoftLayer_Billing_Payment_Card_ManualPayment

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// with a credit card.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty"`

	// CapturePayPalTransaction - The SoftLayer_Billing_Payment_PayPal_Transaction tied to the capture
	// performed as part of this manual payment. This will only exist if the manual payment was performed
	// via PayPal.
	CapturePayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"capturePayPalTransaction,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card manual payment.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount,omitempty"`

	// AuthorizedPayPalTransaction - This is the PayPal transaction data tied to a PayPal manual payment.
	AuthorizedPayPalTransaction *SoftLayer_Billing_Payment_PayPal_Transaction `json:"authorizedPayPalTransaction,omitempty"`

	// AuthorizedCreditCardTransaction - This is the credit card transaction data tied to a credit card
	// manual payment.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty"`

	// TicketAttachmentReferences - These are tickets tied to a credit card manual payment.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences,omitempty"`
}

func (softlayer_billing_payment_card_manualpayment *SoftLayer_Billing_Payment_Card_ManualPayment_Extended) String() string {
	return "SoftLayer_Billing_Payment_Card_ManualPayment"
}
