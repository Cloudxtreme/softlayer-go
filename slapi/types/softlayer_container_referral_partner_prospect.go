package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// SurveyId - <nil>
	SurveyId string `json:"surveyId,omitempty"`

	// Address1 - <nil>
	Address1 string `json:"address1,omitempty"`

	// City - <nil>
	City string `json:"city,omitempty"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone,omitempty"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode,omitempty"`

	// Questions - <nil>
	Questions []string `json:"questions,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// Email - <nil>
	Email string `json:"email,omitempty"`

	// Address2 - <nil>
	Address2 string `json:"address2,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses,omitempty"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
