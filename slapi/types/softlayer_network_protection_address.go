package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Protection_Address - <nil>
type SoftLayer_Network_Protection_Address struct {

	// ManagementMethodType - <nil>
	ManagementMethodType string `json:"managementMethodType,omitempty"`

	// DepartmentId - <nil>
	DepartmentId int `json:"departmentId,omitempty"`

	// IpAddress - <nil>
	IpAddress string `json:"ipAddress,omitempty"`

	// Subnet - <nil>
	Subnet *SoftLayer_Network_Subnet `json:"subnet,omitempty"`

	// UserDepartment - <nil>
	UserDepartment *SoftLayer_User_Employee_Department `json:"userDepartment,omitempty"`

	// UserRecord - <nil>
	UserRecord *SoftLayer_User_Employee `json:"userRecord,omitempty"`

	// ModifiedUser - <nil>
	ModifiedUser *SoftLayer_User_Employee `json:"modifiedUser,omitempty"`

	// PrimaryRouter - <nil>
	PrimaryRouter *SoftLayer_Hardware_Router `json:"primaryRouter,omitempty"`

	// SubnetIpAddress - <nil>
	SubnetIpAddress *SoftLayer_Network_Subnet_IpAddress `json:"subnetIpAddress,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// ServiceProvider - <nil>
	ServiceProvider *SoftLayer_Service_Provider `json:"serviceProvider,omitempty"`

	// TerminatedUser - <nil>
	TerminatedUser *SoftLayer_User_Employee `json:"terminatedUser,omitempty"`

	// Transactions - <nil>
	Transactions []*SoftLayer_Provisioning_Version1_Transaction `json:"transactions,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Location - <nil>
	Location *SoftLayer_Location `json:"location,omitempty"`

	// TransactionCount - no documentation
	TransactionCount uint64 `json:"transactionCount,omitempty"`
}

func (softlayer_network_protection_address *SoftLayer_Network_Protection_Address) String() string {
	return "SoftLayer_Network_Protection_Address"
}
