package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Referral_Partner_Commission - <nil>
type SoftLayer_Container_Referral_Partner_Commission struct {

	// CommissionAmount - <nil>
	CommissionAmount float64 `json:"commissionAmount"`

	// CommissionRate - <nil>
	CommissionRate float64 `json:"commissionRate"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// ReferralAccountId - <nil>
	ReferralAccountId int `json:"referralAccountId"`

	// ReferralCompanyName - <nil>
	ReferralCompanyName string `json:"referralCompanyName"`

	// ReferralPartnerAccountId - <nil>
	ReferralPartnerAccountId int `json:"referralPartnerAccountId"`

	// ReferralRevenue - <nil>
	ReferralRevenue float64 `json:"referralRevenue"`
}
