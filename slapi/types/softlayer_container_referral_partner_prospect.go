package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// LastName - <nil>
	LastName string `json:"lastName,omitempty"`

	// State - <nil>
	State string `json:"state,omitempty"`

	// Country - <nil>
	Country string `json:"country,omitempty"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName,omitempty"`

	// Email - <nil>
	Email string `json:"email,omitempty"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode,omitempty"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses,omitempty"`

	// SurveyId - <nil>
	SurveyId string `json:"surveyId,omitempty"`

	// City - <nil>
	City string `json:"city,omitempty"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone,omitempty"`

	// Address2 - <nil>
	Address2 string `json:"address2,omitempty"`

	// FirstName - <nil>
	FirstName string `json:"firstName,omitempty"`

	// Questions - <nil>
	Questions []string `json:"questions,omitempty"`

	// Address1 - <nil>
	Address1 string `json:"address1,omitempty"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
