package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Network_ContentDelivery_PointsOfPresence - SoftLayer's CDN content delivery
// network offering replicates your data to a number of Points of Presence (POP's) around the world.
// SoftLayer_Container_Network_ContentDelivery_PointsOfPresence models one of these POP locations.
type SoftLayer_Container_Network_ContentDelivery_PointsOfPresence struct {

	// Id - no documentation
	Id int `json:"id"`

	// Name - A CDN Point of Presence's name. This is typically the city that the POP is located in.
	Name string `json:"name"`
}
