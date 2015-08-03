package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand - The SoftLayer_Brand data type contains brand information relating to the single
// SoftLayer customer account. SoftLayer customers are unable to change their brand information in the
// portal or the
type SoftLayer_Brand struct {

	// Name - no documentation
	Name string `json:"name"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId"`

	// Id - <nil>
	Id int `json:"id"`

	// LongName - no documentation
	LongName string `json:"longName"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`
}

// SoftLayer_Brand_Extended is SoftLayer_Brand with all maskable types.
type SoftLayer_Brand_Extended struct {
	SoftLayer_Brand

	// OwnedAccounts - no documentation
	OwnedAccounts []*SoftLayer_Account `json:"ownedAccounts"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// Distributor - <nil>
	Distributor *SoftLayer_Brand `json:"distributor"`

	// Tickets - <nil>
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// TicketGroups - <nil>
	TicketGroups []*SoftLayer_Ticket_Group `json:"ticketGroups"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// TicketGroupCount - no documentation
	TicketGroupCount uint64 `json:"ticketGroupCount"`

	// AllowAccountCreationFlag - This flag indicates if creation of accounts is allowed.
	AllowAccountCreationFlag bool `json:"allowAccountCreationFlag"`

	// DistributorFlag - <nil>
	DistributorFlag string `json:"distributorFlag"`

	// OpenTickets - <nil>
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// AllOwnedAccountCount - no documentation
	AllOwnedAccountCount uint64 `json:"allOwnedAccountCount"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// HasAgentSupportFlag - <nil>
	HasAgentSupportFlag bool `json:"hasAgentSupportFlag"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// AllOwnedAccounts - no documentation
	AllOwnedAccounts []*SoftLayer_Account `json:"allOwnedAccounts"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Brand_Contact `json:"contacts"`

	// DistributorChildFlag - <nil>
	DistributorChildFlag bool `json:"distributorChildFlag"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// OwnedAccountCount - no documentation
	OwnedAccountCount uint64 `json:"ownedAccountCount"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`
}

func (softlayer_brand *SoftLayer_Brand) String() string {
	return "SoftLayer_Brand"
}
