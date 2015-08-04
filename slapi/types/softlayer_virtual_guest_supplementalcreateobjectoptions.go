package types

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions - <nil>
type SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions struct {

	// ImmediateApprovalOnlyFlag - When explicitly set to true, createObject(s) will fail unless the order
	// is started automatically. This can be used by automated systems to fail an order that might
	// otherwise require manual approval. For multi-guest orders via
	// [[SoftLayer_Virtual_Guest/createObjects|createObjects]], this value must be the exact same for every
	// item.
	ImmediateApprovalOnlyFlag bool `json:"immediateApprovalOnlyFlag,omitempty"`

	// PostInstallScriptUri - URI of the script to be downloaded and executed after installation is
	// complete. This can be different for each virtual guest when multiple are sent to
	// [[SoftLayer_Virtual_Guest/createObjects|createObjects]].
	PostInstallScriptUri string `json:"postInstallScriptUri,omitempty"`
}

func (softlayer_virtual_guest_supplementalcreateobjectoptions *SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions) String() string {
	return "SoftLayer_Virtual_Guest_SupplementalCreateObjectOptions"
}
