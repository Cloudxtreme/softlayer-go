package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Agreement - <nil>
type SoftLayer_Account_Agreement struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// AutoRenew - <nil>
	AutoRenew int `json:"autoRenew,omitempty"`

	// AgreementTypeId - no documentation
	AgreementTypeId int `json:"agreementTypeId,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// CancellationFee - <nil>
	CancellationFee int `json:"cancellationFee,omitempty"`

	// DurationMonths - no documentation
	DurationMonths int `json:"durationMonths,omitempty"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// BillingItems - no documentation
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Account_Agreement_Status `json:"status,omitempty"`

	// TopLevelBillingItemCount - A count of the top level billing item associated with an agreement.
	TopLevelBillingItemCount uint64 `json:"topLevelBillingItemCount,omitempty"`

	// AgreementType - no documentation
	AgreementType *SoftLayer_Account_Agreement_Type `json:"agreementType,omitempty"`

	// AttachedBillingAgreementFileCount - no documentation
	AttachedBillingAgreementFileCount uint64 `json:"attachedBillingAgreementFileCount,omitempty"`

	// BillingItemCount - A count of the billing items associated with an agreement.
	BillingItemCount uint64 `json:"billingItemCount,omitempty"`

	// AttachedBillingAgreementFiles - no documentation
	AttachedBillingAgreementFiles []*SoftLayer_Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles,omitempty"`

	// TopLevelBillingItems - The top level billing item associated with an agreement.
	TopLevelBillingItems []*SoftLayer_Billing_Item `json:"topLevelBillingItems,omitempty"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement) String() string {
	return "SoftLayer_Account_Agreement"
}
