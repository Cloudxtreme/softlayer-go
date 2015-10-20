package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand - The SoftLayer_Brand data type contains brand information relating to the single
// SoftLayer customer account. SoftLayer customers are unable to change their brand information in the
// portal or the
type SoftLayer_Brand struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId,omitempty"`

	// LongName - no documentation
	LongName string `json:"longName,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// CustomerCountryLocationRestrictions - This references relationship between brands, locations and
	// countries associated with a user's account that are ineligible when ordering products. For example,
	// the India datacenter may not be available on this brand for customers that live in Great Britain.
	CustomerCountryLocationRestrictions []*SoftLayer_Brand_Restriction_Location_CustomerCountry `json:"customerCountryLocationRestrictions,omitempty"`

	// Tickets - <nil>
	Tickets []*SoftLayer_Ticket `json:"tickets,omitempty"`

	// OwnedAccountCount - no documentation
	OwnedAccountCount uint64 `json:"ownedAccountCount,omitempty"`

	// TicketGroupCount - no documentation
	TicketGroupCount uint64 `json:"ticketGroupCount,omitempty"`

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog,omitempty"`

	// OpenTickets - <nil>
	OpenTickets []*SoftLayer_Ticket `json:"openTickets,omitempty"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`

	// OwnedAccounts - no documentation
	OwnedAccounts []*SoftLayer_Account `json:"ownedAccounts,omitempty"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// DistributorChildFlag - <nil>
	DistributorChildFlag bool `json:"distributorChildFlag,omitempty"`

	// HasAgentSupportFlag - <nil>
	HasAgentSupportFlag bool `json:"hasAgentSupportFlag,omitempty"`

	// TicketGroups - <nil>
	TicketGroups []*SoftLayer_Ticket_Group `json:"ticketGroups,omitempty"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount,omitempty"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount,omitempty"`

	// AllowAccountCreationFlag - This flag indicates if creation of accounts is allowed.
	AllowAccountCreationFlag bool `json:"allowAccountCreationFlag,omitempty"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Brand_Contact `json:"contacts,omitempty"`

	// DistributorFlag - <nil>
	DistributorFlag string `json:"distributorFlag,omitempty"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount,omitempty"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount,omitempty"`

	// AllOwnedAccounts - no documentation
	AllOwnedAccounts []*SoftLayer_Account `json:"allOwnedAccounts,omitempty"`

	// Distributor - <nil>
	Distributor *SoftLayer_Brand `json:"distributor,omitempty"`

	// AllOwnedAccountCount - no documentation
	AllOwnedAccountCount uint64 `json:"allOwnedAccountCount,omitempty"`

	// CustomerCountryLocationRestrictionCount - A count of this references relationship between brands,
	// locations and countries associated with a user's account that are ineligible when ordering products.
	// For example, the India datacenter may not be available on this brand for customers that live in
	// Great Britain.
	CustomerCountryLocationRestrictionCount uint64 `json:"customerCountryLocationRestrictionCount,omitempty"`
}

func (softlayer_brand *SoftLayer_Brand) String() string {
	return "SoftLayer_Brand"
}
