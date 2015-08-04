package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_CustomerPermission_Permission - Each SoftLayer portal account is assigned a
// series of permissions that determine what access the user has to functions within the SoftLayer
// customer portal. This status is reflected in the SoftLayer_User_Customer_Status data type.
// Permissions differ from user status in that user status applies globally to the portal while user
// permissions are applied to specific portal functions.
type SoftLayer_User_Customer_CustomerPermission_Permission struct {

	// Key - no documentation
	Key string `json:"key,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_user_customer_customerpermission_permission *SoftLayer_User_Customer_CustomerPermission_Permission) String() string {
	return "SoftLayer_User_Customer_CustomerPermission_Permission"
}
