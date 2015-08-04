package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial,omitempty"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite,omitempty"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords,omitempty"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId,omitempty"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription,omitempty"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures,omitempty"`

	// ProductName - <nil>
	ProductName string `json:"productName,omitempty"`

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage,omitempty"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits,omitempty"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort,omitempty"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier,omitempty"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription,omitempty"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong,omitempty"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium,omitempty"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall,omitempty"`

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments,omitempty"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp,omitempty"`

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp,omitempty"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount,omitempty"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) String() string {
	return "SoftLayer_Marketplace_Partner"
}
