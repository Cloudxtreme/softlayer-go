package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Hardware_Server_Configuration - The
// SoftLayer_Container_Hardware_Server_Configuration data type contains information relating to a
// server's item price information, and hard drive partition information.
type SoftLayer_Container_Hardware_Server_Configuration struct {

	// AddToSparePoolAfterOsReload - A flag indicating that the server will be moved into the spare pool
	// after an Operating system reload.
	AddToSparePoolAfterOsReload int `json:"addToSparePoolAfterOsReload"`

	// CustomProvisionScriptUri - The customer provision uri will be used to download and execute a
	// customer defined script on the host at the end of provisioning.
	CustomProvisionScriptUri string `json:"customProvisionScriptUri"`

	// DriveRetentionFlag - A flag indicating that the primary drive will be converted to a portable
	// storage volume during an Operating System reload.
	DriveRetentionFlag bool `json:"driveRetentionFlag"`

	// EraseHardDrives - A flag indicating that all data will be erased from drives during an Operating
	// System reload.
	EraseHardDrives int `json:"eraseHardDrives"`

	// HardDrives - The hard drive partitions that a server can be partitioned with.
	HardDrives []*SoftLayer_Hardware_Component `json:"hardDrives"`

	// ImageTemplateId - An Image Template ID [[SoftLayer_Virtual_Guest_Block_Device_Template_Group]] that
	// will be deployed to the host. If provided no item prices are required.
	ImageTemplateId int `json:"imageTemplateId"`

	// ItemPrices - The item prices that a server can be configured with.
	ItemPrices []*SoftLayer_Product_Item_Price `json:"itemPrices"`

	// ResetIpmiPassword - A flag indicating that the remote management cards password will be reset.
	ResetIpmiPassword int `json:"resetIpmiPassword"`

	// SshKeyIds - IDs to SoftLayer_Security_Ssh_Key objects on the current account which will be added to
	// the server for authentication. SSH Keys will not be added to servers with Microsoft Windows.
	SshKeyIds []int `json:"sshKeyIds"`

	// UpgradeBios - A flag indicating that the the will not be updated when installing the operating
	// system.
	UpgradeBios int `json:"upgradeBios"`

	// UpgradeHardDriveFirmware - A flag indicating that the firmware on all hard drives will be updated
	// when installing the operating system.
	UpgradeHardDriveFirmware int `json:"upgradeHardDriveFirmware"`
}
