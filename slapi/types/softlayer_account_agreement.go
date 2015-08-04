package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Agreement - <nil>
type SoftLayer_Account_Agreement struct {

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate,omitempty"`

	// CancellationFee - <nil>
	CancellationFee int `json:"cancellationFee,omitempty"`

	// DurationMonths - no documentation
	DurationMonths int `json:"durationMonths,omitempty"`

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// AutoRenew - <nil>
	AutoRenew int `json:"autoRenew,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// StatusId - no documentation
	StatusId int `json:"statusId,omitempty"`

	// AgreementTypeId - no documentation
	AgreementTypeId int `json:"agreementTypeId,omitempty"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate,omitempty"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement) String() string {
	return "SoftLayer_Account_Agreement"
}

// SoftLayer_Account_Agreement_Extended is SoftLayer_Account_Agreement with all maskable types.
type SoftLayer_Account_Agreement_Extended struct {
	SoftLayer_Account_Agreement

	// AgreementType - no documentation
	AgreementType *SoftLayer_Account_Agreement_Type `json:"agreementType,omitempty"`

	// Status - no documentation
	Status *SoftLayer_Account_Agreement_Status `json:"status,omitempty"`

	// BillingItems - no documentation
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems,omitempty"`

	// TopLevelBillingItemCount - A count of the top level billing item associated with an agreement.
	TopLevelBillingItemCount uint64 `json:"topLevelBillingItemCount,omitempty"`

	// TopLevelBillingItems - The top level billing item associated with an agreement.
	TopLevelBillingItems []*SoftLayer_Billing_Item `json:"topLevelBillingItems,omitempty"`

	// BillingItemCount - A count of the billing items associated with an agreement.
	BillingItemCount uint64 `json:"billingItemCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AttachedBillingAgreementFiles - no documentation
	AttachedBillingAgreementFiles []*SoftLayer_Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles,omitempty"`

	// AttachedBillingAgreementFileCount - no documentation
	AttachedBillingAgreementFileCount uint64 `json:"attachedBillingAgreementFileCount,omitempty"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement_Extended) String() string {
	return "SoftLayer_Account_Agreement"
}
