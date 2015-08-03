package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Certificate_Request - SoftLayer_Security_Certificate_Request data type is used to
// harness your SSL certificate order to a Certificate Authority. This contains data that is required
// by a Certificate Authority to place an SSL certificate order.
type SoftLayer_Security_Certificate_Request struct {

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - The internal identifier of an SSL certificate request
	Id int `json:"id"`

	// ModifyDate - The date a SSL certificate request was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// TechnicalContactEmailAddress - no documentation
	TechnicalContactEmailAddress string `json:"technicalContactEmailAddress"`

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// ApproverEmailAddress - The email address of a person who will approve your SSL certificate order.
	// This is usually an email address of your domain administrator.
	ApproverEmailAddress string `json:"approverEmailAddress"`

	// CertificateSigningRequest - no documentation
	CertificateSigningRequest string `json:"certificateSigningRequest"`

	// CommonName - no documentation
	CommonName string `json:"commonName"`

	// EffectiveDate - no documentation
	EffectiveDate *time.Time `json:"effectiveDate"`

	// StatusId - A status id reflecting the state of a SSL certificate request
	StatusId int `json:"statusId"`
}

// SoftLayer_Security_Certificate_Request_Extended is SoftLayer_Security_Certificate_Request with all maskable types.
type SoftLayer_Security_Certificate_Request_Extended struct {
	SoftLayer_Security_Certificate_Request

	// CertificateAuthorityName - no documentation
	CertificateAuthorityName string `json:"certificateAuthorityName"`

	// Account - The account to which a SSL certificate request belongs.
	Account *SoftLayer_Account `json:"account"`

	// OrderItem - The associated order item for this SSL certificate request.
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem"`

	// Order - The order contains the information related to a SSL certificate request.
	Order *SoftLayer_Billing_Order `json:"order"`

	// Status - no documentation
	Status *SoftLayer_Security_Certificate_Request_Status `json:"status"`
}

func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) String() string {
	return "SoftLayer_Security_Certificate_Request"
}
