package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category - The SoftLayer_Product_Item_Category data type contains general
// category information for prices.
type SoftLayer_Product_Item_Category struct {

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// Name - The friendly, descriptive name of the category as seen on the order forms and on invoices.
	Name string `json:"name,omitempty"`

	// QuantityLimit - Quantity that can be ordered. If 0, it will inherit the quantity from the server
	// quantity ordered. Otherwise it can be specified with the order separately
	QuantityLimit int `json:"quantityLimit,omitempty"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode,omitempty"`
}

func (softlayer_product_item_category *SoftLayer_Product_Item_Category) String() string {
	return "SoftLayer_Product_Item_Category"
}

// SoftLayer_Product_Item_Category_Extended is SoftLayer_Product_Item_Category with all maskable types.
type SoftLayer_Product_Item_Category_Extended struct {
	SoftLayer_Product_Item_Category

	// QuestionReferences - The question references that are associated with an item category.
	QuestionReferences []*SoftLayer_Product_Item_Category_Question_Xref `json:"questionReferences,omitempty"`

	// PackageConfigurations - A list of configuration available in this category.'
	PackageConfigurations []*SoftLayer_Product_Package_Order_Configuration `json:"packageConfigurations,omitempty"`

	// Questions - The questions that are associated with an item category.
	Questions []*SoftLayer_Product_Item_Category_Question `json:"questions,omitempty"`

	// PresetConfigurationCount - A count of a list of preset configurations this category is used in.'
	PresetConfigurationCount uint64 `json:"presetConfigurationCount,omitempty"`

	// Group - no documentation
	Group *SoftLayer_Product_Item_Category_Group `json:"group,omitempty"`

	// OrderOptionCount - A count of any unique options associated with an itme category.
	OrderOptionCount uint64 `json:"orderOptionCount,omitempty"`

	// QuestionCount - A count of the questions that are associated with an item category.
	QuestionCount uint64 `json:"questionCount,omitempty"`

	// QuestionReferenceCount - A count of the question references that are associated with an item
	// category.
	QuestionReferenceCount uint64 `json:"questionReferenceCount,omitempty"`

	// Groups - A collection of service offering category groups. Each group contains a collection of items
	// associated with this category.
	Groups []*SoftLayer_Product_Package_Item_Category_Group `json:"groups,omitempty"`

	// OrderOptions - Any unique options associated with an itme category.
	OrderOptions []*SoftLayer_Product_Item_Category_Order_Option_Type `json:"orderOptions,omitempty"`

	// PresetConfigurations - A list of preset configurations this category is used in.'
	PresetConfigurations []*SoftLayer_Product_Package_Preset_Configuration `json:"presetConfigurations,omitempty"`

	// BillingItemCount - A count of the billing items associated with an account that share a category
	// code with an item category's category code.
	BillingItemCount uint64 `json:"billingItemCount,omitempty"`

	// GroupCount - A count of a collection of service offering category groups. Each group contains a
	// collection of items associated with this category.
	GroupCount uint64 `json:"groupCount,omitempty"`

	// PackageConfigurationCount - A count of a list of configuration available in this category.'
	PackageConfigurationCount uint64 `json:"packageConfigurationCount,omitempty"`

	// BillingItems - The billing items associated with an account that share a category code with an item
	// category's category code.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems,omitempty"`
}

func (softlayer_product_item_category *SoftLayer_Product_Item_Category_Extended) String() string {
	return "SoftLayer_Product_Item_Category"
}
