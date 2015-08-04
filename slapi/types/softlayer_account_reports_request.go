package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Reports_Request - <nil>
type SoftLayer_Account_Reports_Request struct {

	// AccountContactId - <nil>
	AccountContactId int `json:"accountContactId,omitempty"`

	// TicketId - <nil>
	TicketId int `json:"ticketId,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// ComplianceReportTypeId - <nil>
	ComplianceReportTypeId string `json:"complianceReportTypeId,omitempty"`

	// Notes - <nil>
	Notes string `json:"notes,omitempty"`

	// Report - <nil>
	Report string `json:"report,omitempty"`

	// Status - <nil>
	Status string `json:"status,omitempty"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId,omitempty"`

	// AccountId - <nil>
	AccountId int `json:"accountId,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - <nil>
	Id int `json:"id,omitempty"`

	// Nda - <nil>
	Nda string `json:"nda,omitempty"`

	// RequestKey - <nil>
	RequestKey string `json:"requestKey,omitempty"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) String() string {
	return "SoftLayer_Account_Reports_Request"
}

// SoftLayer_Account_Reports_Request_Extended is SoftLayer_Account_Reports_Request with all maskable types.
type SoftLayer_Account_Reports_Request_Extended struct {
	SoftLayer_Account_Reports_Request

	// Account - <nil>
	Account *SoftLayer_Account `json:"account,omitempty"`

	// ReportType - no documentation
	ReportType *SoftLayer_Compliance_Report_Type `json:"reportType,omitempty"`

	// AccountContact - A request's corresponding external contact, if one exists.
	AccountContact *SoftLayer_Account_Contact `json:"accountContact,omitempty"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket,omitempty"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user,omitempty"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request_Extended) String() string {
	return "SoftLayer_Account_Reports_Request"
}
