package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Hardware_Component_RemoteManagement_User - The credentials used for remote management such
// as username, password, etc...
type SoftLayer_Hardware_Component_RemoteManagement_User struct {

	// Password - The password used for this remote management command.
	Password string `json:"password"`

	// Username - The username used for this remote management command.
	Username string `json:"username"`
}

func (softlayer_hardware_component_remotemanagement_user *SoftLayer_Hardware_Component_RemoteManagement_User) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_User"
}

// SoftLayer_Hardware_Component_RemoteManagement_User_Extended is SoftLayer_Hardware_Component_RemoteManagement_User with all maskable types.
type SoftLayer_Hardware_Component_RemoteManagement_User_Extended struct {
	SoftLayer_Hardware_Component_RemoteManagement_User

	// Hardware - <nil>
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// NetworkComponent - <nil>
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`
}

func (softlayer_hardware_component_remotemanagement_user *SoftLayer_Hardware_Component_RemoteManagement_User_Extended) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_User"
}
