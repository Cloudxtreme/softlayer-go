package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_User_Employee_Department - SoftLayer_User_Employee_Department models a department within
// SoftLayer's internal employee hierarchy. Common departments include Support, Sales, Accounting,
// Development, Systems, and Networking.
type SoftLayer_User_Employee_Department struct {

	// Name - The name of one of SoftLayer's employee departments.
	Name string `json:"name"`
}

func (softlayer_user_employee_department *SoftLayer_User_Employee_Department) String() string {
	return "SoftLayer_User_Employee_Department"
}
