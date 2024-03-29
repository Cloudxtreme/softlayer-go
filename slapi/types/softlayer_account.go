package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"

	time "time"
)

// SoftLayer_Account - The SoftLayer_Account data type contains general information relating to a
// single SoftLayer customer account. Personal information in this type such as names, addresses, and
// phone numbers are assigned to the account only and not to users belonging to the account. The
// SoftLayer_Account data type contains a number of relational properties that are used by the
// SoftLayer customer portal to quickly present a variety of account related services to it's users.
// SoftLayer customers are unable to change their company account information in the portal or the If
// you need to change this information please open a sales ticket in our customer portal and our
// account management staff will assist you.
type SoftLayer_Account struct {

	// Email - no documentation
	Email string `json:"email,omitempty"`

	// State - A two-letter abbreviation of the state in the mailing address belonging to an account. If an
	// account does not reside in a province then this is typically blank.
	State string `json:"state,omitempty"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId,omitempty"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate,omitempty"`

	// ClaimedTaxExemptTxFlag - Whether an account is exempt from taxes on their invoices.
	ClaimedTaxExemptTxFlag bool `json:"claimedTaxExemptTxFlag,omitempty"`

	// PostalCode - The postal code of the mailing address belonging to an account.
	PostalCode string `json:"postalCode,omitempty"`

	// Address1 - The first line of the mailing address belonging to an account.
	Address1 string `json:"address1,omitempty"`

	// StatusDate - no documentation
	StatusDate *time.Time `json:"statusDate,omitempty"`

	// City - The city of the mailing address belonging to an account.
	City string `json:"city,omitempty"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone,omitempty"`

	// IsReseller - A flag indicating if an account belongs to a reseller or not.
	IsReseller int `json:"isReseller,omitempty"`

	// LastName - Each customer account is listed under a single individual. This is that individual's last
	// name.
	LastName string `json:"lastName,omitempty"`

	// BrandId - no documentation
	BrandId int `json:"brandId,omitempty"`

	// AllowedPptpVpnQuantity - no documentation
	AllowedPptpVpnQuantity int `json:"allowedPptpVpnQuantity,omitempty"`

	// Address2 - The second line of the mailing address belonging to an account.
	Address2 string `json:"address2,omitempty"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone,omitempty"`

	// FirstName - Each customer account is listed under a single individual. This is that individual's
	// first name.
	FirstName string `json:"firstName,omitempty"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone,omitempty"`

	// Country - A two-letter abbreviation of the country in the mailing address belonging to an account.
	Country string `json:"country,omitempty"`

	// CompanyName - no documentation
	CompanyName string `json:"companyName,omitempty"`

	// Id - A customer account's internal identifier. Account numbers are typically preceded by the string
	// in the customer portal. Every SoftLayer account has at least one portal user whose username follows
	// the + account number naming scheme.
	Id int `json:"id,omitempty"`

	// AccountStatusId - no documentation
	AccountStatusId int `json:"accountStatusId,omitempty"`

	// AccountManagedResourcesFlag - A flag indicating that the account has a managed resource.
	AccountManagedResourcesFlag bool `json:"accountManagedResourcesFlag,omitempty"`

	// LateFeeProtectionFlag - no documentation
	LateFeeProtectionFlag bool `json:"lateFeeProtectionFlag,omitempty"`

	// FlexibleCreditEnrollments - All of the account's current and former Flexible Credit enrollments.
	FlexibleCreditEnrollments []*SoftLayer_FlexibleCredit_Enrollment `json:"flexibleCreditEnrollments,omitempty"`

	// GlobalIpRecordCount - no documentation
	GlobalIpRecordCount uint64 `json:"globalIpRecordCount,omitempty"`

	// HardwareWithPlesk - All hardware associated with an account that has the Plesk web hosting control
	// panel installed.
	HardwareWithPlesk []*SoftLayer_Hardware `json:"hardwareWithPlesk,omitempty"`

	// PrivateIpAddresses - <nil>
	PrivateIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"privateIpAddresses,omitempty"`

	// LastFiveClosedSupportTicketCount - A count of the five most recently closed support tickets
	// associated with an account.
	LastFiveClosedSupportTicketCount uint64 `json:"lastFiveClosedSupportTicketCount,omitempty"`

	// VirtualGuestsWithMcafeeAntivirusRedhatCount - A count of all virtual guests associated with an
	// account that have McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhatCount uint64 `json:"virtualGuestsWithMcafeeAntivirusRedhatCount,omitempty"`

	// EvaultMasterUserCount - A count of an account's master EVault user. This is only used when an
	// account has EVault service.
	EvaultMasterUserCount uint64 `json:"evaultMasterUserCount,omitempty"`

	// NetworkMonitorDownVirtualGuestCount - A count of virtual guest which is currently experiencing a
	// service failure.
	NetworkMonitorDownVirtualGuestCount uint64 `json:"networkMonitorDownVirtualGuestCount,omitempty"`

	// NetworkStorageCount - A count of an account's associated storage volumes. This includes Lockbox,
	// EVault, and iSCSI volumes.
	NetworkStorageCount uint64 `json:"networkStorageCount,omitempty"`

	// PermissionRoles - no documentation
	PermissionRoles []*SoftLayer_User_Permission_Role `json:"permissionRoles,omitempty"`

	// RwhoisData - An account's reverse data. This data is used when making requests.
	RwhoisData *SoftLayer_Network_Subnet_Rwhois_Data `json:"rwhoisData,omitempty"`

	// AttributeCount - A count of the account attribute values for a SoftLayer customer account.
	AttributeCount uint64 `json:"attributeCount,omitempty"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount,omitempty"`

	// VirtualGuestsWithCpanel - All virtual guests associated with an account that has the cPanel web
	// hosting control panel installed.
	VirtualGuestsWithCpanel []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithCpanel,omitempty"`

	// TicketsClosedInTheLastThreeDaysCount - A count of tickets closed within the last 72 hours or last 10
	// tickets, whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDaysCount uint64 `json:"ticketsClosedInTheLastThreeDaysCount,omitempty"`

	// HourlyVirtualGuests - An account's associated hourly virtual guest objects.
	HourlyVirtualGuests []*SoftLayer_Virtual_Guest `json:"hourlyVirtualGuests,omitempty"`

	// NetworkGatewayCount - A count of all network gateway devices on this account.
	NetworkGatewayCount uint64 `json:"networkGatewayCount,omitempty"`

	// RouterCount - A count of all Routers that an accounts VLANs reside on
	RouterCount uint64 `json:"routerCount,omitempty"`

	// BrandKeyName - no documentation
	BrandKeyName string `json:"brandKeyName,omitempty"`

	// HardwareWithPleskCount - A count of all hardware associated with an account that has the Plesk web
	// hosting control panel installed.
	HardwareWithPleskCount uint64 `json:"hardwareWithPleskCount,omitempty"`

	// NextBillingPublicAllotmentHardwareBandwidthDetails - - This information can be pulled directly
	// through tapping keys now - The allotments for this account and their servers for the next billing
	// cycle. The public inbound and outbound bandwidth is calculated for each server in addition to the
	// daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"nextBillingPublicAllotmentHardwareBandwidthDetails,omitempty"`

	// PendingInvoiceTotalRecurringTaxAmount - The total amount of the recurring taxes on an account's
	// pending invoice, if one exists.
	PendingInvoiceTotalRecurringTaxAmount slapi.Float64 `json:"pendingInvoiceTotalRecurringTaxAmount,omitempty"`

	// GlobalIpRecords - <nil>
	GlobalIpRecords []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpRecords,omitempty"`

	// HardwareWithUrchin - All hardware associated with an account that has the Urchin web traffic
	// analytics package installed.
	HardwareWithUrchin []*SoftLayer_Hardware `json:"hardwareWithUrchin,omitempty"`

	// HasIderaBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has an installation
	// of Idera Server Backup otherwise 0.
	HasIderaBareMetalRestorePluginFlag bool `json:"hasIderaBareMetalRestorePluginFlag,omitempty"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount,omitempty"`

	// LastCanceledBillingItem - no documentation
	LastCanceledBillingItem *SoftLayer_Billing_Item `json:"lastCanceledBillingItem,omitempty"`

	// MessageQueueAccounts - no documentation
	MessageQueueAccounts []*SoftLayer_Network_Message_Queue `json:"messageQueueAccounts,omitempty"`

	// NextInvoiceTotalTaxableRecurringAmount - The total recurring charge amount of an account's next
	// invoice measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalTaxableRecurringAmount slapi.Float64 `json:"nextInvoiceTotalTaxableRecurringAmount,omitempty"`

	// AdcLoadBalancerCount - no documentation
	AdcLoadBalancerCount uint64 `json:"adcLoadBalancerCount,omitempty"`

	// AccountLinks - <nil>
	AccountLinks []*SoftLayer_Account_Link `json:"accountLinks,omitempty"`

	// HourlyBareMetalInstances - An account's associated hourly bare metal server objects.
	HourlyBareMetalInstances []*SoftLayer_Hardware `json:"hourlyBareMetalInstances,omitempty"`

	// LockboxNetworkStorage - no documentation
	LockboxNetworkStorage []*SoftLayer_Network_Storage `json:"lockboxNetworkStorage,omitempty"`

	// NetworkMonitorDownVirtualGuests - Virtual guest which is currently experiencing a service failure.
	NetworkMonitorDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorDownVirtualGuests,omitempty"`

	// OpenRecurringInvoices - no documentation
	OpenRecurringInvoices []*SoftLayer_Billing_Invoice `json:"openRecurringInvoices,omitempty"`

	// PreviousRecurringRevenue - The total recurring amount for an accounts previous revenue.
	PreviousRecurringRevenue slapi.Float64 `json:"previousRecurringRevenue,omitempty"`

	// AffiliateId - An affiliate identifier associated with the customer account.
	AffiliateId string `json:"affiliateId,omitempty"`

	// OpenBillingTickets - The open billing tickets associated with an account.
	OpenBillingTickets []*SoftLayer_Ticket `json:"openBillingTickets,omitempty"`

	// AllTopLevelBillingItemsUnfiltered - The billing items that will be on an account's next invoice.
	// Does not consider associated items.
	AllTopLevelBillingItemsUnfiltered []*SoftLayer_Billing_Item `json:"allTopLevelBillingItemsUnfiltered,omitempty"`

	// DomainCount - A count of the DNS domains associated with an account.
	DomainCount uint64 `json:"domainCount,omitempty"`

	// VirtualGuestsWithMcafeeCount - A count of all virtual guests associated with an account that have
	// McAfee Secure software components.
	VirtualGuestsWithMcafeeCount uint64 `json:"virtualGuestsWithMcafeeCount,omitempty"`

	// HourlyServiceBillingItems - Hourly service billing items that will be on an account's next invoice.
	HourlyServiceBillingItems []*SoftLayer_Billing_Item `json:"hourlyServiceBillingItems,omitempty"`

	// UpgradeRequests - no documentation
	UpgradeRequests []*SoftLayer_Product_Upgrade_Request `json:"upgradeRequests,omitempty"`

	// LockboxNetworkStorageCount - A count of an account's associated Lockbox storage volumes.
	LockboxNetworkStorageCount uint64 `json:"lockboxNetworkStorageCount,omitempty"`

	// VirtualGuestsWithQuantastorCount - A count of all virtual guests associated with an account that
	// have the QuantaStor storage system installed.
	VirtualGuestsWithQuantastorCount uint64 `json:"virtualGuestsWithQuantastorCount,omitempty"`

	// NextInvoiceIncubatorExemptTotal - The pre-tax total amount exempt from incubator credit for the
	// account's next invoice. This field is now deprecated and will soon be removed. Please update all
	// references to instead use nextInvoiceTotalAmount
	NextInvoiceIncubatorExemptTotal slapi.Float64 `json:"nextInvoiceIncubatorExemptTotal,omitempty"`

	// BandwidthAllotmentsOverAllocationCount - A count of the bandwidth allotments for an account
	// currently over allocation.
	BandwidthAllotmentsOverAllocationCount uint64 `json:"bandwidthAllotmentsOverAllocationCount,omitempty"`

	// EvaultNetworkStorage - no documentation
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage,omitempty"`

	// FacilityLogs - Logs of who entered a colocation area which is assigned to this account, or when a
	// user under this account enters a datacenter.
	FacilityLogs []*SoftLayer_User_Access_Facility_Log `json:"facilityLogs,omitempty"`

	// GlobalIpv6Records - <nil>
	GlobalIpv6Records []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpv6Records,omitempty"`

	// HourlyServiceBillingItemCount - A count of hourly service billing items that will be on an account's
	// next invoice.
	HourlyServiceBillingItemCount uint64 `json:"hourlyServiceBillingItemCount,omitempty"`

	// PriceRestrictions - no documentation
	PriceRestrictions []*SoftLayer_Product_Item_Price_Account_Restriction `json:"priceRestrictions,omitempty"`

	// MonthlyBareMetalInstances - An account's associated monthly bare metal server objects.
	MonthlyBareMetalInstances []*SoftLayer_Hardware `json:"monthlyBareMetalInstances,omitempty"`

	// OpenTicketsWaitingOnCustomer - All open tickets associated with an account last edited by an
	// employee.
	OpenTicketsWaitingOnCustomer []*SoftLayer_Ticket `json:"openTicketsWaitingOnCustomer,omitempty"`

	// SecurityScanRequests - no documentation
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests,omitempty"`

	// NextInvoiceTotalOneTimeAmount - The total one-time charge amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeAmount slapi.Float64 `json:"nextInvoiceTotalOneTimeAmount,omitempty"`

	// Routers - no documentation
	Routers []*SoftLayer_Hardware `json:"routers,omitempty"`

	// VirtualGuestsWithCpanelCount - A count of all virtual guests associated with an account that has the
	// cPanel web hosting control panel installed.
	VirtualGuestsWithCpanelCount uint64 `json:"virtualGuestsWithCpanelCount,omitempty"`

	// IpAddresses - <nil>
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses,omitempty"`

	// NextInvoiceTotalRecurringAmountBeforeAccountDiscount - The total recurring charge amount of an
	// account's next invoice measured in US Dollars assuming no changes or charges occur between now and
	// time of billing.
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount slapi.Float64 `json:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount,omitempty"`

	// NextInvoiceTotalRecurringTaxAmount - The total recurring tax amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringTaxAmount slapi.Float64 `json:"nextInvoiceTotalRecurringTaxAmount,omitempty"`

	// ResourceGroups - no documentation
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups,omitempty"`

	// ActiveColocationContainerCount - A count of the account's active top level colocation containers.
	ActiveColocationContainerCount uint64 `json:"activeColocationContainerCount,omitempty"`

	// ActiveFlexibleCreditEnrollment - Account's currently active Flexible Credit enrollment.
	ActiveFlexibleCreditEnrollment *SoftLayer_FlexibleCredit_Enrollment `json:"activeFlexibleCreditEnrollment,omitempty"`

	// GlobalLoadBalancerAccounts - The global load balancer accounts for a softlayer customer account.
	GlobalLoadBalancerAccounts []*SoftLayer_Network_LoadBalancer_Global_Account `json:"globalLoadBalancerAccounts,omitempty"`

	// HasPendingOrder - The number of orders in a status for a SoftLayer customer account.
	HasPendingOrder uint `json:"hasPendingOrder,omitempty"`

	// LastCancelledServerBillingItem - no documentation
	LastCancelledServerBillingItem *SoftLayer_Billing_Item `json:"lastCancelledServerBillingItem,omitempty"`

	// LatestRecurringInvoice - no documentation
	LatestRecurringInvoice *SoftLayer_Billing_Invoice `json:"latestRecurringInvoice,omitempty"`

	// NextInvoiceTopLevelBillingItems - The billing items that will be on an account's next invoice.
	NextInvoiceTopLevelBillingItems []*SoftLayer_Billing_Item `json:"nextInvoiceTopLevelBillingItems,omitempty"`

	// BillingAgreements - no documentation
	BillingAgreements []*SoftLayer_Account_Agreement `json:"billingAgreements,omitempty"`

	// OpenStackAccountLinks - <nil>
	OpenStackAccountLinks []*SoftLayer_Account_Link `json:"openStackAccountLinks,omitempty"`

	// PendingInvoice - no documentation
	PendingInvoice *SoftLayer_Billing_Invoice `json:"pendingInvoice,omitempty"`

	// TranscodeAccounts - no documentation
	TranscodeAccounts []*SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccounts,omitempty"`

	// InvoiceCount - A count of an account's associated billing invoices.
	InvoiceCount uint64 `json:"invoiceCount,omitempty"`

	// LastFiveClosedSalesTickets - The five most recently closed sales tickets associated with an account.
	LastFiveClosedSalesTickets []*SoftLayer_Ticket `json:"lastFiveClosedSalesTickets,omitempty"`

	// PriorityOneTickets - All priority one tickets associated with an account.
	PriorityOneTickets []*SoftLayer_Ticket `json:"priorityOneTickets,omitempty"`

	// RecentEventCount - no documentation
	RecentEventCount uint64 `json:"recentEventCount,omitempty"`

	// NextInvoiceTotalRecurringAmount - The total recurring charge amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringAmount slapi.Float64 `json:"nextInvoiceTotalRecurringAmount,omitempty"`

	// ReferralPartner - no documentation
	ReferralPartner *SoftLayer_Account `json:"referralPartner,omitempty"`

	// Balance - The account balance of a SoftLayer customer account. An account's balance is the amount of
	// money owed to SoftLayer by the account holder, returned as a floating point number with two decimal
	// places, measured in US Dollars A negative account balance means the account holder has overpaid and
	// is owed money by SoftLayer.
	Balance slapi.Float64 `json:"balance,omitempty"`

	// HourlyVirtualGuestCount - A count of an account's associated hourly virtual guest objects.
	HourlyVirtualGuestCount uint64 `json:"hourlyVirtualGuestCount,omitempty"`

	// OrphanBillingItems - The billing items that have no parent billing item. These are items that don't
	// necessarily belong to a single server.
	OrphanBillingItems []*SoftLayer_Billing_Item `json:"orphanBillingItems,omitempty"`

	// OrphanBillingItemCount - A count of the billing items that have no parent billing item. These are
	// items that don't necessarily belong to a single server.
	OrphanBillingItemCount uint64 `json:"orphanBillingItemCount,omitempty"`

	// PaymentProcessorCount - no documentation
	PaymentProcessorCount uint64 `json:"paymentProcessorCount,omitempty"`

	// PendingInvoiceTotalOneTimeAmount - The total one-time charges for an account's pending invoice, if
	// one exists. In other words, it is the sum of one-time charges, setup fees, and labor fees. It does
	// not include taxes.
	PendingInvoiceTotalOneTimeAmount slapi.Float64 `json:"pendingInvoiceTotalOneTimeAmount,omitempty"`

	// NetworkHardwareCount - A count of an account's associated network hardware.
	NetworkHardwareCount uint64 `json:"networkHardwareCount,omitempty"`

	// Shipments - no documentation
	Shipments []*SoftLayer_Account_Shipment `json:"shipments,omitempty"`

	// ActiveQuoteCount - no documentation
	ActiveQuoteCount uint64 `json:"activeQuoteCount,omitempty"`

	// OpenSupportTicketCount - A count of the open support tickets associated with an account.
	OpenSupportTicketCount uint64 `json:"openSupportTicketCount,omitempty"`

	// OwnedBrands - <nil>
	OwnedBrands []*SoftLayer_Brand `json:"ownedBrands,omitempty"`

	// SalesforceAccountLink - <nil>
	SalesforceAccountLink *SoftLayer_Account_Link `json:"salesforceAccountLink,omitempty"`

	// NotificationSubscriberCount - no documentation
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount,omitempty"`

	// OrderCount - A count of an account's associated billing orders excluding upgrades.
	OrderCount uint64 `json:"orderCount,omitempty"`

	// SecurityCertificateCount - no documentation
	SecurityCertificateCount uint64 `json:"securityCertificateCount,omitempty"`

	// MasterUser - no documentation
	MasterUser *SoftLayer_User_Customer `json:"masterUser,omitempty"`

	// OpenSalesTicketCount - A count of the open sales tickets associated with an account.
	OpenSalesTicketCount uint64 `json:"openSalesTicketCount,omitempty"`

	// PublicAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information can be pulled
	// directly through tapping keys now - The allotments for this account and their servers. The public
	// inbound and outbound bandwidth is calculated for each server in addition to the daily average
	// network traffic since the last billing date.
	PublicAllotmentHardwareBandwidthDetailCount uint64 `json:"publicAllotmentHardwareBandwidthDetailCount,omitempty"`

	// CanOrderAdditionalVlansFlag - Indicating whether this account can order additional Vlans.
	CanOrderAdditionalVlansFlag bool `json:"canOrderAdditionalVlansFlag,omitempty"`

	// LatestRecurringPendingInvoice - no documentation
	LatestRecurringPendingInvoice *SoftLayer_Billing_Invoice `json:"latestRecurringPendingInvoice,omitempty"`

	// OwnedHardwareGenericComponentModelCount - no documentation
	OwnedHardwareGenericComponentModelCount uint64 `json:"ownedHardwareGenericComponentModelCount,omitempty"`

	// AllRecurringTopLevelBillingItems - The billing items that will be on an account's next invoice.
	AllRecurringTopLevelBillingItems []*SoftLayer_Billing_Item `json:"allRecurringTopLevelBillingItems,omitempty"`

	// NetworkTunnelContextCount - no documentation
	NetworkTunnelContextCount uint64 `json:"networkTunnelContextCount,omitempty"`

	// PublicSubnets - All public network subnets associated with an account.
	PublicSubnets []*SoftLayer_Network_Subnet `json:"publicSubnets,omitempty"`

	// CatalystEnrollmentCount - no documentation
	CatalystEnrollmentCount uint64 `json:"catalystEnrollmentCount,omitempty"`

	// PriorityOneTicketCount - A count of all priority one tickets associated with an account.
	PriorityOneTicketCount uint64 `json:"priorityOneTicketCount,omitempty"`

	// ReplicationEventCount - A count of the Replication events for all Network Storage volumes on an
	// account.
	ReplicationEventCount uint64 `json:"replicationEventCount,omitempty"`

	// HubNetworkStorage - no documentation
	HubNetworkStorage []*SoftLayer_Network_Storage `json:"hubNetworkStorage,omitempty"`

	// VirtualGuestsWithMcafeeAntivirusWindowCount - A count of all virtual guests associated with an
	// account that has McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindowCount uint64 `json:"virtualGuestsWithMcafeeAntivirusWindowCount,omitempty"`

	// DisablePaymentProcessingFlag - A flag indicating whether payments are processed for this account.
	DisablePaymentProcessingFlag bool `json:"disablePaymentProcessingFlag,omitempty"`

	// HardwareWithCpanel - All hardware associated with an account that has the cPanel web hosting control
	// panel installed.
	HardwareWithCpanel []*SoftLayer_Hardware `json:"hardwareWithCpanel,omitempty"`

	// OpenAccountingTickets - The open accounting tickets associated with an account.
	OpenAccountingTickets []*SoftLayer_Ticket `json:"openAccountingTickets,omitempty"`

	// HardwareWithWindowCount - A count of all hardware associated with an account that is running a
	// version of the Microsoft Windows operating system.
	HardwareWithWindowCount uint64 `json:"hardwareWithWindowCount,omitempty"`

	// SslVpnUserCount - A count of an account's associated portal users with SSL VPN access.
	SslVpnUserCount uint64 `json:"sslVpnUserCount,omitempty"`

	// SupportRepresentativeCount - A count of the SoftLayer employees that an account is assigned to.
	SupportRepresentativeCount uint64 `json:"supportRepresentativeCount,omitempty"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// PostProvisioningHookCount - A count of customer specified URIs that are downloaded onto a newly
	// provisioned or reloaded server. If the URI is sent over https it will be executed directly on the
	// server.
	PostProvisioningHookCount uint64 `json:"postProvisioningHookCount,omitempty"`

	// BandwidthAllotmentsProjectedOverAllocation - The bandwidth allotments for an account projected to go
	// over allocation.
	BandwidthAllotmentsProjectedOverAllocation []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsProjectedOverAllocation,omitempty"`

	// NetworkHardware - no documentation
	NetworkHardware []*SoftLayer_Hardware `json:"networkHardware,omitempty"`

	// ActiveQuotes - no documentation
	ActiveQuotes []*SoftLayer_Billing_Order_Quote `json:"activeQuotes,omitempty"`

	// OpenAbuseTickets - no documentation
	OpenAbuseTickets []*SoftLayer_Ticket `json:"openAbuseTickets,omitempty"`

	// RemoteManagementCommandRequestCount - A count of remote management command requests for an account
	RemoteManagementCommandRequestCount uint64 `json:"remoteManagementCommandRequestCount,omitempty"`

	// VirtualStorageArchiveRepositoryCount - A count of an account's associated virtual server archived
	// storage repositories.
	VirtualStorageArchiveRepositoryCount uint64 `json:"virtualStorageArchiveRepositoryCount,omitempty"`

	// HardwareWithMcafeeAntivirusRedhat - All hardware associated with an account that has McAfee Secure
	// AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhat []*SoftLayer_Hardware `json:"hardwareWithMcafeeAntivirusRedhat,omitempty"`

	// HardwareWithWindows - All hardware associated with an account that is running a version of the
	// Microsoft Windows operating system.
	HardwareWithWindows []*SoftLayer_Hardware `json:"hardwareWithWindows,omitempty"`

	// ServiceBillingItems - The service billing items that will be on an account's next invoice.
	ServiceBillingItems []*SoftLayer_Billing_Item `json:"serviceBillingItems,omitempty"`

	// OpenStackAccountLinkCount - no documentation
	OpenStackAccountLinkCount uint64 `json:"openStackAccountLinkCount,omitempty"`

	// PortableStorageVolumeCount - no documentation
	PortableStorageVolumeCount uint64 `json:"portableStorageVolumeCount,omitempty"`

	// TagCount - no documentation
	TagCount uint64 `json:"tagCount,omitempty"`

	// HardwareOverBandwidthAllocation - An account's associated hardware objects currently over bandwidth
	// allocation.
	HardwareOverBandwidthAllocation []*SoftLayer_Hardware `json:"hardwareOverBandwidthAllocation,omitempty"`

	// LastFiveClosedAccountingTickets - The five most recently closed accounting tickets associated with
	// an account.
	LastFiveClosedAccountingTickets []*SoftLayer_Ticket `json:"lastFiveClosedAccountingTickets,omitempty"`

	// NetworkGateways - no documentation
	NetworkGateways []*SoftLayer_Network_Gateway `json:"networkGateways,omitempty"`

	// VirtualGuestsWithMcafee - All virtual guests associated with an account that have McAfee Secure
	// software components.
	VirtualGuestsWithMcafee []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafee,omitempty"`

	// AllBillingItems - The billing items that will be on an account's next invoice.
	AllBillingItems []*SoftLayer_Billing_Item `json:"allBillingItems,omitempty"`

	// CdnAccountCount - no documentation
	CdnAccountCount uint64 `json:"cdnAccountCount,omitempty"`

	// IscsiNetworkStorageCount - A count of an account's associated iSCSI storage volumes.
	IscsiNetworkStorageCount uint64 `json:"iscsiNetworkStorageCount,omitempty"`

	// LastFiveClosedSalesTicketCount - A count of the five most recently closed sales tickets associated
	// with an account.
	LastFiveClosedSalesTicketCount uint64 `json:"lastFiveClosedSalesTicketCount,omitempty"`

	// ServiceBillingItemCount - A count of the service billing items that will be on an account's next
	// invoice.
	ServiceBillingItemCount uint64 `json:"serviceBillingItemCount,omitempty"`

	// LastFiveClosedAbuseTickets - The five most recently closed abuse tickets associated with an account.
	LastFiveClosedAbuseTickets []*SoftLayer_Ticket `json:"lastFiveClosedAbuseTickets,omitempty"`

	// SslVpnUsers - An account's associated portal users with SSL VPN access.
	SslVpnUsers []*SoftLayer_User_Customer `json:"sslVpnUsers,omitempty"`

	// MonthlyBareMetalInstanceCount - A count of an account's associated monthly bare metal server
	// objects.
	MonthlyBareMetalInstanceCount uint64 `json:"monthlyBareMetalInstanceCount,omitempty"`

	// NextBillingPublicAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information
	// can be pulled directly through tapping keys now - The allotments for this account and their servers
	// for the next billing cycle. The public inbound and outbound bandwidth is calculated for each server
	// in addition to the daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetailCount uint64 `json:"nextBillingPublicAllotmentHardwareBandwidthDetailCount,omitempty"`

	// PriceRestrictionCount - A count of the item price that an account is restricted to.
	PriceRestrictionCount uint64 `json:"priceRestrictionCount,omitempty"`

	// PrivateBlockDeviceTemplateGroupCount - A count of private and shared template group objects (parent
	// only) for an account.
	PrivateBlockDeviceTemplateGroupCount uint64 `json:"privateBlockDeviceTemplateGroupCount,omitempty"`

	// ActiveColocationContainers - The account's active top level colocation containers.
	ActiveColocationContainers []*SoftLayer_Billing_Item `json:"activeColocationContainers,omitempty"`

	// LegacyIscsiCapacityGB - The total capacity of Legacy iSCSI Volumes on an account, in
	LegacyIscsiCapacityGB uint `json:"legacyIscsiCapacityGB,omitempty"`

	// NotificationSubscribers - <nil>
	NotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"notificationSubscribers,omitempty"`

	// PortableStorageVolumes - <nil>
	PortableStorageVolumes []*SoftLayer_Virtual_Disk_Image `json:"portableStorageVolumes,omitempty"`

	// PublicNetworkVlanCount - A count of the public network VLANs assigned to an account.
	PublicNetworkVlanCount uint64 `json:"publicNetworkVlanCount,omitempty"`

	// ManualPaymentsUnderReview - <nil>
	ManualPaymentsUnderReview []*SoftLayer_Billing_Payment_Card_ManualPayment `json:"manualPaymentsUnderReview,omitempty"`

	// ShipmentCount - A count of shipments that belong to the customer's account.
	ShipmentCount uint64 `json:"shipmentCount,omitempty"`

	// NetworkMonitorDownHardwareCount - A count of hardware which is currently experiencing a service
	// failure.
	NetworkMonitorDownHardwareCount uint64 `json:"networkMonitorDownHardwareCount,omitempty"`

	// PendingInvoiceTopLevelItemCount - A count of a list of top-level invoice items that are on an
	// account's currently pending invoice.
	PendingInvoiceTopLevelItemCount uint64 `json:"pendingInvoiceTopLevelItemCount,omitempty"`

	// OwnedHardwareGenericComponentModels - <nil>
	OwnedHardwareGenericComponentModels []*SoftLayer_Hardware_Component_Model_Generic `json:"ownedHardwareGenericComponentModels,omitempty"`

	// SecurityCertificates - no documentation
	SecurityCertificates []*SoftLayer_Security_Certificate `json:"securityCertificates,omitempty"`

	// VirtualGuestsWithMcafeeAntivirusRedhat - All virtual guests associated with an account that have
	// McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhat []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusRedhat,omitempty"`

	// ActiveAddressCount - A count of the active address(es) that belong to an account.
	ActiveAddressCount uint64 `json:"activeAddressCount,omitempty"`

	// BareMetalInstanceCount - A count of an account's associated bare metal server objects.
	BareMetalInstanceCount uint64 `json:"bareMetalInstanceCount,omitempty"`

	// DomainRegistrationCount - no documentation
	DomainRegistrationCount uint64 `json:"domainRegistrationCount,omitempty"`

	// PublicIpAddressCount - no documentation
	PublicIpAddressCount uint64 `json:"publicIpAddressCount,omitempty"`

	// VirtualGuestsWithMcafeeIntrusionDetectionSystem - All virtual guests associated with an account that
	// has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystem []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeIntrusionDetectionSystem,omitempty"`

	// AccountContactCount - no documentation
	AccountContactCount uint64 `json:"accountContactCount,omitempty"`

	// AllRecurringTopLevelBillingItemsUnfilteredCount - A count of the billing items that will be on an
	// account's next invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfilteredCount uint64 `json:"allRecurringTopLevelBillingItemsUnfilteredCount,omitempty"`

	// LegacyBandwidthAllotmentCount - A count of the legacy bandwidth allotments for an account.
	LegacyBandwidthAllotmentCount uint64 `json:"legacyBandwidthAllotmentCount,omitempty"`

	// TicketsClosedTodayCount - A count of tickets closed today associated with an account.
	TicketsClosedTodayCount uint64 `json:"ticketsClosedTodayCount,omitempty"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand,omitempty"`

	// VirtualGuestsWithUrchin - All virtual guests associated with an account that has the Urchin web
	// traffic analytics package installed.
	VirtualGuestsWithUrchin []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithUrchin,omitempty"`

	// BlockDeviceTemplateGroupCount - A count of private template group objects (parent and children) and
	// the shared template group objects (parent only) for an account.
	BlockDeviceTemplateGroupCount uint64 `json:"blockDeviceTemplateGroupCount,omitempty"`

	// HasEvaultBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has the EVault Bare
	// Metal Server Restore Plugin otherwise 0.
	HasEvaultBareMetalRestorePluginFlag bool `json:"hasEvaultBareMetalRestorePluginFlag,omitempty"`

	// StandardPoolVirtualGuests - An account's virtual guest objects that are hosted on a user provisioned
	// hypervisor.
	StandardPoolVirtualGuests []*SoftLayer_Virtual_Guest `json:"standardPoolVirtualGuests,omitempty"`

	// BandwidthAllotmentsProjectedOverAllocationCount - A count of the bandwidth allotments for an account
	// projected to go over allocation.
	BandwidthAllotmentsProjectedOverAllocationCount uint64 `json:"bandwidthAllotmentsProjectedOverAllocationCount,omitempty"`

	// EvaultNetworkStorageCount - A count of an account's associated EVault storage volumes.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount,omitempty"`

	// IpAddressCount - no documentation
	IpAddressCount uint64 `json:"ipAddressCount,omitempty"`

	// LoadBalancerCount - no documentation
	LoadBalancerCount uint64 `json:"loadBalancerCount,omitempty"`

	// VirtualGuestsOverBandwidthAllocation - An account's associated virtual guest objects currently over
	// bandwidth allocation.
	VirtualGuestsOverBandwidthAllocation []*SoftLayer_Virtual_Guest `json:"virtualGuestsOverBandwidthAllocation,omitempty"`

	// OpenStackObjectStorageCount - A count of an account's associated Openstack related Object Storage
	// accounts.
	OpenStackObjectStorageCount uint64 `json:"openStackObjectStorageCount,omitempty"`

	// NetworkMonitorUpVirtualGuests - no documentation
	NetworkMonitorUpVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorUpVirtualGuests,omitempty"`

	// NetworkVlans - no documentation
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans,omitempty"`

	// BandwidthAllotments - no documentation
	BandwidthAllotments []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotments,omitempty"`

	// ExpiredSecurityCertificates - Stored security certificates that are expired (ie.
	ExpiredSecurityCertificates []*SoftLayer_Security_Certificate `json:"expiredSecurityCertificates,omitempty"`

	// MonthlyVirtualGuests - An account's associated monthly virtual guest objects.
	MonthlyVirtualGuests []*SoftLayer_Virtual_Guest `json:"monthlyVirtualGuests,omitempty"`

	// HubNetworkStorageCount - A count of an account's associated Virtual Storage volumes.
	HubNetworkStorageCount uint64 `json:"hubNetworkStorageCount,omitempty"`

	// AbuseEmails - An email address that is responsible for abuse and legal inquiries on behalf of an
	// account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmails []*SoftLayer_Account_AbuseEmail `json:"abuseEmails,omitempty"`

	// VirtualPrivateRack - no documentation
	VirtualPrivateRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualPrivateRack,omitempty"`

	// PrivateSubnetCount - A count of all private subnets associated with an account.
	PrivateSubnetCount uint64 `json:"privateSubnetCount,omitempty"`

	// AccountContacts - no documentation
	AccountContacts []*SoftLayer_Account_Contact `json:"accountContacts,omitempty"`

	// HardwareWithMcafee - All hardware associated with an account that has McAfee Secure software
	// components.
	HardwareWithMcafee []*SoftLayer_Hardware `json:"hardwareWithMcafee,omitempty"`

	// AvailablePublicNetworkVlanCount - A count of the public network VLANs assigned to an account.
	AvailablePublicNetworkVlanCount uint64 `json:"availablePublicNetworkVlanCount,omitempty"`

	// ValidSecurityCertificateCount - A count of stored security certificates that are not expired (ie.
	ValidSecurityCertificateCount uint64 `json:"validSecurityCertificateCount,omitempty"`

	// CdnAccounts - no documentation
	CdnAccounts []*SoftLayer_Network_ContentDelivery_Account `json:"cdnAccounts,omitempty"`

	// PrivateNetworkVlans - no documentation
	PrivateNetworkVlans []*SoftLayer_Network_Vlan `json:"privateNetworkVlans,omitempty"`

	// NetworkMonitorRecoveringVirtualGuestCount - A count of virtual guest which is currently recovering
	// from a service failure.
	NetworkMonitorRecoveringVirtualGuestCount uint64 `json:"networkMonitorRecoveringVirtualGuestCount,omitempty"`

	// PptpVpnUserCount - A count of an account's associated portal users with VPN access.
	PptpVpnUserCount uint64 `json:"pptpVpnUserCount,omitempty"`

	// AllTopLevelBillingItemsUnfilteredCount - A count of the billing items that will be on an account's
	// next invoice. Does not consider associated items.
	AllTopLevelBillingItemsUnfilteredCount uint64 `json:"allTopLevelBillingItemsUnfilteredCount,omitempty"`

	// UpgradeRequestCount - A count of an account's associated upgrade requests.
	UpgradeRequestCount uint64 `json:"upgradeRequestCount,omitempty"`

	// OpenTicketsWaitingOnCustomerCount - A count of all open tickets associated with an account last
	// edited by an employee.
	OpenTicketsWaitingOnCustomerCount uint64 `json:"openTicketsWaitingOnCustomerCount,omitempty"`

	// ApplicationDeliveryControllers - An account's associated application delivery controller records.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers,omitempty"`

	// DomainRegistrations - <nil>
	DomainRegistrations []*SoftLayer_Dns_Domain_Registration `json:"domainRegistrations,omitempty"`

	// LastFiveClosedOtherTickets - The five most recently closed tickets that do not belong to the abuse,
	// accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTickets []*SoftLayer_Ticket `json:"lastFiveClosedOtherTickets,omitempty"`

	// FlexibleCreditEnrollmentCount - A count of all of the account's current and former Flexible Credit
	// enrollments.
	FlexibleCreditEnrollmentCount uint64 `json:"flexibleCreditEnrollmentCount,omitempty"`

	// HardwareWithHelmCount - A count of all hardware associated with an account that has the Helm web
	// hosting control panel installed.
	HardwareWithHelmCount uint64 `json:"hardwareWithHelmCount,omitempty"`

	// OpenCancellationRequestCount - A count of an open ticket requesting cancellation of this server, if
	// one exists.
	OpenCancellationRequestCount uint64 `json:"openCancellationRequestCount,omitempty"`

	// GlobalIpv4Records - <nil>
	GlobalIpv4Records []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpv4Records,omitempty"`

	// OpenSalesTickets - no documentation
	OpenSalesTickets []*SoftLayer_Ticket `json:"openSalesTickets,omitempty"`

	// BillingInfo - no documentation
	BillingInfo *SoftLayer_Billing_Info `json:"billingInfo,omitempty"`

	// PublicNetworkVlans - no documentation
	PublicNetworkVlans []*SoftLayer_Network_Vlan `json:"publicNetworkVlans,omitempty"`

	// LastFiveClosedAccountingTicketCount - A count of the five most recently closed accounting tickets
	// associated with an account.
	LastFiveClosedAccountingTicketCount uint64 `json:"lastFiveClosedAccountingTicketCount,omitempty"`

	// AccountStatus - An account's status presented in a more detailed data type.
	AccountStatus *SoftLayer_Account_Status `json:"accountStatus,omitempty"`

	// DisplaySupportRepresentativeAssignments - The SoftLayer employees that an account is assigned to.
	DisplaySupportRepresentativeAssignments []*SoftLayer_Account_Attachment_Employee `json:"displaySupportRepresentativeAssignments,omitempty"`

	// HasR1softBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has an installation
	// of R1Soft CDP otherwise 0.
	HasR1softBareMetalRestorePluginFlag bool `json:"hasR1softBareMetalRestorePluginFlag,omitempty"`

	// HardwareWithQuantastorCount - A count of all hardware associated with an account that has the
	// QuantaStor storage system installed.
	HardwareWithQuantastorCount uint64 `json:"hardwareWithQuantastorCount,omitempty"`

	// VirtualDiskImageCount - A count of an account's associated virtual server virtual disk images.
	VirtualDiskImageCount uint64 `json:"virtualDiskImageCount,omitempty"`

	// NetworkCreationFlag - Whether or not this account can define their own networks.
	NetworkCreationFlag bool `json:"networkCreationFlag,omitempty"`

	// GlobalIpv6RecordCount - no documentation
	GlobalIpv6RecordCount uint64 `json:"globalIpv6RecordCount,omitempty"`

	// MonthlyVirtualGuestCount - A count of an account's associated monthly virtual guest objects.
	MonthlyVirtualGuestCount uint64 `json:"monthlyVirtualGuestCount,omitempty"`

	// IscsiNetworkStorage - no documentation
	IscsiNetworkStorage []*SoftLayer_Network_Storage `json:"iscsiNetworkStorage,omitempty"`

	// OpenSupportTickets - The open support tickets associated with an account.
	OpenSupportTickets []*SoftLayer_Ticket `json:"openSupportTickets,omitempty"`

	// Orders - An account's associated billing orders excluding upgrades.
	Orders []*SoftLayer_Billing_Order `json:"orders,omitempty"`

	// PaymentProcessors - <nil>
	PaymentProcessors []*SoftLayer_Billing_Payment_Processor `json:"paymentProcessors,omitempty"`

	// HardwareWithMcafeeIntrusionDetectionSystemCount - A count of all hardware associated with an account
	// that has McAfee Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystemCount uint64 `json:"hardwareWithMcafeeIntrusionDetectionSystemCount,omitempty"`

	// LegacyBandwidthAllotments - no documentation
	LegacyBandwidthAllotments []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"legacyBandwidthAllotments,omitempty"`

	// AccountLinkCount - no documentation
	AccountLinkCount uint64 `json:"accountLinkCount,omitempty"`

	// LastFiveClosedAbuseTicketCount - A count of the five most recently closed abuse tickets associated
	// with an account.
	LastFiveClosedAbuseTicketCount uint64 `json:"lastFiveClosedAbuseTicketCount,omitempty"`

	// VirtualGuestsWithQuantastor - All virtual guests associated with an account that have the QuantaStor
	// storage system installed.
	VirtualGuestsWithQuantastor []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithQuantastor,omitempty"`

	// EvaultMasterUsers - An account's master EVault user. This is only used when an account has EVault
	// service.
	EvaultMasterUsers []*SoftLayer_Account_Password `json:"evaultMasterUsers,omitempty"`

	// Tickets - no documentation
	Tickets []*SoftLayer_Ticket `json:"tickets,omitempty"`

	// SecurityScanRequestCount - A count of an account's vulnerability scan requests.
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount,omitempty"`

	// Attributes - The account attribute values for a SoftLayer customer account.
	Attributes []*SoftLayer_Account_Attribute `json:"attributes,omitempty"`

	// SubnetRegistrationDetailCount - no documentation
	SubnetRegistrationDetailCount uint64 `json:"subnetRegistrationDetailCount,omitempty"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount,omitempty"`

	// VirtualStoragePublicRepositoryCount - A count of an account's associated virtual server public
	// storage repositories.
	VirtualStoragePublicRepositoryCount uint64 `json:"virtualStoragePublicRepositoryCount,omitempty"`

	// NetworkMessageDeliveryAccounts - <nil>
	NetworkMessageDeliveryAccounts []*SoftLayer_Network_Message_Delivery `json:"networkMessageDeliveryAccounts,omitempty"`

	// PendingInvoiceTopLevelItems - A list of top-level invoice items that are on an account's currently
	// pending invoice.
	PendingInvoiceTopLevelItems []*SoftLayer_Billing_Invoice_Item `json:"pendingInvoiceTopLevelItems,omitempty"`

	// PptpVpnUsers - An account's associated portal users with VPN access.
	PptpVpnUsers []*SoftLayer_User_Customer `json:"pptpVpnUsers,omitempty"`

	// SuppressInvoicesFlag - no documentation
	SuppressInvoicesFlag bool `json:"suppressInvoicesFlag,omitempty"`

	// ValidSecurityCertificates - Stored security certificates that are not expired (ie.
	ValidSecurityCertificates []*SoftLayer_Security_Certificate `json:"validSecurityCertificates,omitempty"`

	// ClosedTicketCount - A count of all closed tickets associated with an account.
	ClosedTicketCount uint64 `json:"closedTicketCount,omitempty"`

	// Addresses - no documentation
	Addresses []*SoftLayer_Account_Address `json:"addresses,omitempty"`

	// CatalystEnrollments - <nil>
	CatalystEnrollments []*SoftLayer_Catalyst_Enrollment `json:"catalystEnrollments,omitempty"`

	// OpenRecurringInvoiceCount - no documentation
	OpenRecurringInvoiceCount uint64 `json:"openRecurringInvoiceCount,omitempty"`

	// HardwareWithHelm - All hardware associated with an account that has the Helm web hosting control
	// panel installed.
	HardwareWithHelm []*SoftLayer_Hardware `json:"hardwareWithHelm,omitempty"`

	// SubnetRegistrationDetails - <nil>
	SubnetRegistrationDetails []*SoftLayer_Account_Regional_Registry_Detail `json:"subnetRegistrationDetails,omitempty"`

	// ActiveVirtualLicenseCount - A count of the virtual software licenses controlled by an account
	ActiveVirtualLicenseCount uint64 `json:"activeVirtualLicenseCount,omitempty"`

	// LockboxCapacityGB - The total capacity of Legacy lockbox Volumes on an account, in
	LockboxCapacityGB uint `json:"lockboxCapacityGB,omitempty"`

	// TicketsClosedToday - no documentation
	TicketsClosedToday []*SoftLayer_Ticket `json:"ticketsClosedToday,omitempty"`

	// MediaDataTransferRequestCount - A count of an account's media transfer service requests.
	MediaDataTransferRequestCount uint64 `json:"mediaDataTransferRequestCount,omitempty"`

	// VirtualGuestsProjectedOverBandwidthAllocationCount - A count of an account's associated virtual
	// guest objects currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocationCount uint64 `json:"virtualGuestsProjectedOverBandwidthAllocationCount,omitempty"`

	// VirtualGuestsWithMcafeeIntrusionDetectionSystemCount - A count of all virtual guests associated with
	// an account that has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount uint64 `json:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount,omitempty"`

	// NetworkStorage - An account's associated storage volumes. This includes Lockbox, EVault, and iSCSI
	// volumes.
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage,omitempty"`

	// PostProvisioningHooks - Customer specified URIs that are downloaded onto a newly provisioned or
	// reloaded server. If the URI is sent over https it will be executed directly on the server.
	PostProvisioningHooks []*SoftLayer_Provisioning_Hook `json:"postProvisioningHooks,omitempty"`

	// ReferredAccounts - If this is a account is a referral partner, the accounts this referral partner
	// has referred
	ReferredAccounts []*SoftLayer_Account `json:"referredAccounts,omitempty"`

	// CartCount - no documentation
	CartCount uint64 `json:"cartCount,omitempty"`

	// ClosedTickets - no documentation
	ClosedTickets []*SoftLayer_Ticket `json:"closedTickets,omitempty"`

	// LastFiveClosedSupportTickets - The five most recently closed support tickets associated with an
	// account.
	LastFiveClosedSupportTickets []*SoftLayer_Ticket `json:"lastFiveClosedSupportTickets,omitempty"`

	// Quotes - no documentation
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes,omitempty"`

	// PendingInvoiceTotalOneTimeTaxAmount - The sum of all the taxes related to one time charges for an
	// account's pending invoice, if one exists.
	PendingInvoiceTotalOneTimeTaxAmount slapi.Float64 `json:"pendingInvoiceTotalOneTimeTaxAmount,omitempty"`

	// BlockDeviceTemplateGroups - Private template group objects (parent and children) and the shared
	// template group objects (parent only) for an account.
	BlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups,omitempty"`

	// PrivateAllotmentHardwareBandwidthDetails - - This information can be pulled directly through tapping
	// keys now - The allotments for this account and their servers. The private inbound and outbound
	// bandwidth is calculated for each server in addition to the daily average network traffic since the
	// last billing date.
	PrivateAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"privateAllotmentHardwareBandwidthDetails,omitempty"`

	// SshKeys - Customer specified SSH keys that can be implemented onto a newly provisioned or reloaded
	// server.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys,omitempty"`

	// VirtualDedicatedRacks - no documentation
	VirtualDedicatedRacks []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualDedicatedRacks,omitempty"`

	// PermissionRoleCount - no documentation
	PermissionRoleCount uint64 `json:"permissionRoleCount,omitempty"`

	// AdcLoadBalancers - no documentation
	AdcLoadBalancers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"adcLoadBalancers,omitempty"`

	// OpenStackObjectStorage - An account's associated Openstack related Object Storage accounts.
	OpenStackObjectStorage []*SoftLayer_Network_Storage `json:"openStackObjectStorage,omitempty"`

	// ActiveBillingAgreementCount - no documentation
	ActiveBillingAgreementCount uint64 `json:"activeBillingAgreementCount,omitempty"`

	// HardwareWithCpanelCount - A count of all hardware associated with an account that has the cPanel web
	// hosting control panel installed.
	HardwareWithCpanelCount uint64 `json:"hardwareWithCpanelCount,omitempty"`

	// OpenAbuseTicketCount - A count of the open abuse tickets associated with an account.
	OpenAbuseTicketCount uint64 `json:"openAbuseTicketCount,omitempty"`

	// PrivateAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information can be
	// pulled directly through tapping keys now - The allotments for this account and their servers. The
	// private inbound and outbound bandwidth is calculated for each server in addition to the daily
	// average network traffic since the last billing date.
	PrivateAllotmentHardwareBandwidthDetailCount uint64 `json:"privateAllotmentHardwareBandwidthDetailCount,omitempty"`

	// NasNetworkStorage - no documentation
	NasNetworkStorage []*SoftLayer_Network_Storage `json:"nasNetworkStorage,omitempty"`

	// NetworkMonitorRecoveringHardware - Hardware which is currently recovering from a service failure.
	NetworkMonitorRecoveringHardware []*SoftLayer_Hardware `json:"networkMonitorRecoveringHardware,omitempty"`

	// ScaleGroups - no documentation
	ScaleGroups []*SoftLayer_Scale_Group `json:"scaleGroups,omitempty"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount,omitempty"`

	// BandwidthAllotmentsOverAllocation - The bandwidth allotments for an account currently over
	// allocation.
	BandwidthAllotmentsOverAllocation []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsOverAllocation,omitempty"`

	// Users - no documentation
	Users []*SoftLayer_User_Customer `json:"users,omitempty"`

	// BillingAgreementCount - no documentation
	BillingAgreementCount uint64 `json:"billingAgreementCount,omitempty"`

	// ManualPaymentsUnderReviewCount - no documentation
	ManualPaymentsUnderReviewCount uint64 `json:"manualPaymentsUnderReviewCount,omitempty"`

	// SubnetRegistrationCount - no documentation
	SubnetRegistrationCount uint64 `json:"subnetRegistrationCount,omitempty"`

	// AvailablePublicNetworkVlans - no documentation
	AvailablePublicNetworkVlans []*SoftLayer_Network_Vlan `json:"availablePublicNetworkVlans,omitempty"`

	// PublicIpAddresses - <nil>
	PublicIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"publicIpAddresses,omitempty"`

	// SamlAuthentication - no documentation
	SamlAuthentication *SoftLayer_Account_Authentication_Saml `json:"samlAuthentication,omitempty"`

	// AllTopLevelBillingItemCount - no documentation
	AllTopLevelBillingItemCount uint64 `json:"allTopLevelBillingItemCount,omitempty"`

	// MessageQueueAccountCount - A count of an account's associated Message Queue accounts.
	MessageQueueAccountCount uint64 `json:"messageQueueAccountCount,omitempty"`

	// TranscodeAccountCount - A count of an account's associated Transcode account.
	TranscodeAccountCount uint64 `json:"transcodeAccountCount,omitempty"`

	// BareMetalInstances - no documentation
	BareMetalInstances []*SoftLayer_Hardware `json:"bareMetalInstances,omitempty"`

	// LastFiveClosedTickets - The five most recently closed tickets associated with an account.
	LastFiveClosedTickets []*SoftLayer_Ticket `json:"lastFiveClosedTickets,omitempty"`

	// LatestBillDate - no documentation
	LatestBillDate *time.Time `json:"latestBillDate,omitempty"`

	// NetworkMessageDeliveryAccountCount - no documentation
	NetworkMessageDeliveryAccountCount uint64 `json:"networkMessageDeliveryAccountCount,omitempty"`

	// NetworkMonitorUpHardwareCount - no documentation
	NetworkMonitorUpHardwareCount uint64 `json:"networkMonitorUpHardwareCount,omitempty"`

	// ScaleGroupCount - no documentation
	ScaleGroupCount uint64 `json:"scaleGroupCount,omitempty"`

	// ActiveAddresses - no documentation
	ActiveAddresses []*SoftLayer_Account_Address `json:"activeAddresses,omitempty"`

	// HardwareWithMcafeeCount - A count of all hardware associated with an account that has McAfee Secure
	// software components.
	HardwareWithMcafeeCount uint64 `json:"hardwareWithMcafeeCount,omitempty"`

	// NasNetworkStorageCount - A count of an account's associated NAS storage volumes.
	NasNetworkStorageCount uint64 `json:"nasNetworkStorageCount,omitempty"`

	// NetworkStorageGroupCount - no documentation
	NetworkStorageGroupCount uint64 `json:"networkStorageGroupCount,omitempty"`

	// OwnedBrandCount - no documentation
	OwnedBrandCount uint64 `json:"ownedBrandCount,omitempty"`

	// NetworkMonitorDownHardware - Hardware which is currently experiencing a service failure.
	NetworkMonitorDownHardware []*SoftLayer_Hardware `json:"networkMonitorDownHardware,omitempty"`

	// PrivateIpAddressCount - no documentation
	PrivateIpAddressCount uint64 `json:"privateIpAddressCount,omitempty"`

	// DatacentersWithSubnetAllocationCount - A count of datacenters which contain subnets that the account
	// has access to route.
	DatacentersWithSubnetAllocationCount uint64 `json:"datacentersWithSubnetAllocationCount,omitempty"`

	// DomainsWithoutSecondaryDnsRecordCount - A count of the DNS domains associated with an account that
	// were not created as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecordCount uint64 `json:"domainsWithoutSecondaryDnsRecordCount,omitempty"`

	// SshKeyCount - A count of customer specified SSH keys that can be implemented onto a newly
	// provisioned or reloaded server.
	SshKeyCount uint64 `json:"sshKeyCount,omitempty"`

	// OpenTickets - no documentation
	OpenTickets []*SoftLayer_Ticket `json:"openTickets,omitempty"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// NetworkMonitorUpVirtualGuestCount - A count of virtual guest which is currently online.
	NetworkMonitorUpVirtualGuestCount uint64 `json:"networkMonitorUpVirtualGuestCount,omitempty"`

	// ActiveAccountDiscountBillingItem - The billing item associated with an account's monthly discount.
	ActiveAccountDiscountBillingItem *SoftLayer_Billing_Item `json:"activeAccountDiscountBillingItem,omitempty"`

	// PermissionGroupCount - no documentation
	PermissionGroupCount uint64 `json:"permissionGroupCount,omitempty"`

	// ReferredAccountCount - A count of if this is a account is a referral partner, the accounts this
	// referral partner has referred
	ReferredAccountCount uint64 `json:"referredAccountCount,omitempty"`

	// OpenCancellationRequests - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationRequests []*SoftLayer_Billing_Item_Cancellation_Request `json:"openCancellationRequests,omitempty"`

	// PendingInvoiceTotalAmount - The total amount of an account's pending invoice, if one exists.
	PendingInvoiceTotalAmount slapi.Float64 `json:"pendingInvoiceTotalAmount,omitempty"`

	// SupportSubscriptions - no documentation
	SupportSubscriptions []*SoftLayer_Billing_Item `json:"supportSubscriptions,omitempty"`

	// SupportSubscriptionCount - A count of the active support subscriptions for this account.
	SupportSubscriptionCount uint64 `json:"supportSubscriptionCount,omitempty"`

	// ActiveCatalystEnrollment - <nil>
	ActiveCatalystEnrollment *SoftLayer_Catalyst_Enrollment `json:"activeCatalystEnrollment,omitempty"`

	// PendingEventCount - no documentation
	PendingEventCount uint64 `json:"pendingEventCount,omitempty"`

	// VirtualGuestsOverBandwidthAllocationCount - A count of an account's associated virtual guest objects
	// currently over bandwidth allocation.
	VirtualGuestsOverBandwidthAllocationCount uint64 `json:"virtualGuestsOverBandwidthAllocationCount,omitempty"`

	// AllRecurringTopLevelBillingItemsUnfiltered - The billing items that will be on an account's next
	// invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfiltered []*SoftLayer_Billing_Item `json:"allRecurringTopLevelBillingItemsUnfiltered,omitempty"`

	// NextInvoiceTotalAmount - The pre-tax total amount of an account's next invoice measured in US
	// Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalAmount slapi.Float64 `json:"nextInvoiceTotalAmount,omitempty"`

	// PendingEvents - <nil>
	PendingEvents []*SoftLayer_Notification_Occurrence_Event `json:"pendingEvents,omitempty"`

	// HardwareWithMcafeeAntivirusRedhatCount - A count of all hardware associated with an account that has
	// McAfee Secure AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhatCount uint64 `json:"hardwareWithMcafeeAntivirusRedhatCount,omitempty"`

	// NetworkVlanCount - A count of all network VLANs assigned to an account.
	NetworkVlanCount uint64 `json:"networkVlanCount,omitempty"`

	// PublicAllotmentHardwareBandwidthDetails - - This information can be pulled directly through tapping
	// keys now - The allotments for this account and their servers. The public inbound and outbound
	// bandwidth is calculated for each server in addition to the daily average network traffic since the
	// last billing date.
	PublicAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"publicAllotmentHardwareBandwidthDetails,omitempty"`

	// AllSubnetBillingItemCount - A count of the billing items that will be on an account's next invoice.
	AllSubnetBillingItemCount uint64 `json:"allSubnetBillingItemCount,omitempty"`

	// BandwidthAllotmentCount - A count of the bandwidth allotments for an account.
	BandwidthAllotmentCount uint64 `json:"bandwidthAllotmentCount,omitempty"`

	// DomainsWithoutSecondaryDnsRecords - The DNS domains associated with an account that were not created
	// as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecords []*SoftLayer_Dns_Domain `json:"domainsWithoutSecondaryDnsRecords,omitempty"`

	// ActiveBillingAgreements - no documentation
	ActiveBillingAgreements []*SoftLayer_Account_Agreement `json:"activeBillingAgreements,omitempty"`

	// NetworkVlanSpan - Whether or not an account has automatic private spanning enabled.
	NetworkVlanSpan *SoftLayer_Account_Network_Vlan_Span `json:"networkVlanSpan,omitempty"`

	// FacilityLogCount - A count of logs of who entered a colocation area which is assigned to this
	// account, or when a user under this account enters a datacenter.
	FacilityLogCount uint64 `json:"facilityLogCount,omitempty"`

	// VirtualDedicatedRackCount - no documentation
	VirtualDedicatedRackCount uint64 `json:"virtualDedicatedRackCount,omitempty"`

	// AllTopLevelBillingItems - no documentation
	AllTopLevelBillingItems []*SoftLayer_Billing_Item `json:"allTopLevelBillingItems,omitempty"`

	// Invoices - no documentation
	Invoices []*SoftLayer_Billing_Invoice `json:"invoices,omitempty"`

	// SecondaryDomainCount - A count of the secondary DNS records for a SoftLayer customer account.
	SecondaryDomainCount uint64 `json:"secondaryDomainCount,omitempty"`

	// VirtualGuestsWithUrchinCount - A count of all virtual guests associated with an account that has the
	// Urchin web traffic analytics package installed.
	VirtualGuestsWithUrchinCount uint64 `json:"virtualGuestsWithUrchinCount,omitempty"`

	// NetworkTunnelContexts - no documentation
	NetworkTunnelContexts []*SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContexts,omitempty"`

	// GlobalLoadBalancerAccountCount - A count of the global load balancer accounts for a softlayer
	// customer account.
	GlobalLoadBalancerAccountCount uint64 `json:"globalLoadBalancerAccountCount,omitempty"`

	// NextInvoiceTopLevelBillingItemCount - A count of the billing items that will be on an account's next
	// invoice.
	NextInvoiceTopLevelBillingItemCount uint64 `json:"nextInvoiceTopLevelBillingItemCount,omitempty"`

	// OpenTicketCount - A count of all open tickets associated with an account.
	OpenTicketCount uint64 `json:"openTicketCount,omitempty"`

	// VirtualGuestsWithMcafeeAntivirusWindows - All virtual guests associated with an account that has
	// McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindows []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusWindows,omitempty"`

	// AbuseEmailCount - A count of an email address that is responsible for abuse and legal inquiries on
	// behalf of an account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmailCount uint64 `json:"abuseEmailCount,omitempty"`

	// AddressCount - A count of all the address(es) that belong to an account.
	AddressCount uint64 `json:"addressCount,omitempty"`

	// HourlyBareMetalInstanceCount - A count of an account's associated hourly bare metal server objects.
	HourlyBareMetalInstanceCount uint64 `json:"hourlyBareMetalInstanceCount,omitempty"`

	// HardwareWithMcafeeAntivirusWindows - All hardware associated with an account that has McAfee Secure
	// AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindows []*SoftLayer_Hardware `json:"hardwareWithMcafeeAntivirusWindows,omitempty"`

	// TicketsClosedInTheLastThreeDays - Tickets closed within the last 72 hours or last 10 tickets,
	// whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDays []*SoftLayer_Ticket `json:"ticketsClosedInTheLastThreeDays,omitempty"`

	// VirtualGuestsProjectedOverBandwidthAllocation - An account's associated virtual guest objects
	// currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocation []*SoftLayer_Virtual_Guest `json:"virtualGuestsProjectedOverBandwidthAllocation,omitempty"`

	// InternalNotes - <nil>
	InternalNotes []*SoftLayer_Account_Note `json:"internalNotes,omitempty"`

	// PrivateBlockDeviceTemplateGroups - Private and shared template group objects (parent only) for an
	// account.
	PrivateBlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"privateBlockDeviceTemplateGroups,omitempty"`

	// ActiveNotificationSubscriberCount - no documentation
	ActiveNotificationSubscriberCount uint64 `json:"activeNotificationSubscriberCount,omitempty"`

	// PermissionGroups - no documentation
	PermissionGroups []*SoftLayer_User_Permission_Group `json:"permissionGroups,omitempty"`

	// RecentEvents - <nil>
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents,omitempty"`

	// DisplaySupportRepresentativeAssignmentCount - A count of the SoftLayer employees that an account is
	// assigned to.
	DisplaySupportRepresentativeAssignmentCount uint64 `json:"displaySupportRepresentativeAssignmentCount,omitempty"`

	// HardwareOverBandwidthAllocationCount - A count of an account's associated hardware objects currently
	// over bandwidth allocation.
	HardwareOverBandwidthAllocationCount uint64 `json:"hardwareOverBandwidthAllocationCount,omitempty"`

	// OpenOtherTicketCount - A count of the open tickets that do not belong to the abuse, accounting,
	// sales, or support groups associated with an account.
	OpenOtherTicketCount uint64 `json:"openOtherTicketCount,omitempty"`

	// BrandAccountFlag - <nil>
	BrandAccountFlag bool `json:"brandAccountFlag,omitempty"`

	// VirtualGuestsWithPleskCount - A count of all virtual guests associated with an account that has the
	// Plesk web hosting control panel installed.
	VirtualGuestsWithPleskCount uint64 `json:"virtualGuestsWithPleskCount,omitempty"`

	// HardwareWithQuantastor - All hardware associated with an account that has the QuantaStor storage
	// system installed.
	HardwareWithQuantastor []*SoftLayer_Hardware `json:"hardwareWithQuantastor,omitempty"`

	// Subnets - no documentation
	Subnets []*SoftLayer_Network_Subnet `json:"subnets,omitempty"`

	// VirtualStoragePublicRepositories - An account's associated virtual server public storage
	// repositories.
	VirtualStoragePublicRepositories []*SoftLayer_Virtual_Storage_Repository `json:"virtualStoragePublicRepositories,omitempty"`

	// NetworkStorageGroups - no documentation
	NetworkStorageGroups []*SoftLayer_Network_Storage_Group `json:"networkStorageGroups,omitempty"`

	// ExpiredSecurityCertificateCount - A count of stored security certificates that are expired (ie.
	ExpiredSecurityCertificateCount uint64 `json:"expiredSecurityCertificateCount,omitempty"`

	// StandardPoolVirtualGuestCount - A count of an account's virtual guest objects that are hosted on a
	// user provisioned hypervisor.
	StandardPoolVirtualGuestCount uint64 `json:"standardPoolVirtualGuestCount,omitempty"`

	// RemoteManagementCommandRequests - no documentation
	RemoteManagementCommandRequests []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"remoteManagementCommandRequests,omitempty"`

	// OpenAccountingTicketCount - A count of the open accounting tickets associated with an account.
	OpenAccountingTicketCount uint64 `json:"openAccountingTicketCount,omitempty"`

	// NetworkMonitorUpHardware - no documentation
	NetworkMonitorUpHardware []*SoftLayer_Hardware `json:"networkMonitorUpHardware,omitempty"`

	// HardwareProjectedOverBandwidthAllocationCount - A count of an account's associated hardware objects
	// projected to go over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocationCount uint64 `json:"hardwareProjectedOverBandwidthAllocationCount,omitempty"`

	// InternalNoteCount - no documentation
	InternalNoteCount uint64 `json:"internalNoteCount,omitempty"`

	// EvaultCapacityGB - The total capacity of Legacy EVault Volumes on an account, in
	EvaultCapacityGB uint `json:"evaultCapacityGB,omitempty"`

	// SupportRepresentatives - The SoftLayer employees that an account is assigned to.
	SupportRepresentatives []*SoftLayer_User_Employee `json:"supportRepresentatives,omitempty"`

	// ApplicationDeliveryControllerCount - A count of an account's associated application delivery
	// controller records.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount,omitempty"`

	// LoadBalancers - no documentation
	LoadBalancers []*SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers,omitempty"`

	// MediaDataTransferRequests - no documentation
	MediaDataTransferRequests []*SoftLayer_Account_Media_Data_Transfer_Request `json:"mediaDataTransferRequests,omitempty"`

	// Carts - no documentation
	Carts []*SoftLayer_Billing_Order_Quote `json:"carts,omitempty"`

	// OpenOtherTickets - The open tickets that do not belong to the abuse, accounting, sales, or support
	// groups associated with an account.
	OpenOtherTickets []*SoftLayer_Ticket `json:"openOtherTickets,omitempty"`

	// PendingInvoiceTotalRecurringAmount - The total recurring amount of an account's pending invoice, if
	// one exists.
	PendingInvoiceTotalRecurringAmount slapi.Float64 `json:"pendingInvoiceTotalRecurringAmount,omitempty"`

	// VirtualGuestsWithPlesk - All virtual guests associated with an account that has the Plesk web
	// hosting control panel installed.
	VirtualGuestsWithPlesk []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithPlesk,omitempty"`

	// AllCommissionBillingItems - The billing items that will be on an account's next invoice.
	AllCommissionBillingItems []*SoftLayer_Billing_Item `json:"allCommissionBillingItems,omitempty"`

	// DatacentersWithSubnetAllocations - Datacenters which contain subnets that the account has access to
	// route.
	DatacentersWithSubnetAllocations []*SoftLayer_Location `json:"datacentersWithSubnetAllocations,omitempty"`

	// Domains - no documentation
	Domains []*SoftLayer_Dns_Domain `json:"domains,omitempty"`

	// PrivateSubnets - no documentation
	PrivateSubnets []*SoftLayer_Network_Subnet `json:"privateSubnets,omitempty"`

	// VirtualStorageArchiveRepositories - An account's associated virtual server archived storage
	// repositories.
	VirtualStorageArchiveRepositories []*SoftLayer_Virtual_Storage_Repository `json:"virtualStorageArchiveRepositories,omitempty"`

	// AllRecurringTopLevelBillingItemCount - A count of the billing items that will be on an account's
	// next invoice.
	AllRecurringTopLevelBillingItemCount uint64 `json:"allRecurringTopLevelBillingItemCount,omitempty"`

	// PrivateNetworkVlanCount - A count of the private network VLANs assigned to an account.
	PrivateNetworkVlanCount uint64 `json:"privateNetworkVlanCount,omitempty"`

	// NextInvoiceTotalOneTimeTaxAmount - The total one-time tax amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeTaxAmount slapi.Float64 `json:"nextInvoiceTotalOneTimeTaxAmount,omitempty"`

	// LastFiveClosedTicketCount - A count of the five most recently closed tickets associated with an
	// account.
	LastFiveClosedTicketCount uint64 `json:"lastFiveClosedTicketCount,omitempty"`

	// PublicSubnetCount - A count of all public network subnets associated with an account.
	PublicSubnetCount uint64 `json:"publicSubnetCount,omitempty"`

	// ResourceGroupCount - A count of an account's associated top-level resource groups.
	ResourceGroupCount uint64 `json:"resourceGroupCount,omitempty"`

	// SubnetRegistrations - <nil>
	SubnetRegistrations []*SoftLayer_Network_Subnet_Registration `json:"subnetRegistrations,omitempty"`

	// VdrUpdatesInProgressFlag - Return 0 if vpn updates are currently in progress on this account
	// otherwise 1.
	VdrUpdatesInProgressFlag bool `json:"vdrUpdatesInProgressFlag,omitempty"`

	// SubnetCount - A count of all network subnets associated with an account.
	SubnetCount uint64 `json:"subnetCount,omitempty"`

	// ActiveVirtualLicenses - The virtual software licenses controlled by an account
	ActiveVirtualLicenses []*SoftLayer_Software_VirtualLicense `json:"activeVirtualLicenses,omitempty"`

	// AllSubnetBillingItems - The billing items that will be on an account's next invoice.
	AllSubnetBillingItems []*SoftLayer_Billing_Item `json:"allSubnetBillingItems,omitempty"`

	// AllCommissionBillingItemCount - A count of the billing items that will be on an account's next
	// invoice.
	AllCommissionBillingItemCount uint64 `json:"allCommissionBillingItemCount,omitempty"`

	// HardwareWithUrchinCount - A count of all hardware associated with an account that has the Urchin web
	// traffic analytics package installed.
	HardwareWithUrchinCount uint64 `json:"hardwareWithUrchinCount,omitempty"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount,omitempty"`

	// NetworkMonitorRecoveringVirtualGuests - Virtual guest which is currently recovering from a service
	// failure.
	NetworkMonitorRecoveringVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorRecoveringVirtualGuests,omitempty"`

	// GlobalIpv4RecordCount - no documentation
	GlobalIpv4RecordCount uint64 `json:"globalIpv4RecordCount,omitempty"`

	// NetworkMonitorRecoveringHardwareCount - A count of hardware which is currently recovering from a
	// service failure.
	NetworkMonitorRecoveringHardwareCount uint64 `json:"networkMonitorRecoveringHardwareCount,omitempty"`

	// ActiveNotificationSubscribers - <nil>
	ActiveNotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"activeNotificationSubscribers,omitempty"`

	// HardwareWithMcafeeIntrusionDetectionSystem - All hardware associated with an account that has McAfee
	// Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystem []*SoftLayer_Hardware `json:"hardwareWithMcafeeIntrusionDetectionSystem,omitempty"`

	// SecondaryDomains - The secondary DNS records for a SoftLayer customer account.
	SecondaryDomains []*SoftLayer_Dns_Secondary `json:"secondaryDomains,omitempty"`

	// Tags - <nil>
	Tags []*SoftLayer_Tag `json:"tags,omitempty"`

	// VirtualDiskImages - An account's associated virtual server virtual disk images.
	VirtualDiskImages []*SoftLayer_Virtual_Disk_Image `json:"virtualDiskImages,omitempty"`

	// AbuseEmail - An email address that is responsible for abuse and legal inquiries on behalf of an
	// account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmail string `json:"abuseEmail,omitempty"`

	// ReplicationEvents - The Replication events for all Network Storage volumes on an account.
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents,omitempty"`

	// LastFiveClosedOtherTicketCount - A count of the five most recently closed tickets that do not belong
	// to the abuse, accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTicketCount uint64 `json:"lastFiveClosedOtherTicketCount,omitempty"`

	// OpenBillingTicketCount - A count of the open billing tickets associated with an account.
	OpenBillingTicketCount uint64 `json:"openBillingTicketCount,omitempty"`

	// HardwareProjectedOverBandwidthAllocation - An account's associated hardware objects projected to go
	// over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocation []*SoftLayer_Hardware `json:"hardwareProjectedOverBandwidthAllocation,omitempty"`

	// HardwareWithMcafeeAntivirusWindowCount - A count of all hardware associated with an account that has
	// McAfee Secure AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindowCount uint64 `json:"hardwareWithMcafeeAntivirusWindowCount,omitempty"`
}

func (softlayer_account *SoftLayer_Account) String() string {
	return "SoftLayer_Account"
}
