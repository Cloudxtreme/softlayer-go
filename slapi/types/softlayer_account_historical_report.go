package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Account_Historical_Report - <nil>
type SoftLayer_Account_Historical_Report struct {
}

// GetAccountHostUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountHostUptimeGraphData(commonOptions *slapi.CommonOptions, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetAccountHostUptimeSummary - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountHostUptimeSummary(commonOptions *slapi.CommonOptions, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary
	return returnValue, nil
}

// GetAccountUrlUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountUrlUptimeGraphData(commonOptions *slapi.CommonOptions, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetAccountUrlUptimeSummary - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountUrlUptimeSummary(commonOptions *slapi.CommonOptions, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary
	return returnValue, nil
}

// GetHostUptimeDetail - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetHostUptimeDetail(commonOptions *slapi.CommonOptions, configurationValueId int, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary_Detail, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary_Detail
	return returnValue, nil
}

// GetHostUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetHostUptimeGraphData(commonOptions *slapi.CommonOptions, configurationValueId int, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetUrlUptimeDetail - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetUrlUptimeDetail(commonOptions *slapi.CommonOptions, configurationValueId int, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary_Detail, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary_Detail
	return returnValue, nil
}

// GetUrlUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetUrlUptimeGraphData(commonOptions *slapi.CommonOptions, configurationValueId int, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}
