package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Dns_Domain_ResourceRecord_MxType - SoftLayer_Dns_Domain_ResourceRecord_MxType is a
// SoftLayer_Dns_Domain_ResourceRecord object whose ''type'' property is set to "mx" and used to
// describe MX resource records. MX records control which hosts are responsible as mail exchangers for
// a domain. For instance, in the domain example.org, an MX record whose host is "@" and data is "mail"
// says that the host "mail.example.org" is responsible for handling mail for example.org. That means
// mail sent to users @example.org are delivered to mail.example.org. Domains can have more than one MX
// record if it uses more than one server to send mail through. Multiple MX records are denoted by
// their priority, defined by the mxPriority property. MX records must be defined for hosts with
// accompanying A or resource records. They may not point mail towards a host defined by a record.
type SoftLayer_Dns_Domain_ResourceRecord_MxType struct {
}

// CreateObject - createObject creates a new MX record. The ''host'' property of the templateObject
// parameter is scrubbed to remove all non-alpha numeric characters except for and The ''data''
// property of the templateObject parameter is scrubbed to remove all non-alphanumeric characters for
// "." and Creating an MX record updates the serial number of the domain the resource record is
// associated with.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) CreateObject(ctx *slapi.RequestContext, templateObject SoftLayer_Dns_Domain_ResourceRecord_MxType) (*SoftLayer_Dns_Domain_ResourceRecord_MxType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_MxType
	return returnValue, nil
}

// CreateObjects - Create multiple MX records on a domain. This follows the same logic as
// ''createObject'. The serial number of the domain associated with this MX record is updated upon
// creation. ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it was
// unable to create a resource record.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) CreateObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Dns_Domain_ResourceRecord) ([]*SoftLayer_Dns_Domain_ResourceRecord, error) {
	var returnValue []*SoftLayer_Dns_Domain_ResourceRecord
	return returnValue, nil
}

// DeleteObject - Delete a domain's MX record. '''This cannot be undone.''' Be wary of running this
// method. If you remove a resource record in error you will need to re-create it by creating a new
// SoftLayer_Dns_Domain_ResourceRecord_MxType object. The serial number of the domain associated with
// this MX record is updated upon deletion. ''deleteObject'' returns Boolean ''true'' on successful
// deletion or ''false'' if it was unable to remove a resource record.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) DeleteObject(ctx *slapi.RequestContext) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteObjects - Remove multiple MX records from a domain. This follows the same logic as
// ''deleteObject'' and '''cannot be undone'''. The serial number of the domain associated with this MX
// record is updated upon deletion. ''deleteObjects'' returns Boolean ''true'' on successful deletion
// or ''false'' if it was unable to remove a resource record.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) DeleteObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Dns_Domain_ResourceRecord_MxType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObject - editObject edits an existing MX resource record. The ''host'' property of the
// templateObject parameter is scrubbed to remove all non-alpha numeric characters except for and The
// ''data'' property of the templateObject parameter is scrubbed to remove all non-alphanumeric
// characters for "." and Editing an MX record updates the serial number of the domain the record is
// associated with. ''editObject'' returns Boolean ''true'' on a successful edit or ''false'' if it was
// unable to edit the resource record.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) EditObject(ctx *slapi.RequestContext, templateObject SoftLayer_Dns_Domain_ResourceRecord_MxType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EditObjects - Edit multiple MX records on a domain. This follows the same logic as ''createObject'.
// The serial number of the domain associated with this MX record is updated upon creation.
// ''createObjects'' returns Boolean ''true'' on successful creation or ''false'' if it was unable to
// create a resource record.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) EditObjects(ctx *slapi.RequestContext, templateObjects []SoftLayer_Dns_Domain_ResourceRecord_MxType) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Dns_Domain_ResourceRecord_MxType object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Dns_Domain_ResourceRecord_MxType service. You can only retrieve resource records belonging
// to domains that are assigned to your SoftLayer account.
func (softlayer_dns_domain_resourcerecord_mxtype *SoftLayer_Dns_Domain_ResourceRecord_MxType) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Dns_Domain_ResourceRecord_MxType, error) {
	var returnValue *SoftLayer_Dns_Domain_ResourceRecord_MxType
	return returnValue, nil
}
