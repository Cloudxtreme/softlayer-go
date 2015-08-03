package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Reports_Request - <nil>
type SoftLayer_Account_Reports_Request struct {

	// Nda - <nil>
	Nda string `json:"nda"`

	// RequestKey - <nil>
	RequestKey string `json:"requestKey"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`

	// ComplianceReportTypeId - <nil>
	ComplianceReportTypeId string `json:"complianceReportTypeId"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Status - <nil>
	Status string `json:"status"`

	// TicketId - <nil>
	TicketId int `json:"ticketId"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// Id - <nil>
	Id int `json:"id"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Report - <nil>
	Report string `json:"report"`

	// AccountContactId - <nil>
	AccountContactId int `json:"accountContactId"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) String() string {
	return "SoftLayer_Account_Reports_Request"
}

// SoftLayer_Account_Reports_Request_Extended is SoftLayer_Account_Reports_Request with all maskable types.
type SoftLayer_Account_Reports_Request_Extended struct {
	SoftLayer_Account_Reports_Request

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// ReportType - no documentation
	ReportType *SoftLayer_Compliance_Report_Type `json:"reportType"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// AccountContact - A request's corresponding external contact, if one exists.
	AccountContact *SoftLayer_Account_Contact `json:"accountContact"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request_Extended) String() string {
	return "SoftLayer_Account_Reports_Request"
}
