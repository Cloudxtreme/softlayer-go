package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// Email - <nil>
	Email string `json:"email"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// State - <nil>
	State string `json:"state"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// Questions - <nil>
	Questions []string `json:"questions"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// SurveyId - <nil>
	SurveyId string `json:"surveyId"`

	// City - <nil>
	City string `json:"city"`

	// Country - <nil>
	Country string `json:"country"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
