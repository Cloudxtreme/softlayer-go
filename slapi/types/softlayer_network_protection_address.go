package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Protection_Address - <nil>
type SoftLayer_Network_Protection_Address struct {

	// ManagementMethodType - <nil>
	ManagementMethodType string `json:"managementMethodType"`

	// DepartmentId - <nil>
	DepartmentId int `json:"departmentId"`

	// IpAddress - <nil>
	IpAddress string `json:"ipAddress"`
}

func (softlayer_network_protection_address *SoftLayer_Network_Protection_Address) String() string {
	return "SoftLayer_Network_Protection_Address"
}

// SoftLayer_Network_Protection_Address_Extended is SoftLayer_Network_Protection_Address with all maskable types.
type SoftLayer_Network_Protection_Address_Extended struct {
	SoftLayer_Network_Protection_Address

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// TerminatedUser - <nil>
	TerminatedUser *SoftLayer_User_Employee `json:"terminatedUser"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TransactionCount - no documentation
	TransactionCount uint64 `json:"transactionCount"`

	// PrimaryRouter - <nil>
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// ModifiedUser - <nil>
	ModifiedUser *SoftLayer_User_Employee `json:"modifiedUser"`

	// SubnetIpAddress - <nil>
	SubnetIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"subnetIpAddress"`

	// Transactions - <nil>
	Transactions []*SoftLayer_Provisioning_Version1_Transaction `json:"transactions"`

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Employee `json:"userRecord"`

	// Subnet - <nil>
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// UserDepartment - <nil>
	UserDepartment *SoftLayer_User_Employee_Department `json:"userDepartment"`
}

func (softlayer_network_protection_address *SoftLayer_Network_Protection_Address_Extended) String() string {
	return "SoftLayer_Network_Protection_Address"
}
