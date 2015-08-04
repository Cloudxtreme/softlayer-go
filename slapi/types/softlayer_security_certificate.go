package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Certificate - <nil>
type SoftLayer_Security_Certificate struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// IntermediateCertificate - The intermediate certificate authorities certificate that completes the
	// certificate chain for the issued certificate. Required when clients will only trust the root
	// certificate. This property may only be modified when no services are associated. See
	// associatedServiceCount.
	IntermediateCertificate string `json:"intermediateCertificate,omitempty"`

	// ValidityEnd - The UTC timestamp representing the end of the certificate's validity period. This
	// property is read only. Changes made will be silently ignored.
	ValidityEnd *time.Time `json:"validityEnd,omitempty"`

	// Certificate - The certificate provided publicly to clients requesting identity credentials. This
	// certificate is usually signed by a source trusted by the client or a signature chain can be
	// established between this certificate and the truested certificate. This property may only be
	// modified when no services are associated. See associatedServiceCount.
	Certificate string `json:"certificate,omitempty"`

	// CertificateSigningRequest - The signing request used to request a certificate authority generate a
	// signed certificate. This property may only be modified when no services are associated. See
	// associatedServiceCount.
	CertificateSigningRequest string `json:"certificateSigningRequest,omitempty"`

	// KeySize - The size (number of bits) of the public key represented by the certificate.
	KeySize int `json:"keySize,omitempty"`

	// ValidityBegin - The UTC timestamp representing the beginning of the certificate's validity This
	// property is read only. Changes made will be silently ignored.
	ValidityBegin *time.Time `json:"validityBegin,omitempty"`

	// CommonName - The common name (usually a domain name) encoded within the certificate. This property
	// is read only. Changes made will be silently ignored.
	CommonName string `json:"commonName,omitempty"`

	// ModifyDate - The date the certificate _record_ was last modified.The contents of the certificate may
	// of changed since the record was created, so this does not represent anything about the certificate
	// itself. This property is read only. Changes made will be silently ignored.
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// OrganizationName - The organizational name encoded in the certificate. This property is read only.
	// Changes made will be silently ignored.
	OrganizationName string `json:"organizationName,omitempty"`

	// ValidityDays - The number of days remaining in the validity period for the certificate. This
	// property is read only. Changes made will be silently ignored.
	ValidityDays int `json:"validityDays,omitempty"`

	// CreateDate - The date the certificate _record_ was created. The contents of the certificate may of
	// changed since the record was created, so this does not represent anything about the certificate
	// itself. This property is read only. Changes made will be silently ignored.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Notes - no documentation
	Notes string `json:"notes,omitempty"`

	// PrivateKey - The private key in the key/certificate pair. This property may only be modified when no
	// services are associated. See associatedServiceCount.
	PrivateKey string `json:"privateKey,omitempty"`
}

func (softlayer_security_certificate *SoftLayer_Security_Certificate) String() string {
	return "SoftLayer_Security_Certificate"
}

// SoftLayer_Security_Certificate_Extended is SoftLayer_Security_Certificate with all maskable types.
type SoftLayer_Security_Certificate_Extended struct {
	SoftLayer_Security_Certificate

	// LoadBalancerVirtualIpAddresses - The load balancers virtual IP addresses currently associated with
	// the certificate.
	LoadBalancerVirtualIpAddresses []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"loadBalancerVirtualIpAddresses,omitempty"`

	// LoadBalancerVirtualIpAddressCount - A count of the load balancers virtual IP addresses currently
	// associated with the certificate.
	LoadBalancerVirtualIpAddressCount uint64 `json:"loadBalancerVirtualIpAddressCount,omitempty"`

	// AssociatedServiceCount - The number of services currently associated with the certificate.
	AssociatedServiceCount int `json:"associatedServiceCount,omitempty"`
}

func (softlayer_security_certificate *SoftLayer_Security_Certificate_Extended) String() string {
	return "SoftLayer_Security_Certificate"
}
