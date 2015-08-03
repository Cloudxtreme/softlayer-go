package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Network_Subnet_IpAddress - The SoftLayer_Network_Subnet_IpAddress data type contains
// general information relating to a single SoftLayer IPv4 address.
type SoftLayer_Network_Subnet_IpAddress struct {

	// AllowedHost - The SoftLayer_Network_Storage_Allowed_Host information to connect this IP Address to
	// Network Storage supporting access control lists.
	AllowedHost *SoftLayer_Network_Storage_Allowed_Host `json:"allowedHost"`

	// ContextTunnelTranslationCount - A count of an IPSec network tunnel's address translations. These
	// translations use a SoftLayer ip address from an assigned static NAT subnet to deliver the packets to
	// the remote (customer) destination.
	ContextTunnelTranslationCount uint64 `json:"contextTunnelTranslationCount"`

	// ContextTunnelTranslations - An IPSec network tunnel's address translations. These translations use a
	// SoftLayer ip address from an assigned static NAT subnet to deliver the packets to the remote
	// (customer) destination.
	ContextTunnelTranslations []*SoftLayer_Network_Tunnel_Module_Context_Address_Translation `json:"contextTunnelTranslations"`

	// EndpointSubnetCount - A count of all the subnets routed to an IP address.
	EndpointSubnetCount uint64 `json:"endpointSubnetCount"`

	// EndpointSubnets - no documentation
	EndpointSubnets []*SoftLayer_Network_Subnet `json:"endpointSubnets"`

	// GuestNetworkComponent - A network component that is statically routed to an IP address.
	GuestNetworkComponent *SoftLayer_Virtual_Guest_Network_Component `json:"guestNetworkComponent"`

	// GuestNetworkComponentBinding - A network component that is statically routed to an IP address.
	GuestNetworkComponentBinding *SoftLayer_Virtual_Guest_Network_Component_IpAddress `json:"guestNetworkComponentBinding"`

	// Hardware - no documentation
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// Id - no documentation
	Id int `json:"id"`

	// IpAddress - no documentation
	IpAddress string `json:"ipAddress"`

	// IsBroadcast - Indicates if an IP address is reserved to be used as the network broadcast address and
	// cannot be assigned to a network interface
	IsBroadcast bool `json:"isBroadcast"`

	// IsGateway - Indicates if an IP address is reserved to a gateway and cannot be assigned to a network
	// interface
	IsGateway bool `json:"isGateway"`

	// IsNetwork - Indicates if an IP address is reserved to a network address and cannot be assigned to a
	// network interface
	IsNetwork bool `json:"isNetwork"`

	// IsReserved - Indicates if an IP address is reserved and cannot be assigned to a network interface
	IsReserved bool `json:"isReserved"`

	// NetworkComponent - A network component that is statically routed to an IP address.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// Note - no documentation
	Note string `json:"note"`

	// PrivateNetworkGateway - The network gateway appliance using this address as the private IP address.
	PrivateNetworkGateway *SoftLayer_Network_Gateway `json:"privateNetworkGateway"`

	// ProtectionAddress - <nil>
	ProtectionAddress []*SoftLayer_Network_Protection_Address `json:"protectionAddress"`

	// ProtectionAddressCount - no documentation
	ProtectionAddressCount uint64 `json:"protectionAddressCount"`

	// PublicNetworkGateway - The network gateway appliance using this address as the public IP address.
	PublicNetworkGateway *SoftLayer_Network_Gateway `json:"publicNetworkGateway"`

	// RemoteManagementNetworkComponent - An IPMI-based management network component of the IP address.
	RemoteManagementNetworkComponent *SoftLayer_Network_Component `json:"remoteManagementNetworkComponent"`

	// Subnet - no documentation
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetId - no documentation
	SubnetId int `json:"subnetId"`

	// SyslogEventsOneDay - All events for this IP address stored in the datacenter syslogs from the last
	// 24 hours
	SyslogEventsOneDay []*SoftLayer_Network_Logging_Syslog `json:"syslogEventsOneDay"`

	// SyslogEventsOneDayCount - A count of all events for this IP address stored in the datacenter syslogs
	// from the last 24 hours
	SyslogEventsOneDayCount uint64 `json:"syslogEventsOneDayCount"`

	// SyslogEventsSevenDayCount - A count of all events for this IP address stored in the datacenter
	// syslogs from the last 7 days
	SyslogEventsSevenDayCount uint64 `json:"syslogEventsSevenDayCount"`

	// SyslogEventsSevenDays - All events for this IP address stored in the datacenter syslogs from the
	// last 7 days
	SyslogEventsSevenDays []*SoftLayer_Network_Logging_Syslog `json:"syslogEventsSevenDays"`

	// TopTenSyslogEventsByDestinationPortOneDay - Top Ten network datacenter syslog events, grouped by
	// destination port, for the last 24 hours
	TopTenSyslogEventsByDestinationPortOneDay []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsByDestinationPortOneDay"`

	// TopTenSyslogEventsByDestinationPortOneDayCount - A count of top Ten network datacenter syslog
	// events, grouped by destination port, for the last 24 hours
	TopTenSyslogEventsByDestinationPortOneDayCount uint64 `json:"topTenSyslogEventsByDestinationPortOneDayCount"`

	// TopTenSyslogEventsByDestinationPortSevenDayCount - A count of top Ten network datacenter syslog
	// events, grouped by destination port, for the last 7 days
	TopTenSyslogEventsByDestinationPortSevenDayCount uint64 `json:"topTenSyslogEventsByDestinationPortSevenDayCount"`

	// TopTenSyslogEventsByDestinationPortSevenDays - Top Ten network datacenter syslog events, grouped by
	// destination port, for the last 7 days
	TopTenSyslogEventsByDestinationPortSevenDays []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsByDestinationPortSevenDays"`

	// TopTenSyslogEventsByProtocolsOneDay - Top Ten network datacenter syslog events, grouped by source
	// port, for the last 24 hours
	TopTenSyslogEventsByProtocolsOneDay []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsByProtocolsOneDay"`

	// TopTenSyslogEventsByProtocolsOneDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source port, for the last 24 hours
	TopTenSyslogEventsByProtocolsOneDayCount uint64 `json:"topTenSyslogEventsByProtocolsOneDayCount"`

	// TopTenSyslogEventsByProtocolsSevenDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source port, for the last 7 days
	TopTenSyslogEventsByProtocolsSevenDayCount uint64 `json:"topTenSyslogEventsByProtocolsSevenDayCount"`

	// TopTenSyslogEventsByProtocolsSevenDays - Top Ten network datacenter syslog events, grouped by source
	// port, for the last 7 days
	TopTenSyslogEventsByProtocolsSevenDays []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsByProtocolsSevenDays"`

	// TopTenSyslogEventsBySourceIpOneDay - Top Ten network datacenter syslog events, grouped by source ip
	// address, for the last 24 hours
	TopTenSyslogEventsBySourceIpOneDay []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsBySourceIpOneDay"`

	// TopTenSyslogEventsBySourceIpOneDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source ip address, for the last 24 hours
	TopTenSyslogEventsBySourceIpOneDayCount uint64 `json:"topTenSyslogEventsBySourceIpOneDayCount"`

	// TopTenSyslogEventsBySourceIpSevenDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source ip address, for the last 7 days
	TopTenSyslogEventsBySourceIpSevenDayCount uint64 `json:"topTenSyslogEventsBySourceIpSevenDayCount"`

	// TopTenSyslogEventsBySourceIpSevenDays - Top Ten network datacenter syslog events, grouped by source
	// ip address, for the last 7 days
	TopTenSyslogEventsBySourceIpSevenDays []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsBySourceIpSevenDays"`

	// TopTenSyslogEventsBySourcePortOneDay - Top Ten network datacenter syslog events, grouped by source
	// port, for the last 24 hours
	TopTenSyslogEventsBySourcePortOneDay []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsBySourcePortOneDay"`

	// TopTenSyslogEventsBySourcePortOneDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source port, for the last 24 hours
	TopTenSyslogEventsBySourcePortOneDayCount uint64 `json:"topTenSyslogEventsBySourcePortOneDayCount"`

	// TopTenSyslogEventsBySourcePortSevenDayCount - A count of top Ten network datacenter syslog events,
	// grouped by source port, for the last 7 days
	TopTenSyslogEventsBySourcePortSevenDayCount uint64 `json:"topTenSyslogEventsBySourcePortSevenDayCount"`

	// TopTenSyslogEventsBySourcePortSevenDays - Top Ten network datacenter syslog events, grouped by
	// source port, for the last 7 days
	TopTenSyslogEventsBySourcePortSevenDays []*SoftLayer_Network_Logging_Syslog `json:"topTenSyslogEventsBySourcePortSevenDays"`

	// VirtualGuest - no documentation
	VirtualGuest *SoftLayer_Virtual_Guest `json:"virtualGuest"`

	// VirtualLicenseCount - A count of virtual licenses allocated for an IP Address.
	VirtualLicenseCount uint64 `json:"virtualLicenseCount"`

	// VirtualLicenses - no documentation
	VirtualLicenses []*SoftLayer_Software_VirtualLicense `json:"virtualLicenses"`
}

func (softlayer_network_subnet_ipaddress *SoftLayer_Network_Subnet_IpAddress) String() string {
	return "SoftLayer_Network_Subnet_IpAddress"
}
