package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_User_Customer_External_Binding_Phone_Mode - This container can be used to
// configure the phone authentication mode. By default, in mode with no Pin number will be used. With
// the default mode, you will have to answer a phone call from a trusted 2 form factor vendor during
// authentication process. You have to answer the call and follow the instruction in order to complete
// the authentication. You can also use SMS text message or PhoneFactor mobile app modes (in case
// you're using PhoneFactor). Additionally, you can set up a Pin number. By requiring you to verify
// your secret you can ensure that you have possession of your phone.
type SoftLayer_Container_User_Customer_External_Binding_Phone_Mode struct {

	// Mode - Authentication mode. Valid modes are: In this mode, users will receive a phone call to
	// authenticate. Using PIN can enhance the security of the phone authentication by requiring the user
	// to enter a PIN during the authentication call. Valid Pin modes are: (default) No PIN is used. 4 to
	// 10 digit numeric value The user's voice will be used to identify the user. SMS Text mode will send a
	// SMS text message to the user's phone to complete the authentication. There are 2 different pin
	// modes: (default) A text message containing a One-Time Passcode is sent to the user. The user must
	// reply to the text message entering this OTP to complete the authentication. This mode enhances the
	// security of the authentication by requiring the user to enter the OTP + their PIN in the text reply.
	// This mode is applicable for PhoneFactor. Phone App mode results in a notification being sent to the
	// user's PhoneFactor phone app. There are 2 different pin modes for the mobile app authentication.
	// (default) The first authentication is when the user signs on using a username and password. The
	// second authentication is when the user receives a notification in the PhoneFactor phone app. In
	// Standard Mode, users will prompted to authenticate, deny, or deny and report fraud. This mode
	// enhances the security of the authentication by requiring the user to enter their PIN in the phone
	// app.
	Mode string `json:"mode"`

	// Pin - no documentation
	Pin string `json:"pin"`

	// PinMode - Available Pin modes are: Default: (Pin is not used)
	PinMode string `json:"pinMode"`
}

func (softlayer_container_user_customer_external_binding_phone_mode *SoftLayer_Container_User_Customer_External_Binding_Phone_Mode) String() string {
	return "SoftLayer_Container_User_Customer_External_Binding_Phone_Mode"
}
