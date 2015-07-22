package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Marketplace_Partner - <nil>
type SoftLayer_Marketplace_Partner struct {

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// AttachedFiles - <nil>
	AttachedFiles []*SoftLayer_Marketplace_Partner_Attachment `json:"attachedFiles"`

	// AttachmentCount - no documentation
	AttachmentCount uint64 `json:"attachmentCount"`

	// Attachments - <nil>
	Attachments []*SoftLayer_Marketplace_Partner_Attachment `json:"attachments"`

	// CompanyDescription - <nil>
	CompanyDescription string `json:"companyDescription"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// HeadlineDescription - <nil>
	HeadlineDescription string `json:"headlineDescription"`

	// Id - <nil>
	Id int `json:"id"`

	// LinkFreeTrial - <nil>
	LinkFreeTrial string `json:"linkFreeTrial"`

	// LinkOrderPage - <nil>
	LinkOrderPage string `json:"linkOrderPage"`

	// LinkWebsite - <nil>
	LinkWebsite string `json:"linkWebsite"`

	// LogoMedium - <nil>
	LogoMedium *SoftLayer_Marketplace_Partner_Attachment `json:"logoMedium"`

	// LogoMediumTemp - <nil>
	LogoMediumTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoMediumTemp"`

	// LogoSmall - <nil>
	LogoSmall *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmall"`

	// LogoSmallTemp - <nil>
	LogoSmallTemp *SoftLayer_Marketplace_Partner_Attachment `json:"logoSmallTemp"`

	// MetaDescription - <nil>
	MetaDescription string `json:"metaDescription"`

	// MetaKeywords - <nil>
	MetaKeywords string `json:"metaKeywords"`

	// ProductBenefits - <nil>
	ProductBenefits string `json:"productBenefits"`

	// ProductCategoryId - <nil>
	ProductCategoryId int `json:"productCategoryId"`

	// ProductDescriptionLong - <nil>
	ProductDescriptionLong string `json:"productDescriptionLong"`

	// ProductDescriptionShort - <nil>
	ProductDescriptionShort string `json:"productDescriptionShort"`

	// ProductFeatures - <nil>
	ProductFeatures string `json:"productFeatures"`

	// ProductName - <nil>
	ProductName string `json:"productName"`

	// ProductTitle - <nil>
	ProductTitle string `json:"productTitle"`

	// UrlIdentifier - <nil>
	UrlIdentifier string `json:"urlIdentifier"`
}

// GetAllObjects - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetAllObjects(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Marketplace_Partner, error) {
	var returnValue []*SoftLayer_Marketplace_Partner
	return returnValue, nil
}

// GetAllPublishedPartners - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetAllPublishedPartners(commonOptions *slapi.CommonOptions, searchTerm string) ([]*SoftLayer_Marketplace_Partner, error) {
	var returnValue []*SoftLayer_Marketplace_Partner
	return returnValue, nil
}

// GetFeaturedPartners - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetFeaturedPartners(commonOptions *slapi.CommonOptions, non bool) ([]*SoftLayer_Marketplace_Partner, error) {
	var returnValue []*SoftLayer_Marketplace_Partner
	return returnValue, nil
}

// GetFile - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetFile(commonOptions *slapi.CommonOptions, name string) (*SoftLayer_Marketplace_Partner_File, error) {
	var returnValue *SoftLayer_Marketplace_Partner_File
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Marketplace_Partner, error) {
	var returnValue *SoftLayer_Marketplace_Partner
	return returnValue, nil
}

// GetPartnerByUrlIdentifier - <nil>
func (softlayer_marketplace_partner *SoftLayer_Marketplace_Partner) GetPartnerByUrlIdentifier(commonOptions *slapi.CommonOptions, urlIdentifier string) (*SoftLayer_Marketplace_Partner, error) {
	var returnValue *SoftLayer_Marketplace_Partner
	return returnValue, nil
}
