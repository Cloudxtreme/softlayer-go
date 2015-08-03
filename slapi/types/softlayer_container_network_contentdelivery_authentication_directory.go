package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Container_Network_ContentDelivery_Authentication_Directory -
// SoftLayer_Container_Network_ContentDelivery_Authentication_Directory represents a token
// authentication directory on your CDN FTP or on your origin server.
type SoftLayer_Container_Network_ContentDelivery_Authentication_Directory struct {

	// CreateDate - The date that a token authentication directory was created.
	CreateDate *time.Time `json:"createDate"`

	// Name - The name of a directory or a file within a directory listing.
	Name string `json:"name"`

	// Type - The type of platform that a token authentication directory is defined for. Possible types are
	// Large, Small, Flash and Windows Media
	Type string `json:"type"`
}

func (softlayer_container_network_contentdelivery_authentication_directory *SoftLayer_Container_Network_ContentDelivery_Authentication_Directory) String() string {
	return "SoftLayer_Container_Network_ContentDelivery_Authentication_Directory"
}
