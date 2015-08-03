package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "github.com/sudorandom/softlayer-go/slapi"
)

// SoftLayer_Account_Historical_Report - <nil>
type SoftLayer_Account_Historical_Report struct {
}

func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) String() string {
	return "SoftLayer_Account_Historical_Report"
}

// GetAccountHostUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountHostUptimeGraphData(ctx *slapi.RequestContext, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetAccountHostUptimeSummary - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountHostUptimeSummary(ctx *slapi.RequestContext, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary
	return returnValue, nil
}

// GetAccountUrlUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountUrlUptimeGraphData(ctx *slapi.RequestContext, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetAccountUrlUptimeSummary - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetAccountUrlUptimeSummary(ctx *slapi.RequestContext, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary
	return returnValue, nil
}

// GetHostUptimeDetail - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetHostUptimeDetail(ctx *slapi.RequestContext, configurationValueId int, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary_Detail, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary_Detail
	return returnValue, nil
}

// GetHostUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetHostUptimeGraphData(ctx *slapi.RequestContext, configurationValueId int, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}

// GetUrlUptimeDetail - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetUrlUptimeDetail(ctx *slapi.RequestContext, configurationValueId int, startDateTime string, endDateTime string) (*SoftLayer_Container_Account_Historical_Summary_Detail, error) {
	var returnValue *SoftLayer_Container_Account_Historical_Summary_Detail
	return returnValue, nil
}

// GetUrlUptimeGraphData - <nil>
func (softlayer_account_historical_report *SoftLayer_Account_Historical_Report) GetUrlUptimeGraphData(ctx *slapi.RequestContext, configurationValueId int, startDate string, endDate string) (*SoftLayer_Container_Graph, error) {
	var returnValue *SoftLayer_Container_Graph
	return returnValue, nil
}
