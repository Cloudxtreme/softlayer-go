package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
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

// ConvertToPrimary - A secondary DNS record may be converted to a primary DNS record. By converting a
// secondary DNS record, the SoftLayer name servers will be the authoritative nameserver for this
// domain and will be directly editable in the SoftLayer API and Portal. Primary DNS record conversion
// performs the following steps: * The SOA record is updated with SoftLayer's primary name server. *
// All NS records are removed and replaced with SoftLayer's NS records. * The secondary DNS record is
// removed. After the DNS records are converted, the following restrictions will apply to the new
// domain record: * You will need to manage the zone record using the [[SoftLayer_Dns_Domain]] service.
// * You may not edit the SOA or NS records. * You may only edit the following resource records: A, MX,
// TX, This change can not be undone, and the record can not be converted back into a secondary DNS
// record once the conversion is complete.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) ConvertToPrimary(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateObject - Create a secondary DNS record. The ''zoneName'', ''masterIpAddress'', and
// ''transferFrequency'' properties in the templateObject parameter are required parameters to create a
// secondary DNS record.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Secondary) (*SoftLayer_Dns_Secondary, error) {
	var returnValue *SoftLayer_Dns_Secondary
	return returnValue, nil
}

// CreateObjects - Create multiple secondary DNS records. Each record passed to ''createObjects''
// follows the logic in the SoftLayer_Dns_Secondary
// [[SoftLayer_Dns_Secondary::createObject|createObject]] method.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) CreateObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Secondary) ([]*SoftLayer_Dns_Secondary, error) {
	var returnValue []*SoftLayer_Dns_Secondary
	return returnValue, nil
}

// DeleteObject - Delete a secondary DNS Record. This will also remove any associated domain records
// and resource records on the SoftLayer nameservers that were created as a result of the zone
// transfers. This action cannot be undone.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - Edit the properties of a secondary DNS record by passing in a modified instance of a
// SoftLayer_Dns_Secondary object. You may only edit the ''masterIpAddress'' and ''transferFrequency''
// properties of your secondary DNS record. ''ZoneName'' may not be altered after a secondary DNS
// record has been created. Please remove and re-create the record if you need to make changes to your
// zone name.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Secondary) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetByDomainName - Search for [[SoftLayer_Dns_Domain_Secondary]] records by domain name.
// getByDomainName() performs an inclusive search for secondary domain records, returning multiple
// records based on partial name matches. Use this method to locate secondary domain records if you
// don't have access to their id numbers.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) GetByDomainName(commonOptions *slapi.CommonOptions, name string) ([]*SoftLayer_Dns_Secondary, error) {
	var returnValue []*SoftLayer_Dns_Secondary
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Secondary object whose ID number corresponds to
// the ID number of the init paramater passed to the SoftLayer_Dns_Secondary service. You can only
// retrieve a secondary DNS record that is assigned to your SoftLayer customer account.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Dns_Secondary, error) {
	var returnValue *SoftLayer_Dns_Secondary
	return returnValue, nil
}

// TransferNow - Force a secondary DNS zone transfer by setting it's status "Transfer Now". A zone
// transfer will be initiated within a minute of receiving this API call.
func (softlayer_dns_secondary *SoftLayer_Dns_Secondary) TransferNow(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
