package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Compliance_Report_Type - <nil>
type SoftLayer_Compliance_Report_Type struct {

	// Id - <nil>
	Id int `json:"id"`

	// KeyName - <nil>
	KeyName string `json:"keyName"`

	// Name - <nil>
	Name string `json:"name"`
}

// GetAllObjects - <nil>
func (softlayer_compliance_report_type *SoftLayer_Compliance_Report_Type) GetAllObjects(ctx *slapi.RequestContext) ([]*SoftLayer_Compliance_Report_Type, error) {
	var returnValue []*SoftLayer_Compliance_Report_Type
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_compliance_report_type *SoftLayer_Compliance_Report_Type) GetObject(ctx *slapi.RequestContext) (*SoftLayer_Compliance_Report_Type, error) {
	var returnValue *SoftLayer_Compliance_Report_Type
	return returnValue, nil
}
