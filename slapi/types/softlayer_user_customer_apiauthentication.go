package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_ApiAuthentication - The SoftLayer_User_Customer_ApiAuthentication type
// contains user's authentication key(s).
type SoftLayer_User_Customer_ApiAuthentication struct {

	// TimestampKey - no documentation
	TimestampKey int `json:"timestampKey"`

	// UserId - no documentation
	UserId int `json:"userId"`

	// AuthenticationKey - no documentation
	AuthenticationKey string `json:"authenticationKey"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddressRestriction - The IP addresses or IP ranges from which this user may access the SoftLayer
	// Specify subnets in format and separate multiple addresses and subnets by commas. You may combine
	// IPv4 and IPv6 addresses and subnets, for example: 192.168.0.0/16,fe80:021b::0/64.
	IpAddressRestriction string `json:"ipAddressRestriction"`
}

// SoftLayer_User_Customer_ApiAuthentication_Extended is SoftLayer_User_Customer_ApiAuthentication with all maskable types.
type SoftLayer_User_Customer_ApiAuthentication_Extended struct {
	SoftLayer_User_Customer_ApiAuthentication

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_apiauthentication *SoftLayer_User_Customer_ApiAuthentication) String() string {
	return "SoftLayer_User_Customer_ApiAuthentication"
}
