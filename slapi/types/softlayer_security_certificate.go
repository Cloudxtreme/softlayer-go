package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Certificate - <nil>
type SoftLayer_Security_Certificate struct {

	// AssociatedServiceCount - The number of services currently associated with the certificate.
	AssociatedServiceCount int `json:"associatedServiceCount"`

	// Certificate - The certificate provided publicly to clients requesting identity credentials. This
	// certificate is usually signed by a source trusted by the client or a signature chain can be
	// established between this certificate and the truested certificate. This property may only be
	// modified when no services are associated. See associatedServiceCount.
	Certificate string `json:"certificate"`

	// CertificateSigningRequest - The signing request used to request a certificate authority generate a
	// signed certificate. This property may only be modified when no services are associated. See
	// associatedServiceCount.
	CertificateSigningRequest string `json:"certificateSigningRequest"`

	// CommonName - The common name (usually a domain name) encoded within the certificate. This property
	// is read only. Changes made will be silently ignored.
	CommonName string `json:"commonName"`

	// CreateDate - The date the certificate _record_ was created. The contents of the certificate may of
	// changed since the record was created, so this does not represent anything about the certificate
	// itself. This property is read only. Changes made will be silently ignored.
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// IntermediateCertificate - The intermediate certificate authorities certificate that completes the
	// certificate chain for the issued certificate. Required when clients will only trust the root
	// certificate. This property may only be modified when no services are associated. See
	// associatedServiceCount.
	IntermediateCertificate string `json:"intermediateCertificate"`

	// KeySize - The size (number of bits) of the public key represented by the certificate.
	KeySize int `json:"keySize"`

	// LoadBalancerVirtualIpAddressCount - A count of the load balancers virtual IP addresses currently
	// associated with the certificate.
	LoadBalancerVirtualIpAddressCount uint64 `json:"loadBalancerVirtualIpAddressCount"`

	// LoadBalancerVirtualIpAddresses - The load balancers virtual IP addresses currently associated with
	// the certificate.
	LoadBalancerVirtualIpAddresses []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"loadBalancerVirtualIpAddresses"`

	// ModifyDate - The date the certificate _record_ was last modified.The contents of the certificate may
	// of changed since the record was created, so this does not represent anything about the certificate
	// itself. This property is read only. Changes made will be silently ignored.
	ModifyDate *time.Time `json:"modifyDate"`

	// Notes - no documentation
	Notes string `json:"notes"`

	// OrganizationName - The organizational name encoded in the certificate. This property is read only.
	// Changes made will be silently ignored.
	OrganizationName string `json:"organizationName"`

	// PrivateKey - The private key in the key/certificate pair. This property may only be modified when no
	// services are associated. See associatedServiceCount.
	PrivateKey string `json:"privateKey"`

	// ValidityBegin - The UTC timestamp representing the beginning of the certificate's validity This
	// property is read only. Changes made will be silently ignored.
	ValidityBegin *time.Time `json:"validityBegin"`

	// ValidityDays - The number of days remaining in the validity period for the certificate. This
	// property is read only. Changes made will be silently ignored.
	ValidityDays int `json:"validityDays"`

	// ValidityEnd - The UTC timestamp representing the end of the certificate's validity period. This
	// property is read only. Changes made will be silently ignored.
	ValidityEnd *time.Time `json:"validityEnd"`
}

func (softlayer_security_certificate *SoftLayer_Security_Certificate) String() string {
	return "SoftLayer_Security_Certificate"
}
