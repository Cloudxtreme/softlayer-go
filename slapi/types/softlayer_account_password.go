package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Password - The SoftLayer_Account_Password contains username, passwords and notes
// for services that may require for external applications such the Webcc interface for the EVault
// Storage service.
type SoftLayer_Account_Password struct {

	// TypeId - An identifier relating to a username/password combinations's associated service.
	TypeId int `json:"typeId"`

	// Username - The username portion of a username/password combination.
	Username string `json:"username"`

	// AccountId - The SoftLayer customer account id that a username/password combination is associated
	// with.
	AccountId int `json:"accountId"`

	// Id - A username/password combination's internal identifier.
	Id int `json:"id"`

	// Notes - A simple description of a username/password combination. These notes don't affect portal
	// functionality.
	Notes string `json:"notes"`

	// Password - The password portion of a username/password combination.
	Password string `json:"password"`
}

func (softlayer_account_password *SoftLayer_Account_Password) String() string {
	return "SoftLayer_Account_Password"
}

// SoftLayer_Account_Password_Extended is SoftLayer_Account_Password with all maskable types.
type SoftLayer_Account_Password_Extended struct {
	SoftLayer_Account_Password

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Type - The service that an account/password combination is tied to.
	Type *SoftLayer_Account_Password_Type `json:"type"`
}

func (softlayer_account_password *SoftLayer_Account_Password_Extended) String() string {
	return "SoftLayer_Account_Password"
}
