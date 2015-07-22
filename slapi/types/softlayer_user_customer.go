package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_User_Customer - The SoftLayer_User_Customer data type contains general information
// relating to a single SoftLayer customer portal user. Personal information in this type such as
// names, addresses, and phone numbers are not necessarily associated with the customer account the
// user is assigned to.
type SoftLayer_User_Customer struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - A portal user's associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// AdditionalEmailCount - A count of a portal user's additional email addresses. These email addresses
	// are contacted when updates are made to support tickets.
	AdditionalEmailCount uint64 `json:"additionalEmailCount"`

	// AdditionalEmails - A portal user's additional email addresses. These email addresses are contacted
	// when updates are made to support tickets.
	AdditionalEmails []*SoftLayer_User_Customer_AdditionalEmail `json:"additionalEmails"`

	// Address1 - The first line of the mailing address belonging to a portal user.
	Address1 string `json:"address1"`

	// Address2 - The second line of the mailing address belonging to a portal user.
	Address2 string `json:"address2"`

	// Aim - no documentation
	Aim string `json:"aim"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone"`

	// ApiAuthenticationKeyCount - A count of a portal user's API Authentication keys. There is a max limit
	// of two API keys per user.
	ApiAuthenticationKeyCount uint64 `json:"apiAuthenticationKeyCount"`

	// ApiAuthenticationKeys - A portal user's API Authentication keys. There is a max limit of two API
	// keys per user.
	ApiAuthenticationKeys []*SoftLayer_User_Customer_ApiAuthentication `json:"apiAuthenticationKeys"`

	// AuthenticationToken - The authentication token used for logging into the SoftLayer customer portal.
	AuthenticationToken *SoftLayer_Container_User_Authentication_Token `json:"authenticationToken"`

	// CdnAccountCount - A count of the CDN accounts associated with a portal user.
	CdnAccountCount uint64 `json:"cdnAccountCount"`

	// CdnAccounts - no documentation
	CdnAccounts []*SoftLayer_Network_ContentDelivery_Account `json:"cdnAccounts"`

	// ChildUserCount - A count of a portal user's child users. Some portal users may not have child users.
	ChildUserCount uint64 `json:"childUserCount"`

	// ChildUsers - A portal user's child users. Some portal users may not have child users.
	ChildUsers []*SoftLayer_User_Customer `json:"childUsers"`

	// City - The city of the mailing address belonging to a portal user.
	City string `json:"city"`

	// ClosedTicketCount - no documentation
	ClosedTicketCount uint64 `json:"closedTicketCount"`

	// ClosedTickets - no documentation
	ClosedTickets []*SoftLayer_Ticket `json:"closedTickets"`

	// CompanyName - A portal user's associated company. This may not be the same company as the customer
	// that owns this portal user.
	CompanyName string `json:"companyName"`

	// Country - A two-letter abbreviation of the country in the mailing address belonging to a portal
	// user.
	Country string `json:"country"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DaylightSavingsTimeFlag - Whether a portal user's time zone is affected by Daylight Savings Time.
	DaylightSavingsTimeFlag bool `json:"daylightSavingsTimeFlag"`

	// DenyAllResourceAccessOnCreateFlag - Flag used to deny access to all hardware and cloud computing
	// instances upon user creation.
	DenyAllResourceAccessOnCreateFlag bool `json:"denyAllResourceAccessOnCreateFlag"`

	// DisplayName - <nil>
	DisplayName string `json:"displayName"`

	// Email - no documentation
	Email string `json:"email"`

	// ExternalBindingCount - A count of the external authentication bindings that link an external
	// identifier to a SoftLayer user.
	ExternalBindingCount uint64 `json:"externalBindingCount"`

	// ExternalBindings - The external authentication bindings that link an external identifier to a
	// SoftLayer user.
	ExternalBindings []*SoftLayer_User_External_Binding `json:"externalBindings"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// ForumPasswordHash - A user's password for the SoftLayer forums, hashed for auto-login capability
	// from the SoftLayer customer portal
	ForumPasswordHash string `json:"forumPasswordHash"`

	// Hardware - A portal user's accessible hardware. These permissions control which hardware a user has
	// access to in the SoftLayer customer portal.
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of a portal user's accessible hardware. These permissions control which
	// hardware a user has access to in the SoftLayer customer portal.
	HardwareCount uint64 `json:"hardwareCount"`

	// HardwareNotificationCount - A count of hardware notifications associated with this user. A hardware
	// notification links a user to a piece of hardware, and that user will be notified if any monitors on
	// that hardware fail, if the monitors have a status of 'Notify User'.
	HardwareNotificationCount uint64 `json:"hardwareNotificationCount"`

	// HardwareNotifications - Hardware notifications associated with this user. A hardware notification
	// links a user to a piece of hardware, and that user will be notified if any monitors on that hardware
	// fail, if the monitors have a status of 'Notify User'.
	HardwareNotifications []*SoftLayer_User_Customer_Notification_Hardware `json:"hardwareNotifications"`

	// HasAcknowledgedSupportPolicyFlag - Whether or not a user has acknowledged the support policy.
	HasAcknowledgedSupportPolicyFlag bool `json:"hasAcknowledgedSupportPolicyFlag"`

	// HasFullHardwareAccessFlag - Whether or not a portal user has access to all hardware on their
	// account.
	HasFullHardwareAccessFlag bool `json:"hasFullHardwareAccessFlag"`

	// HasFullVirtualGuestAccessFlag - Whether or not a portal user has access to all hardware on their
	// account.
	HasFullVirtualGuestAccessFlag bool `json:"hasFullVirtualGuestAccessFlag"`

	// Icq - no documentation
	Icq string `json:"icq"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddressRestriction - The IP addresses or IP ranges from which a user may login to the SoftLayer
	// customer portal. Specify subnets in format and separate multiple addresses and subnets by commas.
	// You may combine IPv4 and IPv6 addresses and subnets, for example: 192.168.0.0/16,fe80:021b::0/64.
	IpAddressRestriction string `json:"ipAddressRestriction"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// LayoutProfileCount - no documentation
	LayoutProfileCount uint64 `json:"layoutProfileCount"`

	// LayoutProfiles - <nil>
	LayoutProfiles []*SoftLayer_Layout_Profile `json:"layoutProfiles"`

	// Locale - A user's locale. Locale holds user's language and region information.
	Locale *SoftLayer_Locale `json:"locale"`

	// LocaleId - A portal user's associated [[SoftLayer_Locale|locale]] id.
	LocaleId int `json:"localeId"`

	// LoginAttemptCount - A count of a user's attempts to log into the SoftLayer customer portal.
	LoginAttemptCount uint64 `json:"loginAttemptCount"`

	// LoginAttempts - A user's attempts to log into the SoftLayer customer portal.
	LoginAttempts []*SoftLayer_User_Customer_Access_Authentication `json:"loginAttempts"`

	// MobileDeviceCount - A count of a portal user's associated mobile device profiles.
	MobileDeviceCount uint64 `json:"mobileDeviceCount"`

	// MobileDevices - no documentation
	MobileDevices []*SoftLayer_User_Customer_MobileDevice `json:"mobileDevices"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Msn - no documentation
	Msn string `json:"msn"`

	// NotificationSubscriberCount - A count of notification subscription records for the user.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount"`

	// NotificationSubscribers - no documentation
	NotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"notificationSubscribers"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount"`

	// OpenTickets - no documentation
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// OverrideCount - no documentation
	OverrideCount uint64 `json:"overrideCount"`

	// Overrides - no documentation
	Overrides []*SoftLayer_Network_Service_Vpn_Overrides `json:"overrides"`

	// Parent - A portal user's parent user. If a SoftLayer_User_Customer has a null parentId property then
	// it doesn't have a parent user.
	Parent *SoftLayer_User_Customer `json:"parent"`

	// ParentId - A portal user's parent user. Id a users parentId is ''null'' then it doesn't have a
	// parent user in the customer portal.
	ParentId int `json:"parentId"`

	// PasswordExpireDate - no documentation
	PasswordExpireDate *time.Time `json:"passwordExpireDate"`

	// PermissionCount - A count of a portal user's permissions. These permissions control that user's
	// access to functions within the SoftLayer customer portal and
	PermissionCount uint64 `json:"permissionCount"`

	// PermissionSystemVersion - <nil>
	PermissionSystemVersion int `json:"permissionSystemVersion"`

	// Permissions - A portal user's permissions. These permissions control that user's access to functions
	// within the SoftLayer customer portal and
	Permissions []*SoftLayer_User_Customer_CustomerPermission_Permission `json:"permissions"`

	// PostalCode - The postal code of the mailing address belonging to an portal user.
	PostalCode string `json:"postalCode"`

	// PptpVpnAllowedFlag - Whether a portal user may connect to the SoftLayer private network via VPN or
	// not.
	PptpVpnAllowedFlag bool `json:"pptpVpnAllowedFlag"`

	// PreferenceCount - no documentation
	PreferenceCount uint64 `json:"preferenceCount"`

	// Preferences - <nil>
	Preferences []*SoftLayer_User_Preference `json:"preferences"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`

	// SalesforceUserLink - <nil>
	SalesforceUserLink *SoftLayer_User_Customer_Link `json:"salesforceUserLink"`

	// SavedId - <nil>
	SavedId string `json:"savedId"`

	// SecondaryLoginManagementFlag - Whether a user may change their security options (IP restriction,
	// password expiration, or enforce security questions on login) which were pre-selected by their
	// account's master user.
	SecondaryLoginManagementFlag bool `json:"secondaryLoginManagementFlag"`

	// SecondaryLoginRequiredFlag - Whether a user is required to answer a security question when logging
	// into the SoftLayer customer portal.
	SecondaryLoginRequiredFlag bool `json:"secondaryLoginRequiredFlag"`

	// SecondaryPasswordModifyDate - no documentation
	SecondaryPasswordModifyDate *time.Time `json:"secondaryPasswordModifyDate"`

	// SecondaryPasswordTimeoutDays - The number of days for which a user's password is active.
	SecondaryPasswordTimeoutDays int `json:"secondaryPasswordTimeoutDays"`

	// SecurityAnswerCount - A count of a portal user's security question answers. Some portal users may
	// not have security answers or may not be configured to require answering a security question on
	// login.
	SecurityAnswerCount uint64 `json:"securityAnswerCount"`

	// SecurityAnswers - A portal user's security question answers. Some portal users may not have security
	// answers or may not be configured to require answering a security question on login.
	SecurityAnswers []*SoftLayer_User_Customer_Security_Answer `json:"securityAnswers"`

	// Sms - A phone number that can receive SMS text messages for this portal user.
	Sms string `json:"sms"`

	// SslVpnAllowedFlag - Whether a portal user may connect to the SoftLayer private network via SSL VPN
	// or not.
	SslVpnAllowedFlag bool `json:"sslVpnAllowedFlag"`

	// State - A two-letter abbreviation of the state in the mailing address belonging to a portal user. If
	// a user does not reside in a province then this is typically blank.
	State string `json:"state"`

	// StatusDate - The date a portal users record's last status change.
	StatusDate *time.Time `json:"statusDate"`

	// SubscriberCount - A count of a user's notification subscription records.
	SubscriberCount uint64 `json:"subscriberCount"`

	// Subscribers - no documentation
	Subscribers []*SoftLayer_Notification_User_Subscriber `json:"subscribers"`

	// SuccessfulLoginCount - A count of a user's successful attempts to log into the SoftLayer customer
	// portal.
	SuccessfulLoginCount uint64 `json:"successfulLoginCount"`

	// SuccessfulLogins - A user's successful attempts to log into the SoftLayer customer portal.
	SuccessfulLogins []*SoftLayer_User_Customer_Access_Authentication `json:"successfulLogins"`

	// SupportPolicyAcknowledgementRequiredFlag - Whether or not a user is required to acknowledge the
	// support policy for portal access.
	SupportPolicyAcknowledgementRequiredFlag int `json:"supportPolicyAcknowledgementRequiredFlag"`

	// SurveyCount - A count of the surveys that a user has taken in the SoftLayer customer portal.
	SurveyCount uint64 `json:"surveyCount"`

	// SurveyRequiredFlag - Whether or not a user must take a brief survey the next time they log into the
	// SoftLayer customer portal.
	SurveyRequiredFlag bool `json:"surveyRequiredFlag"`

	// Surveys - The surveys that a user has taken in the SoftLayer customer portal.
	Surveys []*SoftLayer_Survey `json:"surveys"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// Tickets - no documentation
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// Timezone - no documentation
	Timezone *SoftLayer_Locale_Timezone `json:"timezone"`

	// TimezoneId - no documentation
	TimezoneId int `json:"timezoneId"`

	// UnsuccessfulLoginCount - A count of a user's unsuccessful attempts to log into the SoftLayer
	// customer portal.
	UnsuccessfulLoginCount uint64 `json:"unsuccessfulLoginCount"`

	// UnsuccessfulLogins - A user's unsuccessful attempts to log into the SoftLayer customer portal.
	UnsuccessfulLogins []*SoftLayer_User_Customer_Access_Authentication `json:"unsuccessfulLogins"`

	// UserLinkCount - no documentation
	UserLinkCount uint64 `json:"userLinkCount"`

	// UserLinks - <nil>
	UserLinks []*SoftLayer_User_Customer_Link `json:"userLinks"`

	// UserStatus - A portal user's status, which controls overall access to the SoftLayer customer portal
	// and VPN access to the private network.
	UserStatus *SoftLayer_User_Customer_Status `json:"userStatus"`

	// UserStatusId - no documentation
	UserStatusId int `json:"userStatusId"`

	// Username - no documentation
	Username string `json:"username"`

	// VirtualGuestCount - A count of a portal user's accessible CloudLayer Computing Instances. These
	// permissions control which CloudLayer Computing Instances a user has access to in the SoftLayer
	// customer portal.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - A portal user's accessible CloudLayer Computing Instances. These permissions control
	// which CloudLayer Computing Instances a user has access to in the SoftLayer customer portal.
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// VpnManualConfig - Whether a portal user vpn subnets have been manual configured.
	VpnManualConfig bool `json:"vpnManualConfig"`

	// Yahoo - no documentation
	Yahoo string `json:"yahoo"`
}

// AcknowledgeSupportPolicy - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) AcknowledgeSupportPolicy(ctx *slapi.RequestContext) error {
	return nil
}

// AddApiAuthenticationKey - Create a user's API authentication key, allowing that user access to query
// the SoftLayer addApiAuthenticationKey() returns the users new API key. Each portal user is allowed a
// maximum of two API keys.
func (softlayer_user_customer *SoftLayer_User_Customer) AddApiAuthenticationKey(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// AddBulkHardwareAccess - Add multiple hardware to a portal user's hardware access list. A user's
// hardware access list controls which of an account's hardware objects a user has access to in the
// SoftLayer customer portal and Hardware does not exist in the SoftLayer portal and returns "not
// found" exceptions in the API if the user doesn't have access to it. addBulkHardwareAccess() does not
// attempt to add hardware access if the given user already has access to that hardware object. Users
// can assign hardware access to their child users, but not to themselves. An account's master has
// access to all hardware on their customer account and can set hardware access for any of the other
// users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) AddBulkHardwareAccess(ctx *slapi.RequestContext, hardwareIds []int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddBulkPortalPermission - Add multiple permissions to a portal user's permission set.
// [[Permissions]] control which features in the SoftLayer customer portal and API a user may use.
// addBulkPortalPermission() does not attempt to add permissions already assigned to the user. Users
// can assign permissions to their child users, but not to themselves. An account's master has all
// portal permissions and can set permissions for any of the other users on their account. Use the
// [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list
// of all permissions available in the SoftLayer customer portal and Permissions are removed based on
// the keyName property of the permission objects within the permissions parameter.
func (softlayer_user_customer *SoftLayer_User_Customer) AddBulkPortalPermission(ctx *slapi.RequestContext, permissions []SoftLayer_User_Customer_CustomerPermission_Permission) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddBulkRoles - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) AddBulkRoles(ctx *slapi.RequestContext, roles []SoftLayer_User_Permission_Role) error {
	return nil
}

// AddBulkVirtualGuestAccess - Add multiple CloudLayer Computing Instances to a portal user's access
// list. A user's CloudLayer Computing Instance access list controls which of an account's CloudLayer
// Computing Instance objects a user has access to in the SoftLayer customer portal and CloudLayer
// Computing Instances do not exist in the SoftLayer portal and returns "not found" exceptions in the
// API if the user doesn't have access to it. addBulkVirtualGuestAccess() does not attempt to add
// CloudLayer Computing Instance access if the given user already has access to that CloudLayer
// Computing Instance object. Users can assign CloudLayer Computing Instance access to their child
// users, but not to themselves. An account's master has access to all CloudLayer Computing Instances
// on their customer account and can set CloudLayer Computing Instance access for any of the other
// users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) AddBulkVirtualGuestAccess(ctx *slapi.RequestContext, virtualGuestIds []int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddExternalBinding - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) AddExternalBinding(ctx *slapi.RequestContext, externalBinding SoftLayer_User_External_Binding) (*SoftLayer_User_Customer_External_Binding, error) {
	var returnValue *SoftLayer_User_Customer_External_Binding
	return returnValue, nil
}

// AddHardwareAccess - Add hardware to a portal user's hardware access list. A user's hardware access
// list controls which of an account's hardware objects a user has access to in the SoftLayer customer
// portal and Hardware does not exist in the SoftLayer portal and returns "not found" exceptions in the
// API if the user doesn't have access to it. If a user already has access to the hardware you're
// attempting to add then addHardwareAccess() returns true. Users can assign hardware access to their
// child users, but not to themselves. An account's master has access to all hardware on their customer
// account and can set hardware access for any of the other users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) AddHardwareAccess(ctx *slapi.RequestContext, hardwareId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddNotificationSubscriber - Create a notification subscription record for the user. If a
// subscription record exists for the notification, the record will be set to active, if currently
// inactive.
func (softlayer_user_customer *SoftLayer_User_Customer) AddNotificationSubscriber(ctx *slapi.RequestContext, notificationKeyName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddPortalPermission - Add a permission to a portal user's permission set. [[Permissions]] control
// which features in the SoftLayer customer portal and API a user may use. If the user already has the
// permission you're attempting to add then addPortalPermission() returns true. Users can assign
// permissions to their child users, but not to themselves. An account's master has all portal
// permissions and can set permissions for any of the other users on their account. Use the
// [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list
// of all permissions available in the SoftLayer customer portal and Permissions are added based on the
// keyName property of the permission parameter.
func (softlayer_user_customer *SoftLayer_User_Customer) AddPortalPermission(ctx *slapi.RequestContext, permission SoftLayer_User_Customer_CustomerPermission_Permission) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddRole - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) AddRole(ctx *slapi.RequestContext, role SoftLayer_User_Permission_Role) error {
	return nil
}

// AddVirtualGuestAccess - Add a CloudLayer Computing Instance to a portal user's access list. A user's
// CloudLayer Computing Instance access list controls which of an account's CloudLayer Computing
// Instance objects a user has access to in the SoftLayer customer portal and CloudLayer Computing
// Instances do not exist in the SoftLayer portal and returns "not found" exceptions in the API if the
// user doesn't have access to it. If a user already has access to the CloudLayer Computing Instance
// you're attempting to add then addVirtualGuestAccess() returns true. Users can assign CloudLayer
// Computing Instance access to their child users, but not to themselves. An account's master has
// access to all CloudLayer Computing Instances on their customer account and can set CloudLayer
// Computing Instance access for any of the other users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) AddVirtualGuestAccess(ctx *slapi.RequestContext, virtualGuestId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CheckExternalAuthenticationStatus - This service checks the result of a previously requested
// external authentication. [[SoftLayer_Container_User_Customer_External_Binding_Phone|Phone external
// binding]] container can be used for this service. Make sure to set the
// [[SoftLayer_Container_User_Customer_External_Binding_Phone::authenticationToken|authenticationToken]]
// that is generated by [[SoftLayer_User_Customer|initiateExternalAuthentication]] service.
func (softlayer_user_customer *SoftLayer_User_Customer) CheckExternalAuthenticationStatus(ctx *slapi.RequestContext, authenticationContainer SoftLayer_Container_User_Customer_External_Binding) (*SoftLayer_Container_User_Customer_Portal_Token, error) {
	var returnValue *SoftLayer_Container_User_Customer_Portal_Token
	return returnValue, nil
}

// CreateNotificationSubscriber - no documentation
func (softlayer_user_customer *SoftLayer_User_Customer) CreateNotificationSubscriber(ctx *slapi.RequestContext, keyName string, resourceTableId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - Create a new user in the SoftLayer customer portal. createObject() creates a user's
// portal record and adds them into the SoftLayer community forums. It is no longer possible to set up
// the SSL or enable flag in this call since the manage permissions have not yet been set. You will
// need to make a subsequent call to edit object in order to enable VPN access. An account's master
// user and sub-users who have the User Manage permission can add new users. createObject() creates
// users with a default permission set. After adding a user it may be helpful to set their permissions
// and hardware access.
func (softlayer_user_customer *SoftLayer_User_Customer) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_User_Customer, password string, vpnPassword string) (*SoftLayer_User_Customer, error) {
	var returnValue *SoftLayer_User_Customer
	return returnValue, nil
}

// CreateSubscriberDeliveryMethods - Create delivery methods for a notification that the user is
// subscribed to. Multiple delivery method keyNames can be supplied to create multiple delivery methods
// for the specified notification. Available delivery methods - Available notifications -
func (softlayer_user_customer *SoftLayer_User_Customer) CreateSubscriberDeliveryMethods(ctx *slapi.RequestContext, notificationKeyName string, deliveryMethodKeyNames []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeactivateNotificationSubscriber - no documentation
func (softlayer_user_customer *SoftLayer_User_Customer) DeactivateNotificationSubscriber(ctx *slapi.RequestContext, keyName string, resourceTableId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Account master users and sub-users who have the User Manage permission in the SoftLayer
// customer portal can update other user's information. Use editObject() if you wish to edit a single
// user account. Users who do not have the User Manage permission can only update their own
// information.
func (softlayer_user_customer *SoftLayer_User_Customer) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_User_Customer) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Account master users and sub-users who have the User Manage permission in the
// SoftLayer customer portal can update other user's information. Use editObjects() if you wish to edit
// multiple users at once. Users who do not have the User Manage permission can only update their own
// information.
func (softlayer_user_customer *SoftLayer_User_Customer) EditObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_User_Customer) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// FindUserPreference - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) FindUserPreference(ctx *slapi.RequestContext, profileName string, containerKeyname string, preferenceKeyname string) ([]*SoftLayer_Layout_Profile, error) {
	var returnValue []*SoftLayer_Layout_Profile
	return returnValue, nil
}

// GetActiveExternalAuthenticationVendors - The getActiveExternalAuthenticationVendors method will
// return a list of available external vendors that a SoftLayer user can authenticate against. The list
// will only contain vendors for which the user has at least one active external binding.
func (softlayer_user_customer *SoftLayer_User_Customer) GetActiveExternalAuthenticationVendors(ctx *slapi.RequestContext) ([]*SoftLayer_Container_User_Customer_External_Binding_Vendor, error) {
	var returnValue []*SoftLayer_Container_User_Customer_External_Binding_Vendor
	return returnValue, nil
}

// GetAuthenticationToken - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetAuthenticationToken(ctx *slapi.RequestContext, token SoftLayer_Container_User_Authentication_Token) (*SoftLayer_Container_User_Authentication_Token, error) {
	var returnValue *SoftLayer_Container_User_Authentication_Token
	return returnValue, nil
}

// GetDefaultSecurityQuestions - Retrieve default security questions. The default security questions
// will be used during the password reset process if a user has not set up their own security questions
// and answers.
func (softlayer_user_customer *SoftLayer_User_Customer) GetDefaultSecurityQuestions(ctx *slapi.RequestContext, key string) ([]*SoftLayer_User_Security_Question, error) {
	var returnValue []*SoftLayer_User_Security_Question
	return returnValue, nil
}

// GetHardwareCount - Retrieve the number of servers that a portal user has access to. Portal users can
// have restrictions set to limit services for and to perform actions on hardware. You can set these
// permissions in the portal by clicking the "administrative" then "user admin" links.
func (softlayer_user_customer *SoftLayer_User_Customer) GetHardwareCount(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetImpersonationToken - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetImpersonationToken(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_User_Customer object whose ID number corresponds to
// the ID number of the init parameter passed to the SoftLayer_User_Customer service. You can only
// retrieve users that are assigned to the customer account belonging to the user making the API call.
func (softlayer_user_customer *SoftLayer_User_Customer) GetObject(ctx *slapi.RequestContext) (*SoftLayer_User_Customer, error) {
	var returnValue *SoftLayer_User_Customer
	return returnValue, nil
}

// GetPortalLoginToken - Attempt to authenticate a username and password to the SoftLayer customer
// portal. Many portal user accounts are configured to require answering a security question on login.
// In this case getPortalLoginToken() also verifies the given security question ID and answer. If
// authentication is successful then the API returns a token containing the ID of the authenticated
// user and a hash key used by the SoftLayer customer portal to maintain authentication.
func (softlayer_user_customer *SoftLayer_User_Customer) GetPortalLoginToken(ctx *slapi.RequestContext, username string, password string, securityQuestionId int, securityQuestionAnswer string) (*SoftLayer_Container_User_Customer_Portal_Token, error) {
	var returnValue *SoftLayer_Container_User_Customer_Portal_Token
	return returnValue, nil
}

// GetSupportPolicyDocument - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetSupportPolicyDocument(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetSupportPolicyName - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetSupportPolicyName(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetSupportedLocales - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetSupportedLocales(ctx *slapi.RequestContext) ([]*SoftLayer_Locale, error) {
	var returnValue []*SoftLayer_Locale
	return returnValue, nil
}

// GetUserFromLostPasswordRequest - Retrieve a user object using a password recovery key received in an
// email generated by the [[SoftLayer_User_Customer::lostPassword|lostPassword]] method. The SoftLayer
// customer portal uses getUserFromLostPasswordRequest() to retrieve user security questions. Password
// recovery keys are valid for 24 hours after they're generated.
func (softlayer_user_customer *SoftLayer_User_Customer) GetUserFromLostPasswordRequest(ctx *slapi.RequestContext, key string) ([]*SoftLayer_User_Security_Question, error) {
	var returnValue []*SoftLayer_User_Security_Question
	return returnValue, nil
}

// GetUserPreferences - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) GetUserPreferences(ctx *slapi.RequestContext, profileName string, containerKeyname string) ([]*SoftLayer_Layout_Profile, error) {
	var returnValue []*SoftLayer_Layout_Profile
	return returnValue, nil
}

// GetVirtualGuestCount - Retrieve the number of CloudLayer Computing Instances that a portal user has
// access to. Portal users can have restrictions set to limit services for and to perform actions on
// CloudLayer Computing Instances. You can set these permissions in the portal by clicking the
// "administrative" then "user admin" links.
func (softlayer_user_customer *SoftLayer_User_Customer) GetVirtualGuestCount(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// InTerminalStatus - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) InTerminalStatus(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// InitiateExternalAuthentication - The service initiates an external authentication with the given
// external authentication vendor. The authentication container and its content will be verified before
// an attempt is made to initiate an external authentication.
// [[SoftLayer_Container_User_Customer_External_Binding_Phone|Phone external binding]] container can be
// used for this service. This service returns a unique authentication request token. You can use
// [[SoftLayer_User_Customer::checkExternalAuthenticationStatus|checkExternalAuthenticationStatus]]
// service to check if the authentication request is complete or not.
func (softlayer_user_customer *SoftLayer_User_Customer) InitiateExternalAuthentication(ctx *slapi.RequestContext, authenticationContainer SoftLayer_Container_User_Customer_External_Binding) (string, error) {
	var returnValue string
	return returnValue, nil
}

// IsMasterUser - Portal users are considered master users if they don't have an associated parent
// user. The only users who don't have parent users are users whose username matches their SoftLayer
// account name. Master users have special permissions throughout the SoftLayer customer portal.
func (softlayer_user_customer *SoftLayer_User_Customer) IsMasterUser(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsValidForumPassword - Determine if a string is the given user's login password to the SoftLayer
// community forums.
func (softlayer_user_customer *SoftLayer_User_Customer) IsValidForumPassword(ctx *slapi.RequestContext, password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// IsValidPortalPassword - Determine if a string is the given user's login password to the SoftLayer
// customer portal.
func (softlayer_user_customer *SoftLayer_User_Customer) IsValidPortalPassword(ctx *slapi.RequestContext, password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// LostPassword - SoftLayer provides a way for users of it's customer portal to recover lost passwords.
// The lostPassword() method is the first step in this process. Given a valid username and email
// address, the SoftLayer API will email the address provided with a URL to visit to begin the password
// recovery process. The last part of this URL is a hash key that's used as an identifier throughout
// this process. Use this hash key in the
// [[SoftLayer_User_Customer::setPasswordFromLostPasswordRequest|setPasswordFromLostPasswordRequest]]
// method to reset a user's password. Password recovery hash keys are valid for 24 hours after they're
// generated.
func (softlayer_user_customer *SoftLayer_User_Customer) LostPassword(ctx *slapi.RequestContext, username string, email string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PerformExternalAuthentication - The perform external authentication method will authenticate the
// given external authentication container with an external vendor. The authentication container and
// its contents will be verified before an attempt is made to authenticate the contents of the
// container with an external vendor.
func (softlayer_user_customer *SoftLayer_User_Customer) PerformExternalAuthentication(ctx *slapi.RequestContext, authenticationContainer SoftLayer_Container_User_Customer_External_Binding) (*SoftLayer_Container_User_Customer_Portal_Token, error) {
	var returnValue *SoftLayer_Container_User_Customer_Portal_Token
	return returnValue, nil
}

// RemoveAllHardwareAccessForThisUser - Remove all hardware from a portal user's hardware access list.
// A user's hardware access list controls which of an account's hardware objects a user has access to
// in the SoftLayer customer portal and If the current user does not have administrative privileges
// over this user, an inadequate permissions exception will get thrown. Users can call this function on
// child users, but not to themselves. An account's master has access to all users permissions on their
// account.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveAllHardwareAccessForThisUser(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveAllVirtualAccessForThisUser - Remove all cloud computing instances from a portal user's
// instance access list. A user's instance access list controls which of an account's computing
// instance objects a user has access to in the SoftLayer customer portal and If the current user does
// not have administrative privileges over this user, an inadequate permissions exception will get
// thrown. Users can call this function on child users, but not to themselves. An account's master has
// access to all users permissions on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveAllVirtualAccessForThisUser(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveApiAuthenticationKey - Remove a user's API authentication key, removing that user's access to
// query the SoftLayer
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveApiAuthenticationKey(ctx *slapi.RequestContext, keyId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveBulkHardwareAccess - Remove multiple hardware from a portal user's hardware access list. A
// user's hardware access list controls which of an account's hardware objects a user has access to in
// the SoftLayer customer portal and Hardware does not exist in the SoftLayer portal and returns "not
// found" exceptions in the API if the user doesn't have access to it. If a user does not has access to
// the hardware you're attempting remove add then removeBulkHardwareAccess() returns true. Users can
// assign hardware access to their child users, but not to themselves. An account's master has access
// to all hardware on their customer account and can set hardware access for any of the other users on
// their account. If the user has full hardware access, then it will provide access to but passed in"
// hardware ids.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveBulkHardwareAccess(ctx *slapi.RequestContext, hardwareIds []int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveBulkPortalPermission - Remove multiple permissions from a portal user's permission set.
// [[Permissions]] control which features in the SoftLayer customer portal and API a user may use.
// Removing a user's permission will affect that user's portal and API access. removePortalPermission()
// does not attempt to remove permissions that are not assigned to the user. Users can assign
// permissions to their child users, but not to themselves. An account's master has all portal
// permissions and can set permissions for any of the other users on their account. Use the
// [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list
// of all permissions available in the SoftLayer customer portal and Permissions are removed based on
// the keyName property of the permission objects within the permissions parameter.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveBulkPortalPermission(ctx *slapi.RequestContext, permissions []SoftLayer_User_Customer_CustomerPermission_Permission) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveBulkRoles - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveBulkRoles(ctx *slapi.RequestContext, roles []SoftLayer_User_Permission_Role) error {
	return nil
}

// RemoveBulkVirtualGuestAccess - Remove multiple CloudLayer Computing Instances from a portal user's
// access list. A user's CloudLayer Computing Instance access list controls which of an account's
// CloudLayer Computing Instance objects a user has access to in the SoftLayer customer portal and
// CloudLayer Computing Instances do not exist in the SoftLayer portal and returns "not found"
// exceptions in the API if the user doesn't have access to it. If a user does not has access to the
// CloudLayer Computing Instance you're attempting remove add then removeBulkVirtualGuestAccess()
// returns true. Users can assign CloudLayer Computing Instance access to their child users, but not to
// themselves. An account's master has access to all CloudLayer Computing Instances on their customer
// account and can set hardware access for any of the other users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveBulkVirtualGuestAccess(ctx *slapi.RequestContext, virtualGuestIds []int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveExternalBinding - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveExternalBinding(ctx *slapi.RequestContext, externalBinding SoftLayer_User_External_Binding) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveHardwareAccess - Remove hardware from a portal user's hardware access list. A user's hardware
// access list controls which of an account's hardware objects a user has access to in the SoftLayer
// customer portal and Hardware does not exist in the SoftLayer portal and returns "not found"
// exceptions in the API if the user doesn't have access to it. If a user does not has access to the
// hardware you're attempting remove add then removeHardwareAccess() returns true. Users can assign
// hardware access to their child users, but not to themselves. An account's master has access to all
// hardware on their customer account and can set hardware access for any of the other users on their
// account.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveHardwareAccess(ctx *slapi.RequestContext, hardwareId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemovePortalPermission - Remove a permission from a portal user's permission set. [[Permissions]]
// control which features in the SoftLayer customer portal and API a user may use. Removing a user's
// permission will affect that user's portal and API access. If the user does not have the permission
// you're attempting to remove then removePortalPermission() returns true. Users can assign permissions
// to their child users, but not to themselves. An account's master has all portal permissions and can
// set permissions for any of the other users on their account. Use the
// [[SoftLayer_User_Customer_CustomerPermission_Permission::getAllObjects]] method to retrieve a list
// of all permissions available in the SoftLayer customer portal and Permissions are removed based on
// the keyName property of the permission parameter.
func (softlayer_user_customer *SoftLayer_User_Customer) RemovePortalPermission(ctx *slapi.RequestContext, permission SoftLayer_User_Customer_CustomerPermission_Permission) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveRole - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveRole(ctx *slapi.RequestContext, role SoftLayer_User_Permission_Role) error {
	return nil
}

// RemoveVirtualGuestAccess - Remove a CloudLayer Computing Instance from a portal user's access list.
// A user's CloudLayer Computing Instance access list controls which of an account's computing
// instances a user has access to in the SoftLayer customer portal and CloudLayer Computing Instances
// do not exist in the SoftLayer portal and returns "not found" exceptions in the API if the user
// doesn't have access to it. If a user does not has access to the CloudLayer Computing Instance you're
// attempting remove add then removeVirtualGuestAccess() returns true. Users can assign CloudLayer
// Computing Instance access to their child users, but not to themselves. An account's master has
// access to all CloudLayer Computing Instances on their customer account and can set instance access
// for any of the other users on their account.
func (softlayer_user_customer *SoftLayer_User_Customer) RemoveVirtualGuestAccess(ctx *slapi.RequestContext, virtualGuestId int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ResetExpiredPassword - Attempt to authenticate a username and password to the SoftLayer customer
// portal and reset there password. If authentication and password reset is successful then the API
// returns true.
func (softlayer_user_customer *SoftLayer_User_Customer) ResetExpiredPassword(ctx *slapi.RequestContext, username string, password string, newPassword string, securityQuestionId int, securityQuestionAnswer string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SamlAuthenticate - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) SamlAuthenticate(ctx *slapi.RequestContext, accountId string, samlResponse string) (*SoftLayer_Container_User_Customer_Portal_Token, error) {
	var returnValue *SoftLayer_Container_User_Customer_Portal_Token
	return returnValue, nil
}

// SamlBeginAuthentication - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) SamlBeginAuthentication(ctx *slapi.RequestContext, accountId int) (string, error) {
	var returnValue string
	return returnValue, nil
}

// SamlBeginLogout - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) SamlBeginLogout(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// SamlLogout - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) SamlLogout(ctx *slapi.RequestContext, samlResponse string) error {
	return nil
}

// SetPasswordFromLostPasswordRequest - Set a user's password via the lost password recovery system,
// using a password recovery key received in an email generated by the
// [[SoftLayer_User_Customer::lostPassword|lostPassword]] method. Password recovery keys are valid for
// 24 hours after they're generated. User portal passwords must match the following restrictions.
// Portal passwords must... * ...be over eight characters long. * ...be under twenty characters long. *
// ...contain at least one uppercase letter * ...contain at least one lowercase letter * ...contain at
// least one number * ...contain one of the special characters _ - | @ . , ? / ! ~ # $ % ^ & * ( ) { }
// [ ] \ + = * ...not match your username * ...not match your forum password
func (softlayer_user_customer *SoftLayer_User_Customer) SetPasswordFromLostPasswordRequest(ctx *slapi.RequestContext, key string, password string, securityAnswers []SoftLayer_User_Customer_Security_Answer) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateForumPassword - Update a user's password on the SoftLayer community forums. As with portal
// passwords, user forum passwords must match the following restrictions. Forum passwords must... *
// ...be over eight characters long. * ...be under twenty characters long. * ...contain at least one
// uppercase letter * ...contain at least one lowercase letter * ...contain at least one number *
// ...contain one of the special characters _ - | @ . , ? / ! ~ # $ % ^ & * ( ) { } [ ] \ + = * ...not
// match your username * ...not match your portal password Finally, users can only update their own
// password.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateForumPassword(ctx *slapi.RequestContext, password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateNotificationSubscriber - Update the active status for a notification that the user is
// subscribed to. A notification along with an active flag can be supplied to update the active status
// for a particular notification subscription.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateNotificationSubscriber(ctx *slapi.RequestContext, notificationKeyName string, active int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdatePassword - Update a user's password on the SoftLayer customer portal. As with forum passwords,
// user portal passwords must match the following restrictions. Portal passwords must... * ...be over
// eight characters long. * ...be under twenty characters long. * ...contain at least one uppercase
// letter * ...contain at least one lowercase letter * ...contain at least one number * ...contain one
// of the special characters _ - | @ . , ? / ! ~ # $ % ^ & * ( ) { } [ ] \ + = * ...not match your
// username * ...not match your forum password Finally, users can only update their own password. An
// account's master user can update any of their account users' passwords.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdatePassword(ctx *slapi.RequestContext, password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateSecurityAnswers - Update a user's login security questions and answers on the SoftLayer
// customer portal. These questions and answers are used to optionally log into the SoftLayer customer
// portal using two-factor authentication. Each user must have three distinct questions set with a
// unique answer for each question, and each answer may only contain alphanumeric or the . , - _ ( ) [
// ] : ; > < characters. Existing user security questions and answers are deleted before new ones are
// set, and users may only update their own security questions and answers.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateSecurityAnswers(ctx *slapi.RequestContext, questions []SoftLayer_User_Security_Question, answers []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateSubscriberDeliveryMethod - Update a delivery method for a notification that the user is
// subscribed to. A delivery method keyName along with an active flag can be supplied to update the
// active status of the delivery methods for the specified notification. Available delivery methods -
// Available notifications -
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateSubscriberDeliveryMethod(ctx *slapi.RequestContext, notificationKeyName string, deliveryMethodKeyNames []string, active int) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateVpnPassword - Update a user's VPN password on the SoftLayer customer portal. As with portal
// passwords, VPN passwords must match the following restrictions. VPN passwords must... * ...be over
// eight characters long. * ...be under twenty characters long. * ...contain at least one uppercase
// letter * ...contain at least one lowercase letter * ...contain at least one number * ...contain one
// of the special characters _ - | @ . , ? / ! ~ # $ % ^ & * ( ) { } [ ] \ = * ...not match your
// username * ...not match your forum password Finally, users can only update their own VPN password.
// An account's master user can update any of their account users' VPN passwords.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateVpnPassword(ctx *slapi.RequestContext, password string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateVpnUser - Always call this function to enable changes when manually configuring VPN subnet
// access.
func (softlayer_user_customer *SoftLayer_User_Customer) UpdateVpnUser(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ValidateAuthenticationToken - <nil>
func (softlayer_user_customer *SoftLayer_User_Customer) ValidateAuthenticationToken(ctx *slapi.RequestContext, authenticationToken SoftLayer_Container_User_Authentication_Token) (*SoftLayer_Container_User_Customer_Portal_Token, error) {
	var returnValue *SoftLayer_Container_User_Customer_Portal_Token
	return returnValue, nil
}
