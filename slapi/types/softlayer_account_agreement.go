package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Agreement - <nil>
type SoftLayer_Account_Agreement struct {

	// DurationMonths - no documentation
	DurationMonths int `json:"durationMonths"`

	// Id - no documentation
	Id int `json:"id"`

	// CancellationFee - <nil>
	CancellationFee int `json:"cancellationFee"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// AgreementTypeId - no documentation
	AgreementTypeId int `json:"agreementTypeId"`

	// AutoRenew - <nil>
	AutoRenew int `json:"autoRenew"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Title - no documentation
	Title string `json:"title"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement) String() string {
	return "SoftLayer_Account_Agreement"
}

// SoftLayer_Account_Agreement_Extended is SoftLayer_Account_Agreement with all maskable types.
type SoftLayer_Account_Agreement_Extended struct {
	SoftLayer_Account_Agreement

	// AttachedBillingAgreementFileCount - no documentation
	AttachedBillingAgreementFileCount uint64 `json:"attachedBillingAgreementFileCount"`

	// AgreementType - no documentation
	AgreementType *SoftLayer_Account_Agreement_Type `json:"agreementType"`

	// TopLevelBillingItems - The top level billing item associated with an agreement.
	TopLevelBillingItems []*SoftLayer_Billing_Item `json:"topLevelBillingItems"`

	// TopLevelBillingItemCount - A count of the top level billing item associated with an agreement.
	TopLevelBillingItemCount uint64 `json:"topLevelBillingItemCount"`

	// AttachedBillingAgreementFiles - no documentation
	AttachedBillingAgreementFiles []*SoftLayer_Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles"`

	// BillingItems - no documentation
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// BillingItemCount - A count of the billing items associated with an agreement.
	BillingItemCount uint64 `json:"billingItemCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Status - no documentation
	Status *SoftLayer_Account_Agreement_Status `json:"status"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement_Extended) String() string {
	return "SoftLayer_Account_Agreement"
}
