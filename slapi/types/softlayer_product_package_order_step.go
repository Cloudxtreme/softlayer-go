package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Package_Order_Step - Each package has at least 1 step to the ordering process. For
// server orders, there are many. Each step has certain item categories which are displayed. This type
// describes the steps for each package.
type SoftLayer_Product_Package_Order_Step struct {

	// Id - The unique identifier for this object. It is not used anywhere but in this object.
	Id int `json:"id,omitempty"`

	// Step - The number of the step in the order process for this package. These are sequential and only
	// needed for step-based ordering.
	Step string `json:"step,omitempty"`
}

func (softlayer_product_package_order_step *SoftLayer_Product_Package_Order_Step) String() string {
	return "SoftLayer_Product_Package_Order_Step"
}

// SoftLayer_Product_Package_Order_Step_Extended is SoftLayer_Product_Package_Order_Step with all maskable types.
type SoftLayer_Product_Package_Order_Step_Extended struct {
	SoftLayer_Product_Package_Order_Step

	// PreviousSteps - no documentation
	PreviousSteps []*SoftLayer_Product_Package_Order_Step_Next `json:"previousSteps,omitempty"`

	// InclusivePreviousStepCount - A count of the next steps in the ordering process for the package tied
	// to this object, including this step.
	InclusivePreviousStepCount uint64 `json:"inclusivePreviousStepCount,omitempty"`

	// NextStepCount - A count of the next steps in the ordering process for the package tied to this
	// object.
	NextStepCount uint64 `json:"nextStepCount,omitempty"`

	// PreviousStepCount - no documentation
	PreviousStepCount uint64 `json:"previousStepCount,omitempty"`

	// InclusivePreviousSteps - The next steps in the ordering process for the package tied to this object,
	// including this step.
	InclusivePreviousSteps []*SoftLayer_Product_Package_Order_Step_Next `json:"inclusivePreviousSteps,omitempty"`

	// NextSteps - The next steps in the ordering process for the package tied to this object.
	NextSteps []*SoftLayer_Product_Package_Order_Step_Next `json:"nextSteps,omitempty"`
}

func (softlayer_product_package_order_step *SoftLayer_Product_Package_Order_Step_Extended) String() string {
	return "SoftLayer_Product_Package_Order_Step"
}
