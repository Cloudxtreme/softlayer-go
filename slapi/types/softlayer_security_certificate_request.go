package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Security_Certificate_Request - SoftLayer_Security_Certificate_Request data type is used to
// harness your SSL certificate order to a Certificate Authority. This contains data that is required
// by a Certificate Authority to place an SSL certificate order.
type SoftLayer_Security_Certificate_Request struct {

	// Account - The account to which a SSL certificate request belongs.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// ApproverEmailAddress - The email address of a person who will approve your SSL certificate order.
	// This is usually an email address of your domain administrator.
	ApproverEmailAddress string `json:"approverEmailAddress"`

	// CertificateAuthorityName - no documentation
	CertificateAuthorityName string `json:"certificateAuthorityName"`

	// CertificateSigningRequest - no documentation
	CertificateSigningRequest string `json:"certificateSigningRequest"`

	// CommonName - no documentation
	CommonName string `json:"commonName"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// EffectiveDate - no documentation
	EffectiveDate *time.Time `json:"effectiveDate"`

	// ExpirationDate - no documentation
	ExpirationDate *time.Time `json:"expirationDate"`

	// Id - The internal identifier of an SSL certificate request
	Id int `json:"id"`

	// ModifyDate - The date a SSL certificate request was last modified.
	ModifyDate *time.Time `json:"modifyDate"`

	// Order - The order contains the information related to a SSL certificate request.
	Order *SoftLayer_Billing_Order `json:"order"`

	// OrderItem - The associated order item for this SSL certificate request.
	OrderItem *SoftLayer_Billing_Order_Item `json:"orderItem"`

	// Status - no documentation
	Status *SoftLayer_Security_Certificate_Request_Status `json:"status"`

	// StatusId - A status id reflecting the state of a SSL certificate request
	StatusId int `json:"statusId"`

	// TechnicalContactEmailAddress - no documentation
	TechnicalContactEmailAddress string `json:"technicalContactEmailAddress"`
}

func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) String() string {
	return "SoftLayer_Security_Certificate_Request"
}

// CancelSslOrder - Cancels a pending SSL certificate order at Certificate Authority
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) CancelSslOrder(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAdministratorEmailDomains - Gets the email domains that can be used to validate a certificate to
// a domain.
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) GetAdministratorEmailDomains(ctx *slapi.RequestContext, commonName string) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAdministratorEmailPrefixes - Gets the email accounts that can be used to validate a certificate
// to a domain.
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) GetAdministratorEmailPrefixes(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Security_Certificate_Request, error) {
	var returnValue *SoftLayer_Security_Certificate_Request
	return returnValue, nil
}

// GetPreviousOrderData - Returns previous SSL certificate order data. You can use this data for to
// place a renewal order for a completed SSL certificate.
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) GetPreviousOrderData(ctx *slapi.RequestContext) (*SoftLayer_Container_Product_Order_Security_Certificate, error) {
	var returnValue *SoftLayer_Container_Product_Order_Security_Certificate
	return returnValue, nil
}

// GetSslCertificateRequests - no documentation
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) GetSslCertificateRequests(ctx *slapi.RequestContext, accountId int) ([]*SoftLayer_Security_Certificate_Request, error) {
	var returnValue []*SoftLayer_Security_Certificate_Request
	return returnValue, nil
}

// ResendEmail - A Certificate Authority sends out various emails to your domain administrator or your
// technical contact. Use this service to re-send these emails if you did not receive them initially.
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) ResendEmail(ctx *slapi.RequestContext, emailType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ValidateCsr - Allows you to validate a Certificate Signing Request required for an SSL certificate
// with the certificate authority This method sends the the length of the subscription in months, the
// certificate type, and the server type for validation against requirements of the CA. Returns true if
// valid. More information on CSR generation can be found at:
// [http://en.wikipedia.org/wiki/Certificate_signing_request Wikipedia]
// [https://knowledge.verisign.com/support/ssl-certificates-support/index?page=content&id=AR235&actp=LIST&viewlocale=en_US
// VeriSign]
func (softlayer_security_certificate_request *SoftLayer_Security_Certificate_Request) ValidateCsr(ctx *slapi.RequestContext, csr string, validityMonths int, itemId int, serverType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
