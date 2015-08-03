package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

// SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem -
// SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem models a single Microsoft
// Update as reported by SoftLayer's private Windows Server Update Services services. All servers
// purchased with Microsoft Windows retrieve updates from SoftLayer's servers by default.
type SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem struct {

	// Description - no documentation
	Description string `json:"description"`

	// Failed - Flag indicating that this patch failed to properly install
	Failed bool `json:"failed"`

	// KbArticleNumber - A Windows Update's knowledge base article number. Every Windows Update can be
	// referenced on the Microsoft Help and Support site at the URL
	KbArticleNumber int `json:"kbArticleNumber"`

	// Optional - Flag indicating that the update is entirely optionals
	Optional bool `json:"optional"`

	// RequiresReboot - Flag indicating that a reboot is needed for this update to be fully applied
	RequiresReboot bool `json:"requiresReboot"`
}

func (softlayer_container_utility_microsoft_windows_updateservices_updateitem *SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem) String() string {
	return "SoftLayer_Container_Utility_Microsoft_Windows_UpdateServices_UpdateItem"
}
