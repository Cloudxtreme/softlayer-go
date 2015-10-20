package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Billing_Payment_Card_ChangeRequest - The SoftLayer_Billing_Payment_Card_ChangeRequest data
// type contains general information relating to attempted credit card information changes.
type SoftLayer_Billing_Payment_Card_ChangeRequest struct {

	// BillingAddressLine2 - The second line in the address. Information such as suite number goes here.
	BillingAddressLine2 string `json:"billingAddressLine2,omitempty"`

	// BillingCity - no documentation
	BillingCity string `json:"billingCity,omitempty"`

	// BillingNameCompany - no documentation
	BillingNameCompany string `json:"billingNameCompany,omitempty"`

	// AccountId - The account ID to which the credit card and billing information is associated with.
	AccountId int `json:"accountId,omitempty"`

	// CardAccountLast4 - no documentation
	CardAccountLast4 string `json:"cardAccountLast4,omitempty"`

	// CardExpirationYear - The year in which a customer's payment card will expire.
	CardExpirationYear string `json:"cardExpirationYear,omitempty"`

	// CreditCardVerificationNumber - The credit card verification number submitted in the change request.
	CreditCardVerificationNumber string `json:"creditCardVerificationNumber,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// PaymentType - The description of the type of payment sent in a change transaction.
	PaymentType string `json:"paymentType,omitempty"`

	// BillingEmail - The email address associated with a customer account.
	BillingEmail string `json:"billingEmail,omitempty"`

	// BillingPhoneFax - no documentation
	BillingPhoneFax string `json:"billingPhoneFax,omitempty"`

	// BillingCountryCode - The 2-character Country code for an account's address. (i.e.
	BillingCountryCode string `json:"billingCountryCode,omitempty"`

	// BillingPostalCode - The Zip or Postal Code for the billing address on an account.
	BillingPostalCode string `json:"billingPostalCode,omitempty"`

	// BillingPhoneVoice - The phone number associated with a customer account.
	BillingPhoneVoice string `json:"billingPhoneVoice,omitempty"`

	// CardAccountNumber - no documentation
	CardAccountNumber string `json:"cardAccountNumber,omitempty"`

	// CardNickname - <nil>
	CardNickname string `json:"cardNickname,omitempty"`

	// PaymentRoleId - <nil>
	PaymentRoleId int `json:"paymentRoleId,omitempty"`

	// BillingAddressLine1 - The physical street address. Reserve information such as "apartment #123" or
	// "Suite 2" for line 1.
	BillingAddressLine1 string `json:"billingAddressLine1,omitempty"`

	// BillingNameLast - no documentation
	BillingNameLast string `json:"billingNameLast,omitempty"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId,omitempty"`

	// Amount - The total amount of the attempted transaction, represented in decimal format as US Dollars
	Amount slapi.Float64 `json:"amount,omitempty"`

	// BillingNameFirst - no documentation
	BillingNameFirst string `json:"billingNameFirst,omitempty"`

	// BillingState - no documentation
	BillingState string `json:"billingState,omitempty"`

	// CardExpirationMonth - The month in which a customer's payment card will expire.
	CardExpirationMonth string `json:"cardExpirationMonth,omitempty"`

	// CardType - The type of payment card a customer has. (i.e. Visa, MasterCard, American Express).
	CardType string `json:"cardType,omitempty"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// CaptureCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the capture of
	// funds performed as part of this change request.
	CaptureCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"captureCreditCardTransaction,omitempty"`

	// TicketAttachmentReferences - These are tickets tied to a credit card change request.
	TicketAttachmentReferences []*SoftLayer_Ticket_Attachment `json:"ticketAttachmentReferences,omitempty"`

	// TicketAttachmentReferenceCount - A count of these are tickets tied to a credit card change request.
	TicketAttachmentReferenceCount uint64 `json:"ticketAttachmentReferenceCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AuthorizedCreditCardTransaction - The SoftLayer_Billing_Payment_Card_Transaction tied to the
	// authorization performed as part of this change request.
	AuthorizedCreditCardTransaction *SoftLayer_Billing_Payment_Card_Transaction `json:"authorizedCreditCardTransaction,omitempty"`
}

func (softlayer_billing_payment_card_changerequest *SoftLayer_Billing_Payment_Card_ChangeRequest) String() string {
	return "SoftLayer_Billing_Payment_Card_ChangeRequest"
}
