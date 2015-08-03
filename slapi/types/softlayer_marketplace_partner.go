package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId"`

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles"`

	// Id - <nil>
	Id int `json:"id"`

	// ProductName - <nil>
	ProductName string `json:"productName"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) String() string {
	return "SoftLayer_Marketplace_Partner"
}

// SoftLayer_Marketplace_Partner_Extended is SoftLayer_Marketplace_Partner with all maskable types.
type SoftLayer_Marketplace_Partner_Extended struct {
	SoftLayer_Marketplace_Partner

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp"`

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner_Extended) String() string {
	return "SoftLayer_Marketplace_Partner"
}
