package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	slapi "go-softlayer/slapi"
)

// SoftLayer_Network_Subnet_Swip_Transaction - The SoftLayer_Network_Subnet_Swip_Transaction data type
// contains basic information tracked at SoftLayer to allow automation of Swip creation, update, and
// removal requests. A specific transaction is attached to an accountId and a subnetId. This also
// contains a "Status Name" which tells the customer what the transaction is doing: * Request is queued
// up to be sent to * The email request has been sent to * has confirmed that the request is good, and
// should be available in 24 hours * OK: The subnet has been checked with and it the transaction has
// completed correctly * A subnet is queued to be removed from ARIN's systems * The removal email
// request has been sent to * has confirmed that the removal request is good, and the subnet should be
// clear in in 24 hours * This specific Transaction has been removed from and is no longer in effect *
// Sometimes a request doesn't go through correctly and has to be manually processed by SoftLayer. This
// may take some time.
type SoftLayer_Network_Subnet_Swip_Transaction struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// Id - no documentation
	Id int `json:"id"`

	// StatusName - A Name describing which state a transaction is in.
	StatusName string `json:"statusName"`

	// Subnet - no documentation
	Subnet *SoftLayer_Network_Subnet `json:"subnet"`

	// SubnetId - no documentation
	SubnetId int `json:"subnetId"`
}

// FindMyTransactions - This function will return an array of SoftLayer_Network_Subnet_Swip_Transaction
// objects, one for each that is currently in transaction with This includes all swip registrations,
// swip removal requests, and objects that are currently OK.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) FindMyTransactions(commonOptions *slapi.CommonOptions) ([]*SoftLayer_Network_Subnet_Swip_Transaction, error) {
	var returnValue []*SoftLayer_Network_Subnet_Swip_Transaction
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_Subnet_Swip_Transaction object whose ID number
// corresponds to the ID number of the init parameter passed to the
// SoftLayer_Network_Subnet_Swip_transaction service. You can only retrieve Swip transactions tied to
// the account.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) GetObject(commonOptions *slapi.CommonOptions) (*SoftLayer_Network_Subnet_Swip_Transaction, error) {
	var returnValue *SoftLayer_Network_Subnet_Swip_Transaction
	return returnValue, nil
}

// RemoveAllSubnetSwips - This method finds all subnets attached to your account that are in OK status
// and starts transactions with allowing you to remove your registration information.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) RemoveAllSubnetSwips(commonOptions *slapi.CommonOptions) (int, error) {
	var returnValue int
	return returnValue, nil
}

// RemoveSwipData - This function, when called on an instantiated transaction, will allow you to start
// a transaction with allowing you to remove your registration information.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) RemoveSwipData(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ResendSwipData - This function will allow you to update ARIN's registration data for a subnet to
// your current data.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) ResendSwipData(commonOptions *slapi.CommonOptions) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SwipAllSubnets - swipAllSubnets finds all subnets attached to your account and attempts to create a
// transaction for all subnets that do not already have a transaction in progress.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) SwipAllSubnets(commonOptions *slapi.CommonOptions) (int, error) {
	var returnValue int
	return returnValue, nil
}

// UpdateAllSubnetSwips - This method finds all subnets attached to your account that are in status and
// updates their data with Use this function after you have updated your data if you want to keep up to
// date.
func (softlayer_network_subnet_swip_transaction *SoftLayer_Network_Subnet_Swip_Transaction) UpdateAllSubnetSwips(commonOptions *slapi.CommonOptions) (int, error) {
	var returnValue int
	return returnValue, nil
}
