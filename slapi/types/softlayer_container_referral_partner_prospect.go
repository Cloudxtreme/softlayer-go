package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// City - <nil>
	City string `json:"city,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode,omitempty"`

	// Email - <nil>
	Email string `json:"email,omitempty"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses,omitempty"`

	// Address1 - <nil>
	Address1 string `json:"address1,omitempty"`

	// Address2 - <nil>
	Address2 string `json:"address2,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Questions - <nil>
	Questions []string `json:"questions,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// SurveyId - <nil>
	SurveyId string `json:"surveyId,omitempty"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
