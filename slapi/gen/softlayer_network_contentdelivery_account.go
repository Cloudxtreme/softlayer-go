package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_ContentDelivery_Account - The SoftLayer_Network_ContentDelivery_Account data type
// models an individual CDN account. CDN accounts contain references to the SoftLayer customer account
// they belong to, login credentials for upload services, and a CDN account's status. Please contact
// SoftLayer sales to purchase or cancel a CDN account
type SoftLayer_Network_ContentDelivery_Account struct {

	// Account - The customer account that a CDN account belongs to.
	Account *SoftLayer_Account `json:"account"`

	// AccountId - The internal identifier of the customer account that a CDN account belongs to.
	AccountId int `json:"accountId"`

	// AssociatedCdnAccountId - The CDN account id that this CDN account is associated with.
	AssociatedCdnAccountId string `json:"associatedCdnAccountId"`

	// AuthenticationIpAddressCount - A count of the IP addresses that are used for the content
	// authentication service.
	AuthenticationIpAddressCount uint64 `json:"authenticationIpAddressCount"`

	// AuthenticationIpAddresses - The IP addresses that are used for the content authentication service.
	AuthenticationIpAddresses []*SoftLayer_Network_ContentDelivery_Authentication_Address `json:"authenticationIpAddresses"`

	// BillingItem - no documentation
	BillingItem *SoftLayer_Billing_Item `json:"billingItem"`

	// CdnAccountName - no documentation
	CdnAccountName string `json:"cdnAccountName"`

	// CdnAccountNote - no documentation
	CdnAccountNote string `json:"cdnAccountNote"`

	// CdnSolutionName - no documentation
	CdnSolutionName string `json:"cdnSolutionName"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// DependantServiceFlag - Indicates if CDN account is dependent on other service. If set, this CDN
	// account is limited to these services: createOriginPullMapping, deleteOriginPullRule,
	// getOriginPullMappingInformation, getCdnUrls, purgeCache, loadContent, manageHttpCompression
	DependantServiceFlag bool `json:"dependantServiceFlag"`

	// Id - no documentation
	Id int `json:"id"`

	// LegacyCdnFlag - no documentation
	LegacyCdnFlag bool `json:"legacyCdnFlag"`

	// LogEnabledFlag - no documentation
	LogEnabledFlag string `json:"logEnabledFlag"`

	// ProviderPortalAccessFlag - Indicates if customer is allowed to access the CDN provider's management
	// portal.
	ProviderPortalAccessFlag bool `json:"providerPortalAccessFlag"`

	// Status - A CDN account's status presented in a more detailed data type.
	Status *SoftLayer_Network_ContentDelivery_Account_Status `json:"status"`

	// StatusId - no documentation
	StatusId int `json:"statusId"`

	// TokenAuthenticationEnabledFlag - Indicates if the token authentication service is enabled or not.
	TokenAuthenticationEnabledFlag bool `json:"tokenAuthenticationEnabledFlag"`
}

// AuthenticateResourceRequest - Internap servers attempts to validate a token before serving a
// protected content. SoftLayer customer does not need to invoke this method. Please refer to
// [[SoftLayer_Network_ContentDelivery_Authentication_Token|Authentication Token]] object for more
// details on Content Authentication Service.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) AuthenticateResourceRequest(parameter SoftLayer_Container_Network_ContentDelivery_Authentication_Parameter) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateDirectory - You can further organize your contents on the CDN FTP server by creating sub
// directories. This method creates a directory on the CDN FTP server. A user must have privilege to
// use this method. A directory name must be an absolute path and you can only create sub directories
// in /media folder.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) CreateDirectory(directoryName string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateFtpUser - This method allows you to create a default CDN FTP user record on the
// ftp.cdnlayer.service.softlayer.com server. As with a CDN FTP user account, you can upload contents
// to the CDN host server through the SoftLayer private network. SoftLayer currently allows only one
// FTP user for each CDN account. Your default CDN FTP user record is created upon successful creation
// of a CDN account. You may not need to use this method at all. This is provided in support of the
// previous CDN customers. SoftLayer may offer multiple CDN FTP users for a single CDN account in the
// future. Optionally, you can provide a new password when invoking this method and a new password must
// follow the rules below: * ...must be between 8 and 20 characters long * ...must be an alphanumeric
// value * ...can contain these characters: - _ ! % # $ ^ &
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) CreateFtpUser(newPassword string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateOriginPullMapping - With Origin Pull, content is pulled from your origin server as needed and
// then delivered to visitors. You do not need to upload your files to the CDN you can utilize the
// files that currently exist on your origin server. It will take 10 to 15 minutes for this to take
// effect after you create an Origin Pull rule. Origin Pull is only supported for protocol and you
// would continue to use the CDN FTP for Flash and Windows Media streaming services. A valid origin
// host can include a directory information. You may include an authentication username and password
// along with an origin host. If you set an authentication username and password, CDN servers will
// include "Authorization:" header in every request. You may use the "Authorization:" header to grant
// access to CDN servers or you may use it to distinguish CDN servers from normal visitors. Here is a
// list of valid origin hosts: * www.website.com * www.website.com/cdn_content *
// cdn_user:password@www.website.com * cdn_user:password@www.website.com/images An authentication
// username should be an alphanumeric string and allowed special characters are . - An authentication
// password should be an alphanumeric string and allowed special characters are . - _ ! # $ % ^ & Both
// username and password must be between 3 to 10 characters long. CDN nodes will cache your contents
// and you can control cache lifetime by modifying your web server's configuration. This method also
// creates a FTP directory restriction upon successful Origin Pull set up. You will not be able to
// access /media/http directory since contents will be pulled from your origin server. An origin domain
// must be a valid domain name and it can contain path information. This can help you organize contents
// on your origin server. For example, you could set an origin domain as: mydomain.com/cdn_contents A
// record allows you to have a customized You can get rid of your CDN account name from the A valid for
// the Origin Pull method must point to .http.cdn.softlayer.net. There are 2 types of origin pull
// mappings. The one with a record or the one without a record and they work very differently. gzip is
// supported if your web server sends a proper gzip header. For more details, visit our
// [http://knowledgelayer.softlayer.com/topic/cdn KnowledgeLayer]
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) CreateOriginPullMapping(mappingObject SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateOriginPullRule - This method is deprecated, please use
// [[[[SoftLayer_Network_ContentDelivery_Account::createOriginPullMapping|createOriginPullMapping]]
// method instead.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) CreateOriginPullRule(originDomain string, cnameRecord string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateTokenAuthenticationDirectory - You need to specify a directory on your CDN FTP or on your
// origin host in which your secure content resides to enable the token authentication . It will take
// about about 30 minutes for a newly configured token authentication directory to take effect.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) CreateTokenAuthenticationDirectory(directory string, mediaType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteFtpUser - This method deletes your FTP user record on the ftp.cdnlayer.service.softlayer.com
// server. Refer to the service overview of
// [[SoftLayer_Network_ContentDelivery_Account::createFtpUser|createFtpUser]] method for more
// information on the CDN FTP server.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) DeleteFtpUser() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DeleteOriginPullRule - This method removes an Origin Pull domain rule. Once an Origin Pull rule is
// removed, you will be able to access the /media/http directory. It will take 10 to 15 minutes for
// this to take effect after you remove your Origin Pull rule. Cached contents on CDN POPs may live
// longer than 15 minutes.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) DeleteOriginPullRule(originMappingId string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// DisableLogging - no documentation
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) DisableLogging() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// EnableLogging - no documentation
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) EnableLogging() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetAllPopsBandwidthData - This method returns bandwidth data for each
// [[SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary|POP Bandwidth]]
// object contains a starting time, ending time, total bytes, POP name and bandwidth unit. POP
// bandwidth data is updated everyday at 22:50 CST (or It queries and stores POP data from the day
// before. It is a more resource intensive process than a regular CDN bandwidth update thus we run this
// once a day. Since the POP bandwidth data is delayed for a day, there is no correction process for
// POP data. The POP bandwidth is not associated with any billing process and is mainly used to
// generate a POP bandwidth graph.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetAllPopsBandwidthData(beginDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_Bandwidth_PointsOfPresence_Summary
	return returnValue, nil
}

// GetAllPopsBandwidthImage - This method returns a bandwidth graph for every POP wrapped in
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object. A POP bandwidth graph
// shows bandwidth consumption per each POP in a bar graph.
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object contains a begin time,
// end time, title of the graph, binary date, in and outbound total bandwidth in bytes
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetAllPopsBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time, unit string) (*SoftLayer_Container_Bandwidth_GraphOutputsExtended, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputsExtended
	return returnValue, nil
}

// GetAuthenticationServiceEndpoints - CDN servers will invoke a Web Service method to validate a
// content authentication token. This method returns all token validation web service endpoints set for
// a CDN account. You can override the default web service by calling
// [[SoftLayer_Network_ContentDelivery_Authentication_Token|setContentAuthenticationWsdl
// setContentAuthenticationWsdl]] method.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetAuthenticationServiceEndpoints() ([]*SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_Authentication_ServiceEndpoint
	return returnValue, nil
}

// GetBandwidthData - This method returns bandwidth data for a given time range. It returns an array of
// [[SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary|bandwidth summary]] objects.
// [[SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary|Bandwidth summary]] object contains
// a beginning time, ending time, and bandwidth in bytes. A Beginning and ending date parameters have
// to be a timestamp in "yyyy-mm-dd HH24:mi:ss" format and it assumes the time is in Central Standard
// Time or Central Daylight Time time zone. CDN bandwidth data is stored in Greenwich Mean Time
// internally and converts a beginning and ending time to GMT before querying. Unlike server bandwidth,
// CDN bandwidth returns total bytes consumed within an hour. For example, if you pass "2008-10-10
// 00:00:00" for a beginning time and "2008-10-10 05:00:00" for an ending time, your return value will
// have 6 elements of bandwidth summary objects. The first bandwidth summary object will have the total
// bytes consumed between 2008-10-10 00:00:00 and 2008-10-10 05:00:00. And the last object will have
// the bandwidth consumed between 2008-10-10 05:00:00 and 2008-10-10 00:59:59. The bandwidth data is
// updated at 10 minutes after every hour. The queried data is on a two hour time delay. The two hour
// delay is required to gather bandwidth data from each POP and that is the minimum delay required to
// create a feasible graph. It usually takes about 8 hours to reconcile all the data from every CDN
// This hourly data is corrected after 24 hours if necessary. If you consume a large amount of
// bandwidth, your bandwidth data will be updated the next day.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetBandwidthData(beginDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_Bandwidth_Summary
	return returnValue, nil
}

// GetBandwidthDataWithTypes - This method returns bandwidth data for a given time range. It returns an
// array of [[SoftLayer_Container_Network_ContentDelivery_Report_Usage|bandwidth usage report]]
// objects. These will be first sorted by timestamp, and there will be one entry with that timestamp
// for each enabled region. The region type is provided only when non-region-specific data is returned.
// [[SoftLayer_Container_Network_ContentDelivery_Report_Usage|bandwidth usage report]] objects with a
// region will never contain non-region-specific data. Non-region-specific values are standardTotal and
// sslTotal; standardTotal is computed by adding the Large, Windows Media, Flash and Application
// Delivery Network bandwidth. The sslTotal is computed by adding the Large SSL bandwidth and the
// Application Delivery Network SSL bandwidth. A Beginning and ending date parameters have to be a
// timestamp in "yyyy-mm-dd HH24:mi:ss" format and it assumes the time is in Central Standard Time or
// Central Daylight Time time zone. CDN bandwidth data is stored in Greenwich Mean Time internally and
// converts a beginning and ending time to GMT before querying. Unlike server bandwidth, CDN bandwidth
// returns total bytes consumed within an hour. For example, if you pass "2008-10-10 00:00:00" for a
// beginning time and "2008-10-10 05:00:00" for an ending time, your return value will have 6 elements
// of bandwidth summary objects. The first bandwidth summary object will have the total bytes consumed
// between 2008-10-10 00:00:00 and 2008-10-10 05:00:00. And the last object will have the bandwidth
// consumed between 2008-10-10 05:00:00 and 2008-10-10 00:59:59. The bandwidth data is updated at 10
// minutes after every hour. The queried data is on a two hour time delay. The two hour delay is
// required to gather bandwidth data from each POP and that is the minimum delay required to create a
// feasible graph. It usually takes about 8 hours to reconcile all the data from every CDN This hourly
// data is corrected after 24 hours if necessary. If you consume a large amount of bandwidth, your
// bandwidth data will be updated the next day.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetBandwidthDataWithTypes(beginDateTime time.Time, endDateTime time.Time, period string) ([]*SoftLayer_Container_Network_ContentDelivery_Report_Usage, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_Report_Usage
	return returnValue, nil
}

// GetBandwidthImage - This method returns a bandwidth graph wrapped in
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object.
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object contains a starting
// time, ending time, graph title, graph binary data, and in and outbound total bytes.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetBandwidthImage(title string, beginDateTime time.Time, endDateTime time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputsExtended, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputsExtended
	return returnValue, nil
}

// GetCustomerOrigins - An origin pull mapping is a combination of your customer origin record and a
// (optional) record. You can now keep track of your customer origin records separate from your
// records. This service returns your customer origin records.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetCustomerOrigins(mediaType string) ([]*SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping
	return returnValue, nil
}

// GetDirectoryInformation - This method returns an array of
// [[SoftLayer_Container_Network_Directory_Listing|Directory Listing]] objects. You must have privilege
// and you can only retrieve directory information within /media directory. A
// [[SoftLayer_Container_Network_Directory_Listing|Directory Listing]] object contains type (indicating
// whether it is a file or a directory), name and file count if it is a directory.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetDirectoryInformation(directoryName string) ([]*SoftLayer_Container_Network_Directory_Listing, error) {
	var returnValue []*SoftLayer_Container_Network_Directory_Listing
	return returnValue, nil
}

// GetDiskSpaceUsageDataByDate - This method returns disk space usage data for your CDN
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetDiskSpaceUsageDataByDate(beginDateTime time.Time, endDateTime time.Time) ([]*SoftLayer_Metric_Tracking_Object_Data, error) {
	var returnValue []*SoftLayer_Metric_Tracking_Object_Data
	return returnValue, nil
}

// GetDiskSpaceUsageImageByDate - This method returns a disk usage graph wrapped in
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object.
// [[SoftLayer_Container_Bandwidth_GraphOutputsExtended|Bandwidth Graph]] object contains a starting
// time, ending time, graph title, graph binary data, and in and outbound total bytes.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetDiskSpaceUsageImageByDate(beginDateTime time.Time, endDateTime time.Time) (*SoftLayer_Container_Bandwidth_GraphOutputs, error) {
	var returnValue *SoftLayer_Container_Bandwidth_GraphOutputs
	return returnValue, nil
}

// GetFtpAttributes - This method returns your login credentials to the CDN FTP server
// (ftp.cdnlayer.service.softlayer.com server). You must have privilege. Refer to the service overview
// of [[SoftLayer_Network_ContentDelivery_Account::createFtpUser|createFtpUser]] method for more
// information on the CDN FTP server. If you want to download raw log files, prefix the username with
// (without quotes) when logging in. SoftLayer designed CDN accounts so they can have multiple CDN FTP
// users. However, this method returns the default CDN FTP user information: multi user support may be
// implemented in the future.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetFtpAttributes() (*SoftLayer_Container_Network_Authentication_Data, error) {
	var returnValue *SoftLayer_Container_Network_Authentication_Data
	return returnValue, nil
}

// GetMediaUrls - This method returns CDN URLs for static file (http), Flash streaming (rtmp) and
// Window Media (mms) streaming services. You can generate your CDN URLs based on the information
// retrieved by this method.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetMediaUrls() ([]*SoftLayer_Container_Network_ContentDelivery_SupportedProtocol, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_SupportedProtocol
	return returnValue, nil
}

// GetObject - getObject retrieves the SoftLayer_Network_ContentDelivery_Account object whose ID number
// corresponds to the ID number of the initial parameter passed to the
// SoftLayer_Network_ContentDelivery_Account service. You can only retrieve CDN accounts assigned to
// your SoftLayer customer account.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetObject() (*SoftLayer_Network_ContentDelivery_Account, error) {
	var returnValue *SoftLayer_Network_ContentDelivery_Account
	return returnValue, nil
}

// GetOriginPullMappingInformation - This method returns a list of origin pull configuration data.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetOriginPullMappingInformation() ([]*SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_OriginPull_Mapping
	return returnValue, nil
}

// GetOriginPullSupportedMediaUrls - This method returns CDN URLs that supports Origin Pull mappings.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetOriginPullSupportedMediaUrls() ([]*SoftLayer_Container_Network_ContentDelivery_SupportedProtocol, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_SupportedProtocol
	return returnValue, nil
}

// GetOriginPullUrl - This method returns the domain name of your Origin Pull rule. It assumes you have
// already setup an Origin Pull rule. Otherwise, it will throw an exception. A returning value is the
// value of the first parameter (origin pull domain) you provided to
// [[SoftLayer_Network_ContentDelivery_Account::createOriginPullRule|createOriginPullRule]] method. See
// Error Handling section below for possible errors.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetOriginPullUrl() (string, error) {
	var returnValue string
	return returnValue, nil
}

// GetPopNames - This method returns an array of CDN POPs (Points of Presence) object.
// [[SoftLayer_Container_Network_ContentDelivery_PointsOfPresence|POP object]] object contains the POP
// id and name.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetPopNames() ([]*SoftLayer_Container_Network_ContentDelivery_PointsOfPresence, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_PointsOfPresence
	return returnValue, nil
}

// GetProviderPortalCredentials - This method returns your login credentials to the CDN provider
// portal.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetProviderPortalCredentials() (*SoftLayer_Container_Network_Authentication_Data, error) {
	var returnValue *SoftLayer_Container_Network_Authentication_Data
	return returnValue, nil
}

// GetTokenAuthenticationDirectories - This method returns all token authentication directories.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetTokenAuthenticationDirectories() ([]*SoftLayer_Container_Network_Directory_Listing, error) {
	var returnValue []*SoftLayer_Container_Network_Directory_Listing
	return returnValue, nil
}

// GetVendorFtpAttributes - This method returns your login credentials to the public CDN
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) GetVendorFtpAttributes() (*SoftLayer_Container_Network_Authentication_Data, error) {
	var returnValue *SoftLayer_Container_Network_Authentication_Data
	return returnValue, nil
}

// LoadContent - Whether you are using Origin Pull or POP Pull, your content will be transferred and
// cached on CDN POP (node) on the initial request. If you wish to load your content to all CDN POPs,
// you may use this service to accomplish that. Please keep in mind, it will take about 10 to 15
// minutes to load content to all CDN POPs depending on the load. You can only specify 5 URLs at a
// time.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) LoadContent(objectUrls []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// ManageHttpCompression - Compression is used to reduce the bandwidth used to deliver an object. You
// can specify list of content types that needs to be compressed. If you omit the content type
// parameter, these values will be used by default: * text/plain * text/html * text/css *
// application/x-javascript * text/javascript Note that files larger than 1MB will never be served with
// compression regardless of whether their content-type is enabled for compression.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) ManageHttpCompression(enableFlag bool, mimeTypes []string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// PurgeCache - CDN's cache mechanism works similar to that of web browsers. When CDN pulls a file from
// your origin server or from your CDN FTP directory for the first time, it creates a cache file on
// itself. CDN re-uses cache files to save trips to the origin server and thus it speeds up delivering
// content to visitors. This method removes cached objects on every server in the CDN network. If you
// see a stale content or a file that send an incorrect header, purging cache will correct the issue.
// CDN will pull a fresh content from your origin server or your CDN This method takes an array of
// URLs. A URL must be exact as it is being requested by clients. An example URLs may look like this: *
// http:// .http.cdn.softlayer.net/mycdnname/some_file.txt If you created a that points to CDN host,
// use your URL instead. * http://image.mydomain.com/some_file.txt It takes approximately 3-5 minutes
// for the system to delete the requested object on every CDN server from submission .
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) PurgeCache(objectUrls []string) ([]*SoftLayer_Container_Network_ContentDelivery_PurgeService_Response, error) {
	var returnValue []*SoftLayer_Container_Network_ContentDelivery_PurgeService_Response
	return returnValue, nil
}

// RemoveAuthenticationDirectory - If you want to turn off the token authentication, use this method to
// remove a directory from the token authentication directory.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) RemoveAuthenticationDirectory(directory string, mediaType string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// RemoveFile - With this method you can remove a file or a directory on the CDN FTP server. If a
// source name ends with a slash this method assumes it is a directory. A source name must be an
// absolute path. It does not check to see if a file or directory exists before deletion. You can only
// remove files and directories that are in /media folder. Be sure to catch an exception for the detail
// on an error.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) RemoveFile(source string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetAuthenticationServiceEndpoint - CDN servers will invoke a Web Service method to validate a
// content authentication token. CDN uses the default Web Service provided by SoftLayer to validate a
// token. A customer can use their own implementation of the token authentication Web Service. A valid
// will look similar [https://manage.softlayer.com/CdnService/authenticationWsdlExample/wsdl this].
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) SetAuthenticationServiceEndpoint(webserviceEndpoint string, protocol string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// SetFtpPassword - With a CDN you can upload contents to CDN host server. Once you uploaded contents,
// your contents will be fetched by the CDN POP (Points of Presence) servers as needed. CDN supports
// three protocols: Flash streaming (rtmp), Window Media streaming (mms) and Once you log in to the CDN
// FTP server, you will see three directories under /media directory. You have to upload your contents
// to a proper directory to use the different services. Refer to
// [[SoftLayer_Network_ContentDelivery_Account|CDN Account]] service overview for details on the CDN
// FTP server. "gzip" is supported if you compress your content before uploading and you have to change
// its extension to ".gz". [SoftLayer_Network_ContentDelivery_Account::createOriginPullRule|Origin
// Pull] also supports "gzip" contents and you don't have to modify file extension with Origin Pull.
// Once uploaded, your contents should be available almost immediately to visitors. However, it may
// take about 30 minutes to propagate files to the entire CDN network after uploading. For more
// details, visit our [hhttp://knowledgelayer.softlayer.com/topic/cdn KnowledgeLayer] This method
// updates the password for your CDN FTP account on the ftp.cdnlayer.service.softlayer.com server. You
// must provide an alphanumeric value for a new password. - _ ! % # $ ^ & * characters are allowed
// beside an alphanumeric string.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) SetFtpPassword(newPassword string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UpdateNote - This method allows you to edit CDN account note. The maximum length for CDN account
// note is 30 characters.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) UpdateNote(note string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// UploadStream - With this method, you can upload binary data to the CDN FTP server. This method
// supports files up to 20 Mega Bytes. You need to use the CDN FTP (ftp.cdnlayer.service.softlayer.com)
// to upload files larger than 20 MB. This method takes [[SoftLayer_Container_Utility_File_Attachment]]
// a first parameter. A target name must be an absolute path and you can only upload a file to a
// directory that is in /media folder.
func (softlayer_network_contentdelivery_account *SoftLayer_Network_ContentDelivery_Account) UploadStream(source SoftLayer_Container_Utility_File_Attachment, target string) (bool, error) {
	var returnValue bool
	return returnValue, nil
}
