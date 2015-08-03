package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner_Attachment - <nil>
type SoftLayer_Marketplace_Partner_Attachment struct {

	// AttachmentType - <nil>
	AttachmentType *SoftLayer_Marketplace_Partner_Attachment_Type `json:"attachmentType"`

	// AttachmentTypeId - <nil>
	AttachmentTypeId int `json:"attachmentTypeId"`

	// BaseName - <nil>
	BaseName string `json:"baseName"`

	// DisplayName - <nil>
	DisplayName string `json:"displayName"`

	// FileName - <nil>
	FileName string `json:"fileName"`

	// Id - <nil>
	Id int `json:"id"`

	// MarketplacePartnerId - <nil>
	MarketplacePartnerId int `json:"marketplacePartnerId"`

	// SaveAsName - <nil>
	SaveAsName string `json:"saveAsName"`
}

func (softlayer_marketplace_partner_attachment *SoftLayer_Marketplace_Partner_Attachment) String() string {
	return "SoftLayer_Marketplace_Partner_Attachment"
}
