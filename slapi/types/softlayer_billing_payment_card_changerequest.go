package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Billing_Payment_Card_ChangeRequest - The SoftLayer_Billing_Payment_Card_ChangeRequest data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ChangeRequest struct {

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId"`

	// Id - no documentation
	Id int `json:"id"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName"`

	// CardNickname - <nil>
	CardNickname string `json:"cardNickname"`

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany"`

	// CardType - The type of payment card a customer has. (i.e. Visa, MasterCard, American Express).
	CardType string `json:"cardType"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount float64 `json:"amount"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth"`

	// PaymentRoleId - <nil>
	PaymentRoleId int `json:"paymentRoleId"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice"`

	// BillingState - no documentation
	BillingState string `json:"billingState"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4"`
}

func (softlayer_billing_payment_card_changerequest *SoftLayer_Billing_Payment_Card_ChangeRequest) String() string {
	return "SoftLayer_Billing_Payment_Card_ChangeRequest"
}

// SoftLayer_Billing_Payment_Card_ChangeRequest_Extended is SoftLayer_Billing_Payment_Card_ChangeRequest with all maskable types.
type SoftLayer_Billing_Payment_Card_ChangeRequest_Extended struct {
	SoftLayer_Billing_Payment_Card_ChangeRequest

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card change request.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount"`

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture of
	// funds performed as part of this change request.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction"`

	// AuthorizedCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the
	// authorization performed as part of this change request.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction"`

	// TicketAttachmentReferences - These are tickets tied to a credit card change request.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences"`
}

func (softlayer_billing_payment_card_changerequest *SoftLayer_Billing_Payment_Card_ChangeRequest_Extended) String() string {
	return "SoftLayer_Billing_Payment_Card_ChangeRequest"
}
