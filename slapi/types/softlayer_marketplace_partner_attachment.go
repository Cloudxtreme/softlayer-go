package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Marketplace_Partner_Attachment - <nil>
type SoftLayer_Marketplace_Partner_Attachment struct {

	// AttachmentTypeId - <nil>
	AttachmentTypeId int `json:"attachmentTypeId,omitempty"`

	// BaseName - <nil>
	BaseName string `json:"baseName,omitempty"`

	// DisplayName - <nil>
	DisplayName string `json:"displayName,omitempty"`

	// FileName - <nil>
	FileName string `json:"fileName,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// MarketplacePartnerId - <nil>
	MarketplacePartnerId int `json:"marketplacePartnerId,omitempty"`

	// SaveAsName - <nil>
	SaveAsName string `json:"saveAsName,omitempty"`

	// AttachmentType - <nil>
	AttachmentType *SoftLayer_Marketplace_Partner_Attachment_Type `json:"attachmentType,omitempty"`
}

func (softlayer_marketplace_partner_attachment *SoftLayer_Marketplace_Partner_Attachment) String() string {
	return "SoftLayer_Marketplace_Partner_Attachment"
}
