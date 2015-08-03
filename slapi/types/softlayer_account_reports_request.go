package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Reports_Request - <nil>
type SoftLayer_Account_Reports_Request struct {

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Nda - <nil>
	Nda string `json:"nda"`

	// RequestKey - <nil>
	RequestKey string `json:"requestKey"`

	// Status - <nil>
	Status string `json:"status"`

	// AccountContactId - <nil>
	AccountContactId int `json:"accountContactId"`

	// ComplianceReportTypeId - <nil>
	ComplianceReportTypeId string `json:"complianceReportTypeId"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Report - <nil>
	Report string `json:"report"`

	// TicketId - <nil>
	TicketId int `json:"ticketId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`
}

// SoftLayer_Account_Reports_Request_Extended is SoftLayer_Account_Reports_Request with all maskable types.
type SoftLayer_Account_Reports_Request_Extended struct {
	SoftLayer_Account_Reports_Request

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// AccountContact - A request's corresponding external contact, if one exists.
	AccountContact *SoftLayer_Account_Contact `json:"accountContact"`

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ReportType - no documentation
	ReportType *SoftLayer_Compliance_Report_Type `json:"reportType"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) String() string {
	return "SoftLayer_Account_Reports_Request"
}
