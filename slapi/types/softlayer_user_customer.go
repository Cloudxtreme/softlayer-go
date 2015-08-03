package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_User_Customer - The SoftLayer_User_Customer data type contains general information
// relating to a single SoftLayer customer portal user. Personal information in this type such as
// names, addresses, and phone numbers are not necessarily associated with the customer account the
// user is assigned to.
type SoftLayer_User_Customer struct {

	// DisplayName - <nil>
	DisplayName string `json:"displayName"`

	// LocaleId - A portal user's associated [[SoftLayer_Locale|locale]] id.
	LocaleId int `json:"localeId"`

	// PermissionSystemVersion - <nil>
	PermissionSystemVersion int `json:"permissionSystemVersion"`

	// CompanyName - A portal user's associated company. This may not be the same company as the customer
	// that owns this portal user.
	CompanyName string `json:"companyName"`

	// PostalCode - The postal code of the mailing address belonging to an portal user.
	PostalCode string `json:"postalCode"`

	// Msn - no documentation
	Msn string `json:"msn"`

	// Aim - no documentation
	Aim string `json:"aim"`

	// StatusDate - The date a portal users record's last status change.
	StatusDate *time.Time `json:"statusDate"`

	// AuthenticationToken - The authentication token used for logging into the SoftLayer customer portal.
	AuthenticationToken *SoftLayer_Container_User_Authentication_Token `json:"authenticationToken"`

	// SslVpnAllowedFlag - Whether a portal user may connect to the SoftLayer private network via SSL VPN
	// or not.
	SslVpnAllowedFlag bool `json:"sslVpnAllowedFlag"`

	// Email - no documentation
	Email string `json:"email"`

	// Icq - no documentation
	Icq string `json:"icq"`

	// LastName - no documentation
	LastName string `json:"lastName"`

	// UserStatusId - no documentation
	UserStatusId int `json:"userStatusId"`

	// Address1 - The first line of the mailing address belonging to a portal user.
	Address1 string `json:"address1"`

	// SecondaryPasswordTimeoutDays - The number of days for which a user's password is active.
	SecondaryPasswordTimeoutDays int `json:"secondaryPasswordTimeoutDays"`

	// PptpVpnAllowedFlag - Whether a portal user may connect to the SoftLayer private network via VPN or
	// not.
	PptpVpnAllowedFlag bool `json:"pptpVpnAllowedFlag"`

	// SecondaryPasswordModifyDate - no documentation
	SecondaryPasswordModifyDate *time.Time `json:"secondaryPasswordModifyDate"`

	// PasswordExpireDate - no documentation
	PasswordExpireDate *time.Time `json:"passwordExpireDate"`

	// FirstName - no documentation
	FirstName string `json:"firstName"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// ParentId - A portal user's parent user. Id a users parentId is ''null'' then it doesn't have a
	// parent user in the customer portal.
	ParentId int `json:"parentId"`

	// State - A two-letter abbreviation of the state in the mailing address belonging to a portal user. If
	// a user does not reside in a province then this is typically blank.
	State string `json:"state"`

	// ForumPasswordHash - A user's password for the SoftLayer forums, hashed for auto-login capability
	// from the SoftLayer customer portal
	ForumPasswordHash string `json:"forumPasswordHash"`

	// Yahoo - no documentation
	Yahoo string `json:"yahoo"`

	// Sms - A phone number that can receive SMS text messages for this portal user.
	Sms string `json:"sms"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DenyAllResourceAccessOnCreateFlag - Flag used to deny access to all hardware and cloud computing
	// instances upon user creation.
	DenyAllResourceAccessOnCreateFlag bool `json:"denyAllResourceAccessOnCreateFlag"`

	// SavedId - <nil>
	SavedId string `json:"savedId"`

	// Address2 - The second line of the mailing address belonging to a portal user.
	Address2 string `json:"address2"`

	// SecondaryLoginManagementFlag - Whether a user may change their security options (IP restriction,
	// password expiration, or enforce security questions on login) which were pre-selected by their
	// account's master user.
	SecondaryLoginManagementFlag bool `json:"secondaryLoginManagementFlag"`

	// AccountId - A portal user's associated [[SoftLayer_Account|customer account]] id.
	AccountId int `json:"accountId"`

	// Country - A two-letter abbreviation of the country in the mailing address belonging to a portal
	// user.
	Country string `json:"country"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone"`

	// TimezoneId - no documentation
	TimezoneId int `json:"timezoneId"`

	// IpAddressRestriction - The IP addresses or IP ranges from which a user may login to the SoftLayer
	// customer portal. Specify subnets in format and separate multiple addresses and subnets by commas.
	// You may combine IPv4 and IPv6 addresses and subnets, for example: 192.168.0.0/16,fe80:021b::0/64.
	IpAddressRestriction string `json:"ipAddressRestriction"`

	// Username - no documentation
	Username string `json:"username"`

	// SecondaryLoginRequiredFlag - Whether a user is required to answer a security question when logging
	// into the SoftLayer customer portal.
	SecondaryLoginRequiredFlag bool `json:"secondaryLoginRequiredFlag"`

	// VpnManualConfig - Whether a portal user vpn subnets have been manual configured.
	VpnManualConfig bool `json:"vpnManualConfig"`

	// City - The city of the mailing address belonging to a portal user.
	City string `json:"city"`

	// DaylightSavingsTimeFlag - Whether a portal user's time zone is affected by Daylight Savings Time.
	DaylightSavingsTimeFlag bool `json:"daylightSavingsTimeFlag"`
}

// SoftLayer_User_Customer_Extended is SoftLayer_User_Customer with all maskable types.
type SoftLayer_User_Customer_Extended struct {
	SoftLayer_User_Customer

	// AdditionalEmailCount - A count of a portal user's additional email addresses. These email addresses
	// are contacted when updates are made to support tickets.
	AdditionalEmailCount uint64 `json:"additionalEmailCount"`

	// LoginAttemptCount - A count of a user's attempts to log into the SoftLayer customer portal.
	LoginAttemptCount uint64 `json:"loginAttemptCount"`

	// OverrideCount - no documentation
	OverrideCount uint64 `json:"overrideCount"`

	// SecurityAnswerCount - A count of a portal user's security question answers. Some portal users may
	// not have security answers or may not be configured to require answering a security question on
	// login.
	SecurityAnswerCount uint64 `json:"securityAnswerCount"`

	// Timezone - no documentation
	Timezone *SoftLayer_Locale_Timezone `json:"timezone"`

	// ApiAuthenticationKeys - A portal user's API Authentication keys. There is a max limit of two API
	// keys per user.
	ApiAuthenticationKeys []*SoftLayer_User_Customer_ApiAuthentication `json:"apiAuthenticationKeys"`

	// UserLinkCount - no documentation
	UserLinkCount uint64 `json:"userLinkCount"`

	// SuccessfulLogins - A user's successful attempts to log into the SoftLayer customer portal.
	SuccessfulLogins []*SoftLayer_User_Customer_Access_Authentication `json:"successfulLogins"`

	// LoginAttempts - A user's attempts to log into the SoftLayer customer portal.
	LoginAttempts []*SoftLayer_User_Customer_Access_Authentication `json:"loginAttempts"`

	// Overrides - no documentation
	Overrides []*SoftLayer_Network_Service_Vpn_Overrides `json:"overrides"`

	// Permissions - A portal user's permissions. These permissions control that user's access to functions
	// within the SoftLayer customer portal and
	Permissions []*SoftLayer_User_Customer_CustomerPermission_Permission `json:"permissions"`

	// SubscriberCount - A count of a user's notification subscription records.
	SubscriberCount uint64 `json:"subscriberCount"`

	// Roles - <nil>
	Roles []*SoftLayer_User_Permission_Role `json:"roles"`

	// Subscribers - no documentation
	Subscribers []*SoftLayer_Notification_User_Subscriber `json:"subscribers"`

	// SurveyRequiredFlag - Whether or not a user must take a brief survey the next time they log into the
	// SoftLayer customer portal.
	SurveyRequiredFlag bool `json:"surveyRequiredFlag"`

	// Parent - A portal user's parent user. If a SoftLayer_User_Customer has a null parentId property then
	// it doesn't have a parent user.
	Parent *SoftLayer_User_Customer `json:"parent"`

	// UserStatus - A portal user's status, which controls overall access to the SoftLayer customer portal
	// and VPN access to the private network.
	UserStatus *SoftLayer_User_Customer_Status `json:"userStatus"`

	// LayoutProfileCount - no documentation
	LayoutProfileCount uint64 `json:"layoutProfileCount"`

	// VirtualGuestCount - A count of a portal user's accessible CloudLayer Computing Instances. These
	// permissions control which CloudLayer Computing Instances a user has access to in the SoftLayer
	// customer portal.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// CdnAccounts - no documentation
	CdnAccounts []*SoftLayer_Network_ContentDelivery_Account `json:"cdnAccounts"`

	// ClosedTickets - no documentation
	ClosedTickets []*SoftLayer_Ticket `json:"closedTickets"`

	// Hardware - A portal user's accessible hardware. These permissions control which hardware a user has
	// access to in the SoftLayer customer portal.
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// UnsuccessfulLogins - A user's unsuccessful attempts to log into the SoftLayer customer portal.
	UnsuccessfulLogins []*SoftLayer_User_Customer_Access_Authentication `json:"unsuccessfulLogins"`

	// VirtualGuests - A portal user's accessible CloudLayer Computing Instances. These permissions control
	// which CloudLayer Computing Instances a user has access to in the SoftLayer customer portal.
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// ApiAuthenticationKeyCount - A count of a portal user's API Authentication keys. There is a max limit
	// of two API keys per user.
	ApiAuthenticationKeyCount uint64 `json:"apiAuthenticationKeyCount"`

	// HardwareNotificationCount - A count of hardware notifications associated with this user. A hardware
	// notification links a user to a piece of hardware, and that user will be notified if any monitors on
	// that hardware fail, if the monitors have a status of 'Notify User'.
	HardwareNotificationCount uint64 `json:"hardwareNotificationCount"`

	// PermissionCount - A count of a portal user's permissions. These permissions control that user's
	// access to functions within the SoftLayer customer portal and
	PermissionCount uint64 `json:"permissionCount"`

	// HardwareNotifications - Hardware notifications associated with this user. A hardware notification
	// links a user to a piece of hardware, and that user will be notified if any monitors on that hardware
	// fail, if the monitors have a status of 'Notify User'.
	HardwareNotifications []*SoftLayer_User_Customer_Notification_Hardware `json:"hardwareNotifications"`

	// ExternalBindingCount - A count of the external authentication bindings that link an external
	// identifier to a SoftLayer user.
	ExternalBindingCount uint64 `json:"externalBindingCount"`

	// HardwareCount - A count of a portal user's accessible hardware. These permissions control which
	// hardware a user has access to in the SoftLayer customer portal.
	HardwareCount uint64 `json:"hardwareCount"`

	// PreferenceCount - no documentation
	PreferenceCount uint64 `json:"preferenceCount"`

	// AdditionalEmails - A portal user's additional email addresses. These email addresses are contacted
	// when updates are made to support tickets.
	AdditionalEmails []*SoftLayer_User_Customer_AdditionalEmail `json:"additionalEmails"`

	// HasFullHardwareAccessFlag - Whether or not a portal user has access to all hardware on their
	// account.
	HasFullHardwareAccessFlag bool `json:"hasFullHardwareAccessFlag"`

	// ClosedTicketCount - no documentation
	ClosedTicketCount uint64 `json:"closedTicketCount"`

	// OpenTicketCount - no documentation
	OpenTicketCount uint64 `json:"openTicketCount"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// NotificationSubscribers - no documentation
	NotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"notificationSubscribers"`

	// OpenTickets - no documentation
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// ChildUserCount - A count of a portal user's child users. Some portal users may not have child users.
	ChildUserCount uint64 `json:"childUserCount"`

	// ExternalBindings - The external authentication bindings that link an external identifier to a
	// SoftLayer user.
	ExternalBindings []*SoftLayer_User_External_Binding `json:"externalBindings"`

	// Surveys - The surveys that a user has taken in the SoftLayer customer portal.
	Surveys []*SoftLayer_Survey `json:"surveys"`

	// MobileDeviceCount - A count of a portal user's associated mobile device profiles.
	MobileDeviceCount uint64 `json:"mobileDeviceCount"`

	// SuccessfulLoginCount - A count of a user's successful attempts to log into the SoftLayer customer
	// portal.
	SuccessfulLoginCount uint64 `json:"successfulLoginCount"`

	// RoleCount - no documentation
	RoleCount uint64 `json:"roleCount"`

	// ActionCount - no documentation
	ActionCount uint64 `json:"actionCount"`

	// SalesforceUserLink - <nil>
	SalesforceUserLink *SoftLayer_User_Customer_Link `json:"salesforceUserLink"`

	// Preferences - <nil>
	Preferences []*SoftLayer_User_Preference `json:"preferences"`

	// Actions - <nil>
	Actions []*SoftLayer_User_Permission_Action `json:"actions"`

	// Locale - A user's locale. Locale holds user's language and region information.
	Locale *SoftLayer_Locale `json:"locale"`

	// MobileDevices - no documentation
	MobileDevices []*SoftLayer_User_Customer_MobileDevice `json:"mobileDevices"`

	// CdnAccountCount - A count of the CDN accounts associated with a portal user.
	CdnAccountCount uint64 `json:"cdnAccountCount"`

	// SurveyCount - A count of the surveys that a user has taken in the SoftLayer customer portal.
	SurveyCount uint64 `json:"surveyCount"`

	// LayoutProfiles - <nil>
	LayoutProfiles []*SoftLayer_Layout_Profile `json:"layoutProfiles"`

	// Tickets - no documentation
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// NotificationSubscriberCount - A count of notification subscription records for the user.
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount"`

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// SupportPolicyAcknowledgementRequiredFlag - Whether or not a user is required to acknowledge the
	// support policy for portal access.
	SupportPolicyAcknowledgementRequiredFlag int `json:"supportPolicyAcknowledgementRequiredFlag"`

	// UnsuccessfulLoginCount - A count of a user's unsuccessful attempts to log into the SoftLayer
	// customer portal.
	UnsuccessfulLoginCount uint64 `json:"unsuccessfulLoginCount"`

	// HasAcknowledgedSupportPolicyFlag - Whether or not a user has acknowledged the support policy.
	HasAcknowledgedSupportPolicyFlag bool `json:"hasAcknowledgedSupportPolicyFlag"`

	// SecurityAnswers - A portal user's security question answers. Some portal users may not have security
	// answers or may not be configured to require answering a security question on login.
	SecurityAnswers []*SoftLayer_User_Customer_Security_Answer `json:"securityAnswers"`

	// HasFullVirtualGuestAccessFlag - Whether or not a portal user has access to all hardware on their
	// account.
	HasFullVirtualGuestAccessFlag bool `json:"hasFullVirtualGuestAccessFlag"`

	// UserLinks - <nil>
	UserLinks []*SoftLayer_User_Customer_Link `json:"userLinks"`

	// ChildUsers - A portal user's child users. Some portal users may not have child users.
	ChildUsers []*SoftLayer_User_Customer `json:"childUsers"`
}

func (softlayer_user_customer *SoftLayer_User_Customer) String() string {
	return "SoftLayer_User_Customer"
}
