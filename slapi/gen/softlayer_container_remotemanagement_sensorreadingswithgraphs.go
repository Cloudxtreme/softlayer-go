package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs - The
// SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs contains the raw data retrieved from a
// server's remote management card. Along with the raw data, two sets of graphs will be returned. One
// set of graphs is used to display, using thermometer graphs, the temperatures (cpu(s) and system)
// retrieved from the management card. The other set is used to display speed for each of the server's
// fans.
type SoftLayer_Container_RemoteManagement_SensorReadingsWithGraphs struct {

	// RawData - The raw data returned from the server's remote management card.
	RawData []*SoftLayer_Container_RemoteManagement_SensorReading `json:"rawData"`

	// SpeedGraphs - no documentation
	SpeedGraphs []*SoftLayer_Container_RemoteManagement_Graphs_SensorSpeed `json:"speedGraphs"`

	// TemperatureGraphs - The graph(s) to display the server's cpu(s) and system temperatures.
	TemperatureGraphs []*SoftLayer_Container_RemoteManagement_Graphs_SensorTemperature `json:"temperatureGraphs"`
}
