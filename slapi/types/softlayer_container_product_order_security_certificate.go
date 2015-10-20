package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Container_Product_Order_Security_Certificate - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype contains everything
// required to place a secure certificate order with SoftLayer.
type SoftLayer_Container_Product_Order_Security_Certificate struct {

	// RegionalGroup - Specifying a regional group name allows you to not worry about placing your server
	// or service at a specific datacenter, but to any datacenter within that regional group. See
	// [[SoftLayer_Location_Group_Regional]] to get a list of available regional group names. location and
	// regionalGroup are mutually exclusive on an order container. If both location and regionalGroup are
	// provided, an exception will be thrown indicating that only 1 is allowed. If a regional group is
	// provided and VLANs are specified (within the hardware or virtualGuests properties), we will use the
	// datacenter where the VLANs are located. If no VLANs are specified, we will use the preferred
	// datacenter on the regional group object.
	RegionalGroup string `json:"regionalGroup,omitempty"`

	// Prices - This is a collection of [[SoftLayer_Product_Item_Price]] objects. The only required
	// property to populate for an item price object when ordering is its id - all other supplied
	// information about the price (e.g., recurringFee, setupFee, etc.) will be ignored. Unless the
	// [[SoftLayer_Product_Package]] associated with the order allows for preset prices, this property is
	// required to place an order.
	Prices []*SoftLayer_Product_Item_Price `json:"prices,omitempty"`

	// AdministrativeContact - The administrator contact associated with a SSL certificate. If the contact
	// is not provided the technical contact will be used. If the address is not provided the organization
	// information address will be used.
	AdministrativeContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"administrativeContact,omitempty"`

	// ItemCategoryQuestionAnswers - The collection of
	// [[SoftLayer_Container_Product_Item_Category_Question_Answer]] for any product category that has
	// additional questions requiring user input.
	ItemCategoryQuestionAnswers []*SoftLayer_Container_Product_Item_Category_Question_Answer `json:"itemCategoryQuestionAnswers,omitempty"`

	// TotalSetupTax - no documentation
	TotalSetupTax slapi.Float64 `json:"totalSetupTax,omitempty"`

	// TotalRecurringTax - no documentation
	TotalRecurringTax slapi.Float64 `json:"totalRecurringTax,omitempty"`

	// OrganizationInformation - The organization information associated with a SSL certificate.
	OrganizationInformation *SoftLayer_Container_Product_Order_Attribute_Organization `json:"organizationInformation,omitempty"`

	// ServerType - The server type. This is the name from a
	// [[SoftLayer_Security_Certificate_Request_ServerType]] object.
	ServerType string `json:"serverType,omitempty"`

	// PresetId - A preset configuration id for the package. Is required if not submitting any prices.
	PresetId int `json:"presetId,omitempty"`

	// BigDataOrderFlag - Flag for identifying an order for Big Data Deployment.
	BigDataOrderFlag bool `json:"bigDataOrderFlag,omitempty"`

	// IsManagedOrder - Flag to identify a "managed" order. This value is set internally.
	IsManagedOrder int `json:"isManagedOrder,omitempty"`

	// BillingContact - The billing contact associated with a SSL certificate. If the contact is not
	// provided the technical contact will be used. If the address is not provided the organization
	// information address will be used.
	BillingContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"billingContact,omitempty"`

	// PrivateCloudOrderType - Type of Virtual Server (Private Node) order. Potential values:
	PrivateCloudOrderType string `json:"privateCloudOrderType,omitempty"`

	// PackageId - The [[SoftLayer_Product_Package]] id for an order container. This is required to place
	// an order.
	PackageId int `json:"packageId,omitempty"`

	// Message - A generic message about the order. Does not need to be sent in with any orders.
	Message string `json:"message,omitempty"`

	// PostTaxRecurring - The post-tax recurring charge for the order. This is the sum of preTaxRecurring +
	// totalRecurringTax.
	PostTaxRecurring slapi.Float64 `json:"postTaxRecurring,omitempty"`

	// VirtualGuests - For virtual guest (virtual server) orders, this property is required if you did not
	// specify data in the hardware property. This is an array of [[SoftLayer_Virtual_Guest]] objects. The
	// hostname and domain properties are required for each virtual guest object. There is no need to
	// specify data in this property and the hardware property - only one is required for virtual server
	// orders.
	VirtualGuests []*SoftLayer_Virtual_Guest `json:"virtualGuests,omitempty"`

	// CancelUrl - The URL to which PayPal redirects browser after checkout has been canceled before
	// completion of a payment.
	CancelUrl string `json:"cancelUrl,omitempty"`

	// CertificateSigningRequest - The base64 encoded string that sent from an applicant to a certificate
	// authority. The CSR contains information identifying the applicant and the public key chosen by the
	// applicant. The corresponding private key should not be included.
	CertificateSigningRequest string `json:"certificateSigningRequest,omitempty"`

	// SendQuoteEmailFlag - This flag indicates that the quote should be sent to the email address
	// associated with the account or order.
	SendQuoteEmailFlag bool `json:"sendQuoteEmailFlag,omitempty"`

	// PreTaxRecurringMonthly - The pre-tax monthly recurring total of the order. If there are only hourly
	// prices on the order, this value will be 0.
	PreTaxRecurringMonthly slapi.Float64 `json:"preTaxRecurringMonthly,omitempty"`

	// TaxCacheHash - The order container may not contain the final tax rates when it is returned from
	// [[SoftLayer_Product_Order/verifyOrder|verifyOrder]]. This hash will facilitate checking if the tax
	// rates have finished being calculated and retrieving the accurate tax rate values.
	TaxCacheHash string `json:"taxCacheHash,omitempty"`

	// ContainerIdentifier - User-specified description to identify a particular order container. This is
	// useful if you have a multi-configuration order (multiple orderContainers ) and you want to be able
	// to easily determine one from another. Populating this value may be helpful if an exception is thrown
	// when placing an order and it's tied to a specific order container.
	ContainerIdentifier string `json:"containerIdentifier,omitempty"`

	// LocationObject - This [[SoftLayer_Location]] object will be determined from the location property
	// and will be returned in the order verification or placement response. Any value specified here will
	// get overwritten by the verification process.
	LocationObject *SoftLayer_Location `json:"locationObject,omitempty"`

	// OrderApproverEmailAddress - The email address that can approve a secure certificate order.
	OrderApproverEmailAddress string `json:"orderApproverEmailAddress,omitempty"`

	// ProratedOrderTotal - This is the same as the proratedInitialCharge, except the balance on the
	// account is ignored. This is the prorated total amount of the order.
	ProratedOrderTotal slapi.Float64 `json:"proratedOrderTotal,omitempty"`

	// ResourceGroupId - An optional resource group identifier specifying the resource group to attach the
	// order to
	ResourceGroupId int `json:"resourceGroupId,omitempty"`

	// Quantity - no documentation
	Quantity int `json:"quantity,omitempty"`

	// SshKeys - The containers which hold SoftLayer_Security_Ssh_Key IDs to add to their respective
	// servers. The order of containers passed in needs to match the order they are assigned to either
	// hardware or virtualGuests. SSH Keys will not be assigned for servers with Microsoft Windows.
	SshKeys []*SoftLayer_Container_Product_Order_SshKeys `json:"sshKeys,omitempty"`

	// OrderContainers - Orders may contain an array of configurations. Populating this property allows you
	// to purchase multiple configurations within a single order. Each order container will have its own
	// individual settings independent of the other order containers. For example, it is possible to order
	// a bare metal server in one configuration and a virtual server in another. If orderContainers is
	// populated on the base order container, most of the configuration-specific properties are ignored on
	// the base container. For example, prices , location and packageId will be ignored on the base
	// container, but since the billingInformation is a property that's not specific to a single order
	// container (but the order as a whole) it must be populated on the base container.
	OrderContainers []*SoftLayer_Container_Product_Order `json:"orderContainers,omitempty"`

	// StorageGroups - For orders that want to add storage groups such as across multiple disks, simply add
	// [[SoftLayer_Container_Product_Order_Storage_Group]] objects to this array. Storage groups will only
	// be used if the disk controller price is selected. Any other disk controller types will ignore the
	// storage groups set here. The first storage group in this array will be considered the primary
	// storage group, which is used for the OS. Any other storage groups will act as data storage.
	StorageGroups []*SoftLayer_Container_Product_Order_Storage_Group `json:"storageGroups,omitempty"`

	// ContainerSplHash - This hash is internally-generated and is used to for tracking order containers.
	ContainerSplHash string `json:"containerSplHash,omitempty"`

	// DisplayLayerSessionId - This is the configuration identifier for tracking orders on the order forms.
	DisplayLayerSessionId string `json:"displayLayerSessionId,omitempty"`

	// PrivateCloudOrderFlag - Flag for identifying a container as Virtual Server (Private Node).
	PrivateCloudOrderFlag bool `json:"privateCloudOrderFlag,omitempty"`

	// UseHourlyPricing - An optional flag to use hourly pricing instead of standard monthly pricing.
	UseHourlyPricing bool `json:"useHourlyPricing,omitempty"`

	// ServerCoreCount - The number of cores for the server being ordered. This value is set internally.
	ServerCoreCount int `json:"serverCoreCount,omitempty"`

	// StepId - An optional parameter for step-based order processing.
	StepId int `json:"stepId,omitempty"`

	// FlexibleCreditProgramPrice - The [[SoftLayer_Product_Item_Price]] for the Flexible Credit Program
	// discount. The oneTimeFee field contains the calculated discount being applied to the order.
	FlexibleCreditProgramPrice *SoftLayer_Product_Item_Price `json:"flexibleCreditProgramPrice,omitempty"`

	// Priorities - no documentation
	Priorities []string `json:"priorities,omitempty"`

	// Properties - no documentation
	Properties []*SoftLayer_Container_Product_Order_Property `json:"properties,omitempty"`

	// OrderVerificationExceptions - Collection of exceptions resulting from the verification of the order.
	// This value is set internally and is not required for end users when placing an order. When placing
	// API orders, users can use this value to determine the container-specific exception that was thrown.
	OrderVerificationExceptions []*SoftLayer_Container_Exception `json:"orderVerificationExceptions,omitempty"`

	// TechIncubatorItemPrice - The SoftLayer_Product_Item_Price for the Tech Incubator discount. The
	// oneTimeFee field contain the calculated discount being applied to the order.
	TechIncubatorItemPrice *SoftLayer_Product_Item_Price `json:"techIncubatorItemPrice,omitempty"`

	// ExtendedHardwareTesting - <nil>
	ExtendedHardwareTesting bool `json:"extendedHardwareTesting,omitempty"`

	// BillingOrderItemId - This is the ID of the [[SoftLayer_Billing_Order_Item]] of this
	// configuration/container. It is used for rebuilding an order container from a quote and is set
	// automatically.
	BillingOrderItemId int `json:"billingOrderItemId,omitempty"`

	// PreTaxSetup - no documentation
	PreTaxSetup slapi.Float64 `json:"preTaxSetup,omitempty"`

	// SourceVirtualGuestId - An optional computing instance identifier to be used as an installation base
	// for a computing instance order
	SourceVirtualGuestId int `json:"sourceVirtualGuestId,omitempty"`

	// QuoteName - no documentation
	QuoteName string `json:"quoteName,omitempty"`

	// PreTaxRecurring - The pre-tax recurring total of the order. If there are mixed monthly and hourly
	// prices on the order, this will be the sum of preTaxRecurringHourly and preTaxRecurringMonthly.
	PreTaxRecurring slapi.Float64 `json:"preTaxRecurring,omitempty"`

	// PresaleEvent - If there are any presale events available for an order, this value will be populated.
	// It is set internally and is not required for end users when placing an order. See
	// [[SoftLayer_Sales_Presale_Event]] for more info.
	PresaleEvent *SoftLayer_Sales_Presale_Event `json:"presaleEvent,omitempty"`

	// ImageTemplateId - An optional virtual disk image template identifier to be used as an installation
	// base for a computing instance order
	ImageTemplateId int `json:"imageTemplateId,omitempty"`

	// RenewalFlag - Indicates if it is an renewal order of an existing SSL certificate.
	RenewalFlag bool `json:"renewalFlag,omitempty"`

	// TechnicalContact - The technical contact associated with a SSL certificate. If the address is not
	// provided the organization information address will be used.
	TechnicalContact *SoftLayer_Container_Product_Order_Attribute_Contact `json:"technicalContact,omitempty"`

	// CurrencyShortName - no documentation
	CurrencyShortName string `json:"currencyShortName,omitempty"`

	// PaymentType - The Payment Type is Optional. If nothing is sent in, then the normal method of payment
	// will be used. For paypal customers, this means a paypalToken will be returned in the receipt. This
	// token is to be used on the paypal website to complete the order. For Credit Card customers, the card
	// on file in our system will be used to make an initial authorization. To force the order to use a
	// payment type, use one of the following: or
	PaymentType string `json:"paymentType,omitempty"`

	// BillingInformation - Billing Information associated with an order. For existing customers this
	// information is completely ignored. Do not send this information for existing customers.
	BillingInformation *SoftLayer_Container_Product_Order_Billing_Information `json:"billingInformation,omitempty"`

	// PostTaxSetup - The post-tax setup fees of the order. This is the sum of preTaxSetup + totalSetupTax;
	PostTaxSetup slapi.Float64 `json:"postTaxSetup,omitempty"`

	// PrimaryDiskPartitionId - The id of a [[SoftLayer_Hardware_Component_Partition_Template]]. This
	// property is optional. If no partition template is provided, a default will be used according to the
	// operating system chosen with the order. Using the
	// [[SoftLayer_Hardware_Component_Partition_OperatingSystem]] service, getPartitionTemplates will
	// return those available for the particular operating system.
	PrimaryDiskPartitionId int `json:"primaryDiskPartitionId,omitempty"`

	// ImageTemplateGlobalIdentifier - An optional virtual disk image template identifier to be used as an
	// installation base for a computing instance order
	ImageTemplateGlobalIdentifier string `json:"imageTemplateGlobalIdentifier,omitempty"`

	// PostTaxRecurringMonthly - The post-tax recurring monthly charge for the order. This is the sum of
	// preTaxRecurringMonthly + totalRecurringTax.
	PostTaxRecurringMonthly slapi.Float64 `json:"postTaxRecurringMonthly,omitempty"`

	// ResourceGroupName - This variable specifies the name of the resource group the server configuration
	// belongs to. For MongoDB Replica sets, it would be the replica set name.
	ResourceGroupName string `json:"resourceGroupName,omitempty"`

	// ValidityMonths - The period that a SSL certificate is valid for. For example, 12, 24, 36
	ValidityMonths int `json:"validityMonths,omitempty"`

	// OrderHostnames - no documentation
	OrderHostnames []string `json:"orderHostnames,omitempty"`

	// PromotionCode - no documentation
	PromotionCode string `json:"promotionCode,omitempty"`

	// PostTaxRecurringHourly - The post-tax recurring hourly charge for the order. Since taxes are not
	// calculated for hourly orders, this value will be the same as preTaxRecurringHourly.
	PostTaxRecurringHourly slapi.Float64 `json:"postTaxRecurringHourly,omitempty"`

	// ReturnUrl - The URL to which PayPal redirects browser after a payment is completed.
	ReturnUrl string `json:"returnUrl,omitempty"`

	// ServerCount - no documentation
	ServerCount int `json:"serverCount,omitempty"`

	// ResourceGroupTemplateId - An optional resource group template identifier to be used as a deployment
	// base for a Virtual Server (Private Node) order.
	ResourceGroupTemplateId int `json:"resourceGroupTemplateId,omitempty"`

	// ProratedInitialCharge - The Prorated Initial Charge plus the balance on the account. Only the
	// recurring fees are prorated. Here's how the calculation works: We take the postTaxRecurring value
	// and we prorate it based on the time between now and the next bill date for this account. After this,
	// we add in the setup fee since this is not prorated. Then, if there is a balance on the account, we
	// add that to the account. In the event that there is a credit balance on the account, we will
	// subtract this amount from the order total. If the credit balance on the account is greater than the
	// prorated initial charge, the order will go through without a charge to the credit card on the
	// account or requiring a paypal payment. The credit on the account will be reduced by the order total,
	// and the order will await approval from sales, as normal. If there is a pending order already in the
	// system, We will ignore the balance on the account completely, in the calculation of the initial
	// charge. This is to protect against two orders coming into the system and getting the benefit of a
	// credit balance, or worse, both orders being charged the order amount + the balance on the account.
	ProratedInitialCharge slapi.Float64 `json:"proratedInitialCharge,omitempty"`

	// Location - The [[SoftLayer_Location_Region]] keyname or specific [[SoftLayer_Location_Datacenter]]
	// id where the order should be provisioned. If this value is provided and the regionalGroup property
	// is also specified, an exception will be thrown indicating that only 1 is allowed.
	Location string `json:"location,omitempty"`

	// Hardware - For orders that contain servers (bare metal, virtual server, big data, etc.), the
	// hardware property is required. This property is an array of [[SoftLayer_Hardware]] objects. The
	// hostname and domain properties are required for each hardware object. Note that virtual server
	// ([[SoftLayer_Container_Product_Order_Virtual_Guest]]) orders may populate this field instead of the
	// virtualGuests property.
	Hardware []*SoftLayer_Hardware `json:"hardware,omitempty"`

	// PreTaxRecurringHourly - The pre-tax hourly recurring total of the order. If there are only monthly
	// prices on the order, this value will be 0.
	PreTaxRecurringHourly slapi.Float64 `json:"preTaxRecurringHourly,omitempty"`

	// TaxCompletedFlag - Flag to indicate if the order container has the final tax rates for the order.
	// Some tax rates are calculated in the background because they take longer, and they might not be
	// finished when the container is returned from [[SoftLayer_Product_Order/verifyOrder|verifyOrder]].
	TaxCompletedFlag bool `json:"taxCompletedFlag,omitempty"`

	// DeviceFingerprintId - no documentation
	DeviceFingerprintId string `json:"deviceFingerprintId,omitempty"`

	// ProvisionScripts - The URLs for scripts to execute on their respective servers after they have been
	// provisioned. Provision scripts are not available for Microsoft Windows servers.
	ProvisionScripts []string `json:"provisionScripts,omitempty"`
}

func (softlayer_container_product_order_security_certificate *SoftLayer_Container_Product_Order_Security_Certificate) String() string {
	return "SoftLayer_Container_Product_Order_Security_Certificate"
}
