package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Reports_Request - <nil>
type SoftLayer_Account_Reports_Request struct {

	// TicketId - <nil>
	TicketId int `json:"ticketId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Nda - <nil>
	Nda string `json:"nda,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Report - <nil>
	Report string `json:"report,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// AccountContactId - <nil>
	AccountContactId int `json:"accountContactId,omitempty"`

	// ComplianceReportTypeId - <nil>
	ComplianceReportTypeId string `json:"complianceReportTypeId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// RequestKey - <nil>
	RequestKey string `json:"requestKey,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// AccountContact - A request's corresponding external contact, if one exists.
	AccountContact *SoftLayer_Account_Contact `json:"accountContact,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// ReportType - no documentation
	ReportType *SoftLayer_Compliance_Report_Type `json:"reportType,omitempty"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) String() string {
	return "SoftLayer_Account_Reports_Request"
}
