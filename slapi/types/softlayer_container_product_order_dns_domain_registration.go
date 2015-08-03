package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Dns_Domain_Registration - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. The
// SoftLayer_Container_Product_Order_Dns_Domain_Registration datatype contains everything required to
// place a domain registration order with SoftLayer.
type SoftLayer_Container_Product_Order_Dns_Domain_Registration struct {

	// TechnicalContact - Technical contact information associated with an registraton or transfer. This is
	// required if registration type is 'new' or 'transfer'.
	TechnicalContact *SoftLayer_Container_Dns_Domain_Registration_Contact `json:"technicalContact"`

	// AdministrativeContact - Administrative contact information associated with an registraton or
	// transfer. This is required if registration type is 'new' or 'transfer'.
	AdministrativeContact *SoftLayer_Container_Dns_Domain_Registration_Contact `json:"administrativeContact"`

	// BillingContact - Billing contact information associated with an registraton or transfer. This is
	// required if registration type is 'new' or 'transfer'.
	BillingContact *SoftLayer_Container_Dns_Domain_Registration_Contact `json:"billingContact"`

	// DomainRegistrationList - The list of domains to be registered. This is required if registration type
	// is 'new', 'renew', or 'transfer'.
	DomainRegistrationList []*SoftLayer_Container_Dns_Domain_Registration_List `json:"domainRegistrationList"`

	// OwnerContact - Owner contact information associated with an registraton or transfer. This is
	// required if registration type is 'new' or 'transfer'.
	OwnerContact *SoftLayer_Container_Dns_Domain_Registration_Contact `json:"ownerContact"`

	// RegistrationType - The type of a domain registration order. The registration type is Required.
	// Allowed values are new, transfer, and renew
	RegistrationType string `json:"registrationType"`
}

func (softlayer_container_product_order_dns_domain_registration *SoftLayer_Container_Product_Order_Dns_Domain_Registration) String() string {
	return "SoftLayer_Container_Product_Order_Dns_Domain_Registration"
}
