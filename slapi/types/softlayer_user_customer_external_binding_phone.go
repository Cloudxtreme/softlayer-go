package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Customer_External_Binding_Phone - The SoftLayer_User_Customer_External_Binding_Phone
// data type contains information about an external binding that uses a phone call, SMS or mobile app
// for 2 form factor authentication. The external binding information is used when a SoftLayer customer
// logs into the SoftLayer customer portal or VPN to authenticate them against a trusted 3rd party, in
// this case using a mobile phone, mobile phone application or land-line phone. SoftLayer users with an
// active external binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Phone struct {
}

// SoftLayer_User_Customer_External_Binding_Phone_Extended is SoftLayer_User_Customer_External_Binding_Phone with all maskable types.
type SoftLayer_User_Customer_External_Binding_Phone_Extended struct {
	SoftLayer_User_Customer_External_Binding_Phone

	// BindingStatus - The current external binding status. It can be or
	BindingStatus string `json:"bindingStatus"`

	// PinLength - <nil>
	PinLength string `json:"pinLength"`
}

func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) String() string {
	return "SoftLayer_User_Customer_External_Binding_Phone"
}
