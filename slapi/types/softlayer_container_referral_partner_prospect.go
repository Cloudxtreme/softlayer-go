package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Referral_Partner_Prospect - <nil>
type SoftLayer_Container_Referral_Partner_Prospect struct {

	// Address1 - <nil>
	Address1 string `json:"address1"`

	// Address2 - <nil>
	Address2 string `json:"address2"`

	// City - <nil>
	City string `json:"city"`

	// CompanyName - <nil>
	CompanyName string `json:"companyName"`

	// Country - <nil>
	Country string `json:"country"`

	// Email - <nil>
	Email string `json:"email"`

	// FirstName - <nil>
	FirstName string `json:"firstName"`

	// LastName - <nil>
	LastName string `json:"lastName"`

	// OfficePhone - <nil>
	OfficePhone string `json:"officePhone"`

	// PostalCode - <nil>
	PostalCode string `json:"postalCode"`

	// Questions - <nil>
	Questions []string `json:"questions"`

	// Responses - <nil>
	Responses []*SoftLayer_Survey_Response `json:"responses"`

	// State - <nil>
	State string `json:"state"`

	// SurveyId - <nil>
	SurveyId string `json:"surveyId"`
}

func (softlayer_container_referral_partner_prospect *SoftLayer_Container_Referral_Partner_Prospect) String() string {
	return "SoftLayer_Container_Referral_Partner_Prospect"
}
