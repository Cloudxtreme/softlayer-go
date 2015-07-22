package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"

	slapi "go-softlayer/slapi"
)

// SoftLayer_Hardware_Benchmark_Certification - The SoftLayer_Hardware_Benchmark_Certification data
// type contains general information relating to a single SoftLayer hardware benchmark certification
// document.
type SoftLayer_Hardware_Benchmark_Certification struct {

	// Account - Information regarding a benchmark certification result's associated SoftLayer customer
	// account.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the SoftLayer customer account associated with a benchmark
	// certification result.
	AccountId int `json:"accountId"`

	// CreateDate - The date that a benchmark certification result was generated.
	CreateDate *time.Time `json:"createDate"`

	// Hardware - Information regarding the piece of hardware on which a benchmark certification test was
	// performed.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// HardwareId - A benchmark certification results's associated hardware's internal identification
	// number.
	HardwareId int `json:"hardwareId"`
}

// GetObject - getObject retrieves the SoftLayer_Hardware_Benchmark_Certification object whose ID
// number corresponds to the ID number of the init parameter passed to the
// SoftLayer_Hardware_Benchmark_Certification service.
func (softlayer_hardware_benchmark_certification *SoftLayer_Hardware_Benchmark_Certification) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Hardware_Benchmark_Certification, error) {
	var returnValue *SoftLayer_Hardware_Benchmark_Certification
	return returnValue, nil
}

// GetResultFile - Attempt to retrieve the file associated with a benchmark certification result, if
// such a file exists. If there is no file for this benchmark certification result, calling this method
// throws an exception. The "getResultFile" method attempts to retrieve the file associated with a
// benchmark certification result, if such a file exists. If no file exists for the benchmark
// certification, an exception is thrown.
func (softlayer_hardware_benchmark_certification *SoftLayer_Hardware_Benchmark_Certification) GetResultFile(commonOptions *slapi.CommonOptions) (string, error) {
	var returnValue string
	return returnValue, nil
}
