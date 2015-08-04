package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Virtual_Guest_Configuration - The guest configuration container is used to
// provide configuration options for creating computing instances. Each configuration option will
// include both an itemPrice and a template . The itemPrice value will provide hourly and monthly costs
// (if either are applicable), and a description of the option. The template will provide a fragment of
// the request with the properties and values that must be sent when creating a computing instance with
// the option. The [[SoftLayer_Virtual_Guest/getCreateObjectOptions|getCreateObjectOptions]] method
// returns this data structure.
type SoftLayer_Container_Virtual_Guest_Configuration struct {

	// OperatingSystems - Available operating system options. The operatingSystemReferenceCode value in the
	// template is an identifier for a particular operating system. When provided exactly as shown in the
	// template, that operating system will be used. A reference code is structured as three tokens
	// separated by underscores. The first token represents the product, the second is the version of the
	// product, and the third is whether the OS is 32 or 64bit. When providing an
	// operatingSystemReferenceCode while ordering a computing instance the only token required to match
	// exactly is the product. The version token may be given as else it will require an exact match as
	// well. When the bits token is not provided, 64 bits will be assumed. Providing the value of for a
	// version will select the latest release of that product for the operating system. As this may change
	// over time, you should be sure that the release version is irrelevant for your applications. For
	// Windows based operating systems the version will represent both the release version (2008, 2012,
	// etc) and the edition (Standard, Enterprise, etc). For all other operating systems the version will
	// represent the major version (Centos 6, Ubuntu 12, etc) of that operating system, minor versions are
	// not represented in a reference code. Notice - Some operating systems are charged based on the value
	// specified in startCpus . The price which is used can be determined by calling
	// [[SoftLayer_Virtual_Guest/generateOrderTemplate|generateOrderTemplate]] with your desired device
	// specifications.
	OperatingSystems []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"operatingSystems,omitempty"`

	// Processors - Available processor options. The startCpus value in the template represents the number
	// of cores allocated to the computing instance. The dedicatedAccountHostOnlyFlag value in the template
	// represents whether the instance will run on hosts with instances belonging to other accounts.
	Processors []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"processors,omitempty"`

	// BlockDevices - Available block device options. A computing instance will have at least one block
	// device represented by a device number of '0' . The blockDevices.device value in the template
	// represents which device the option is for. The blockDevices.diskImage.capacity value in the template
	// represents the size, in gigabytes, of the disk. The localDiskFlag value in the template represents
	// whether the option is a local or SAN based disk. Note: The block device number '1' is reserved for
	// the disk attached to the computing instance.
	BlockDevices []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"blockDevices,omitempty"`

	// Datacenters - Available datacenter options. The datacenter.name value in the template represents
	// which datacenter the computing instance will be provisioned in.
	Datacenters []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"datacenters,omitempty"`

	// Memory - Available memory options. The maxMemory value in the template represents the amount of
	// memory, in megabytes, allocated to the computing instance.
	Memory []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"memory,omitempty"`

	// NetworkComponents - Available network component options. The networkComponent.maxSpeed value in the
	// template represents the link speed, in megabits per second, of the network connections for a
	// computing instance.
	NetworkComponents []*SoftLayer_Container_Virtual_Guest_Configuration_Option `json:"networkComponents,omitempty"`
}

func (softlayer_container_virtual_guest_configuration *SoftLayer_Container_Virtual_Guest_Configuration) String() string {
	return "SoftLayer_Container_Virtual_Guest_Configuration"
}
