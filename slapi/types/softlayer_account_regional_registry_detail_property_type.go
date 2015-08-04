package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Account_Regional_Registry_Detail_Property_Type - Subnet Registration Detail Property Type
// objects describe the nature of a [[SoftLayer_Account_Regional_Registry_Detail_Property]] object.
// These types use [http://php.net/pcre.pattern.php Perl-Compatible Regular Expressions] to validate
// the value of a property object.
type SoftLayer_Account_Regional_Registry_Detail_Property_Type struct {

	// ValueExpression - A Perl-compatible regular expression used to describe the valid format of the
	// property
	ValueExpression string `json:"valueExpression,omitempty"`

	// CreateDate - <nil>
	CreateDate *time.Time `json:"createDate,omitempty"`

	// Id - no documentation
	Id int `json:"id,omitempty"`

	// KeyName - no documentation
	KeyName string `json:"keyName,omitempty"`

	// ModifyDate - <nil>
	ModifyDate *time.Time `json:"modifyDate,omitempty"`

	// Name - no documentation
	Name string `json:"name,omitempty"`
}

func (softlayer_account_regional_registry_detail_property_type *SoftLayer_Account_Regional_Registry_Detail_Property_Type) String() string {
	return "SoftLayer_Account_Regional_Registry_Detail_Property_Type"
}
