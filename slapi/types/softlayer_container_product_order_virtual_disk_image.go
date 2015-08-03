package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Product_Order_Virtual_Disk_Image - This is the datatype that needs to be
// populated and sent to SoftLayer_Product_Order::placeOrder. This datatype has everything required to
// place a Portable Storage order with SoftLayer.
type SoftLayer_Container_Product_Order_Virtual_Disk_Image struct {

	// DiskDescription - no documentation
	DiskDescription string `json:"diskDescription"`
}

func (softlayer_container_product_order_virtual_disk_image *SoftLayer_Container_Product_Order_Virtual_Disk_Image) String() string {
	return "SoftLayer_Container_Product_Order_Virtual_Disk_Image"
}
