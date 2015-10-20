package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature - The
// SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature contains graphs to display the cpu(s)
// and system temperatures retrieved from the management card using thermometer graphs.
type SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature struct {

	// Title - no documentation
	Title string `json:"title,omitempty"`

	// Graph - The graph to display the server's cpu(s) and system temperatures.
	Graph string `json:"graph,omitempty"`
}

func (softlayer_container_remotemanagement_graphs_sensortemperature *SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature) String() string {
	return "SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature"
}
