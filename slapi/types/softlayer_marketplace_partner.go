package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite,omitempty"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription,omitempty"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits,omitempty"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong,omitempty"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle,omitempty"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles,omitempty"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage,omitempty"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId,omitempty"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort,omitempty"`

	// ProductName - <nil>
	ProductName string `json:"productName,omitempty"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription,omitempty"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial,omitempty"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures,omitempty"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) String() string {
	return "SoftLayer_Marketplace_Partner"
}

// SoftLayer_Marketplace_Partner_Extended is SoftLayer_Marketplace_Partner with all maskable types.
type SoftLayer_Marketplace_Partner_Extended struct {
	SoftLayer_Marketplace_Partner

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments,omitempty"`

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp,omitempty"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium,omitempty"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall,omitempty"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp,omitempty"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount,omitempty"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner_Extended) String() string {
	return "SoftLayer_Marketplace_Partner"
}
