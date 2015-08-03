package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_User_Customer_External_Binding_Phone - The SoftLayer_User_Customer_External_Binding_Phone
// data type contains information about an external binding that uses a phone call, SMS or mobile app
// for 2 form factor authentication. The external binding information is used when a SoftLayer customer
// logs into the SoftLayer customer portal or VPN to authenticate them against a trusted 3rd party, in
// this case using a mobile phone, mobile phone application or land-line phone. SoftLayer users with an
// active external binding will be prohibited from using the API for security reasons.
type SoftLayer_User_Customer_External_Binding_Phone struct {

	// BindingStatus - The current external binding status. It can be or
	BindingStatus string `json:"bindingStatus"`

	// PinLength - <nil>
	PinLength string `json:"pinLength"`
}

func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) String() string {
	return "SoftLayer_User_Customer_External_Binding_Phone"
}

// CheckPhoneValidationResult - no documentation
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) CheckPhoneValidationResult(ctx *slapi.RequestContext, token string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Disable - Disabling an external binding will allow you to keep the external binding on your
// SoftLayer account, but will not require you to authentication with our trusted 2 form factor vendor
// when logging into the SoftLayer customer portal. You may supply one of the following reason when you
// disable an external binding: *Unspecified *TemporarilyUnavailable *Lost *Stolen
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) Disable(ctx *slapi.RequestContext, reason string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// Enable - Enabling an external binding will activate the binding on your account and require you to
// authenticate with our trusted 3rd party 2 form factor vendor when logging into the SoftLayer
// customer portal. Please note that API access will be disabled for users that have an active external
// binding.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) Enable(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllAuthenticationModes - This service returns key names of all available authentication modes.
// See [[SoftLayer_Container_User_Customer_External_Binding_Phone_Mode|authentication mode]] container
// for details.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetAllAuthenticationModes(ctx *slapi.RequestContext) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAllAuthenticationPinModes - This service returns key names of all available authentication modes.
// Refer to
// [[SoftLayer_User_Customer_External_Binding_Phone::getAllAuthenticationModes|getAllAuthenticationModes]]
// to retrieve authentication mode key names.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetAllAuthenticationPinModes(ctx *slapi.RequestContext, authenticationModeKeyName string) ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetAuthenticationMode - <nil>
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetAuthenticationMode(ctx *slapi.RequestContext) (*SoftLayer_Container_User_Customer_External_Binding_Phone_Mode, error) {
	var returnValue *SoftLayer_Container_User_Customer_External_Binding_Phone_Mode
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer_External_Binding_Phone, error) {
	var returnValue *SoftLayer_User_Customer_External_Binding_Phone
	return returnValue, nil
}

// GetPhoneAppActivationCode - Some vendor's mobile app requires an activation code. Use this method to
// get an activation data.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetPhoneAppActivationCode(ctx *slapi.RequestContext) ([]*SoftLayer_User_External_Binding_Attribute, error) {
	var returnValue []*SoftLayer_User_External_Binding_Attribute
	return returnValue, nil
}

// GetPhoneData - <nil>
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) GetPhoneData(ctx *slapi.RequestContext) ([]*SoftLayer_Container_User_Data_Phone, error) {
	var returnValue []*SoftLayer_Container_User_Data_Phone
	return returnValue, nil
}

// RequestPhoneValidation - Initiates a phone validation requests and returns a unique token. Use
// [[SoftLayer_User_Customer_External_Binding_Phone::checkPhoneValidationResult|checkPhoneValidationResult]]
// to find the phone validation result.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) RequestPhoneValidation(ctx *slapi.RequestContext, phoneData SoftLayer_Container_User_Data_Phone) (string, error) {
	var returnValue string
	return returnValue, nil
}

// UpdateAuthenticationMode - This service allow you to change your phone authentication mode. See
// [[SoftLayer_Container_User_Customer_External_Binding_Phone_Mode|authentication mode]] container for
// available modes.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) UpdateAuthenticationMode(ctx *slapi.RequestContext, mode SoftLayer_Container_User_Customer_External_Binding_Phone_Mode) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdatePhone - Phone external binding supports a primary and a backup phone number. You can use this
// method to update your phone number used for the phone authentication. You can provide an array of
// [[SoftLayer_Container_User_Data_Phone|User Phone]] objects. You have to mark one as the primary
// phone number by setting "phoneType" to *countryCode: Country code number for the phone number.
// Default: 1 (United States & Canada +1) *phone: Phone number that 2 Form Factor system will call or
// text for user authentication. The phone number format must match the format selected in the Country
// Code. *extension: Specify the extension that will be dialed after the call is answered. Digits,
// commas, *, and # are allowed. Commas can be used for a one second pause to navigate phone system
// menus. *phoneType: Specify the primary and backup phone number by setting this value to or If
// omitted, it will be considered to be the primary phone number. If you are passing two Phone objects,
// you must specify the phone type of each phone number.
func (softlayer_user_customer_external_binding_phone *SoftLayer_User_Customer_External_Binding_Phone) UpdatePhone(ctx *slapi.RequestContext, phoneData []SoftLayer_Container_User_Data_Phone) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
