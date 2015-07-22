package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Resource_Metadata - <nil>
type SoftLayer_Resource_Metadata struct {
}

// GetBackendMacAddresses - The getBackendMacAddresses method retrieves a list of backend MAC addresses
// for the resource
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetBackendMacAddresses() ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetDatacenter - The getDatacenter method retrieves the name of the datacenter in which the resource
// is located.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetDatacenter() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetDatacenterId - The getDatacenterId retrieves the ID for the datacenter in which the resource is
// located.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetDatacenterId() (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetDomain - The getDomain method retrieves the hostname for the resource.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetDomain() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetFrontendMacAddresses - The getFrontendMacAddresses method retrieves a list of frontend MAC
// addresses for the resource
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetFrontendMacAddresses() ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetFullyQualifiedDomainName - The getFullyQualifiedDomainName method provides the user with a
// combined return which includes the hostname and domain for the resource. Because this method returns
// multiple pieces of information, it avoids the need to use multiple methods to return the desired
// information.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetFullyQualifiedDomainName() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetHostname - The getHostname method retrieves the hostname for the resource.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetHostname() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetId - no documentation
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetId() (int, error) {
	var returnValue int
	return returnValue, nil
}

// GetPrimaryBackendIpAddress - The getPrimaryBackendIpAddress method retrieves the primary backend IP
// address for the resource
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetPrimaryBackendIpAddress() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPrimaryIpAddress - The getPrimaryIpAddress method retrieves the primary IP address for the
// resource. For resources with a frontend network, the frontend IP address will be returned. For
// resources that have been provisioned with only a backend network, the backend IP address will be
// returned, as a frontend address will not exist.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetPrimaryIpAddress() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetProvisionState - The getProvisionState method retrieves the provision state of the resource. The
// provision state may be used to determine when it is considered safe to perform additional setup
// operations. The method returns to indicate the provision is in progress and when the provision is
// complete.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetProvisionState() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetRouter - The getRouter method will return the router associated with a network component. When
// the router is redundant, the hostname of the redundant group will be returned, rather than the
// router hostname.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetRouter(macAddress string) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetServiceResource - The getServiceResource method retrieves a specific service resource associated
// with the resource. Service resources are additional resources that may be used by this resource.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetServiceResource(serviceName string, index int) (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetServiceResources - The getServiceResources method retrieves all service resources associated with
// the resource. Service resources are additional resources that may be used by this resource. The
// output format is
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetServiceResources() ([]*SoftLayer_Network_Service_Resource, error) {
	var returnValue []*SoftLayer_Network_Service_Resource
	return returnValue, nil
}

// GetTags - The getTags method retrieves all tags associated with the resource. Tags are single
// keywords assigned to a resource that assist the user in identifying the resource and its roles when
// performing a simple search. Tags are assigned by any user with access to the resource.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetTags() ([]string, error) {
	var returnValue []string
	return returnValue, nil
}

// GetUserMetadata - The getUserMetadata method retrieves metadata completed by users who interact with
// the resource. Metadata gathered using this method is unique to parameters set using the
// '''setUserMetadata''' method, which must be executed prior to completing this method. User metadata
// may also be provided while placing an order for a resource.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetUserMetadata() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetVlanIds - The getVlanIds method returns a list of IDs for the network component matching the
// provided MAC address associated with the resource. For each return, the native will appear first,
// followed by any trunked VLANs associated with the network component.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetVlanIds(macAddress string) ([]int, error) {
	var returnValue []int
	return returnValue, nil
}

// GetVlans - The getVlans method returns a list of numbers for the network component matching the
// provided MAC address associated with the resource. For each return, the native will appear first,
// followed by any trunked VLANs associated with the network component.
func (softlayer_resource_metadata *SoftLayer_Resource_Metadata) GetVlans(macAddress string) ([]int, error) {
	var returnValue []int
	return returnValue, nil
}
