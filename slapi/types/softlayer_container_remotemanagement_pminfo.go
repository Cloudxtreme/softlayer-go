package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_RemoteManagement_PmInfo - The SoftLayer_Container_RemoteManagement_PmInfo
// contains pminfo information retrieved from a server's remote management card.
type SoftLayer_Container_RemoteManagement_PmInfo struct {

	// PmInfoId - no documentation
	PmInfoId string `json:"pmInfoId,omitempty"`

	// PmInfoReading - no documentation
	PmInfoReading string `json:"pmInfoReading,omitempty"`
}

func (softlayer_container_remotemanagement_pminfo *SoftLayer_Container_RemoteManagement_PmInfo) String() string {
	return "SoftLayer_Container_RemoteManagement_PmInfo"
}
