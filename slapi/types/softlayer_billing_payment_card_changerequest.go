package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Card_ChangeRequest - The SoftLayer_Billing_Payment_Card_ChangeRequest data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ChangeRequest struct {

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode,omitempty"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst,omitempty"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast,omitempty"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// BillingState - no documentation
	BillingState string `json:"billingState,omitempty"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear,omitempty"`

	// CardNickname - <nil>
	CardNickname string `json:"cardNickname,omitempty"`

	// PaymentRoleId - <nil>
	PaymentRoleId int `json:"paymentRoleId,omitempty"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount,omitempty"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1,omitempty"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail,omitempty"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax,omitempty"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId,omitempty"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId,omitempty"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity,omitempty"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany,omitempty"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType,omitempty"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice,omitempty"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName,omitempty"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2,omitempty"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4,omitempty"`

	// CardType - The type of payment card a customer has. (i.e. Visa, MasterCard, American Express).
	CardType string `json:"cardType,omitempty"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber,omitempty"`
}

func (softlayer_billing_payment_card_changerequest *SoftLayer_Billing_Payment_Card_ChangeRequest) String() string {
	return "SoftLayer_Billing_Payment_Card_ChangeRequest"
}

// SoftLayer_Billing_Payment_Card_ChangeRequest_Extended is SoftLayer_Billing_Payment_Card_ChangeRequest with all maskable types.
type SoftLayer_Billing_Payment_Card_ChangeRequest_Extended struct {
	SoftLayer_Billing_Payment_Card_ChangeRequest

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture of
	// funds performed as part of this change request.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card change request.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount,omitempty"`

	// TicketAttachmentReferences - These are tickets tied to a credit card change request.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences,omitempty"`

	// AuthorizedCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the
	// authorization performed as part of this change request.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_billing_payment_card_changerequest *SoftLayer_Billing_Payment_Card_ChangeRequest_Extended) String() string {
	return "SoftLayer_Billing_Payment_Card_ChangeRequest"
}
