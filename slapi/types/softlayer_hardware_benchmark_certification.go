package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Benchmark_Certification - The SoftLayer_Hardware_Benchmark_Certification data
// type contains general information relating to a single SoftLayer hardware benchmark certification
// document.
type SoftLayer_Hardware_Benchmark_Certification struct {

	// CreateDate - The date that a benchmark certification result was generated.
	CreateDate *time.Time `json:"createDate,omitempty"`

	// HardwareId - A benchmark certification results's associated hardware's internal identification
	// number.
	HardwareId int `json:"hardwareId,omitempty"`

	// AccountId - The internal identifier of the SoftLayer customer account associated with a benchmark
	// certification result.
	AccountId int `json:"accountId,omitempty"`

	// Account - Information regarding a benchmark certification result's associated SoftLayer customer
	// account.
	Account *SoftLayer_Account `json:"account,omitempty"`

	// Hardware - Information regarding the piece of hardware on which a benchmark certification test was
	// performed.
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`
}

func (softlayer_hardware_benchmark_certification *SoftLayer_Hardware_Benchmark_Certification) String() string {
	return "SoftLayer_Hardware_Benchmark_Certification"
}
