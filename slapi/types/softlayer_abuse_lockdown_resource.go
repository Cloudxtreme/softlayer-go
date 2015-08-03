package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Abuse_Lockdown_Resource - <nil>
type SoftLayer_Abuse_Lockdown_Resource struct {
}

func (softlayer_abuse_lockdown_resource *SoftLayer_Abuse_Lockdown_Resource) String() string {
	return "SoftLayer_Abuse_Lockdown_Resource"
}

// SoftLayer_Abuse_Lockdown_Resource_Extended is SoftLayer_Abuse_Lockdown_Resource with all maskable types.
type SoftLayer_Abuse_Lockdown_Resource_Extended struct {
	SoftLayer_Abuse_Lockdown_Resource

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// InvoiceItem - <nil>
	InvoiceItem *SoftLayer_Billing_Invoice_Item `json:"invoiceItem"`
}

func (softlayer_abuse_lockdown_resource *SoftLayer_Abuse_Lockdown_Resource_Extended) String() string {
	return "SoftLayer_Abuse_Lockdown_Resource"
}
