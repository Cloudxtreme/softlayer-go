package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Billing_Info - Every SoftLayer customer account has billing specific information which is
// kept in the SoftLayer_Billing_Info data type. This information is used by the SoftLayer accounting
// group when sending invoices and making billing inquiries.
type SoftLayer_Billing_Info struct {

	// LastFourPaymentCardDigits - The last four digits of the credit card currently on the account. This
	// is the only portion of the card that we store. For Paypal customers, this value will be empty.
	LastFourPaymentCardDigits int `json:"lastFourPaymentCardDigits"`

	// CardExpirationMonth - no documentation
	CardExpirationMonth int `json:"cardExpirationMonth"`

	// CardExpirationYear - no documentation
	CardExpirationYear int `json:"cardExpirationYear"`

	// CreateDate - The date a customer's billing information was created.
	CreateDate *time.Time `json:"createDate"`

	// Id - A SoftLayer customer's billing information identifier.
	Id int `json:"id"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// AnniversaryDayOfMonth - The day of the month that a SoftLayer customer is billed.
	AnniversaryDayOfMonth int `json:"anniversaryDayOfMonth"`

	// LastPaymentDate - The date of the last payment received by SoftLayer from the account holder.
	LastPaymentDate *time.Time `json:"lastPaymentDate"`

	// PercentDiscountRecurring - The percentage discount received on all recurring charges on a customer's
	// monthly bill.
	PercentDiscountRecurring int `json:"percentDiscountRecurring"`

	// SparePoolAmount - The total recurring fee amount for servers that are in the spare pool status.
	SparePoolAmount int `json:"sparePoolAmount"`

	// VatId - <nil>
	VatId string `json:"vatId"`

	// CardType - no documentation
	CardType string `json:"cardType"`

	// PaymentTerms - no documentation
	PaymentTerms int `json:"paymentTerms"`

	// PercentDiscountOnetime - The percentage discount received on all one-time charges on a customer's
	// monthly bill.
	PercentDiscountOnetime int `json:"percentDiscountOnetime"`

	// CardAccountNumber - This value doesn't persist to this object. It's used as part of the account
	// creation process only;
	CardAccountNumber string `json:"cardAccountNumber"`

	// CardNickname - <nil>
	CardNickname string `json:"cardNickname"`

	// CardVerificationNumber - This value doesn't persist to this object. It's used as part of the account
	// creation process only.
	CardVerificationNumber string `json:"cardVerificationNumber"`

	// ModifyDate - The date a customer's billing information was last modified.
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_billing_info *SoftLayer_Billing_Info) String() string {
	return "SoftLayer_Billing_Info"
}

// SoftLayer_Billing_Info_Extended is SoftLayer_Billing_Info with all maskable types.
type SoftLayer_Billing_Info_Extended struct {
	SoftLayer_Billing_Info

	// Account - The SoftLayer customer account associated with this billing information.
	Account *SoftLayer_Account `json:"account"`

	// AchInformationCount - no documentation
	AchInformationCount uint64 `json:"achInformationCount"`

	// Currency - no documentation
	Currency *SoftLayer_Billing_Currency `json:"currency"`

	// LastBillDate - no documentation
	LastBillDate *time.Time `json:"lastBillDate"`

	// AchInformation - <nil>
	AchInformation []*SoftLayer_Billing_Info_Ach `json:"achInformation"`

	// CurrentBillingCycle - Information related to an account's current and previous billing cycles.
	CurrentBillingCycle *SoftLayer_Billing_Info_Cycle `json:"currentBillingCycle"`

	// NextBillDate - no documentation
	NextBillDate *time.Time `json:"nextBillDate"`
}

func (softlayer_billing_info *SoftLayer_Billing_Info_Extended) String() string {
	return "SoftLayer_Billing_Info"
}
