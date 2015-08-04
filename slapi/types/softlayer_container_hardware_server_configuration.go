package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Hardware_Server_Configuration - The
// SoftLayer_Container_Hardware_Server_Configuration data type contains information relating to a
// server's item price information, and hard drive partition information.
type SoftLayer_Container_Hardware_Server_Configuration struct {

	// SshKeyIds - IDs to SoftLayer_Security_Ssh_Key objects on the current account which will be added to
	// the server for authentication. SSH Keys will not be added to servers with Microsoft Windows.
	SshKeyIds []int `json:"sshKeyIds,omitempty"`

	// UpgradeHardDriveFirmware - A flag indicating that the firmware on all hard drives will be updated
	// when installing the operating system.
	UpgradeHardDriveFirmware int `json:"upgradeHardDriveFirmware,omitempty"`

	// DriveRetentionFlag - A flag indicating that the primary drive will be converted to a portable
	// storage volume during an Operating System reload.
	DriveRetentionFlag bool `json:"driveRetentionFlag,omitempty"`

	// HardDrives - The hard drive partitions that a server can be partitioned with.
	HardDrives []*SoftLayer_Hardware_Component `json:"hardDrives,omitempty"`

	// ImageTemplateId - An Image Template ID [[SoftLayer_Virtual_Guest_Block_Device_Template_Group]] that
	// will be deployed to the host. If provided no item prices are required.
	ImageTemplateId int `json:"imageTemplateId,omitempty"`

	// ItemPrices - The item prices that a server can be configured with.
	ItemPrices []*SoftLayer_Product_Item_Price `json:"itemPrices,omitempty"`

	// UpgradeBios - A flag indicating that the the will not be updated when installing the operating
	// system.
	UpgradeBios int `json:"upgradeBios,omitempty"`

	// AddToSparePoolAfterOsReload - A flag indicating that the server will be moved into the spare pool
	// after an Operating system reload.
	AddToSparePoolAfterOsReload int `json:"addToSparePoolAfterOsReload,omitempty"`

	// CustomProvisionScriptUri - The customer provision uri will be used to download and execute a
	// customer defined script on the host at the end of provisioning.
	CustomProvisionScriptUri string `json:"customProvisionScriptUri,omitempty"`

	// EraseHardDrives - A flag indicating that all data will be erased from drives during an Operating
	// System reload.
	EraseHardDrives int `json:"eraseHardDrives,omitempty"`

	// ResetIpmiPassword - A flag indicating that the remote management cards password will be reset.
	ResetIpmiPassword int `json:"resetIpmiPassword,omitempty"`
}

func (softlayer_container_hardware_server_configuration *SoftLayer_Container_Hardware_Server_Configuration) String() string {
	return "SoftLayer_Container_Hardware_Server_Configuration"
}
