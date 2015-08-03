package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// Email - <nil>
	Email string `json:"email"`

	// State - <nil>
	State string `json:"state"`

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// Questions - <nil>
	Questions []string `json:"questions"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// City - <nil>
	City string `json:"city"`

	// Country - <nil>
	Country string `json:"country"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses"`

	// SurveyId - <nil>
	SurveyId string `json:"surveyId"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
