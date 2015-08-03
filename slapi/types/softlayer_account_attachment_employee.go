package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_Attachment_Employee - A SoftLayer_Account_Attachment_Employee models an assignment
// of a single [[SoftLayer_User_Employee|employee]] with a single [[SoftLayer_Account|account]]
type SoftLayer_Account_Attachment_Employee struct {

	// Account - A [[SoftLayer_Account|account]] that is assigned to a
	// [[SoftLayer_User_Employee|employee]].
	Account *SoftLayer_Account `json:"account"`

	// Employee - A [[SoftLayer_User_Employee|employee]] that is assigned to a
	// [[SoftLayer_Account|account]].
	Employee *SoftLayer_User_Employee `json:"employee"`

	// EmployeeRole - A [[SoftLayer_User_Employee|employee]] that is assigned to a
	// [[SoftLayer_Account|account]].
	EmployeeRole *SoftLayer_Account_Attachment_Employee_Role `json:"employeeRole"`

	// RoleId - no documentation
	RoleId int `json:"roleId"`
}

func (softlayer_account_attachment_employee *SoftLayer_Account_Attachment_Employee) String() string {
	return "SoftLayer_Account_Attachment_Employee"
}
