package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Dns_Domain_ResourceRecord_SrvType - SoftLayer_Dns_Domain_ResourceRecord_SrvType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "srv" and defines a DNS
// SRV record on a SoftLayer hosted domain.
type SoftLayer_Dns_Domain_ResourceRecord_SrvType struct {

	// Port - The TCP or UDP port on which the service is to be found.
	Port int `json:"port"`

	// Priority - The priority of the target host, lower value means more preferred.
	Priority int `json:"priority"`

	// Protocol - The protocol of the desired service; this is usually either TCP or
	Protocol string `json:"protocol"`

	// Service - no documentation
	Service string `json:"service"`

	// Weight - A relative weight for records with the same priority.
	Weight int `json:"weight"`
}

// CreateObject - createObject creates a new SRV record. The ''host'' property of the templateObject
// parameter is scrubbed to remove all non-alpha numeric characters except for and The ''data''
// property of the templateObject parameter is scrubbed to remove all non-alphanumeric characters for
// "." and Creating an SRV record updates the serial number of the domain the resource record is
// associated with.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) CreateObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Domain_ResourceRecord_SrvType) (*SoftLayer_Dns_Domain_ResourceRecord_SrvType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_SrvType
	return returnValue, nil
}

// CreateObjects - Create multiple SRV records on a domain. This follows the same logic as
// ''createObject'. The serial number of the domain associated with this SRV record is updated upon
// creation. ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it was
// unable to create a resource record.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) CreateObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord) ([]*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue []*SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

// DeleteObject - Delete a domain's SRV record. '''This cannot be undone.''' Be wary of running this
// method. If you remove a resource record in error you will need to re-create it by creating a new
// SoftLayer_Dns_Domain_ResourceRecord_SrvType object. The serial number of the domain associated with
// this SRV record is updated upon deletion. ''deleteObject'' returns Boolean ''true'' on successful
// deletion or ''false'' if it was unable to remove a resource record.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) DeleteObject(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObjects - Remove multiple SRV records from a domain. This follows the same logic as
// ''deleteObject'' and '''cannot be undone'''. The serial number of the domain associated with this
// SRV record is updated upon deletion. ''deleteObjects'' returns Boolean ''true'' on successful
// deletion or ''false'' if it was unable to remove a resource record.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) DeleteObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord_SrvType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - editObject edits an existing SRV resource record. The ''host'' property of the
// templateObject parameter is scrubbed to remove all non-alpha numeric characters except for and The
// ''data'' property of the templateObject parameter is scrubbed to remove all non-alphanumeric
// characters for "." and Editing an SRV record updates the serial number of the domain the record is
// associated with. ''editObject'' returns Boolean ''true'' on a successful edit or ''false'' if it was
// unable to edit the resource record.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) EditObject(commonOptions *slapi.CommonOptions, templateObject SoftLayer_Dns_Domain_ResourceRecord_SrvType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Edit multiple SRV records on a domain. This follows the same logic as ''createObject'.
// The serial number of the domain associated with this SRV record is updated upon creation.
// ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it was unable to
// create a resource record.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) EditObjects(commonOptions *slapi.CommonOptions, templateObjects []SoftLayer_Dns_Domain_ResourceRecord_SrvType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Domain_ResourceRecord_SrvType object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Dns_Domain_ResourceRecord_SrvType service. You can only retrieve resource records
// belonging to domains that are assigned to your SoftLayer account.
func (softlayer_dns_domain_resourcerecord_srvtype *SoftLayer_Dns_Domain_ResourceRecord_SrvType) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Dns_Domain_ResourceRecord_SrvType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_SrvType
	return returnValue, nil
}
