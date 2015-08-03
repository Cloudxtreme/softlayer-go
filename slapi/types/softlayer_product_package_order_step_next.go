package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Order_Step_Next - This datatype simply describes which steps are next in
// line for ordering.
type SoftLayer_Product_Package_Order_Step_Next struct {

	// NextOrderStepId - The unique identifier for SoftLayer_Product_Package_Order_Step for the next step
	// in the process.
	NextOrderStepId int `json:"nextOrderStepId"`

	// OrderStepId - The unique identifier for SoftLayer_Product_Package_Order_Step for the current step.
	OrderStepId int `json:"orderStepId"`

	// Id - no documentation
	Id int `json:"id"`
}

func (softlayer_product_package_order_step_next *SoftLayer_Product_Package_Order_Step_Next) String() string {
	return "SoftLayer_Product_Package_Order_Step_Next"
}

// SoftLayer_Product_Package_Order_Step_Next_Extended is SoftLayer_Product_Package_Order_Step_Next with all maskable types.
type SoftLayer_Product_Package_Order_Step_Next_Extended struct {
	SoftLayer_Product_Package_Order_Step_Next

	// Step - The SoftLayer_Product_Package_Order_Step to which this object belongs.
	Step *SoftLayer_Product_Package_Order_Step `json:"step"`
}

func (softlayer_product_package_order_step_next *SoftLayer_Product_Package_Order_Step_Next_Extended) String() string {
	return "SoftLayer_Product_Package_Order_Step_Next"
}
