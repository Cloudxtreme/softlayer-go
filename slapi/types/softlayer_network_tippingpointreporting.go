package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_TippingPointReporting - <nil>
type SoftLayer_Network_TippingPointReporting struct {
}

// DrillDownAttack - This method, when given an attack signature ID (available in the return values of
// getReportForIpAddressOrSubnet and getSubnetReportForEntireAccount) and an IP Address and subnet
// mask, returns all attacks for that subnet in the specified time frame and direction. Once the
// results have been filtered, additional data is available, including starting and ending times for
// the attack, originating IP address and port, and destination IP address and port. CVE and Bugtraq
// information is not available at this level.
func (softlayer_network_tippingpointreporting *SoftLayer_Network_TippingPointReporting) DrillDownAttack(commonOptions *slapi.CommonOptions, signatureId string, IpAddress string, subnetMask int, timeFrame int, direction string) (*SoftLayer_Container_Network_IntrusionProtection_SubnetReport, error) {
	var returnValue *SoftLayer_Container_Network_IntrusionProtection_SubnetReport
	return returnValue, nil
}

// GetMainStatistics - This method returns the attack statistics for the current user's account and for
// the entire SoftLayer network. These attacks are recorded and monitored at the entry point to the
// network, and represent attacks in both directions. The data returned is: * Top attacks (by attack
// name) on datacenter Dal01 in the last hour (and last 24 hours) * Top attacks (by attack name) on IPs
// you own in the last hour (and last 24 hours) * Top IPs attacking IPs you own in the last hour (and
// last 24 hours) Each one of these lists can contain any number of items, the default is 5. The usable
// limit is less than 10, but setting the limit to an abnormally high value will effectively return all
// records. The data is returned as a collection of
// SoftLayer_Container_Network_IntrusionProtection_Statistics objects.
func (softlayer_network_tippingpointreporting *SoftLayer_Network_TippingPointReporting) GetMainStatistics(commonOptions *slapi.CommonOptions, numberOfAttacks int) ([]*SoftLayer_Container_Network_IntrusionProtection_Statistics, error) {
	var returnValue []*SoftLayer_Container_Network_IntrusionProtection_Statistics
	return returnValue, nil
}

// GetReportForIpAddressOrSubnet - This method expands on the getSubnetReportForEntireAccount method by
// offering the ability to filter by subnet or IP address. This method is identical to
// getSubnetReportForEntireAccount, but allows filtering by subnet. Like in the
// getSubnetReportForEntireAccount method, CVE and BugTraq IDs are provided, if available. This method
// should be called once an attack has been identified using getSubnetReportForEntireAccount (in which
// case "All Subnets" is the subnet) or getReportForIpAddressOrSubnet.
func (softlayer_network_tippingpointreporting *SoftLayer_Network_TippingPointReporting) GetReportForIpAddressOrSubnet(commonOptions *slapi.CommonOptions, IpAddress string, subnetMask int, timeFrame int, orderBy string, orderDirection string) ([]*SoftLayer_Container_Network_IntrusionProtection_SubnetReport, error) {
	var returnValue []*SoftLayer_Container_Network_IntrusionProtection_SubnetReport
	return returnValue, nil
}

// GetSubnetReportForEntireAccount - This method returns specific attacks by name for all subnets on
// the current user's account. The data returned is stored in
// SoftLayer_Container_Network_IntrusionProtection_SubnetReport objects, with the "subnet" value set to
// "All Subnets" The data is separated into "Inbound" and "Outbound" traffic. A significant amount of
// outbound attack traffic could indicate that your servers have been compromised. The data returned
// includes Attack Count, attack name, extended attack description, and IDs that correspond with the
// BugTraq or CVE databases. BugTraq can be accessed at [http://www.securityfocus.com/vulnerabilities]
// The CVE database is located at [http://cve.mitre.org/find/index.html] For more detailed information,
// use the getReportForIpAddressOrSubnet method
func (softlayer_network_tippingpointreporting *SoftLayer_Network_TippingPointReporting) GetSubnetReportForEntireAccount(commonOptions *slapi.CommonOptions, timeFrame int, orderBy string, orderDirection string, returnSubnetGroups bool) ([]*SoftLayer_Container_Network_IntrusionProtection_SubnetReport, error) {
	var returnValue []*SoftLayer_Container_Network_IntrusionProtection_SubnetReport
	return returnValue, nil
}
