package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand - The SoftLayer_Brand data type contains brand information relating to the single
// SoftLayer customer account. SoftLayer customers are unable to change their brand information in the
// portal or the
type SoftLayer_Brand struct {

	// Name - no documentation
	Name string `json:"name,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// LongName - no documentation
	LongName string `json:"longName,omitempty"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId,omitempty"`
}

func (softlayer_brand *SoftLayer_Brand) String() string {
	return "SoftLayer_Brand"
}

// SoftLayer_Brand_Extended is SoftLayer_Brand with all maskable types.
type SoftLayer_Brand_Extended struct {
	SoftLayer_Brand

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog,omitempty"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// Tickets - <nil>
	Tickets []*SoftLayer_Ticket `json:"tickets,omitempty"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// OwnedAccounts - no documentation
	OwnedAccounts []*SoftLayer_Account `json:"ownedAccounts,omitempty"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// Distributor - <nil>
	Distributor *SoftLayer_Brand `json:"distributor,omitempty"`

	// AllOwnedAccountCount - no documentation
	AllOwnedAccountCount uint64 `json:"allOwnedAccountCount,omitempty"`

	// OwnedAccountCount - no documentation
	OwnedAccountCount uint64 `json:"ownedAccountCount,omitempty"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Brand_Contact `json:"contacts,omitempty"`

	// HasAgentSupportFlag - <nil>
	HasAgentSupportFlag bool `json:"hasAgentSupportFlag,omitempty"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount,omitempty"`

	// TicketGroupCount - no documentation
	TicketGroupCount uint64 `json:"ticketGroupCount,omitempty"`

	// OpenTickets - <nil>
	OpenTickets []*SoftLayer_Ticket `json:"openTickets,omitempty"`

	// TicketGroups - <nil>
	TicketGroups []*SoftLayer_Ticket_Group `json:"ticketGroups,omitempty"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount,omitempty"`

	// AllOwnedAccounts - no documentation
	AllOwnedAccounts []*SoftLayer_Account `json:"allOwnedAccounts,omitempty"`

	// AllowAccountCreationFlag - This flag indicates if creation of accounts is allowed.
	AllowAccountCreationFlag bool `json:"allowAccountCreationFlag,omitempty"`

	// DistributorChildFlag - <nil>
	DistributorChildFlag bool `json:"distributorChildFlag,omitempty"`

	// DistributorFlag - <nil>
	DistributorFlag string `json:"distributorFlag,omitempty"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount,omitempty"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`
}

func (softlayer_brand *SoftLayer_Brand_Extended) String() string {
	return "SoftLayer_Brand"
}
