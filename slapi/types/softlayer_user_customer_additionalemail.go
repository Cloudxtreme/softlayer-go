package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_AdditionalEmail - The SoftLayer_User_Customer_AdditionalEmail data type
// contains the additional email for use in ticket update notifications.
type SoftLayer_User_Customer_AdditionalEmail struct {

	// Email - Email assigned to user for use in ticket update notifications.
	Email string `json:"email"`

	// UserId - An internal identifier for the portal user who this additional email belongs to.
	UserId int `json:"userId"`
}

// SoftLayer_User_Customer_AdditionalEmail_Extended is SoftLayer_User_Customer_AdditionalEmail with all maskable types.
type SoftLayer_User_Customer_AdditionalEmail_Extended struct {
	SoftLayer_User_Customer_AdditionalEmail

	// User - The portal user that owns this additional email address.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_user_customer_additionalemail *SoftLayer_User_Customer_AdditionalEmail) String() string {
	return "SoftLayer_User_Customer_AdditionalEmail"
}
