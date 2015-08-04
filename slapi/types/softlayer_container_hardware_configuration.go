package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Hardware_Configuration - The hardware configuration container is used to provide
// configuration options for servers. Each configuration option will include both an itemPrice and a
// template . The itemPrice value will provide hourly and monthly costs (if either are applicable), and
// a description of the option. The template will provide a fragment of the request with the properties
// and values that must be sent when creating a server with the option. The
// [[SoftLayer_Hardware/getCreateObjectOptions|getCreateObjectOptions]] method returns this data
// structure.
type SoftLayer_Container_Hardware_Configuration struct {

	// FixedConfigurationPresets - Available fixed configuration preset options. The
	// fixedConfigurationPreset.keyName value in the template is an identifier for a particular fixed
	// configuration. When provided exactly as shown in the template, that fixed configuration will be
	// used. When providing a fixedConfigurationPreset.keyName while ordering a server the processors and
	// hardDrives configuration options cannot be used.
	FixedConfigurationPresets []*SoftLayer_Container_Hardware_Configuration_Option `json:"fixedConfigurationPresets,omitempty"`

	// HardDrives - Available hard drive options. A server will have at least one hard drive. The
	// hardDrives.capacity value in the template represents the size, in gigabytes, of the disk.
	HardDrives []*SoftLayer_Container_Hardware_Configuration_Option `json:"hardDrives,omitempty"`

	// NetworkComponents - Available network component options. The networkComponent.maxSpeed value in the
	// template represents the link speed, in megabits per second, of the network connections for a server.
	NetworkComponents []*SoftLayer_Container_Hardware_Configuration_Option `json:"networkComponents,omitempty"`

	// OperatingSystems - Available operating system options. The operatingSystemReferenceCode value in the
	// template is an identifier for a particular operating system. When provided exactly as shown in the
	// template, that operating system will be used. A reference code is structured as three tokens
	// separated by underscores. The first token represents the product, the second is the version of the
	// product, and the third is whether the OS is 32 or 64bit. When providing an
	// operatingSystemReferenceCode while ordering a server the only token required to match exactly is the
	// product. The version token may be given as else it will require an exact match as well. When the
	// bits token is not provided, 64 bits will be assumed. Providing the value of for a version will
	// select the latest release of that product for the operating system. As this may change over time,
	// you should be sure that the release version is irrelevant for your applications. For Windows based
	// operating systems the version will represent both the release version (2008, 2012, etc) and the
	// edition (Standard, Enterprise, etc). For all other operating systems the version will represent the
	// major version (Centos 6, Ubuntu 12, etc) of that operating system, minor versions are represented in
	// few reference codes where they are significant.
	OperatingSystems []*SoftLayer_Container_Hardware_Configuration_Option `json:"operatingSystems,omitempty"`

	// Processors - Available processor options. The processorCoreAmount value in the template represents
	// the number of cores allocated to the server. The memoryCapacity value in the template represents the
	// amount of memory, in gigabytes, allocated to the server.
	Processors []*SoftLayer_Container_Hardware_Configuration_Option `json:"processors,omitempty"`

	// Datacenters - Available datacenter options. The datacenter.name value in the template represents
	// which datacenter the server will be provisioned in.
	Datacenters []*SoftLayer_Container_Hardware_Configuration_Option `json:"datacenters,omitempty"`
}

func (softlayer_container_hardware_configuration *SoftLayer_Container_Hardware_Configuration) String() string {
	return "SoftLayer_Container_Hardware_Configuration"
}
