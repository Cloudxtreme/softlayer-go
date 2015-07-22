package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Password_Type - Every username and password combination associated with a
// SoftLayer customer account belongs to a service that SoftLayer provides. The relationship between a
// username/password and it's service is provided by the SoftLayer_Account_Password_Type data type.
// Each username/password belongs to a single service type.
type SoftLayer_Account_Password_Type struct {

	// Description - A description of the use for the account username/password combination.
	Description string `json:"description"`
}
