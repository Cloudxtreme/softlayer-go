package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Monitoring_Robot_Status - Your monitoring robot will be in "Active" status under normal
// circumstances. If you perform an OS reload, your robot will be in "Reclaim" status until it's
// reloaded on your server or virtual server. Advanced monitoring system requires "Nimsoft Monitoring
// (Advanced)" service running and TCP ports 48000 - 48020 to be open on your server or virtual server.
// Monitoring agents cannot be managed nor can the usage data be updated if these ports are closed.
// Your monitoring robot will be in "Limited Connectivity" status if our monitoring management system
// cannot communicate with your system. See [[SoftLayer_Monitoring_Robot::resetStatus|resetStatus]]
// service for more details.
type SoftLayer_Monitoring_Robot_Status struct {

	// Description - no documentation
	Description string `json:"description"`

	// Id - no documentation
	Id int `json:"id"`

	// Name - no documentation
	Name string `json:"name"`
}

func (softlayer_monitoring_robot_status *SoftLayer_Monitoring_Robot_Status) String() string {
	return "SoftLayer_Monitoring_Robot_Status"
}
