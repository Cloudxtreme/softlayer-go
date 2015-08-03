package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Agreement - <nil>
type SoftLayer_Account_Agreement struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AgreementType - no documentation
	AgreementType *SoftLayer_Account_Agreement_Type `json:"agreementType"`

	// AgreementTypeId - no documentation
	AgreementTypeId int `json:"agreementTypeId"`

	// AttachedBillingAgreementFileCount - no documentation
	AttachedBillingAgreementFileCount uint64 `json:"attachedBillingAgreementFileCount"`

	// AttachedBillingAgreementFiles - no documentation
	AttachedBillingAgreementFiles []*SoftLayer_Account_MasterServiceAgreement `json:"attachedBillingAgreementFiles"`

	// AutoRenew - <nil>
	AutoRenew int `json:"autoRenew"`

	// BillingItemCount - A count of the billing items associated with an agreement.
	BillingItemCount uint64 `json:"billingItemCount"`

	// BillingItems - no documentation
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// CancellationFee - <nil>
	CancellationFee int `json:"cancellationFee"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DurationMonths - no documentation
	DurationMonths int `json:"durationMonths"`

	// EndDate - no documentation
	EndDate *time.Time `json:"endDate"`

	// Id - no documentation
	Id int `json:"id"`

	// StartDate - no documentation
	StartDate *time.Time `json:"startDate"`

	// Status - no documentation
	Status *SoftLayer_Account_Agreement_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// Title - no documentation
	Title string `json:"title"`

	// TopLevelBillingItemCount - A count of the top level billing item associated with an agreement.
	TopLevelBillingItemCount uint64 `json:"topLevelBillingItemCount"`

	// TopLevelBillingItems - The top level billing item associated with an agreement.
	TopLevelBillingItems []*SoftLayer_Billing_Item `json:"topLevelBillingItems"`
}

func (softlayer_account_agreement *SoftLayer_Account_Agreement) String() string {
	return "SoftLayer_Account_Agreement"
}

// GetObject - <nil>
func (softlayer_account_agreement *SoftLayer_Account_Agreement) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Agreement, error) {
	var returnValue *SoftLayer_Account_Agreement
	return returnValue, nil
}
