package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand - The SoftLayer_Brand data type contains brand information relating to the single
// SoftLayer customer account. SoftLayer customers are unable to change their brand information in the
// portal or the
type SoftLayer_Brand struct {

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// LongName - no documentation
	LongName string `json:"longName"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId"`

	// Id - <nil>
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_brand *SoftLayer_Brand) String() string {
	return "SoftLayer_Brand"
}

// SoftLayer_Brand_Extended is SoftLayer_Brand with all maskable types.
type SoftLayer_Brand_Extended struct {
	SoftLayer_Brand

	// OwnedAccountCount - no documentation
	OwnedAccountCount uint64 `json:"ownedAccountCount"`

	// TicketGroupCount - no documentation
	TicketGroupCount uint64 `json:"ticketGroupCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// DistributorChildFlag - <nil>
	DistributorChildFlag bool `json:"distributorChildFlag"`

	// OwnedAccounts - no documentation
	OwnedAccounts []*SoftLayer_Account `json:"ownedAccounts"`

	// AllOwnedAccounts - no documentation
	AllOwnedAccounts []*SoftLayer_Account `json:"allOwnedAccounts"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount"`

	// TicketGroups - <nil>
	TicketGroups []*SoftLayer_Ticket_Group `json:"ticketGroups"`

	// OpenTickets - <nil>
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// Tickets - <nil>
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount"`

	// Distributor - <nil>
	Distributor *SoftLayer_Brand `json:"distributor"`

	// HasAgentSupportFlag - <nil>
	HasAgentSupportFlag bool `json:"hasAgentSupportFlag"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`

	// DistributorFlag - <nil>
	DistributorFlag string `json:"distributorFlag"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Brand_Contact `json:"contacts"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// AllowAccountCreationFlag - This flag indicates if creation of accounts is allowed.
	AllowAccountCreationFlag bool `json:"allowAccountCreationFlag"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// AllOwnedAccountCount - no documentation
	AllOwnedAccountCount uint64 `json:"allOwnedAccountCount"`
}

func (softlayer_brand *SoftLayer_Brand_Extended) String() string {
	return "SoftLayer_Brand"
}
