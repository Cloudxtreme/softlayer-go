package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Brand - The SoftLayer_Brand data type contains brand information relating to the single
// SoftLayer customer account. SoftLayer customers are unable to change their brand information in the
// portal or the
type SoftLayer_Brand struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AllOwnedAccountCount - no documentation
	AllOwnedAccountCount uint64 `json:"allOwnedAccountCount"`

	// AllOwnedAccounts - no documentation
	AllOwnedAccounts []*SoftLayer_Account `json:"allOwnedAccounts"`

	// AllowAccountCreationFlag - This flag indicates if creation of accounts is allowed.
	AllowAccountCreationFlag bool `json:"allowAccountCreationFlag"`

	// Catalog - no documentation
	Catalog *SoftLayer_Product_Catalog `json:"catalog"`

	// CatalogId - no documentation
	CatalogId int `json:"catalogId"`

	// ContactCount - no documentation
	ContactCount uint64 `json:"contactCount"`

	// Contacts - no documentation
	Contacts []*SoftLayer_Brand_Contact `json:"contacts"`

	// Distributor - <nil>
	Distributor *SoftLayer_Brand `json:"distributor"`

	// DistributorChildFlag - <nil>
	DistributorChildFlag bool `json:"distributorChildFlag"`

	// DistributorFlag - <nil>
	DistributorFlag string `json:"distributorFlag"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount"`

	// HasAgentSupportFlag - <nil>
	HasAgentSupportFlag bool `json:"hasAgentSupportFlag"`

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - no documentation
	KeyName string `json:"keyName"`

	// LongName - no documentation
	LongName string `json:"longName"`

	// Name - no documentation
	Name string `json:"name"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount"`

	// OpenTickets - <nil>
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// OwnedAccountCount - no documentation
	OwnedAccountCount uint64 `json:"ownedAccountCount"`

	// OwnedAccounts - no documentation
	OwnedAccounts []*SoftLayer_Account `json:"ownedAccounts"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// TicketGroupCount - no documentation
	TicketGroupCount uint64 `json:"ticketGroupCount"`

	// TicketGroups - <nil>
	TicketGroups []*SoftLayer_Ticket_Group `json:"ticketGroups"`

	// Tickets - <nil>
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`

	// Users - <nil>
	Users []*SoftLayer_User_Customer `json:"users"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`
}

func (softlayer_brand *SoftLayer_Brand) String() string {
	return "SoftLayer_Brand"
}
