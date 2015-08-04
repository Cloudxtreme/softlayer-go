package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Container_Referral_Partner_Commission - <nil>
type SoftLayer_Container_Referral_Partner_Commission struct {

	// ReferralCompanyName - <nil>
	ReferralCompanyName string `json:"referralCompanyName,omitempty"`

	// ReferralPartnerAccountId - <nil>
	ReferralPartnerAccountId int `json:"referralPartnerAccountId,omitempty"`

	// ReferralRevenue - <nil>
	ReferralRevenue slapi.Float64 `json:"referralRevenue,omitempty"`

	// CommissionAmount - <nil>
	CommissionAmount slapi.Float64 `json:"commissionAmount,omitempty"`

	// CommissionRate - <nil>
	CommissionRate slapi.Float64 `json:"commissionRate,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ReferralAccountId - <nil>
	ReferralAccountId int `json:"referralAccountId,omitempty"`
}

func (softlayer_container_referral_partner_commission *SoftLayer_Container_Referral_Partner_Commission) String() string {
	return "SoftLayer_Container_Referral_Partner_Commission"
}
