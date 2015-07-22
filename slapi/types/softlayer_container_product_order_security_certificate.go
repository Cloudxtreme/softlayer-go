package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Security_Certificate - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype contains everything
// required to place a secure certificate order with SoftLayer.
type SoftLayer_Container_Product_Order_Security_Certificate struct {

	// AdministrativeContact - The administrator contact associated with a SSL certificate. If the contact
	// is not provided the technical contact will be used. If the address is not provided the organization
	// information address will be used.
	AdministrativeContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"administrativeContact"`

	// BillingContact - The billing contact associated with a SSL certificate. If the contact is not
	// provided the technical contact will be used. If the address is not provided the organization
	// information address will be used.
	BillingContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"billingContact"`

	// CertificateSigningRequest - The base64 encoded string that sent from an applicant to a certificate
	// authority. The CSR contains information identifying the applicant and the public key chosen by the
	// applicant. The corresponding private key should not be included.
	CertificateSigningRequest string `json:"certificateSigningRequest"`

	// OrderApproverEmailAddress - The email address that can approve a secure certificate order.
	OrderApproverEmailAddress string `json:"orderApproverEmailAddress"`

	// OrganizationInformation - The organization information associated with a SSL certificate.
	OrganizationInformation *SoftLayer_Container_Product_Order_Attribute_Organization `json:"organizationInformation"`

	// RenewalFlag - Indicates if it is an renewal order of an existing SSL certificate.
	RenewalFlag bool `json:"renewalFlag"`

	// ServerCount - no documentation
	ServerCount int `json:"serverCount"`

	// ServerType - The server type. This is the name from a
	// [[SoftLayer_Security_Certificate_Request_ServerType]] object.
	ServerType string `json:"serverType"`

	// TechnicalContact - The technical contact associated with a SSL certificate. If the address is not
	// provided the organization information address will be used.
	TechnicalContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"technicalContact"`

	// ValidityMonths - The period that a SSL certificate is valid for. For example, 12, 24, 36
	ValidityMonths int `json:"validityMonths"`
}
