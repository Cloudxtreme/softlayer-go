package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement_User - The credentials used for remote management such
// as username, password, etc...
type SoftLayer_Hardware_Component_RemoteManagement_User struct {

	// Password - The password used for this remote management command.
	Password string `json:"password,omitempty"`

	// Username - The username used for this remote management command.
	Username string `json:"username,omitempty"`

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware,omitempty"`

	// NetworkComponent - <nil>
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent,omitempty"`
}

func (softlayer_hardware_component_remotemanagement_user *SoftLayer_Hardware_Component_RemoteManagement_User) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_User"
}
