package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Referral_Partner_Commission - <nil>
type SoftLayer_Container_Referral_Partner_Commission struct {

	// ReferralRevenue - <nil>
	ReferralRevenue float64 `json:"referralRevenue,omitempty"`

	// CommissionAmount - <nil>
	CommissionAmount float64 `json:"commissionAmount,omitempty"`

	// CommissionRate - <nil>
	CommissionRate float64 `json:"commissionRate,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ReferralAccountId - <nil>
	ReferralAccountId int `json:"referralAccountId,omitempty"`

	// ReferralCompanyName - <nil>
	ReferralCompanyName string `json:"referralCompanyName,omitempty"`

	// ReferralPartnerAccountId - <nil>
	ReferralPartnerAccountId int `json:"referralPartnerAccountId,omitempty"`
}

func (softlayer_container_referral_partner_commission *SoftLayer_Container_Referral_Partner_Commission) String() string {
	return "SoftLayer_Container_Referral_Partner_Commission"
}
