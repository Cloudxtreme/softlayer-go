package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Referral_Partner_Commission - <nil>
type SoftLayer_Container_Referral_Partner_Commission struct {

	// ReferralAccountId - <nil>
	ReferralAccountId int `json:"referralAccountId"`

	// ReferralCompanyName - <nil>
	ReferralCompanyName string `json:"referralCompanyName"`

	// ReferralPartnerAccountId - <nil>
	ReferralPartnerAccountId int `json:"referralPartnerAccountId"`

	// ReferralRevenue - <nil>
	ReferralRevenue float64 `json:"referralRevenue"`

	// CommissionAmount - <nil>
	CommissionAmount float64 `json:"commissionAmount"`

	// CommissionRate - <nil>
	CommissionRate float64 `json:"commissionRate"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`
}

func (softlayer_container_referral_partner_commission *SoftLayer_Container_Referral_Partner_Commission) String() string {
	return "SoftLayer_Container_Referral_Partner_Commission"
}
