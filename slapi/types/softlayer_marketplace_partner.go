package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures,omitempty"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier,omitempty"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription,omitempty"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite,omitempty"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords,omitempty"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits,omitempty"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong,omitempty"`

	// ProductName - <nil>
	ProductName string `json:"productName,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription,omitempty"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId,omitempty"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort,omitempty"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles,omitempty"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial,omitempty"`

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage,omitempty"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription,omitempty"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle,omitempty"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount,omitempty"`

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp,omitempty"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium,omitempty"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall,omitempty"`

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments,omitempty"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp,omitempty"`
}

func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) String() string {
	return "SoftLayer_Marketplace_Partner"
}
