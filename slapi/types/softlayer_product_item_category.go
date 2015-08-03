package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Product_Item_Category - The SoftLayer_Product_Item_Category data type contains general
// category information for prices.
type SoftLayer_Product_Item_Category struct {

	// Name - The friendly, descriptive name of the category as seen on the order forms and on invoices.
	Name string `json:"name"`

	// Id - no documentation
	Id int `json:"id"`

	// QuantityLimit - Quantity that can be ordered. If 0, it will inherit the quantity from the server
	// quantity ordered. Otherwise it can be specified with the order separately
	QuantityLimit int `json:"quantityLimit"`

	// CategoryCode - no documentation
	CategoryCode string `json:"categoryCode"`
}

// SoftLayer_Product_Item_Category_Extended is SoftLayer_Product_Item_Category with all maskable types.
type SoftLayer_Product_Item_Category_Extended struct {
	SoftLayer_Product_Item_Category

	// OrderOptionCount - A count of any unique options associated with an itme category.
	OrderOptionCount uint64 `json:"orderOptionCount"`

	// PackageConfigurationCount - A count of a list of configuration available in this category.'
	PackageConfigurationCount uint64 `json:"packageConfigurationCount"`

	// Group - no documentation
	Group *SoftLayer_Product_Item_Category_Group `json:"group"`

	// OrderOptions - Any unique options associated with an itme category.
	OrderOptions []*SoftLayer_Product_Item_Category_Order_Option_Type `json:"orderOptions"`

	// PackageConfigurations - A list of configuration available in this category.'
	PackageConfigurations []*SoftLayer_Product_Package_Order_Configuration `json:"packageConfigurations"`

	// QuestionReferences - The question references that are associated with an item category.
	QuestionReferences []*SoftLayer_Product_Item_Category_Question_Xref `json:"questionReferences"`

	// Questions - The questions that are associated with an item category.
	Questions []*SoftLayer_Product_Item_Category_Question `json:"questions"`

	// PresetConfigurationCount - A count of a list of preset configurations this category is used in.'
	PresetConfigurationCount uint64 `json:"presetConfigurationCount"`

	// QuestionCount - A count of the questions that are associated with an item category.
	QuestionCount uint64 `json:"questionCount"`

	// Groups - A collection of service offering category groups. Each group contains a collection of items
	// associated with this category.
	Groups []*SoftLayer_Product_Package_Item_Category_Group `json:"groups"`

	// GroupCount - A count of a collection of service offering category groups. Each group contains a
	// collection of items associated with this category.
	GroupCount uint64 `json:"groupCount"`

	// BillingItems - The billing items associated with an account that share a category code with an item
	// category's category code.
	BillingItems []*SoftLayer_Billing_Item `json:"billingItems"`

	// PresetConfigurations - A list of preset configurations this category is used in.'
	PresetConfigurations []*SoftLayer_Product_Package_Preset_Configuration `json:"presetConfigurations"`

	// QuestionReferenceCount - A count of the question references that are associated with an item
	// category.
	QuestionReferenceCount uint64 `json:"questionReferenceCount"`

	// BillingItemCount - A count of the billing items associated with an account that share a category
	// code with an item category's category code.
	BillingItemCount uint64 `json:"billingItemCount"`
}

func (softlayer_product_item_category *SoftLayer_Product_Item_Category) String() string {
	return "SoftLayer_Product_Item_Category"
}
