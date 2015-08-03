package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures"`

	// Id - <nil>
	Id int `json:"id"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId"`

	// ProductName - <nil>
	ProductName string `json:"productName"`
}

// SoftLayer_Marketplace_Partner_Extended is SoftLayer_Marketplace_Partner with all maskable types.
type SoftLayer_Marketplace_Partner_Extended struct {
	SoftLayer_Marketplace_Partner

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall"`

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) String() string {
	return "SoftLayer_Marketplace_Partner"
}
