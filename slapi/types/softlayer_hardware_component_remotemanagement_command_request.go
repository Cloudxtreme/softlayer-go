package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Hardware_Component_RemoteManagement_Command_Request - The
// SoftLayer_Hardware_Component_RemoteManagement_Command_Request contains details for remote management
// commands issued to a server's remote management card. Details for remote management commands such as
// powerOn, powerOff, powerCycle, rebootDefault, rebootSoft, rebootHard can be retrieved. Details such
// as the user who issued the command, the id of the remote management card the command was issued,
// when the command was issued may be retrieved.
type SoftLayer_Hardware_Component_RemoteManagement_Command_Request struct {

	// Processed - Execution status of the remote management command. True is successful. False is failure.
	Processed bool `json:"processed"`

	// CreateDate - The timestamp the remote management command was issued.
	CreateDate *time.Time `json:"createDate"`

	// HardwareId - no documentation
	HardwareId int `json:"hardwareId"`

	// ModifyDate - The timestamp recorded when the remote management command returned a status of the
	// command issued.
	ModifyDate *time.Time `json:"modifyDate"`
}

func (softlayer_hardware_component_remotemanagement_command_request *SoftLayer_Hardware_Component_RemoteManagement_Command_Request) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_Command_Request"
}

// SoftLayer_Hardware_Component_RemoteManagement_Command_Request_Extended is SoftLayer_Hardware_Component_RemoteManagement_Command_Request with all maskable types.
type SoftLayer_Hardware_Component_RemoteManagement_Command_Request_Extended struct {
	SoftLayer_Hardware_Component_RemoteManagement_Command_Request

	// Hardware - The id of the hardware to perform the remote management or powerstrip command on.
	Hardware *SoftLayer_Hardware `json:"hardware"`

	// NetworkComponent - A hardware's network components. Network components are hardware components such
	// as cards or Ethernet cards.
	NetworkComponent *SoftLayer_Network_Component `json:"networkComponent"`

	// RemoteManagementCommand - no documentation
	RemoteManagementCommand *SoftLayer_Hardware_Component_RemoteManagement_Command `json:"remoteManagementCommand"`

	// User - Information regarding the user who issued the remote management command.
	User *SoftLayer_User_Customer `json:"user"`
}

func (softlayer_hardware_component_remotemanagement_command_request *SoftLayer_Hardware_Component_RemoteManagement_Command_Request_Extended) String() string {
	return "SoftLayer_Hardware_Component_RemoteManagement_Command_Request"
}
