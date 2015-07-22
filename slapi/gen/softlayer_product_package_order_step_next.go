package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Order_Step_Next - This datatype simply describes which steps are next in
// line for ordering.
type SoftLayer_Product_Package_Order_Step_Next struct {

	// Id - no documentation
	Id int `json:"id"`

	// NextOrderStepId - The unique identifier for SoftLayer_Product_Package_Order_Step for the next step
	// in the process.
	NextOrderStepId int `json:"nextOrderStepId"`

	// OrderStepId - The unique identifier for SoftLayer_Product_Package_Order_Step for the current step.
	OrderStepId int `json:"orderStepId"`

	// Step - The SoftLayer_Product_Package_Order_Step to which this object belongs.
	Step *SoftLayer_Product_Package_Order_Step `json:"step"`
}
