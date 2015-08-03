package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Protection_Address - <nil>
type SoftLayer_Network_Protection_Address struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// DepartmentId - <nil>
	DepartmentId int `json:"departmentId"`

	// IpAddress - <nil>
	IpAddress string `json:"ipAddress"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location"`

	// ManagementMethodType - <nil>
	ManagementMethodType string `json:"managementMethodType"`

	// ModifiedUser - <nil>
	ModifiedUser *SoftLayer_User_Employee `json:"modifiedUser"`

	// PrimaryRouter - <nil>
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider"`

	// Subnet - <nil>
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetIpAddress - <nil>
	SubnetIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"subnetIpAddress"`

	// TerminatedUser - <nil>
	TerminatedUser *SoftLayer_User_Employee `json:"terminatedUser"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TransactionCount - no documentation
	TransactionCount uint64 `json:"transactionCount"`

	// Transactions - <nil>
	Transactions []*SoftLayer_Provisioning_Version1_Transaction `json:"transactions"`

	// UserDepartment - <nil>
	UserDepartment *SoftLayer_User_Employee_Department `json:"userDepartment"`

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Employee `json:"userRecord"`
}

func (softlayer_network_protection_address *SoftLayer_Network_Protection_Address) String() string {
	return "SoftLayer_Network_Protection_Address"
}
