package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Reports_Request - <nil>
type SoftLayer_Account_Reports_Request struct {

	// Account - <nil>
	Account *SoftLayer_Account `json:"account"`

	// AccountContact - A request's corresponding external contact, if one exists.
	AccountContact *SoftLayer_Account_Contact `json:"accountContact"`

	// AccountContactId - <nil>
	AccountContactId int `json:"accountContactId"`

	// AccountId - <nil>
	AccountId int `json:"accountId"`

	// ComplianceReportTypeId - <nil>
	ComplianceReportTypeId string `json:"complianceReportTypeId"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate"`

	// Id - <nil>
	Id int `json:"id"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate"`

	// Nda - <nil>
	Nda string `json:"nda"`

	// Notes - <nil>
	Notes string `json:"notes"`

	// Report - <nil>
	Report string `json:"report"`

	// ReportType - no documentation
	ReportType *SoftLayer_Compliance_Report_Type `json:"reportType"`

	// RequestKey - <nil>
	RequestKey string `json:"requestKey"`

	// Status - <nil>
	Status string `json:"status"`

	// Ticket - <nil>
	Ticket *SoftLayer_Ticket `json:"ticket"`

	// TicketId - <nil>
	TicketId int `json:"ticketId"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UsrRecordId - <nil>
	UsrRecordId int `json:"usrRecordId"`
}

// CreateRequest - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) CreateRequest(commonOptions *slapi.CommonOptions, contact SoftLayer_Account_Contact, reason string, reportType string) (*SoftLayer_Account_Reports_Request, error) {
	var returnValue *SoftLayer_Account_Reports_Request
	return returnValue, nil
}

// GetAllObjects - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) GetAllObjects(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_Reports_Request, error) {
	var returnValue *SoftLayer_Account_Reports_Request
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Account_Reports_Request, error) {
	var returnValue *SoftLayer_Account_Reports_Request
	return returnValue, nil
}

// GetRequestByRequestKey - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) GetRequestByRequestKey(commonOptions *slapi.CommonOptions, requestKey string) (*SoftLayer_Account_Reports_Request, error) {
	var returnValue *SoftLayer_Account_Reports_Request
	return returnValue, nil
}

// SendReportEmail - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) SendReportEmail(commonOptions *slapi.CommonOptions, request SoftLayer_Account_Reports_Request) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateTicketOnDecline - <nil>
func (softlayer_account_reports_request *SoftLayer_Account_Reports_Request) UpdateTicketOnDecline(commonOptions *slapi.CommonOptions, request SoftLayer_Account_Reports_Request) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
