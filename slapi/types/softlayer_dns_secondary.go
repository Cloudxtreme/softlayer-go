package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Dns_Secondary - The SoftLayer_Dns_Secondary data type contains information on a single
// secondary DNS zone which is managed through SoftLayer's zone transfer service. Domains created via
// zone transfer may not be modified by the SoftLayer portal or
type SoftLayer_Dns_Secondary struct {

	// Account - The SoftLayer account that owns a secondary DNS record.
	Account *SoftLayer_Account `json:"account"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Domain - The domain record created by zone transfer from a secondary DNS record.
	Domain *SoftLayer_Dns_Domain `json:"domain"`

	// ErrorMessageCount - A count of the error messages created during secondary DNS record transfer.
	ErrorMessageCount uint64 `json:"errorMessageCount"`

	// ErrorMessages - The error messages created during secondary DNS record transfer.
	ErrorMessages []*SoftLayer_Dns_Message `json:"errorMessages"`

	// Id - The internal identifier for a secondary DNS record.
	Id int `json:"id"`

	// LastUpdate - The date when the most recent secondary DNS zone transfer took place.
	LastUpdate *time.Time `json:"lastUpdate"`

	// MasterIpAddress - The IP address of the master name server where a secondary DNS zone is transferred
	// from.
	MasterIpAddress string `json:"masterIpAddress"`

	// Status - no documentation
	Status *SoftLayer_Dns_Status `json:"status"`

	// StatusId - The current status of a secondary DNS record. The status may be one of the following:
	// :*'''0''': Disabled :*'''1''': Active :*'''2''': Transfer Now :*'''3''': An error occurred that
	// prevented the zone transfer from being completed.
	StatusId int `json:"statusId"`

	// StatusText - The textual representation of a secondary DNS zone's status.
	StatusText string `json:"statusText"`

	// TransferFrequency - How often a secondary DNS zone should be transferred in minutes.
	TransferFrequency int `json:"transferFrequency"`

	// ZoneName - no documentation
	ZoneName string `json:"zoneName"`
}

func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) String() string {
	return "SoftLayer_Dns_Secondary"
}
