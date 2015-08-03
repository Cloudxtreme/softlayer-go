package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Security_Certificate_Entry - <nil>
type SoftLayer_Security_Certificate_Entry struct {

	// CertificateId - no documentation
	CertificateId int `json:"certificateId"`

	// CommonName - The common name (usually a domain name) encoded within the certificate.
	CommonName string `json:"commonName"`

	// KeySize - The size (number of bits) of the public key represented by the certificate.
	KeySize int `json:"keySize"`

	// OrganizationName - The organizational name encoded in the certificate.
	OrganizationName string `json:"organizationName"`

	// ValidityBegin - The UTC timestamp representing the beginning of the certificate's validity
	ValidityBegin *time.Time `json:"validityBegin"`

	// ValidityDays - The number of days remaining in the validity period for the certificate.
	ValidityDays int `json:"validityDays"`

	// ValidityEnd - The UTC timestamp representing the end of the certificate's validity period.
	ValidityEnd *time.Time `json:"validityEnd"`
}

func (softlayer_security_certificate_entry *SoftLayer_Security_Certificate_Entry) String() string {
	return "SoftLayer_Security_Certificate_Entry"
}
