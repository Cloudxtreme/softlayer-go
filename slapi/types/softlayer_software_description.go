package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Software_Description - This class holds a description for a specific installation of a
// Software Component. SoftLayer_Software_Licenses tie a Software Component (A specific installation on
// a piece of hardware) to it's description. The "Manufacturer" and "Name" properties of a
// SoftLayer_Software_Description are used by the framework to factory specific objects, objects that
// may have special methods for that specific piece of software, or objects that contain application
// specific data, such as default ports. For example, if you create a SoftLayer_Software_Component
// who's SoftLayer_Software_License points to the SoftLayer_Software_Description for "Swsoft" "Plesk",
// you'll actually get a SoftLayer_Software_Component_Swsoft_Plesk object.
type SoftLayer_Software_Description struct {

	// Id - An ID number to identify this Software Description.
	Id int `json:"id"`

	// VirtualizationPlatform - This is set to '1' if this Software Description a platform for hosting
	// virtual servers.
	VirtualizationPlatform int `json:"virtualizationPlatform"`

	// LicenseTermUnit - The unit of measurement (day, month, or year) for license registration. Used in
	// conjunction with licenseTermValue to determine overall license registration length of a new license.
	LicenseTermUnit string `json:"licenseTermUnit"`

	// UpgradeSoftwareDescriptionId - Contains the ID of the suggested upgrade from this
	// Software_Description to a more powerful software installation.
	UpgradeSoftwareDescriptionId int `json:"upgradeSoftwareDescriptionId"`

	// VirtualLicense - This is set to '1' if this Software Description can be licensed to a Virtual
	// Machine (an IP address).
	VirtualLicense int `json:"virtualLicense"`

	// LicenseTermValue - The number of units (licenseTermUnit) a new license is valid for at the time of
	// registration.
	LicenseTermValue int `json:"licenseTermValue"`

	// LongDescription - The manufacturer, name and version of a piece of software.
	LongDescription string `json:"longDescription"`

	// Manufacturer - The name of the manufacturer for this specific piece of software. This name is used
	// by SoftLayer_Software_Component to tailor make (factory) specific types of Software Components that
	// know details like default ports.
	Manufacturer string `json:"manufacturer"`

	// ControlPanel - This is set to '1' if this Software Description describes a Control Panel.
	ControlPanel int `json:"controlPanel"`

	// Name - The name of this specific piece of software. This name is used by
	// SoftLayer_Software_Component to tailor make (factory) specific types of Software Components that
	// know details like default ports.
	Name string `json:"name"`

	// ReferenceCode - A reference code is structured as three tokens separated by underscores. The first
	// token represents the product, the second is the version of the product, and the third is whether the
	// software is 32 or 64bit.
	ReferenceCode string `json:"referenceCode"`

	// UpgradeSwDescId - Contains the ID of the suggested upgrade from this Software_Description to a more
	// powerful software installation. (Deprecated - Use upgradeSoftwareDescriptionId)
	UpgradeSwDescId int `json:"upgradeSwDescId"`

	// OperatingSystem - This is set to '1' if this Software Description describes an Operating System.
	OperatingSystem int `json:"operatingSystem"`

	// Version - no documentation
	Version string `json:"version"`
}

func (softlayer_software_description *SoftLayer_Software_Description) String() string {
	return "SoftLayer_Software_Description"
}

// SoftLayer_Software_Description_Extended is SoftLayer_Software_Description with all maskable types.
type SoftLayer_Software_Description_Extended struct {
	SoftLayer_Software_Description

	// ValidFilesystemTypes - <nil>
	ValidFilesystemTypes []*SoftLayer_Configuration_Storage_Filesystem_Type `json:"validFilesystemTypes"`

	// AttributeCount - no documentation
	AttributeCount uint64 `json:"attributeCount"`

	// FeatureCount - A count of the feature attributes of a software description.
	FeatureCount uint64 `json:"featureCount"`

	// LatestVersionCount - A count of the latest version of a software description.
	LatestVersionCount uint64 `json:"latestVersionCount"`

	// Attributes - <nil>
	Attributes []*SoftLayer_Software_Description_Attribute `json:"attributes"`

	// AverageInstallationDuration - The average amount of time that a software description takes to
	// install.
	AverageInstallationDuration int `json:"averageInstallationDuration"`

	// ReloadTransactionGroup - The transaction group that a software description belongs to. A transaction
	// group is a sequence of transactions that must be performed in a specific order for the installation
	// of software.
	ReloadTransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"reloadTransactionGroup"`

	// SoftwareLicenseCount - A count of software Licenses that govern this Software Description.
	SoftwareLicenseCount uint64 `json:"softwareLicenseCount"`

	// ValidFilesystemTypeCount - no documentation
	ValidFilesystemTypeCount uint64 `json:"validFilesystemTypeCount"`

	// CompatibleSoftwareDescriptionCount - A count of a list of the software descriptions that are
	// compatible with this software description.
	CompatibleSoftwareDescriptionCount uint64 `json:"compatibleSoftwareDescriptionCount"`

	// ProductItemCount - A count of the various product items to which this software description is
	// linked.
	ProductItemCount uint64 `json:"productItemCount"`

	// CompatibleSoftwareDescriptions - A list of the software descriptions that are compatible with this
	// software description.
	CompatibleSoftwareDescriptions []*SoftLayer_Software_Description `json:"compatibleSoftwareDescriptions"`

	// LatestVersion - no documentation
	LatestVersion []*SoftLayer_Software_Description `json:"latestVersion"`

	// ProvisionTransactionGroup - This details the provisioning transaction group for this software. This
	// is only valid for Operating System software.
	ProvisionTransactionGroup *SoftLayer_Provisioning_Version1_Transaction_Group `json:"provisionTransactionGroup"`

	// ProductItems - The various product items to which this software description is linked.
	ProductItems []*SoftLayer_Product_Item `json:"productItems"`

	// RequiredUser - The default user created for a given a software description.
	RequiredUser string `json:"requiredUser"`

	// SoftwareLicenses - Software Licenses that govern this Software Description.
	SoftwareLicenses []*SoftLayer_Software_License `json:"softwareLicenses"`

	// UpgradeSoftwareDescription - A suggestion for an upgrade path from this Software Description
	UpgradeSoftwareDescription *SoftLayer_Software_Description `json:"upgradeSoftwareDescription"`

	// Features - no documentation
	Features []*SoftLayer_Software_Description_Feature `json:"features"`

	// UpgradeSwDesc - A suggestion for an upgrade path from this Software Description (Deprecated - Use
	// upgradeSoftwareDescription)
	UpgradeSwDesc *SoftLayer_Software_Description `json:"upgradeSwDesc"`
}

func (softlayer_software_description *SoftLayer_Software_Description_Extended) String() string {
	return "SoftLayer_Software_Description"
}
