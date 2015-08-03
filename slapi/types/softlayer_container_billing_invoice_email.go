package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Billing_Invoice_Email - This container is used to provide all the options for
// [[SoftLayer_Billing_Invoice/emailInvoices|emailInvoices]] in order to have the nessary invoices sent
// to the user's email.
type SoftLayer_Container_Billing_Invoice_Email struct {

	// ExcelInvoiceIds - no documentation
	ExcelInvoiceIds []int `json:"excelInvoiceIds"`

	// PdfDetailedInvoiceIds - no documentation
	PdfDetailedInvoiceIds []int `json:"pdfDetailedInvoiceIds"`

	// PdfInvoiceIds - no documentation
	PdfInvoiceIds []int `json:"pdfInvoiceIds"`

	// Type - The type of Invoices to be emailed [current|next]. If next is selected, the account id will
	// be used.
	Type string `json:"type"`
}

func (softlayer_container_billing_invoice_email *SoftLayer_Container_Billing_Invoice_Email) String() string {
	return "SoftLayer_Container_Billing_Invoice_Email"
}
