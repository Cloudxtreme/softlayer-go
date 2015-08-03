package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "github.com/sudorandom/softlayer-go/slapi"
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

	// AbuseEmail - An email address that is responsible for abuse and legal inquiries on behalf of an
	// account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmail string `json:"abuseEmail"`

	// AbuseEmailCount - A count of an email address that is responsible for abuse and legal inquiries on
	// behalf of an account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmailCount uint64 `json:"abuseEmailCount"`

	// AbuseEmails - An email address that is responsible for abuse and legal inquiries on behalf of an
	// account. For instance, new legal and abuse tickets are sent to this address.
	AbuseEmails []*SoftLayer_Account_AbuseEmail `json:"abuseEmails"`

	// AccountContactCount - no documentation
	AccountContactCount uint64 `json:"accountContactCount"`

	// AccountContacts - <nil>
	AccountContacts []*SoftLayer_Account_Contact `json:"accountContacts"`

	// AccountLinkCount - no documentation
	AccountLinkCount uint64 `json:"accountLinkCount"`

	// AccountLinks - <nil>
	AccountLinks []*SoftLayer_Account_Link `json:"accountLinks"`

	// AccountManagedResourcesFlag - A flag indicating that the account has a managed resource.
	AccountManagedResourcesFlag bool `json:"accountManagedResourcesFlag"`

	// AccountStatus - An account's status presented in a more detailed data type.
	AccountStatus *SoftLayer_Account_Status `json:"accountStatus"`

	// AccountStatusId - no documentation
	AccountStatusId int `json:"accountStatusId"`

	// ActiveAccountDiscountBillingItem - The billing item associated with an account's monthly discount.
	ActiveAccountDiscountBillingItem *SoftLayer_Billing_Item `json:"activeAccountDiscountBillingItem"`

	// ActiveAddressCount - A count of the active address(es) that belong to an account.
	ActiveAddressCount uint64 `json:"activeAddressCount"`

	// ActiveAddresses - no documentation
	ActiveAddresses []*SoftLayer_Account_Address `json:"activeAddresses"`

	// ActiveBillingAgreementCount - no documentation
	ActiveBillingAgreementCount uint64 `json:"activeBillingAgreementCount"`

	// ActiveBillingAgreements - no documentation
	ActiveBillingAgreements []*SoftLayer_Account_Agreement `json:"activeBillingAgreements"`

	// ActiveCatalystEnrollment - <nil>
	ActiveCatalystEnrollment *SoftLayer_Catalyst_Enrollment `json:"activeCatalystEnrollment"`

	// ActiveColocationContainerCount - A count of the account's active top level colocation containers.
	ActiveColocationContainerCount uint64 `json:"activeColocationContainerCount"`

	// ActiveColocationContainers - The account's active top level colocation containers.
	ActiveColocationContainers []*SoftLayer_Billing_Item `json:"activeColocationContainers"`

	// ActiveFlexibleCreditEnrollment - Account's currently active Flexible Credit enrollment.
	ActiveFlexibleCreditEnrollment *SoftLayer_FlexibleCredit_Enrollment `json:"activeFlexibleCreditEnrollment"`

	// ActiveNotificationSubscriberCount - no documentation
	ActiveNotificationSubscriberCount uint64 `json:"activeNotificationSubscriberCount"`

	// ActiveNotificationSubscribers - <nil>
	ActiveNotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"activeNotificationSubscribers"`

	// ActiveQuoteCount - no documentation
	ActiveQuoteCount uint64 `json:"activeQuoteCount"`

	// ActiveQuotes - no documentation
	ActiveQuotes []*SoftLayer_Billing_Order_Quote `json:"activeQuotes"`

	// ActiveVirtualLicenseCount - A count of the virtual software licenses controlled by an account
	ActiveVirtualLicenseCount uint64 `json:"activeVirtualLicenseCount"`

	// ActiveVirtualLicenses - The virtual software licenses controlled by an account
	ActiveVirtualLicenses []*SoftLayer_Software_VirtualLicense `json:"activeVirtualLicenses"`

	// AdcLoadBalancerCount - no documentation
	AdcLoadBalancerCount uint64 `json:"adcLoadBalancerCount"`

	// AdcLoadBalancers - no documentation
	AdcLoadBalancers []*SoftLayer_Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"adcLoadBalancers"`

	// Address1 - The first line of the mailing address belonging to an account.
	Address1 string `json:"address1"`

	// Address2 - The second line of the mailing address belonging to an account.
	Address2 string `json:"address2"`

	// AddressCount - A count of all the address(es) that belong to an account.
	AddressCount uint64 `json:"addressCount"`

	// Addresses - no documentation
	Addresses []*SoftLayer_Account_Address `json:"addresses"`

	// AffiliateId - An affiliate identifier associated with the customer account.
	AffiliateId string `json:"affiliateId"`

	// AllBillingItems - The billing items that will be on an account's next invoice.
	AllBillingItems []*SoftLayer_Billing_Item `json:"allBillingItems"`

	// AllCommissionBillingItemCount - A count of the billing items that will be on an account's next
	// invoice.
	AllCommissionBillingItemCount uint64 `json:"allCommissionBillingItemCount"`

	// AllCommissionBillingItems - The billing items that will be on an account's next invoice.
	AllCommissionBillingItems []*SoftLayer_Billing_Item `json:"allCommissionBillingItems"`

	// AllRecurringTopLevelBillingItemCount - A count of the billing items that will be on an account's
	// next invoice.
	AllRecurringTopLevelBillingItemCount uint64 `json:"allRecurringTopLevelBillingItemCount"`

	// AllRecurringTopLevelBillingItems - The billing items that will be on an account's next invoice.
	AllRecurringTopLevelBillingItems []*SoftLayer_Billing_Item `json:"allRecurringTopLevelBillingItems"`

	// AllRecurringTopLevelBillingItemsUnfiltered - The billing items that will be on an account's next
	// invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfiltered []*SoftLayer_Billing_Item `json:"allRecurringTopLevelBillingItemsUnfiltered"`

	// AllRecurringTopLevelBillingItemsUnfilteredCount - A count of the billing items that will be on an
	// account's next invoice. Does not consider associated items.
	AllRecurringTopLevelBillingItemsUnfilteredCount uint64 `json:"allRecurringTopLevelBillingItemsUnfilteredCount"`

	// AllSubnetBillingItemCount - A count of the billing items that will be on an account's next invoice.
	AllSubnetBillingItemCount uint64 `json:"allSubnetBillingItemCount"`

	// AllSubnetBillingItems - The billing items that will be on an account's next invoice.
	AllSubnetBillingItems []*SoftLayer_Billing_Item `json:"allSubnetBillingItems"`

	// AllTopLevelBillingItemCount - no documentation
	AllTopLevelBillingItemCount uint64 `json:"allTopLevelBillingItemCount"`

	// AllTopLevelBillingItems - no documentation
	AllTopLevelBillingItems []*SoftLayer_Billing_Item `json:"allTopLevelBillingItems"`

	// AllTopLevelBillingItemsUnfiltered - The billing items that will be on an account's next invoice.
	// Does not consider associated items.
	AllTopLevelBillingItemsUnfiltered []*SoftLayer_Billing_Item `json:"allTopLevelBillingItemsUnfiltered"`

	// AllTopLevelBillingItemsUnfilteredCount - A count of the billing items that will be on an account's
	// next invoice. Does not consider associated items.
	AllTopLevelBillingItemsUnfilteredCount uint64 `json:"allTopLevelBillingItemsUnfilteredCount"`

	// AllowedPptpVpnQuantity - no documentation
	AllowedPptpVpnQuantity int `json:"allowedPptpVpnQuantity"`

	// AlternatePhone - no documentation
	AlternatePhone string `json:"alternatePhone"`

	// ApplicationDeliveryControllerCount - A count of an account's associated application delivery
	// controller records.
	ApplicationDeliveryControllerCount uint64 `json:"applicationDeliveryControllerCount"`

	// ApplicationDeliveryControllers - An account's associated application delivery controller records.
	ApplicationDeliveryControllers []*SoftLayer_Network_Application_Delivery_Controller `json:"applicationDeliveryControllers"`

	// AttributeCount - A count of the account attribute values for a SoftLayer customer account.
	AttributeCount uint64 `json:"attributeCount"`

	// Attributes - The account attribute values for a SoftLayer customer account.
	Attributes []*SoftLayer_Account_Attribute `json:"attributes"`

	// AvailablePublicNetworkVlanCount - A count of the public network VLANs assigned to an account.
	AvailablePublicNetworkVlanCount uint64 `json:"availablePublicNetworkVlanCount"`

	// AvailablePublicNetworkVlans - no documentation
	AvailablePublicNetworkVlans []*SoftLayer_Network_Vlan `json:"availablePublicNetworkVlans"`

	// Balance - The account balance of a SoftLayer customer account. An account's balance is the amount of
	// money owed to SoftLayer by the account holder, returned as a floating point number with two decimal
	// places, measured in US Dollars A negative account balance means the account holder has overpaid and
	// is owed money by SoftLayer.
	Balance float64 `json:"balance"`

	// BandwidthAllotmentCount - A count of the bandwidth allotments for an account.
	BandwidthAllotmentCount uint64 `json:"bandwidthAllotmentCount"`

	// BandwidthAllotments - no documentation
	BandwidthAllotments []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotments"`

	// BandwidthAllotmentsOverAllocation - The bandwidth allotments for an account currently over
	// allocation.
	BandwidthAllotmentsOverAllocation []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsOverAllocation"`

	// BandwidthAllotmentsOverAllocationCount - A count of the bandwidth allotments for an account
	// currently over allocation.
	BandwidthAllotmentsOverAllocationCount uint64 `json:"bandwidthAllotmentsOverAllocationCount"`

	// BandwidthAllotmentsProjectedOverAllocation - The bandwidth allotments for an account projected to go
	// over allocation.
	BandwidthAllotmentsProjectedOverAllocation []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"bandwidthAllotmentsProjectedOverAllocation"`

	// BandwidthAllotmentsProjectedOverAllocationCount - A count of the bandwidth allotments for an account
	// projected to go over allocation.
	BandwidthAllotmentsProjectedOverAllocationCount uint64 `json:"bandwidthAllotmentsProjectedOverAllocationCount"`

	// BareMetalInstanceCount - A count of an account's associated bare metal server objects.
	BareMetalInstanceCount uint64 `json:"bareMetalInstanceCount"`

	// BareMetalInstances - no documentation
	BareMetalInstances []*SoftLayer_Hardware `json:"bareMetalInstances"`

	// BillingAgreementCount - no documentation
	BillingAgreementCount uint64 `json:"billingAgreementCount"`

	// BillingAgreements - no documentation
	BillingAgreements []*SoftLayer_Account_Agreement `json:"billingAgreements"`

	// BillingInfo - no documentation
	BillingInfo *SoftLayer_Billing_Info `json:"billingInfo"`

	// BlockDeviceTemplateGroupCount - A count of private template group objects (parent and children) and
	// the shared template group objects (parent only) for an account.
	BlockDeviceTemplateGroupCount uint64 `json:"blockDeviceTemplateGroupCount"`

	// BlockDeviceTemplateGroups - Private template group objects (parent and children) and the shared
	// template group objects (parent only) for an account.
	BlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups"`

	// Brand - <nil>
	Brand *SoftLayer_Brand `json:"brand"`

	// BrandAccountFlag - <nil>
	BrandAccountFlag bool `json:"brandAccountFlag"`

	// BrandId - no documentation
	BrandId int `json:"brandId"`

	// BrandKeyName - no documentation
	BrandKeyName string `json:"brandKeyName"`

	// CanOrderAdditionalVlansFlag - Indicating whether this account can order additional Vlans.
	CanOrderAdditionalVlansFlag bool `json:"canOrderAdditionalVlansFlag"`

	// CartCount - no documentation
	CartCount uint64 `json:"cartCount"`

	// Carts - no documentation
	Carts []*SoftLayer_Billing_Order_Quote `json:"carts"`

	// CatalystEnrollmentCount - no documentation
	CatalystEnrollmentCount uint64 `json:"catalystEnrollmentCount"`

	// CatalystEnrollments - <nil>
	CatalystEnrollments []*SoftLayer_Catalyst_Enrollment `json:"catalystEnrollments"`

	// CdnAccountCount - no documentation
	CdnAccountCount uint64 `json:"cdnAccountCount"`

	// CdnAccounts - no documentation
	CdnAccounts []*SoftLayer_Network_ContentDelivery_Account `json:"cdnAccounts"`

	// City - The city of the mailing address belonging to an account.
	City string `json:"city"`

	// ClaimedTaxExemptTxFlag - Whether an account is exempt from taxes on their invoices.
	ClaimedTaxExemptTxFlag bool `json:"claimedTaxExemptTxFlag"`

	// ClosedTicketCount - A count of all closed tickets associated with an account.
	ClosedTicketCount uint64 `json:"closedTicketCount"`

	// ClosedTickets - no documentation
	ClosedTickets []*SoftLayer_Ticket `json:"closedTickets"`

	// CompanyName - no documentation
	CompanyName string `json:"companyName"`

	// Country - A two-letter abbreviation of the country in the mailing address belonging to an account.
	Country string `json:"country"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DatacentersWithSubnetAllocationCount - A count of datacenters which contain subnets that the account
	// has access to route.
	DatacentersWithSubnetAllocationCount uint64 `json:"datacentersWithSubnetAllocationCount"`

	// DatacentersWithSubnetAllocations - Datacenters which contain subnets that the account has access to
	// route.
	DatacentersWithSubnetAllocations []*SoftLayer_Location `json:"datacentersWithSubnetAllocations"`

	// DeviceFingerprintId - Device Fingerprint Identifier - Used internally and can safely be ignored.
	DeviceFingerprintId string `json:"deviceFingerprintId"`

	// DisablePaymentProcessingFlag - A flag indicating whether payments are processed for this account.
	DisablePaymentProcessingFlag bool `json:"disablePaymentProcessingFlag"`

	// DisplaySupportRepresentativeAssignmentCount - A count of the SoftLayer employees that an account is
	// assigned to.
	DisplaySupportRepresentativeAssignmentCount uint64 `json:"displaySupportRepresentativeAssignmentCount"`

	// DisplaySupportRepresentativeAssignments - The SoftLayer employees that an account is assigned to.
	DisplaySupportRepresentativeAssignments []*SoftLayer_Account_Attachment_Employee `json:"displaySupportRepresentativeAssignments"`

	// DomainCount - A count of the DNS domains associated with an account.
	DomainCount uint64 `json:"domainCount"`

	// DomainRegistrationCount - no documentation
	DomainRegistrationCount uint64 `json:"domainRegistrationCount"`

	// DomainRegistrations - <nil>
	DomainRegistrations []*SoftLayer_Dns_Domain_Registration `json:"domainRegistrations"`

	// Domains - no documentation
	Domains []*SoftLayer_Dns_Domain `json:"domains"`

	// DomainsWithoutSecondaryDnsRecordCount - A count of the DNS domains associated with an account that
	// were not created as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecordCount uint64 `json:"domainsWithoutSecondaryDnsRecordCount"`

	// DomainsWithoutSecondaryDnsRecords - The DNS domains associated with an account that were not created
	// as a result of a secondary DNS zone transfer.
	DomainsWithoutSecondaryDnsRecords []*SoftLayer_Dns_Domain `json:"domainsWithoutSecondaryDnsRecords"`

	// Email - no documentation
	Email string `json:"email"`

	// EvaultCapacityGB - The total capacity of Legacy EVault Volumes on an account, in
	EvaultCapacityGB uint `json:"evaultCapacityGB"`

	// EvaultMasterUserCount - A count of an account's master EVault user. This is only used when an
	// account has EVault service.
	EvaultMasterUserCount uint64 `json:"evaultMasterUserCount"`

	// EvaultMasterUsers - An account's master EVault user. This is only used when an account has EVault
	// service.
	EvaultMasterUsers []*SoftLayer_Account_Password `json:"evaultMasterUsers"`

	// EvaultNetworkStorage - no documentation
	EvaultNetworkStorage []*SoftLayer_Network_Storage `json:"evaultNetworkStorage"`

	// EvaultNetworkStorageCount - A count of an account's associated EVault storage volumes.
	EvaultNetworkStorageCount uint64 `json:"evaultNetworkStorageCount"`

	// ExpiredSecurityCertificateCount - A count of stored security certificates that are expired (ie.
	ExpiredSecurityCertificateCount uint64 `json:"expiredSecurityCertificateCount"`

	// ExpiredSecurityCertificates - Stored security certificates that are expired (ie.
	ExpiredSecurityCertificates []*SoftLayer_Security_Certificate `json:"expiredSecurityCertificates"`

	// FacilityLogCount - A count of logs of who entered a colocation area which is assigned to this
	// account, or when a user under this account enters a datacenter.
	FacilityLogCount uint64 `json:"facilityLogCount"`

	// FacilityLogs - Logs of who entered a colocation area which is assigned to this account, or when a
	// user under this account enters a datacenter.
	FacilityLogs []*SoftLayer_User_Access_Facility_Log `json:"facilityLogs"`

	// FaxPhone - no documentation
	FaxPhone string `json:"faxPhone"`

	// FirstName - Each customer account is listed under a single individual. This is that individual's
	// first name.
	FirstName string `json:"firstName"`

	// FlexibleCreditEnrollmentCount - A count of all of the account's current and former Flexible Credit
	// enrollments.
	FlexibleCreditEnrollmentCount uint64 `json:"flexibleCreditEnrollmentCount"`

	// FlexibleCreditEnrollments - All of the account's current and former Flexible Credit enrollments.
	FlexibleCreditEnrollments []*SoftLayer_FlexibleCredit_Enrollment `json:"flexibleCreditEnrollments"`

	// GlobalIpRecordCount - no documentation
	GlobalIpRecordCount uint64 `json:"globalIpRecordCount"`

	// GlobalIpRecords - <nil>
	GlobalIpRecords []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpRecords"`

	// GlobalIpv4RecordCount - no documentation
	GlobalIpv4RecordCount uint64 `json:"globalIpv4RecordCount"`

	// GlobalIpv4Records - <nil>
	GlobalIpv4Records []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpv4Records"`

	// GlobalIpv6RecordCount - no documentation
	GlobalIpv6RecordCount uint64 `json:"globalIpv6RecordCount"`

	// GlobalIpv6Records - <nil>
	GlobalIpv6Records []*SoftLayer_Network_Subnet_IpAddress_Global `json:"globalIpv6Records"`

	// GlobalLoadBalancerAccountCount - A count of the global load balancer accounts for a softlayer
	// customer account.
	GlobalLoadBalancerAccountCount uint64 `json:"globalLoadBalancerAccountCount"`

	// GlobalLoadBalancerAccounts - The global load balancer accounts for a softlayer customer account.
	GlobalLoadBalancerAccounts []*SoftLayer_Network_LoadBalancer_Global_Account `json:"globalLoadBalancerAccounts"`

	// Hardware - no documentation
	Hardware []*SoftLayer_Hardware `json:"hardware"`

	// HardwareCount - A count of an account's associated hardware objects.
	HardwareCount uint64 `json:"hardwareCount"`

	// HardwareOverBandwidthAllocation - An account's associated hardware objects currently over bandwidth
	// allocation.
	HardwareOverBandwidthAllocation []*SoftLayer_Hardware `json:"hardwareOverBandwidthAllocation"`

	// HardwareOverBandwidthAllocationCount - A count of an account's associated hardware objects currently
	// over bandwidth allocation.
	HardwareOverBandwidthAllocationCount uint64 `json:"hardwareOverBandwidthAllocationCount"`

	// HardwareProjectedOverBandwidthAllocation - An account's associated hardware objects projected to go
	// over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocation []*SoftLayer_Hardware `json:"hardwareProjectedOverBandwidthAllocation"`

	// HardwareProjectedOverBandwidthAllocationCount - A count of an account's associated hardware objects
	// projected to go over bandwidth allocation.
	HardwareProjectedOverBandwidthAllocationCount uint64 `json:"hardwareProjectedOverBandwidthAllocationCount"`

	// HardwareWithCpanel - All hardware associated with an account that has the cPanel web hosting control
	// panel installed.
	HardwareWithCpanel []*SoftLayer_Hardware `json:"hardwareWithCpanel"`

	// HardwareWithCpanelCount - A count of all hardware associated with an account that has the cPanel web
	// hosting control panel installed.
	HardwareWithCpanelCount uint64 `json:"hardwareWithCpanelCount"`

	// HardwareWithHelm - All hardware associated with an account that has the Helm web hosting control
	// panel installed.
	HardwareWithHelm []*SoftLayer_Hardware `json:"hardwareWithHelm"`

	// HardwareWithHelmCount - A count of all hardware associated with an account that has the Helm web
	// hosting control panel installed.
	HardwareWithHelmCount uint64 `json:"hardwareWithHelmCount"`

	// HardwareWithMcafee - All hardware associated with an account that has McAfee Secure software
	// components.
	HardwareWithMcafee []*SoftLayer_Hardware `json:"hardwareWithMcafee"`

	// HardwareWithMcafeeAntivirusRedhat - All hardware associated with an account that has McAfee Secure
	// AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhat []*SoftLayer_Hardware `json:"hardwareWithMcafeeAntivirusRedhat"`

	// HardwareWithMcafeeAntivirusRedhatCount - A count of all hardware associated with an account that has
	// McAfee Secure AntiVirus for Redhat software components.
	HardwareWithMcafeeAntivirusRedhatCount uint64 `json:"hardwareWithMcafeeAntivirusRedhatCount"`

	// HardwareWithMcafeeAntivirusWindowCount - A count of all hardware associated with an account that has
	// McAfee Secure AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindowCount uint64 `json:"hardwareWithMcafeeAntivirusWindowCount"`

	// HardwareWithMcafeeAntivirusWindows - All hardware associated with an account that has McAfee Secure
	// AntiVirus for Windows software components.
	HardwareWithMcafeeAntivirusWindows []*SoftLayer_Hardware `json:"hardwareWithMcafeeAntivirusWindows"`

	// HardwareWithMcafeeCount - A count of all hardware associated with an account that has McAfee Secure
	// software components.
	HardwareWithMcafeeCount uint64 `json:"hardwareWithMcafeeCount"`

	// HardwareWithMcafeeIntrusionDetectionSystem - All hardware associated with an account that has McAfee
	// Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystem []*SoftLayer_Hardware `json:"hardwareWithMcafeeIntrusionDetectionSystem"`

	// HardwareWithMcafeeIntrusionDetectionSystemCount - A count of all hardware associated with an account
	// that has McAfee Secure Intrusion Detection System software components.
	HardwareWithMcafeeIntrusionDetectionSystemCount uint64 `json:"hardwareWithMcafeeIntrusionDetectionSystemCount"`

	// HardwareWithPlesk - All hardware associated with an account that has the Plesk web hosting control
	// panel installed.
	HardwareWithPlesk []*SoftLayer_Hardware `json:"hardwareWithPlesk"`

	// HardwareWithPleskCount - A count of all hardware associated with an account that has the Plesk web
	// hosting control panel installed.
	HardwareWithPleskCount uint64 `json:"hardwareWithPleskCount"`

	// HardwareWithQuantastor - All hardware associated with an account that has the QuantaStor storage
	// system installed.
	HardwareWithQuantastor []*SoftLayer_Hardware `json:"hardwareWithQuantastor"`

	// HardwareWithQuantastorCount - A count of all hardware associated with an account that has the
	// QuantaStor storage system installed.
	HardwareWithQuantastorCount uint64 `json:"hardwareWithQuantastorCount"`

	// HardwareWithUrchin - All hardware associated with an account that has the Urchin web traffic
	// analytics package installed.
	HardwareWithUrchin []*SoftLayer_Hardware `json:"hardwareWithUrchin"`

	// HardwareWithUrchinCount - A count of all hardware associated with an account that has the Urchin web
	// traffic analytics package installed.
	HardwareWithUrchinCount uint64 `json:"hardwareWithUrchinCount"`

	// HardwareWithWindowCount - A count of all hardware associated with an account that is running a
	// version of the Microsoft Windows operating system.
	HardwareWithWindowCount uint64 `json:"hardwareWithWindowCount"`

	// HardwareWithWindows - All hardware associated with an account that is running a version of the
	// Microsoft Windows operating system.
	HardwareWithWindows []*SoftLayer_Hardware `json:"hardwareWithWindows"`

	// HasEvaultBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has the EVault Bare
	// Metal Server Restore Plugin otherwise 0.
	HasEvaultBareMetalRestorePluginFlag bool `json:"hasEvaultBareMetalRestorePluginFlag"`

	// HasIderaBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has an installation
	// of Idera Server Backup otherwise 0.
	HasIderaBareMetalRestorePluginFlag bool `json:"hasIderaBareMetalRestorePluginFlag"`

	// HasPendingOrder - The number of orders in a status for a SoftLayer customer account.
	HasPendingOrder uint `json:"hasPendingOrder"`

	// HasR1softBareMetalRestorePluginFlag - Return 1 if one of the account's hardware has an installation
	// of R1Soft CDP otherwise 0.
	HasR1softBareMetalRestorePluginFlag bool `json:"hasR1softBareMetalRestorePluginFlag"`

	// HourlyBareMetalInstanceCount - A count of an account's associated hourly bare metal server objects.
	HourlyBareMetalInstanceCount uint64 `json:"hourlyBareMetalInstanceCount"`

	// HourlyBareMetalInstances - An account's associated hourly bare metal server objects.
	HourlyBareMetalInstances []*SoftLayer_Hardware `json:"hourlyBareMetalInstances"`

	// HourlyServiceBillingItemCount - A count of hourly service billing items that will be on an account's
	// next invoice.
	HourlyServiceBillingItemCount uint64 `json:"hourlyServiceBillingItemCount"`

	// HourlyServiceBillingItems - Hourly service billing items that will be on an account's next invoice.
	HourlyServiceBillingItems []*SoftLayer_Billing_Item `json:"hourlyServiceBillingItems"`

	// HourlyVirtualGuestCount - A count of an account's associated hourly virtual guest objects.
	HourlyVirtualGuestCount uint64 `json:"hourlyVirtualGuestCount"`

	// HourlyVirtualGuests - An account's associated hourly virtual guest objects.
	HourlyVirtualGuests []*SoftLayer_Virtual_Guest `json:"hourlyVirtualGuests"`

	// HubNetworkStorage - no documentation
	HubNetworkStorage []*SoftLayer_Network_Storage `json:"hubNetworkStorage"`

	// HubNetworkStorageCount - A count of an account's associated Virtual Storage volumes.
	HubNetworkStorageCount uint64 `json:"hubNetworkStorageCount"`

	// Id - A customer account's internal identifier. Account numbers are typically preceded by the string
	// in the customer portal. Every SoftLayer account has at least one portal user whose username follows
	// the + account number naming scheme.
	Id int `json:"id"`

	// InternalNoteCount - no documentation
	InternalNoteCount uint64 `json:"internalNoteCount"`

	// InternalNotes - <nil>
	InternalNotes []*SoftLayer_Account_Note `json:"internalNotes"`

	// InvoiceCount - A count of an account's associated billing invoices.
	InvoiceCount uint64 `json:"invoiceCount"`

	// Invoices - no documentation
	Invoices []*SoftLayer_Billing_Invoice `json:"invoices"`

	// IpAddressCount - no documentation
	IpAddressCount uint64 `json:"ipAddressCount"`

	// IpAddresses - <nil>
	IpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"ipAddresses"`

	// IsReseller - A flag indicating if an account belongs to a reseller or not.
	IsReseller int `json:"isReseller"`

	// IscsiNetworkStorage - no documentation
	IscsiNetworkStorage []*SoftLayer_Network_Storage `json:"iscsiNetworkStorage"`

	// IscsiNetworkStorageCount - A count of an account's associated iSCSI storage volumes.
	IscsiNetworkStorageCount uint64 `json:"iscsiNetworkStorageCount"`

	// LastCanceledBillingItem - no documentation
	LastCanceledBillingItem *SoftLayer_Billing_Item `json:"lastCanceledBillingItem"`

	// LastCancelledServerBillingItem - no documentation
	LastCancelledServerBillingItem *SoftLayer_Billing_Item `json:"lastCancelledServerBillingItem"`

	// LastFiveClosedAbuseTicketCount - A count of the five most recently closed abuse tickets associated
	// with an account.
	LastFiveClosedAbuseTicketCount uint64 `json:"lastFiveClosedAbuseTicketCount"`

	// LastFiveClosedAbuseTickets - The five most recently closed abuse tickets associated with an account.
	LastFiveClosedAbuseTickets []*SoftLayer_Ticket `json:"lastFiveClosedAbuseTickets"`

	// LastFiveClosedAccountingTicketCount - A count of the five most recently closed accounting tickets
	// associated with an account.
	LastFiveClosedAccountingTicketCount uint64 `json:"lastFiveClosedAccountingTicketCount"`

	// LastFiveClosedAccountingTickets - The five most recently closed accounting tickets associated with
	// an account.
	LastFiveClosedAccountingTickets []*SoftLayer_Ticket `json:"lastFiveClosedAccountingTickets"`

	// LastFiveClosedOtherTicketCount - A count of the five most recently closed tickets that do not belong
	// to the abuse, accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTicketCount uint64 `json:"lastFiveClosedOtherTicketCount"`

	// LastFiveClosedOtherTickets - The five most recently closed tickets that do not belong to the abuse,
	// accounting, sales, or support groups associated with an account.
	LastFiveClosedOtherTickets []*SoftLayer_Ticket `json:"lastFiveClosedOtherTickets"`

	// LastFiveClosedSalesTicketCount - A count of the five most recently closed sales tickets associated
	// with an account.
	LastFiveClosedSalesTicketCount uint64 `json:"lastFiveClosedSalesTicketCount"`

	// LastFiveClosedSalesTickets - The five most recently closed sales tickets associated with an account.
	LastFiveClosedSalesTickets []*SoftLayer_Ticket `json:"lastFiveClosedSalesTickets"`

	// LastFiveClosedSupportTicketCount - A count of the five most recently closed support tickets
	// associated with an account.
	LastFiveClosedSupportTicketCount uint64 `json:"lastFiveClosedSupportTicketCount"`

	// LastFiveClosedSupportTickets - The five most recently closed support tickets associated with an
	// account.
	LastFiveClosedSupportTickets []*SoftLayer_Ticket `json:"lastFiveClosedSupportTickets"`

	// LastFiveClosedTicketCount - A count of the five most recently closed tickets associated with an
	// account.
	LastFiveClosedTicketCount uint64 `json:"lastFiveClosedTicketCount"`

	// LastFiveClosedTickets - The five most recently closed tickets associated with an account.
	LastFiveClosedTickets []*SoftLayer_Ticket `json:"lastFiveClosedTickets"`

	// LastName - Each customer account is listed under a single individual. This is that individual's last
	// name.
	LastName string `json:"lastName"`

	// LateFeeProtectionFlag - no documentation
	LateFeeProtectionFlag bool `json:"lateFeeProtectionFlag"`

	// LatestBillDate - no documentation
	LatestBillDate *time.Time `json:"latestBillDate"`

	// LatestRecurringInvoice - no documentation
	LatestRecurringInvoice *SoftLayer_Billing_Invoice `json:"latestRecurringInvoice"`

	// LatestRecurringPendingInvoice - no documentation
	LatestRecurringPendingInvoice *SoftLayer_Billing_Invoice `json:"latestRecurringPendingInvoice"`

	// LegacyBandwidthAllotmentCount - A count of the legacy bandwidth allotments for an account.
	LegacyBandwidthAllotmentCount uint64 `json:"legacyBandwidthAllotmentCount"`

	// LegacyBandwidthAllotments - no documentation
	LegacyBandwidthAllotments []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"legacyBandwidthAllotments"`

	// LegacyIscsiCapacityGB - The total capacity of Legacy iSCSI Volumes on an account, in
	LegacyIscsiCapacityGB uint `json:"legacyIscsiCapacityGB"`

	// LoadBalancerCount - no documentation
	LoadBalancerCount uint64 `json:"loadBalancerCount"`

	// LoadBalancers - no documentation
	LoadBalancers []*SoftLayer_Network_LoadBalancer_VirtualIpAddress `json:"loadBalancers"`

	// LockboxCapacityGB - The total capacity of Legacy lockbox Volumes on an account, in
	LockboxCapacityGB uint `json:"lockboxCapacityGB"`

	// LockboxNetworkStorage - no documentation
	LockboxNetworkStorage []*SoftLayer_Network_Storage `json:"lockboxNetworkStorage"`

	// LockboxNetworkStorageCount - A count of an account's associated Lockbox storage volumes.
	LockboxNetworkStorageCount uint64 `json:"lockboxNetworkStorageCount"`

	// ManualPaymentsUnderReview - <nil>
	ManualPaymentsUnderReview []*SoftLayer_Billing_Payment_Card_ManualPayment `json:"manualPaymentsUnderReview"`

	// ManualPaymentsUnderReviewCount - no documentation
	ManualPaymentsUnderReviewCount uint64 `json:"manualPaymentsUnderReviewCount"`

	// MasterUser - no documentation
	MasterUser *SoftLayer_User_Customer `json:"masterUser"`

	// MediaDataTransferRequestCount - A count of an account's media transfer service requests.
	MediaDataTransferRequestCount uint64 `json:"mediaDataTransferRequestCount"`

	// MediaDataTransferRequests - no documentation
	MediaDataTransferRequests []*SoftLayer_Account_Media_Data_Transfer_Request `json:"mediaDataTransferRequests"`

	// MessageQueueAccountCount - A count of an account's associated Message Queue accounts.
	MessageQueueAccountCount uint64 `json:"messageQueueAccountCount"`

	// MessageQueueAccounts - no documentation
	MessageQueueAccounts []*SoftLayer_Network_Message_Queue `json:"messageQueueAccounts"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// MonthlyBareMetalInstanceCount - A count of an account's associated monthly bare metal server
	// objects.
	MonthlyBareMetalInstanceCount uint64 `json:"monthlyBareMetalInstanceCount"`

	// MonthlyBareMetalInstances - An account's associated monthly bare metal server objects.
	MonthlyBareMetalInstances []*SoftLayer_Hardware `json:"monthlyBareMetalInstances"`

	// MonthlyVirtualGuestCount - A count of an account's associated monthly virtual guest objects.
	MonthlyVirtualGuestCount uint64 `json:"monthlyVirtualGuestCount"`

	// MonthlyVirtualGuests - An account's associated monthly virtual guest objects.
	MonthlyVirtualGuests []*SoftLayer_Virtual_Guest `json:"monthlyVirtualGuests"`

	// NasNetworkStorage - no documentation
	NasNetworkStorage []*SoftLayer_Network_Storage `json:"nasNetworkStorage"`

	// NasNetworkStorageCount - A count of an account's associated NAS storage volumes.
	NasNetworkStorageCount uint64 `json:"nasNetworkStorageCount"`

	// NetworkGatewayCount - A count of all network gateway devices on this account.
	NetworkGatewayCount uint64 `json:"networkGatewayCount"`

	// NetworkGateways - no documentation
	NetworkGateways []*SoftLayer_Network_Gateway `json:"networkGateways"`

	// NetworkHardware - no documentation
	NetworkHardware []*SoftLayer_Hardware `json:"networkHardware"`

	// NetworkHardwareCount - A count of an account's associated network hardware.
	NetworkHardwareCount uint64 `json:"networkHardwareCount"`

	// NetworkMessageDeliveryAccountCount - no documentation
	NetworkMessageDeliveryAccountCount uint64 `json:"networkMessageDeliveryAccountCount"`

	// NetworkMessageDeliveryAccounts - <nil>
	NetworkMessageDeliveryAccounts []*SoftLayer_Network_Message_Delivery `json:"networkMessageDeliveryAccounts"`

	// NetworkMonitorDownHardware - Hardware which is currently experiencing a service failure.
	NetworkMonitorDownHardware []*SoftLayer_Hardware `json:"networkMonitorDownHardware"`

	// NetworkMonitorDownHardwareCount - A count of hardware which is currently experiencing a service
	// failure.
	NetworkMonitorDownHardwareCount uint64 `json:"networkMonitorDownHardwareCount"`

	// NetworkMonitorDownVirtualGuestCount - A count of virtual guest which is currently experiencing a
	// service failure.
	NetworkMonitorDownVirtualGuestCount uint64 `json:"networkMonitorDownVirtualGuestCount"`

	// NetworkMonitorDownVirtualGuests - Virtual guest which is currently experiencing a service failure.
	NetworkMonitorDownVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorDownVirtualGuests"`

	// NetworkMonitorRecoveringHardware - Hardware which is currently recovering from a service failure.
	NetworkMonitorRecoveringHardware []*SoftLayer_Hardware `json:"networkMonitorRecoveringHardware"`

	// NetworkMonitorRecoveringHardwareCount - A count of hardware which is currently recovering from a
	// service failure.
	NetworkMonitorRecoveringHardwareCount uint64 `json:"networkMonitorRecoveringHardwareCount"`

	// NetworkMonitorRecoveringVirtualGuestCount - A count of virtual guest which is currently recovering
	// from a service failure.
	NetworkMonitorRecoveringVirtualGuestCount uint64 `json:"networkMonitorRecoveringVirtualGuestCount"`

	// NetworkMonitorRecoveringVirtualGuests - Virtual guest which is currently recovering from a service
	// failure.
	NetworkMonitorRecoveringVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorRecoveringVirtualGuests"`

	// NetworkMonitorUpHardware - no documentation
	NetworkMonitorUpHardware []*SoftLayer_Hardware `json:"networkMonitorUpHardware"`

	// NetworkMonitorUpHardwareCount - no documentation
	NetworkMonitorUpHardwareCount uint64 `json:"networkMonitorUpHardwareCount"`

	// NetworkMonitorUpVirtualGuestCount - A count of virtual guest which is currently online.
	NetworkMonitorUpVirtualGuestCount uint64 `json:"networkMonitorUpVirtualGuestCount"`

	// NetworkMonitorUpVirtualGuests - no documentation
	NetworkMonitorUpVirtualGuests []*SoftLayer_Virtual_Guest `json:"networkMonitorUpVirtualGuests"`

	// NetworkStorage - An account's associated storage volumes. This includes Lockbox, EVault, and iSCSI
	// volumes.
	NetworkStorage []*SoftLayer_Network_Storage `json:"networkStorage"`

	// NetworkStorageCount - A count of an account's associated storage volumes. This includes Lockbox,
	// EVault, and iSCSI volumes.
	NetworkStorageCount uint64 `json:"networkStorageCount"`

	// NetworkStorageGroupCount - no documentation
	NetworkStorageGroupCount uint64 `json:"networkStorageGroupCount"`

	// NetworkStorageGroups - no documentation
	NetworkStorageGroups []*SoftLayer_Network_Storage_Group `json:"networkStorageGroups"`

	// NetworkTunnelContextCount - no documentation
	NetworkTunnelContextCount uint64 `json:"networkTunnelContextCount"`

	// NetworkTunnelContexts - no documentation
	NetworkTunnelContexts []*SoftLayer_Network_Tunnel_Module_Context `json:"networkTunnelContexts"`

	// NetworkVlanCount - A count of all network VLANs assigned to an account.
	NetworkVlanCount uint64 `json:"networkVlanCount"`

	// NetworkVlanSpan - Whether or not an account has automatic private spanning enabled.
	NetworkVlanSpan *SoftLayer_Account_Network_Vlan_Span `json:"networkVlanSpan"`

	// NetworkVlans - no documentation
	NetworkVlans []*SoftLayer_Network_Vlan `json:"networkVlans"`

	// NextBillingPublicAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information
	// can be pulled directly through tapping keys now - The allotments for this account and their servers
	// for the next billing cycle. The public inbound and outbound bandwidth is calculated for each server
	// in addition to the daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetailCount uint64 `json:"nextBillingPublicAllotmentHardwareBandwidthDetailCount"`

	// NextBillingPublicAllotmentHardwareBandwidthDetails - - This information can be pulled directly
	// through tapping keys now - The allotments for this account and their servers for the next billing
	// cycle. The public inbound and outbound bandwidth is calculated for each server in addition to the
	// daily average network traffic since the last billing date.
	NextBillingPublicAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"nextBillingPublicAllotmentHardwareBandwidthDetails"`

	// NextInvoiceIncubatorExemptTotal - The pre-tax total amount exempt from incubator credit for the
	// account's next invoice.
	NextInvoiceIncubatorExemptTotal float64 `json:"nextInvoiceIncubatorExemptTotal"`

	// NextInvoiceTopLevelBillingItemCount - A count of the billing items that will be on an account's next
	// invoice.
	NextInvoiceTopLevelBillingItemCount uint64 `json:"nextInvoiceTopLevelBillingItemCount"`

	// NextInvoiceTopLevelBillingItems - The billing items that will be on an account's next invoice.
	NextInvoiceTopLevelBillingItems []*SoftLayer_Billing_Item `json:"nextInvoiceTopLevelBillingItems"`

	// NextInvoiceTotalAmount - The pre-tax total amount of an account's next invoice measured in US
	// Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalAmount float32 `json:"nextInvoiceTotalAmount"`

	// NextInvoiceTotalOneTimeAmount - The total one-time charge amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeAmount float64 `json:"nextInvoiceTotalOneTimeAmount"`

	// NextInvoiceTotalOneTimeTaxAmount - The total one-time tax amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalOneTimeTaxAmount float64 `json:"nextInvoiceTotalOneTimeTaxAmount"`

	// NextInvoiceTotalRecurringAmount - The total recurring charge amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringAmount float64 `json:"nextInvoiceTotalRecurringAmount"`

	// NextInvoiceTotalRecurringAmountBeforeAccountDiscount - The total recurring charge amount of an
	// account's next invoice measured in US Dollars assuming no changes or charges occur between now and
	// time of billing.
	NextInvoiceTotalRecurringAmountBeforeAccountDiscount float64 `json:"nextInvoiceTotalRecurringAmountBeforeAccountDiscount"`

	// NextInvoiceTotalRecurringTaxAmount - The total recurring tax amount of an account's next invoice
	// measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalRecurringTaxAmount float32 `json:"nextInvoiceTotalRecurringTaxAmount"`

	// NextInvoiceTotalTaxableRecurringAmount - The total recurring charge amount of an account's next
	// invoice measured in US Dollars assuming no changes or charges occur between now and time of billing.
	NextInvoiceTotalTaxableRecurringAmount float64 `json:"nextInvoiceTotalTaxableRecurringAmount"`

	// NotificationSubscriberCount - no documentation
	NotificationSubscriberCount uint64 `json:"notificationSubscriberCount"`

	// NotificationSubscribers - <nil>
	NotificationSubscribers []*SoftLayer_Notification_Subscriber `json:"notificationSubscribers"`

	// OfficePhone - no documentation
	OfficePhone string `json:"officePhone"`

	// OpenAbuseTicketCount - A count of the open abuse tickets associated with an account.
	OpenAbuseTicketCount uint64 `json:"openAbuseTicketCount"`

	// OpenAbuseTickets - no documentation
	OpenAbuseTickets []*SoftLayer_Ticket `json:"openAbuseTickets"`

	// OpenAccountingTicketCount - A count of the open accounting tickets associated with an account.
	OpenAccountingTicketCount uint64 `json:"openAccountingTicketCount"`

	// OpenAccountingTickets - The open accounting tickets associated with an account.
	OpenAccountingTickets []*SoftLayer_Ticket `json:"openAccountingTickets"`

	// OpenBillingTicketCount - A count of the open billing tickets associated with an account.
	OpenBillingTicketCount uint64 `json:"openBillingTicketCount"`

	// OpenBillingTickets - The open billing tickets associated with an account.
	OpenBillingTickets []*SoftLayer_Ticket `json:"openBillingTickets"`

	// OpenCancellationRequestCount - A count of an open ticket requesting cancellation of this server, if
	// one exists.
	OpenCancellationRequestCount uint64 `json:"openCancellationRequestCount"`

	// OpenCancellationRequests - An open ticket requesting cancellation of this server, if one exists.
	OpenCancellationRequests []*SoftLayer_Billing_Item_Cancellation_Request `json:"openCancellationRequests"`

	// OpenOtherTicketCount - A count of the open tickets that do not belong to the abuse, accounting,
	// sales, or support groups associated with an account.
	OpenOtherTicketCount uint64 `json:"openOtherTicketCount"`

	// OpenOtherTickets - The open tickets that do not belong to the abuse, accounting, sales, or support
	// groups associated with an account.
	OpenOtherTickets []*SoftLayer_Ticket `json:"openOtherTickets"`

	// OpenRecurringInvoiceCount - no documentation
	OpenRecurringInvoiceCount uint64 `json:"openRecurringInvoiceCount"`

	// OpenRecurringInvoices - no documentation
	OpenRecurringInvoices []*SoftLayer_Billing_Invoice `json:"openRecurringInvoices"`

	// OpenSalesTicketCount - A count of the open sales tickets associated with an account.
	OpenSalesTicketCount uint64 `json:"openSalesTicketCount"`

	// OpenSalesTickets - no documentation
	OpenSalesTickets []*SoftLayer_Ticket `json:"openSalesTickets"`

	// OpenStackAccountLinkCount - no documentation
	OpenStackAccountLinkCount uint64 `json:"openStackAccountLinkCount"`

	// OpenStackAccountLinks - <nil>
	OpenStackAccountLinks []*SoftLayer_Account_Link `json:"openStackAccountLinks"`

	// OpenSupportTicketCount - A count of the open support tickets associated with an account.
	OpenSupportTicketCount uint64 `json:"openSupportTicketCount"`

	// OpenSupportTickets - The open support tickets associated with an account.
	OpenSupportTickets []*SoftLayer_Ticket `json:"openSupportTickets"`

	// OpenTicketCount - A count of all open tickets associated with an account.
	OpenTicketCount uint64 `json:"openTicketCount"`

	// OpenTickets - no documentation
	OpenTickets []*SoftLayer_Ticket `json:"openTickets"`

	// OpenTicketsWaitingOnCustomer - All open tickets associated with an account last edited by an
	// employee.
	OpenTicketsWaitingOnCustomer []*SoftLayer_Ticket `json:"openTicketsWaitingOnCustomer"`

	// OpenTicketsWaitingOnCustomerCount - A count of all open tickets associated with an account last
	// edited by an employee.
	OpenTicketsWaitingOnCustomerCount uint64 `json:"openTicketsWaitingOnCustomerCount"`

	// OrderCount - A count of an account's associated billing orders excluding upgrades.
	OrderCount uint64 `json:"orderCount"`

	// Orders - An account's associated billing orders excluding upgrades.
	Orders []*SoftLayer_Billing_Order `json:"orders"`

	// OrphanBillingItemCount - A count of the billing items that have no parent billing item. These are
	// items that don't necessarily belong to a single server.
	OrphanBillingItemCount uint64 `json:"orphanBillingItemCount"`

	// OrphanBillingItems - The billing items that have no parent billing item. These are items that don't
	// necessarily belong to a single server.
	OrphanBillingItems []*SoftLayer_Billing_Item `json:"orphanBillingItems"`

	// OwnedBrandCount - no documentation
	OwnedBrandCount uint64 `json:"ownedBrandCount"`

	// OwnedBrands - <nil>
	OwnedBrands []*SoftLayer_Brand `json:"ownedBrands"`

	// OwnedHardwareGenericComponentModelCount - no documentation
	OwnedHardwareGenericComponentModelCount uint64 `json:"ownedHardwareGenericComponentModelCount"`

	// OwnedHardwareGenericComponentModels - <nil>
	OwnedHardwareGenericComponentModels []*SoftLayer_Hardware_Component_Model_Generic `json:"ownedHardwareGenericComponentModels"`

	// PaymentProcessorCount - no documentation
	PaymentProcessorCount uint64 `json:"paymentProcessorCount"`

	// PaymentProcessors - <nil>
	PaymentProcessors []*SoftLayer_Billing_Payment_Processor `json:"paymentProcessors"`

	// PendingEventCount - no documentation
	PendingEventCount uint64 `json:"pendingEventCount"`

	// PendingEvents - <nil>
	PendingEvents []*SoftLayer_Notification_Occurrence_Event `json:"pendingEvents"`

	// PendingInvoice - no documentation
	PendingInvoice *SoftLayer_Billing_Invoice `json:"pendingInvoice"`

	// PendingInvoiceTopLevelItemCount - A count of a list of top-level invoice items that are on an
	// account's currently pending invoice.
	PendingInvoiceTopLevelItemCount uint64 `json:"pendingInvoiceTopLevelItemCount"`

	// PendingInvoiceTopLevelItems - A list of top-level invoice items that are on an account's currently
	// pending invoice.
	PendingInvoiceTopLevelItems []*SoftLayer_Billing_Invoice_Item `json:"pendingInvoiceTopLevelItems"`

	// PendingInvoiceTotalAmount - The total amount of an account's pending invoice, if one exists.
	PendingInvoiceTotalAmount float64 `json:"pendingInvoiceTotalAmount"`

	// PendingInvoiceTotalOneTimeAmount - The total one-time charges for an account's pending invoice, if
	// one exists. In other words, it is the sum of one-time charges, setup fees, and labor fees. It does
	// not include taxes.
	PendingInvoiceTotalOneTimeAmount float64 `json:"pendingInvoiceTotalOneTimeAmount"`

	// PendingInvoiceTotalOneTimeTaxAmount - The sum of all the taxes related to one time charges for an
	// account's pending invoice, if one exists.
	PendingInvoiceTotalOneTimeTaxAmount float64 `json:"pendingInvoiceTotalOneTimeTaxAmount"`

	// PendingInvoiceTotalRecurringAmount - The total recurring amount of an account's pending invoice, if
	// one exists.
	PendingInvoiceTotalRecurringAmount float64 `json:"pendingInvoiceTotalRecurringAmount"`

	// PendingInvoiceTotalRecurringTaxAmount - The total amount of the recurring taxes on an account's
	// pending invoice, if one exists.
	PendingInvoiceTotalRecurringTaxAmount float64 `json:"pendingInvoiceTotalRecurringTaxAmount"`

	// PermissionGroupCount - no documentation
	PermissionGroupCount uint64 `json:"permissionGroupCount"`

	// PermissionGroups - no documentation
	PermissionGroups []*SoftLayer_User_Permission_Group `json:"permissionGroups"`

	// PermissionRoleCount - no documentation
	PermissionRoleCount uint64 `json:"permissionRoleCount"`

	// PermissionRoles - no documentation
	PermissionRoles []*SoftLayer_User_Permission_Role `json:"permissionRoles"`

	// PortableStorageVolumeCount - no documentation
	PortableStorageVolumeCount uint64 `json:"portableStorageVolumeCount"`

	// PortableStorageVolumes - <nil>
	PortableStorageVolumes []*SoftLayer_Virtual_Disk_Image `json:"portableStorageVolumes"`

	// PostProvisioningHookCount - A count of customer specified URIs that are downloaded onto a newly
	// provisioned or reloaded server. If the URI is sent over https it will be executed directly on the
	// server.
	PostProvisioningHookCount uint64 `json:"postProvisioningHookCount"`

	// PostProvisioningHooks - Customer specified URIs that are downloaded onto a newly provisioned or
	// reloaded server. If the URI is sent over https it will be executed directly on the server.
	PostProvisioningHooks []*SoftLayer_Provisioning_Hook `json:"postProvisioningHooks"`

	// PostalCode - The postal code of the mailing address belonging to an account.
	PostalCode string `json:"postalCode"`

	// PptpVpnUserCount - A count of an account's associated portal users with VPN access.
	PptpVpnUserCount uint64 `json:"pptpVpnUserCount"`

	// PptpVpnUsers - An account's associated portal users with VPN access.
	PptpVpnUsers []*SoftLayer_User_Customer `json:"pptpVpnUsers"`

	// PreviousRecurringRevenue - The total recurring amount for an accounts previous revenue.
	PreviousRecurringRevenue float64 `json:"previousRecurringRevenue"`

	// PriceRestrictionCount - A count of the item price that an account is restricted to.
	PriceRestrictionCount uint64 `json:"priceRestrictionCount"`

	// PriceRestrictions - no documentation
	PriceRestrictions []*SoftLayer_Product_Item_Price_Account_Restriction `json:"priceRestrictions"`

	// PriorityOneTicketCount - A count of all priority one tickets associated with an account.
	PriorityOneTicketCount uint64 `json:"priorityOneTicketCount"`

	// PriorityOneTickets - All priority one tickets associated with an account.
	PriorityOneTickets []*SoftLayer_Ticket `json:"priorityOneTickets"`

	// PrivateAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information can be
	// pulled directly through tapping keys now - The allotments for this account and their servers. The
	// private inbound and outbound bandwidth is calculated for each server in addition to the daily
	// average network traffic since the last billing date.
	PrivateAllotmentHardwareBandwidthDetailCount uint64 `json:"privateAllotmentHardwareBandwidthDetailCount"`

	// PrivateAllotmentHardwareBandwidthDetails - - This information can be pulled directly through tapping
	// keys now - The allotments for this account and their servers. The private inbound and outbound
	// bandwidth is calculated for each server in addition to the daily average network traffic since the
	// last billing date.
	PrivateAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"privateAllotmentHardwareBandwidthDetails"`

	// PrivateBlockDeviceTemplateGroupCount - A count of private and shared template group objects (parent
	// only) for an account.
	PrivateBlockDeviceTemplateGroupCount uint64 `json:"privateBlockDeviceTemplateGroupCount"`

	// PrivateBlockDeviceTemplateGroups - Private and shared template group objects (parent only) for an
	// account.
	PrivateBlockDeviceTemplateGroups []*SoftLayer_Virtual_Guest_Block_Device_Template_Group `json:"privateBlockDeviceTemplateGroups"`

	// PrivateIpAddressCount - no documentation
	PrivateIpAddressCount uint64 `json:"privateIpAddressCount"`

	// PrivateIpAddresses - <nil>
	PrivateIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"privateIpAddresses"`

	// PrivateNetworkVlanCount - A count of the private network VLANs assigned to an account.
	PrivateNetworkVlanCount uint64 `json:"privateNetworkVlanCount"`

	// PrivateNetworkVlans - no documentation
	PrivateNetworkVlans []*SoftLayer_Network_Vlan `json:"privateNetworkVlans"`

	// PrivateSubnetCount - A count of all private subnets associated with an account.
	PrivateSubnetCount uint64 `json:"privateSubnetCount"`

	// PrivateSubnets - no documentation
	PrivateSubnets []*SoftLayer_Network_Subnet `json:"privateSubnets"`

	// PublicAllotmentHardwareBandwidthDetailCount - A count of dEPRECATED - This information can be pulled
	// directly through tapping keys now - The allotments for this account and their servers. The public
	// inbound and outbound bandwidth is calculated for each server in addition to the daily average
	// network traffic since the last billing date.
	PublicAllotmentHardwareBandwidthDetailCount uint64 `json:"publicAllotmentHardwareBandwidthDetailCount"`

	// PublicAllotmentHardwareBandwidthDetails - - This information can be pulled directly through tapping
	// keys now - The allotments for this account and their servers. The public inbound and outbound
	// bandwidth is calculated for each server in addition to the daily average network traffic since the
	// last billing date.
	PublicAllotmentHardwareBandwidthDetails []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"publicAllotmentHardwareBandwidthDetails"`

	// PublicIpAddressCount - no documentation
	PublicIpAddressCount uint64 `json:"publicIpAddressCount"`

	// PublicIpAddresses - <nil>
	PublicIpAddresses []*SoftLayer_Network_Subnet_IpAddress `json:"publicIpAddresses"`

	// PublicNetworkVlanCount - A count of the public network VLANs assigned to an account.
	PublicNetworkVlanCount uint64 `json:"publicNetworkVlanCount"`

	// PublicNetworkVlans - no documentation
	PublicNetworkVlans []*SoftLayer_Network_Vlan `json:"publicNetworkVlans"`

	// PublicSubnetCount - A count of all public network subnets associated with an account.
	PublicSubnetCount uint64 `json:"publicSubnetCount"`

	// PublicSubnets - All public network subnets associated with an account.
	PublicSubnets []*SoftLayer_Network_Subnet `json:"publicSubnets"`

	// QuoteCount - no documentation
	QuoteCount uint64 `json:"quoteCount"`

	// Quotes - no documentation
	Quotes []*SoftLayer_Billing_Order_Quote `json:"quotes"`

	// RecentEventCount - no documentation
	RecentEventCount uint64 `json:"recentEventCount"`

	// RecentEvents - <nil>
	RecentEvents []*SoftLayer_Notification_Occurrence_Event `json:"recentEvents"`

	// ReferralPartner - no documentation
	ReferralPartner *SoftLayer_Account `json:"referralPartner"`

	// ReferredAccountCount - A count of if this is a account is a referral partner, the accounts this
	// referral partner has referred
	ReferredAccountCount uint64 `json:"referredAccountCount"`

	// ReferredAccounts - If this is a account is a referral partner, the accounts this referral partner
	// has referred
	ReferredAccounts []*SoftLayer_Account `json:"referredAccounts"`

	// RemoteManagementCommandRequestCount - A count of remote management command requests for an account
	RemoteManagementCommandRequestCount uint64 `json:"remoteManagementCommandRequestCount"`

	// RemoteManagementCommandRequests - no documentation
	RemoteManagementCommandRequests []*SoftLayer_Hardware_Component_RemoteManagement_Command_Request `json:"remoteManagementCommandRequests"`

	// ReplicationEventCount - A count of the Replication events for all Network Storage volumes on an
	// account.
	ReplicationEventCount uint64 `json:"replicationEventCount"`

	// ReplicationEvents - The Replication events for all Network Storage volumes on an account.
	ReplicationEvents []*SoftLayer_Network_Storage_Event `json:"replicationEvents"`

	// ResourceGroupCount - A count of an account's associated top-level resource groups.
	ResourceGroupCount uint64 `json:"resourceGroupCount"`

	// ResourceGroups - no documentation
	ResourceGroups []*SoftLayer_Resource_Group `json:"resourceGroups"`

	// RouterCount - A count of all Routers that an accounts VLANs reside on
	RouterCount uint64 `json:"routerCount"`

	// Routers - no documentation
	Routers []*SoftLayer_Hardware `json:"routers"`

	// RwhoisData - An account's reverse data. This data is used when making requests.
	RwhoisData *SoftLayer_Network_Subnet_Rwhois_Data `json:"rwhoisData"`

	// SalesforceAccountLink - <nil>
	SalesforceAccountLink *SoftLayer_Account_Link `json:"salesforceAccountLink"`

	// SamlAuthentication - no documentation
	SamlAuthentication *SoftLayer_Account_Authentication_Saml `json:"samlAuthentication"`

	// ScaleGroupCount - no documentation
	ScaleGroupCount uint64 `json:"scaleGroupCount"`

	// ScaleGroups - no documentation
	ScaleGroups []*SoftLayer_Scale_Group `json:"scaleGroups"`

	// SecondaryDomainCount - A count of the secondary DNS records for a SoftLayer customer account.
	SecondaryDomainCount uint64 `json:"secondaryDomainCount"`

	// SecondaryDomains - The secondary DNS records for a SoftLayer customer account.
	SecondaryDomains []*SoftLayer_Dns_Secondary `json:"secondaryDomains"`

	// SecurityCertificateCount - no documentation
	SecurityCertificateCount uint64 `json:"securityCertificateCount"`

	// SecurityCertificates - no documentation
	SecurityCertificates []*SoftLayer_Security_Certificate `json:"securityCertificates"`

	// SecurityScanRequestCount - A count of an account's vulnerability scan requests.
	SecurityScanRequestCount uint64 `json:"securityScanRequestCount"`

	// SecurityScanRequests - no documentation
	SecurityScanRequests []*SoftLayer_Network_Security_Scanner_Request `json:"securityScanRequests"`

	// ServiceBillingItemCount - A count of the service billing items that will be on an account's next
	// invoice.
	ServiceBillingItemCount uint64 `json:"serviceBillingItemCount"`

	// ServiceBillingItems - The service billing items that will be on an account's next invoice.
	ServiceBillingItems []*SoftLayer_Billing_Item `json:"serviceBillingItems"`

	// ShipmentCount - A count of shipments that belong to the customer's account.
	ShipmentCount uint64 `json:"shipmentCount"`

	// Shipments - no documentation
	Shipments []*SoftLayer_Account_Shipment `json:"shipments"`

	// SshKeyCount - A count of customer specified SSH keys that can be implemented onto a newly
	// provisioned or reloaded server.
	SshKeyCount uint64 `json:"sshKeyCount"`

	// SshKeys - Customer specified SSH keys that can be implemented onto a newly provisioned or reloaded
	// server.
	SshKeys []*SoftLayer_Security_Ssh_Key `json:"sshKeys"`

	// SslVpnUserCount - A count of an account's associated portal users with SSL VPN access.
	SslVpnUserCount uint64 `json:"sslVpnUserCount"`

	// SslVpnUsers - An account's associated portal users with SSL VPN access.
	SslVpnUsers []*SoftLayer_User_Customer `json:"sslVpnUsers"`

	// StandardPoolVirtualGuestCount - A count of an account's virtual guest objects that are hosted on a
	// user provisioned hypervisor.
	StandardPoolVirtualGuestCount uint64 `json:"standardPoolVirtualGuestCount"`

	// StandardPoolVirtualGuests - An account's virtual guest objects that are hosted on a user provisioned
	// hypervisor.
	StandardPoolVirtualGuests []*SoftLayer_Virtual_Guest `json:"standardPoolVirtualGuests"`

	// State - A two-letter abbreviation of the state in the mailing address belonging to an account. If an
	// account does not reside in a province then this is typically blank.
	State string `json:"state"`

	// StatusDate - no documentation
	StatusDate *time.Time `json:"statusDate"`

	// SubnetCount - A count of all network subnets associated with an account.
	SubnetCount uint64 `json:"subnetCount"`

	// SubnetRegistrationCount - no documentation
	SubnetRegistrationCount uint64 `json:"subnetRegistrationCount"`

	// SubnetRegistrationDetailCount - no documentation
	SubnetRegistrationDetailCount uint64 `json:"subnetRegistrationDetailCount"`

	// SubnetRegistrationDetails - <nil>
	SubnetRegistrationDetails []*SoftLayer_Account_Regional_Registry_Detail `json:"subnetRegistrationDetails"`

	// SubnetRegistrations - <nil>
	SubnetRegistrations []*SoftLayer_Network_Subnet_Registration `json:"subnetRegistrations"`

	// Subnets - no documentation
	Subnets []*SoftLayer_Network_Subnet `json:"subnets"`

	// SupportRepresentativeCount - A count of the SoftLayer employees that an account is assigned to.
	SupportRepresentativeCount uint64 `json:"supportRepresentativeCount"`

	// SupportRepresentatives - The SoftLayer employees that an account is assigned to.
	SupportRepresentatives []*SoftLayer_User_Employee `json:"supportRepresentatives"`

	// SupportSubscriptionCount - A count of the active support subscriptions for this account.
	SupportSubscriptionCount uint64 `json:"supportSubscriptionCount"`

	// SupportSubscriptions - no documentation
	SupportSubscriptions []*SoftLayer_Billing_Item `json:"supportSubscriptions"`

	// SuppressInvoicesFlag - no documentation
	SuppressInvoicesFlag bool `json:"suppressInvoicesFlag"`

	// TagCount - no documentation
	TagCount uint64 `json:"tagCount"`

	// Tags - <nil>
	Tags []*SoftLayer_Tag `json:"tags"`

	// TicketCount - no documentation
	TicketCount uint64 `json:"ticketCount"`

	// Tickets - no documentation
	Tickets []*SoftLayer_Ticket `json:"tickets"`

	// TicketsClosedInTheLastThreeDays - Tickets closed within the last 72 hours or last 10 tickets,
	// whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDays []*SoftLayer_Ticket `json:"ticketsClosedInTheLastThreeDays"`

	// TicketsClosedInTheLastThreeDaysCount - A count of tickets closed within the last 72 hours or last 10
	// tickets, whichever is less, associated with an account.
	TicketsClosedInTheLastThreeDaysCount uint64 `json:"ticketsClosedInTheLastThreeDaysCount"`

	// TicketsClosedToday - no documentation
	TicketsClosedToday []*SoftLayer_Ticket `json:"ticketsClosedToday"`

	// TicketsClosedTodayCount - A count of tickets closed today associated with an account.
	TicketsClosedTodayCount uint64 `json:"ticketsClosedTodayCount"`

	// TranscodeAccountCount - A count of an account's associated Transcode account.
	TranscodeAccountCount uint64 `json:"transcodeAccountCount"`

	// TranscodeAccounts - no documentation
	TranscodeAccounts []*SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccounts"`

	// UpgradeRequestCount - A count of an account's associated upgrade requests.
	UpgradeRequestCount uint64 `json:"upgradeRequestCount"`

	// UpgradeRequests - no documentation
	UpgradeRequests []*SoftLayer_Product_Upgrade_Request `json:"upgradeRequests"`

	// UserCount - no documentation
	UserCount uint64 `json:"userCount"`

	// Users - no documentation
	Users []*SoftLayer_User_Customer `json:"users"`

	// ValidSecurityCertificateCount - A count of stored security certificates that are not expired (ie.
	ValidSecurityCertificateCount uint64 `json:"validSecurityCertificateCount"`

	// ValidSecurityCertificates - Stored security certificates that are not expired (ie.
	ValidSecurityCertificates []*SoftLayer_Security_Certificate `json:"validSecurityCertificates"`

	// VdrUpdatesInProgressFlag - Return 0 if vpn updates are currently in progress on this account
	// otherwise 1.
	VdrUpdatesInProgressFlag bool `json:"vdrUpdatesInProgressFlag"`

	// VirtualDedicatedRackCount - no documentation
	VirtualDedicatedRackCount uint64 `json:"virtualDedicatedRackCount"`

	// VirtualDedicatedRacks - no documentation
	VirtualDedicatedRacks []*SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualDedicatedRacks"`

	// VirtualDiskImageCount - A count of an account's associated virtual server virtual disk images.
	VirtualDiskImageCount uint64 `json:"virtualDiskImageCount"`

	// VirtualDiskImages - An account's associated virtual server virtual disk images.
	VirtualDiskImages []*SoftLayer_Virtual_Disk_Image `json:"virtualDiskImages"`

	// VirtualGuestCount - A count of an account's associated virtual guest objects.
	VirtualGuestCount uint64 `json:"virtualGuestCount"`

	// VirtualGuests - no documentation
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests"`

	// VirtualGuestsOverBandwidthAllocation - An account's associated virtual guest objects currently over
	// bandwidth allocation.
	VirtualGuestsOverBandwidthAllocation []*SoftLayer_Virtual_Guest `json:"virtualGuestsOverBandwidthAllocation"`

	// VirtualGuestsOverBandwidthAllocationCount - A count of an account's associated virtual guest objects
	// currently over bandwidth allocation.
	VirtualGuestsOverBandwidthAllocationCount uint64 `json:"virtualGuestsOverBandwidthAllocationCount"`

	// VirtualGuestsProjectedOverBandwidthAllocation - An account's associated virtual guest objects
	// currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocation []*SoftLayer_Virtual_Guest `json:"virtualGuestsProjectedOverBandwidthAllocation"`

	// VirtualGuestsProjectedOverBandwidthAllocationCount - A count of an account's associated virtual
	// guest objects currently over bandwidth allocation.
	VirtualGuestsProjectedOverBandwidthAllocationCount uint64 `json:"virtualGuestsProjectedOverBandwidthAllocationCount"`

	// VirtualGuestsWithCpanel - All virtual guests associated with an account that has the cPanel web
	// hosting control panel installed.
	VirtualGuestsWithCpanel []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithCpanel"`

	// VirtualGuestsWithCpanelCount - A count of all virtual guests associated with an account that has the
	// cPanel web hosting control panel installed.
	VirtualGuestsWithCpanelCount uint64 `json:"virtualGuestsWithCpanelCount"`

	// VirtualGuestsWithMcafee - All virtual guests associated with an account that have McAfee Secure
	// software components.
	VirtualGuestsWithMcafee []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafee"`

	// VirtualGuestsWithMcafeeAntivirusRedhat - All virtual guests associated with an account that have
	// McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhat []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusRedhat"`

	// VirtualGuestsWithMcafeeAntivirusRedhatCount - A count of all virtual guests associated with an
	// account that have McAfee Secure AntiVirus for Redhat software components.
	VirtualGuestsWithMcafeeAntivirusRedhatCount uint64 `json:"virtualGuestsWithMcafeeAntivirusRedhatCount"`

	// VirtualGuestsWithMcafeeAntivirusWindowCount - A count of all virtual guests associated with an
	// account that has McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindowCount uint64 `json:"virtualGuestsWithMcafeeAntivirusWindowCount"`

	// VirtualGuestsWithMcafeeAntivirusWindows - All virtual guests associated with an account that has
	// McAfee Secure AntiVirus for Windows software components.
	VirtualGuestsWithMcafeeAntivirusWindows []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeAntivirusWindows"`

	// VirtualGuestsWithMcafeeCount - A count of all virtual guests associated with an account that have
	// McAfee Secure software components.
	VirtualGuestsWithMcafeeCount uint64 `json:"virtualGuestsWithMcafeeCount"`

	// VirtualGuestsWithMcafeeIntrusionDetectionSystem - All virtual guests associated with an account that
	// has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystem []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithMcafeeIntrusionDetectionSystem"`

	// VirtualGuestsWithMcafeeIntrusionDetectionSystemCount - A count of all virtual guests associated with
	// an account that has McAfee Secure Intrusion Detection System software components.
	VirtualGuestsWithMcafeeIntrusionDetectionSystemCount uint64 `json:"virtualGuestsWithMcafeeIntrusionDetectionSystemCount"`

	// VirtualGuestsWithPlesk - All virtual guests associated with an account that has the Plesk web
	// hosting control panel installed.
	VirtualGuestsWithPlesk []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithPlesk"`

	// VirtualGuestsWithPleskCount - A count of all virtual guests associated with an account that has the
	// Plesk web hosting control panel installed.
	VirtualGuestsWithPleskCount uint64 `json:"virtualGuestsWithPleskCount"`

	// VirtualGuestsWithQuantastor - All virtual guests associated with an account that have the QuantaStor
	// storage system installed.
	VirtualGuestsWithQuantastor []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithQuantastor"`

	// VirtualGuestsWithQuantastorCount - A count of all virtual guests associated with an account that
	// have the QuantaStor storage system installed.
	VirtualGuestsWithQuantastorCount uint64 `json:"virtualGuestsWithQuantastorCount"`

	// VirtualGuestsWithUrchin - All virtual guests associated with an account that has the Urchin web
	// traffic analytics package installed.
	VirtualGuestsWithUrchin []*SoftLayer_Virtual_Guest `json:"virtualGuestsWithUrchin"`

	// VirtualGuestsWithUrchinCount - A count of all virtual guests associated with an account that has the
	// Urchin web traffic analytics package installed.
	VirtualGuestsWithUrchinCount uint64 `json:"virtualGuestsWithUrchinCount"`

	// VirtualPrivateRack - no documentation
	VirtualPrivateRack *SoftLayer_Network_Bandwidth_Version1_Allotment `json:"virtualPrivateRack"`

	// VirtualStorageArchiveRepositories - An account's associated virtual server archived storage
	// repositories.
	VirtualStorageArchiveRepositories []*SoftLayer_Virtual_Storage_Repository `json:"virtualStorageArchiveRepositories"`

	// VirtualStorageArchiveRepositoryCount - A count of an account's associated virtual server archived
	// storage repositories.
	VirtualStorageArchiveRepositoryCount uint64 `json:"virtualStorageArchiveRepositoryCount"`

	// VirtualStoragePublicRepositories - An account's associated virtual server public storage
	// repositories.
	VirtualStoragePublicRepositories []*SoftLayer_Virtual_Storage_Repository `json:"virtualStoragePublicRepositories"`

	// VirtualStoragePublicRepositoryCount - A count of an account's associated virtual server public
	// storage repositories.
	VirtualStoragePublicRepositoryCount uint64 `json:"virtualStoragePublicRepositoryCount"`
}

func (softlayer_account *SoftLayer_Account) String() string {
	return "SoftLayer_Account"
}

// ActivatePartner - <nil>
func (softlayer_account *SoftLayer_Account) ActivatePartner(ctx *slapi.RequestContext, accountId string, hashCode string) (*SoftLayer_Account, error) {
	var returnValue *SoftLayer_Account
	return returnValue, nil
}

// AddAchInformation - <nil>
func (softlayer_account *SoftLayer_Account) AddAchInformation(ctx *slapi.RequestContext, achInformation SoftLayer_Container_Billing_Info_Ach) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AddReferralPartnerPaymentOption - <nil>
func (softlayer_account *SoftLayer_Account) AddReferralPartnerPaymentOption(ctx *slapi.RequestContext, paymentOption SoftLayer_Container_Referral_Partner_Payment_Option) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// AreVdrUpdatesBlockedForBilling - This method indicates whether or not Bandwidth Pooling updates are
// blocked for the account so the billing cycle can run. Generally, accounts are restricted from moving
// servers in or out of Bandwidth Pools from 12:00 CST on the day prior to billing, until the billing
// batch completes, sometime after midnight the day of actual billing for the account.
func (softlayer_account *SoftLayer_Account) AreVdrUpdatesBlockedForBilling(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CancelPayPalTransaction - Cancel the PayPal Payment Request process. During the process of
// submitting a PayPal payment request, the customer is redirected to PayPal to confirm the request. If
// the customer elects to cancel the payment from PayPal, they are returned to SoftLayer where the
// manual payment record is updated to a status of canceled.
func (softlayer_account *SoftLayer_Account) CancelPayPalTransaction(ctx *slapi.RequestContext, token string, payerId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CompletePayPalTransaction - Complete the PayPal Payment Request process and receive confirmation
// message. During the process of submitting a PayPal payment request, the customer is redirected to
// PayPal to confirm the request. Once confirmed, PayPal returns the customer to SoftLayer where an
// attempt is made to finalize the transaction. A status message regarding the attempt is returned to
// the calling function.
func (softlayer_account *SoftLayer_Account) CompletePayPalTransaction(ctx *slapi.RequestContext, token string, payerId string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// CountHourlyInstances - Retrieve the number of hourly services on an account that are active, plus
// any pending orders with hourly services attached.
func (softlayer_account *SoftLayer_Account) CountHourlyInstances(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetAccountBackupHistory - This method returns an array of
// SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails objects for the given start and end
// dates. Start and end dates should be be valid ISO 8601 dates. The backupStatus can be one of null,
// 'success', 'failed', or 'conflict'. The 'success' backupStatus returns jobs with a status of the
// 'failed' backupStatus returns jobs with a status of while the 'conflict' backupStatus will return
// jobs that are not or
func (softlayer_account *SoftLayer_Account) GetAccountBackupHistory(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time, backupStatus string) ([]*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails, error) {
	var returnValue []*SoftLayer_Container_Network_Storage_Evault_WebCc_JobDetails
	return returnValue, nil
}

// GetAccountTraitValue - no documentation
func (softlayer_account *SoftLayer_Account) GetAccountTraitValue(ctx *slapi.RequestContext, keyName string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetActiveAlarms - Return all currently active alarms on this account. Only alarms on hardware and
// virtual servers accessible to the current user will be returned.
func (softlayer_account *SoftLayer_Account) GetActiveAlarms(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Monitoring_Alarm_History, error) {
	var returnValue []*SoftLayer_Container_Monitoring_Alarm_History
	return returnValue, nil
}

// GetActiveOutletPackages - This method pulls all the active packages. This will give you a basic
// description of the packages within the SoftLayer Outlet store that are currently active and from
// which you can order a server or additional services.
func (softlayer_account *SoftLayer_Account) GetActiveOutletPackages(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetActivePackages - This method will return the [[SoftLayer_Product_Package]] objects from which you
// can order a bare metal server, virtual server, service (such as CDN or Object Storage) or other
// software. Once you have the package you want to order from, you may query one of various endpoints
// from that package to get specific information about its products and pricing. See
// [[SoftLayer_Product_Package/getCategories|getCategories]] or
// [[SoftLayer_Product_Package/getItems|getItems]] for more information. Packages that have been
// retired will not appear in this result set.
func (softlayer_account *SoftLayer_Account) GetActivePackages(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetActivePackagesByAttribute - This method is deprecated and should not be used in production code.
// This method will return the [[SoftLayer_Product_Package]] objects from which you can order a bare
// metal server, virtual server, service (such as CDN or Object Storage) or other software filtered by
// an attribute type associated with the package. Once you have the package you want to order from, you
// may query one of various endpoints from that package to get specific information about its products
// and pricing. See [[SoftLayer_Product_Package/getCategories|getCategories]] or
// [[SoftLayer_Product_Package/getItems|getItems]] for more information.
func (softlayer_account *SoftLayer_Account) GetActivePackagesByAttribute(ctx *slapi.RequestContext, attributeKeyName string) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetActivePrivateHostedCloudPackages - This method pulls all the active private hosted cloud
// packages. This will give you a basic description of the packages that are currently active and from
// which you can order private hosted cloud configurations.
func (softlayer_account *SoftLayer_Account) GetActivePrivateHostedCloudPackages(ctx *slapi.RequestContext) ([]*SoftLayer_Product_Package, error) {
	var returnValue []*SoftLayer_Product_Package
	return returnValue, nil
}

// GetAggregatedUptimeGraph - no documentation
func (softlayer_account *SoftLayer_Account) GetAggregatedUptimeGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetAlternateCreditCardData - <nil>
func (softlayer_account *SoftLayer_Account) GetAlternateCreditCardData(ctx *slapi.RequestContext) (*SoftLayer_Container_Account_Payment_Method_CreditCard, error) {
	var returnValue *SoftLayer_Container_Account_Payment_Method_CreditCard
	return returnValue, nil
}

// GetAttributeByType - Retrieve a single [[SoftLayer_Account_Attribute]] record by it's type's name.
func (softlayer_account *SoftLayer_Account) GetAttributeByType(ctx *slapi.RequestContext, attributeType string) (*SoftLayer_Account_Attribute, error) {
	var returnValue *SoftLayer_Account_Attribute
	return returnValue, nil
}

// GetAuxiliaryNotifications - <nil>
func (softlayer_account *SoftLayer_Account) GetAuxiliaryNotifications(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_Message, error) {
	var returnValue []*SoftLayer_Container_Utility_Message
	return returnValue, nil
}

// GetAverageArchiveUsageMetricDataByDate - Returns the average disk space usage for all archive
// repositories.
func (softlayer_account *SoftLayer_Account) GetAverageArchiveUsageMetricDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetAveragePublicUsageMetricDataByDate - Returns the average disk space usage for all public
// repositories.
func (softlayer_account *SoftLayer_Account) GetAveragePublicUsageMetricDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) (float32, error) {
	var returnValue float32
	return returnValue, nil
}

// GetCurrentBackupStatisticsGraph - This method returns a SoftLayer_Container_Account_Graph_Outputs
// containing a base64 string PNG image. The optional parameter, detailedGraph, can be passed to get a
// more detailed graph.
func (softlayer_account *SoftLayer_Account) GetCurrentBackupStatisticsGraph(ctx *slapi.RequestContext, detailedGraph bool) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetCurrentTicketStatisticsGraph - <nil>
func (softlayer_account *SoftLayer_Account) GetCurrentTicketStatisticsGraph(ctx *slapi.RequestContext, detailedGraph bool) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetCurrentUser - Retrieve the user record of the user calling the SoftLayer
func (softlayer_account *SoftLayer_Account) GetCurrentUser(ctx *slapi.RequestContext) (*SoftLayer_User_Customer, error) {
	var returnValue *SoftLayer_User_Customer
	return returnValue, nil
}

// GetDiskUsageMetricDataByDate - Retrieve disk usage data on a [[SoftLayer_Virtual_Guest|Cloud
// Computing Instance]] image for the time range you provide from the Metric Tracking Object System and
// Legacy Data Warehouse. Each data entry objects contain ''dateTime'' and ''counter'' properties.
// ''dateTime'' property indicates the time that the disk usage data was measured and ''counter''
// property holds the disk usage in bytes.
func (softlayer_account *SoftLayer_Account) GetDiskUsageMetricDataByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetDiskUsageMetricDataFromLegacyByDate - Retrieve disk usage data on a
// [[SoftLayer_Virtual_Guest|Cloud Computing Instance]] image for the time range you provide from the
// Legacy Data Warehouse. Each data entry objects contain ''dateTime'' and ''counter'' properties.
// ''dateTime'' property indicates the time that the disk usage data was measured and ''counter''
// property holds the disk usage in bytes.
func (softlayer_account *SoftLayer_Account) GetDiskUsageMetricDataFromLegacyByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetDiskUsageMetricDataFromMetricTrackingObjectSystemByDate - Retrieve disk usage data on a
// [[SoftLayer_Virtual_Guest|Cloud Computing Instance]] image for the time range you provide from the
// Metric Tracking Object System. Each data entry object contains ''dateTime'' and ''counter''
// properties. ''dateTime'' property indicates the time that the disk usage data was measured and
// ''counter'' property holds the disk usage in bytes.
func (softlayer_account *SoftLayer_Account) GetDiskUsageMetricDataFromMetricTrackingObjectSystemByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetDiskUsageMetricImageByDate - Returns a disk usage image based on disk usage specified by the
// input parameters.
func (softlayer_account *SoftLayer_Account) GetDiskUsageMetricImageByDate(ctx *slapi.RequestContext, startDateTime time.Time, endDateTime time.Time) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetExecutiveSummaryPdf - This method will return a PDF of the specified report, with the specified
// period within the start and end dates. The pdfType must be one of 'snapshot', or 'historical'.
// Possible historicalType parameters are 'monthly', 'yearly', and 'quarterly'. Start and end dates
// should be in ISO 8601 date format.
func (softlayer_account *SoftLayer_Account) GetExecutiveSummaryPdf(ctx *slapi.RequestContext, pdfType string, historicalType string, startDate string, endDate string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetFlexibleCreditProgramInfo - This method will return a
// [[SoftLayer_Container_Account_Discount_Program]] object containing the Flexible Credit Program
// information for this account. To be considered an active participant, the account must have an
// enrollment record with a monthly credit amount set and the current date must be within the range
// defined by the enrollment and graduation date. The forNextBillCycle parameter can be set to true to
// return a SoftLayer_Container_Account_Discount_Program object with information with relation to the
// next bill cycle. The forNextBillCycle parameter defaults to false. Please note that all discount
// amount entries are reported as pre-tax amounts and the legacy tax fields in the
// [[SoftLayer_Container_Account_Discount_Program]] are deprecated.
func (softlayer_account *SoftLayer_Account) GetFlexibleCreditProgramInfo(ctx *slapi.RequestContext, forNextBillCycle bool) (*SoftLayer_Container_Account_Discount_Program, error) {
	var returnValue *SoftLayer_Container_Account_Discount_Program
	return returnValue, nil
}

// GetHistoricalBackupGraph - <nil>
func (softlayer_account *SoftLayer_Account) GetHistoricalBackupGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetHistoricalBandwidthGraph - This method will return a SoftLayer_Container_Account_Graph_Outputs
// object containing a base64 string PNG image of a line graph of bandwidth statistics given the start
// and end dates. The start and end dates should be valid ISO 8601 date formatted strings.
func (softlayer_account *SoftLayer_Account) GetHistoricalBandwidthGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetHistoricalTicketGraph - Given the start and end dates, this method will return a pie chart of
// ticket statistics in the form of SoftLayer_Container_Account_Graph_Outputs object with a base64 PNG
// string. If an error occurs the graphError parameter will be populated. Possible errors include:
// SoftLayer_Exception_Public Thrown if an invalid start or end date is provided. SoftLayer_Exception
// Thrown if there is an error connecting to HBase. SoftLayer_Exception Thrown if there is no data
// available for the specified date range. SoftLayer_Exception Thrown if there is an error retrieving
// data or generating the graph.
func (softlayer_account *SoftLayer_Account) GetHistoricalTicketGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetHistoricalUptimeGraph - The graph image is returned as a base64 PNG string. Start and end dates
// should be formatted using the ISO 8601 date standard. If there is an error retrieving graph data or
// generating the graph string a graphError attribute will be returned. The graphError attribute may
// contain any of the following error messages: SoftLayer_Exception_Public Thrown if an invalid start
// or end date is provided. SoftLayer_Exception Thrown if there is an error connecting to HBase.
// SoftLayer_Exception Thrown if there is no data available for the specified date range.
// SoftLayer_Exception Thrown if there is an error retrieving data or generating the graph.
func (softlayer_account *SoftLayer_Account) GetHistoricalUptimeGraph(ctx *slapi.RequestContext, startDate time.Time, endDate time.Time) (*SoftLayer_Container_Account_Graph_Outputs, error) {
	var returnValue *SoftLayer_Container_Account_Graph_Outputs
	return returnValue, nil
}

// GetLargestAllowedSubnetCidr - <nil>
func (softlayer_account *SoftLayer_Account) GetLargestAllowedSubnetCidr(ctx *slapi.RequestContext, numberOfHosts int, locationId int) (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetNextInvoiceExcel - Return an account's next invoice in a Microsoft excel format. The "next
// invoice" is what a customer will be billed on their next invoice, assuming no changes are made.
// Currently this does not include Bandwidth Pooling charges.
func (softlayer_account *SoftLayer_Account) GetNextInvoiceExcel(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetNextInvoicePdf - Return an account's next invoice in PDF format. The "next invoice" is what a
// customer will be billed on their next invoice, assuming no changes are made. Currently this does not
// include Bandwidth Pooling charges.
func (softlayer_account *SoftLayer_Account) GetNextInvoicePdf(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetNextInvoicePdfDetailed - Return an account's next invoice detailed portion in PDF format. The
// "next invoice" is what a customer will be billed on their next invoice, assuming no changes are
// made. Currently this does not include Bandwidth Pooling charges.
func (softlayer_account *SoftLayer_Account) GetNextInvoicePdfDetailed(ctx *slapi.RequestContext) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetNextInvoiceZeroFeeItemCounts - <nil>
func (softlayer_account *SoftLayer_Account) GetNextInvoiceZeroFeeItemCounts(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Product_Item_Category_ZeroFee_Count, error) {
	var returnValue []*SoftLayer_Container_Product_Item_Category_ZeroFee_Count
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Account object whose ID number corresponds to the ID
// number of the init parameter passed to the SoftLayer_Account service. You can only retrieve the
// account that your portal user is assigned to.
func (softlayer_account *SoftLayer_Account) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Account, error) {
	var returnValue *SoftLayer_Account
	return returnValue, nil
}

// GetPendingCreditCardChangeRequestData - Before being approved for general use, a credit card must be
// approved by a SoftLayer agent. Once a credit card change request has been either approved or denied,
// the change request will no longer appear in the list of pending change requests. This method will
// return a list of all pending change requests as well as a portion of the data from the original
// request.
func (softlayer_account *SoftLayer_Account) GetPendingCreditCardChangeRequestData(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Account_Payment_Method_CreditCard, error) {
	var returnValue []*SoftLayer_Container_Account_Payment_Method_CreditCard
	return returnValue, nil
}

// GetReferralPartnerCommissionForecast - <nil>
func (softlayer_account *SoftLayer_Account) GetReferralPartnerCommissionForecast(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Referral_Partner_Commission, error) {
	var returnValue []*SoftLayer_Container_Referral_Partner_Commission
	return returnValue, nil
}

// GetReferralPartnerCommissionHistory - <nil>
func (softlayer_account *SoftLayer_Account) GetReferralPartnerCommissionHistory(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Referral_Partner_Commission, error) {
	var returnValue []*SoftLayer_Container_Referral_Partner_Commission
	return returnValue, nil
}

// GetReferralPartnerCommissionPending - <nil>
func (softlayer_account *SoftLayer_Account) GetReferralPartnerCommissionPending(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Referral_Partner_Commission, error) {
	var returnValue []*SoftLayer_Container_Referral_Partner_Commission
	return returnValue, nil
}

// GetSharedBlockDeviceTemplateGroups - This method returns the
// [[SoftLayer_Virtual_Guest_Block_Device_Template_Group]] objects that have been shared with this
// account
func (softlayer_account *SoftLayer_Account) GetSharedBlockDeviceTemplateGroups(ctx *slapi.RequestContext) ([]*SoftLayer_Virtual_Guest_Block_Device_Template_Group, error) {
	var returnValue []*SoftLayer_Virtual_Guest_Block_Device_Template_Group
	return returnValue, nil
}

// GetTechIncubatorProgramInfo - This method will return a SoftLayer_Container_Account_Discount_Program
// object containing the Technology Incubator Program information for this account. To be considered an
// active participant, the account must have an enrollment record with a monthly credit amount set and
// the current date must be within the range defined by the enrollment and graduation date. The
// forNextBillCycle parameter can be set to true to return a
// SoftLayer_Container_Account_Discount_Program object with information with relation to the next bill
// cycle. The forNextBillCycle parameter defaults to false.
func (softlayer_account *SoftLayer_Account) GetTechIncubatorProgramInfo(ctx *slapi.RequestContext, forNextBillCycle bool) (*SoftLayer_Container_Account_Discount_Program, error) {
	var returnValue *SoftLayer_Container_Account_Discount_Program
	return returnValue, nil
}

// GetValidSecurityCertificateEntries - Retrieve a list of valid (non-expired) security certificates
// without the sensitive certificate information. This allows non-privileged users to view and select
// security certificates when configuring associated services.
func (softlayer_account *SoftLayer_Account) GetValidSecurityCertificateEntries(ctx *slapi.RequestContext) ([]*SoftLayer_Security_Certificate_Entry, error) {
	var returnValue []*SoftLayer_Security_Certificate_Entry
	return returnValue, nil
}

// GetWindowsUpdateStatus - Retrieve a list of an account's hardware's Windows Update status. This list
// includes which servers have available updates, which servers require rebooting due to updates, which
// servers have failed retrieving updates, and which servers have failed to communicate with the
// SoftLayer private Windows Software Update Services server.
func (softlayer_account *SoftLayer_Account) GetWindowsUpdateStatus(ctx *slapi.RequestContext) ([]*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status, error) {
	var returnValue []*SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_Status
	return returnValue, nil
}

// HasAttribute - Determine if an account has an [[SoftLayer_Account_Attribute|attribute]] associated
// with it. hasAttribute() returns false if the attribute does not exist or if it does not have a
// value.
func (softlayer_account *SoftLayer_Account) HasAttribute(ctx *slapi.RequestContext, attributeType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// HourlyInstanceLimit - This method will return the limit (number) of hourly services the account is
// allowed to have.
func (softlayer_account *SoftLayer_Account) HourlyInstanceLimit(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// HourlyServerLimit - This method will return the limit (number) of hourly bare metal servers the
// account is allowed to have.
func (softlayer_account *SoftLayer_Account) HourlyServerLimit(ctx *slapi.RequestContext) (int, error) {
	var returnValue int
	return returnValue, nil
}

// RemoveAlternateCreditCard - <nil>
func (softlayer_account *SoftLayer_Account) RemoveAlternateCreditCard(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RequestCreditCardChange - Retrieve the record data associated with the submission of a Credit Card
// Change Request. Softlayer customers are permitted to request a change in Credit Card information.
// Part of the process calls for an attempt by SoftLayer to submit at $1.00 charge to the financial
// institution backing the credit card as a means of verifying that the information provided in the
// change request is valid. The data associated with this change request returned to the calling
// function. If the onlyChangeNicknameFlag parameter is set to true, the nickname of the credit card
// will be changed immediately without requiring approval by an agent. To change the nickname of the
// active payment method, pass the empty string for paymentRoleName. To change the nickname for the
// alternate credit card, pass as the paymentRoleName. vatId must be set, but the value will not be
// used and the empty string is acceptable.
func (softlayer_account *SoftLayer_Account) RequestCreditCardChange(ctx *slapi.RequestContext, request SoftLayer_Billing_Payment_Card_ChangeRequest, vatId string, paymentRoleName string, onlyChangeNicknameFlag string) (*SoftLayer_Billing_Payment_Card_ChangeRequest, error) {
	var returnValue *SoftLayer_Billing_Payment_Card_ChangeRequest
	return returnValue, nil
}

// RequestManualPayment - Retrieve the record data associated with the submission of a Manual Payment
// Request. Softlayer customers are permitted to request a manual one-time payment at a minimum amount
// of $2.00. Customers may submit a Credit Card Payment (Mastercard, Visa, American Express) or a
// PayPal payment. For Credit Card Payments, SoftLayer engages the credit card financial institution to
// submit the payment request. The financial institution's response and other data associated with the
// transaction are returned to the calling function. In the case of PayPal Payments, SoftLayer engages
// the PayPal system to initiate the PayPal payment sequence. The applicable data generated during the
// request is returned to the calling function.
func (softlayer_account *SoftLayer_Account) RequestManualPayment(ctx *slapi.RequestContext, request SoftLayer_Billing_Payment_Card_ManualPayment) (*SoftLayer_Billing_Payment_Card_ManualPayment, error) {
	var returnValue *SoftLayer_Billing_Payment_Card_ManualPayment
	return returnValue, nil
}

// RequestManualPaymentUsingCreditCardOnFile - Retrieve the record data associated with the submission
// of a Manual Payment Request for a manual payment using a credit card which is on file and does not
// require an approval process. Softlayer customers are permitted to request a manual one-time payment
// at a minimum amount of $2.00. Customers may use an existing Credit Card on file (Mastercard, Visa,
// American Express). SoftLayer engages the credit card financial institution to submit the payment
// request. The financial institution's response and other data associated with the transaction are
// returned to the calling function. The applicable data generated during the request is returned to
// the calling function.
func (softlayer_account *SoftLayer_Account) RequestManualPaymentUsingCreditCardOnFile(ctx *slapi.RequestContext, amount string, payWithAlternateCardFlag bool, note string) (*SoftLayer_Billing_Payment_Card_ManualPayment, error) {
	var returnValue *SoftLayer_Billing_Payment_Card_ManualPayment
	return returnValue, nil
}

// SetVlanSpan - Set the flag that enables or disables automatic private network spanning for a
// SoftLayer customer account. Enabling spanning allows an account's servers to talk on the same
// broadcast domain even if they reside within different private vlans.
func (softlayer_account *SoftLayer_Account) SetVlanSpan(ctx *slapi.RequestContext, enabled bool) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SwapCreditCards - <nil>
func (softlayer_account *SoftLayer_Account) SwapCreditCards(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateVpnUsersForResource - Some larger SoftLayer customer accounts may have servers and virtual
// servers on more subnets than SoftLayer's private network VPN devices can assign routes for. In those
// cases routes for individual servers and virtual servers may be assigned individually to an account's
// servers via this method. Always call this method to enable changes when manually configuring VPN
// subnet access.
func (softlayer_account *SoftLayer_Account) UpdateVpnUsersForResource(ctx *slapi.RequestContext, objectId int, objectType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ValidateManualPaymentAmount - This method checks global and account specific requirements and
// returns true if the dollar amount entered is acceptable for this account and false otherwise. Please
// note the dollar amount is in
func (softlayer_account *SoftLayer_Account) ValidateManualPaymentAmount(ctx *slapi.RequestContext, amount string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
