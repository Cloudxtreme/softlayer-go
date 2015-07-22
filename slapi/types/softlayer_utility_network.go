package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Utility_Network - <nil>
type SoftLayer_Utility_Network struct {
}

// IsPingable - Send a single ping from SoftLayer's application servers to the given IP address or
// hostname and return whether or not the remote host was pingable. Pinging is a good way to determine
// if a particular host is alive on the Internet. A host that's unreachable to ping may not necessarily
// be down, as many providers employ filtering to deny traffic. Pinging a hostname instead of an IP
// address may fail if DNS lookup for that hostname also fails. isPingable() differs from
// [[SoftLayer_Utility_Network::ping|ping()]] in that it returns whether an address was pingable while
// ping() retrieves the full results of the ping command.
func (softlayer_utility_network *SoftLayer_Utility_Network) IsPingable(commonOptions *slapi.CommonOptions, address string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// NsLookup - A method used to return the nameserver information for a given address
func (softlayer_utility_network *SoftLayer_Utility_Network) NsLookup(commonOptions *slapi.CommonOptions, address string, type_ string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// Ping - Send a single ping from SoftLayer's application servers to the given IP address or hostname
// and return the raw results of that command. The returned result is similar to the result received
// from running the command `ping -c 1` from a command shell. Pinging is a good way to determine if a
// particular host is alive on the Internet. A host that's unreachable to ping may not necessarily be
// down, as many providers employ filtering to deny traffic. Running ping on a hostname instead of an
// IP address may fail if DNS lookup for that hostname also fails.
func (softlayer_utility_network *SoftLayer_Utility_Network) Ping(commonOptions *slapi.CommonOptions, address string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// Traceroute - Perform an traceroute from SoftLayer's application servers to the given IP address or
// hostname and return the raw results of that command. The returned result is similar to the result
// received from running the command `traceroute` from a command shell. A traceroute sends small
// diagnostic packets to every hop along the network route to a given host. Traceroutes are useful
// tools for debugging network connectivity to a host on the Internet. Routing loops and intermediate
// hop timeouts help to narrow down problematic providers along the network chain. Some providers elect
// to deny on their networks, which may cause a traceroute to show skewed results. Furthermore, many
// providers assign a low priority to traceroutes in their infrastructure which may lead to inaccurate
// hop response times. Running traceroute on a hostname instead of an IP address may fail if DNS lookup
// for that hostname also fails.
func (softlayer_utility_network *SoftLayer_Utility_Network) Traceroute(commonOptions *slapi.CommonOptions, address string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// Whois - Perform a lookup from SoftLayer's application servers on the given IP address or hostname
// and return the raw results of that command. The returned result is similar to the result received
// from running the command `whois` from a command shell. A lookup queries a host's registrar to
// retrieve domain registrant information including registration date, expiry date, and the
// administrative, technical, billing, and abuse contacts responsible for a domain. lookups are useful
// for determining a physical contact responsible for a particular domain. lookups are also useful for
// determining domain availability. Running a lookup on an IP address queries for that IP block's
// ownership, and is helpful for determining a physical entity responsible for a certain IP address.
func (softlayer_utility_network *SoftLayer_Utility_Network) Whois(commonOptions *slapi.CommonOptions, address string) (string, error) {
	var returnValue string
	return returnValue, nil
}
