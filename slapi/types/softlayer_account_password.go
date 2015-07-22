package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Password - The SoftLayer_Account_Password contains username, passwords and notes
// for services that may require for external applications such the Webcc interface for the EVault
// Storage service.
type SoftLayer_Account_Password struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

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

	// Type - The service that an account/password combination is tied to.
	Type *SoftLayer_Account_Password_Type `json:"type"`

	// TypeId - An identifier relating to a username/password combinations's associated service.
	TypeId int `json:"typeId"`

	// Username - The username portion of a username/password combination.
	Username string `json:"username"`
}

// EditObject - The password and/or notes may be modified. Modifying the EVault passwords here will
// also update the password the Webcc interface will use.
func (softlayer_account_password *SoftLayer_Account_Password) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Account_Password) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Account_Password object whose ID corresponds to the ID
// number of the init parameter passed to the SoftLayer_Account_Password service.
func (softlayer_account_password *SoftLayer_Account_Password) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account_Password, error) {
	var returnValue *SoftLayer_Account_Password
	return returnValue, nil
}
