package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Account_AbuseEmail - An unfortunate facet of the hosting business is the necessity of with
// legal and network abuse inquiries. As these types of inquiries frequently contain sensitive
// information SoftLayer keeps a separate account contact email address for direct contact about legal
// and abuse matters, modeled by the SoftLayer_Account_AbuseEmail data type. SoftLayer will typically
// email an account's abuse email address in these types of cases, and an email is automatically sent
// to an account's abuse email address when a legal or abuse ticket is created or updated.
type SoftLayer_Account_AbuseEmail struct {

	// Email - no documentation
	Email string `json:"email,omitempty"`
}

func (softlayer_account_abuseemail *SoftLayer_Account_AbuseEmail) String() string {
	return "SoftLayer_Account_AbuseEmail"
}

// SoftLayer_Account_AbuseEmail_Extended is SoftLayer_Account_AbuseEmail with all maskable types.
type SoftLayer_Account_AbuseEmail_Extended struct {
	SoftLayer_Account_AbuseEmail

	// Account - The account associated with an abuse email address.
	Account *SoftLayer_Account `json:"account,omitempty"`
}

func (softlayer_account_abuseemail *SoftLayer_Account_AbuseEmail_Extended) String() string {
	return "SoftLayer_Account_AbuseEmail"
}
