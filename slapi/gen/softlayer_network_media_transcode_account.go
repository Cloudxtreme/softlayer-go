package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Account - The SoftLayer_Network_Media_Transcode_Account contains
// information regarding a transcode account.
type SoftLayer_Network_Media_Transcode_Account struct {

	// Account - no documentation
	Account *SoftLayer_Account `json:"account"`

	// AccountId - no documentation
	AccountId int `json:"accountId"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// Id - no documentation
	Id int `json:"id"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// TranscodeJobCount - no documentation
	TranscodeJobCount uint64 `json:"transcodeJobCount"`

	// TranscodeJobs - no documentation
	TranscodeJobs []*SoftLayer_Network_Media_Transcode_Job `json:"transcodeJobs"`
}

// CreateTranscodeAccount - With this method, you can create a transcode account. Individual SoftLayer
// account can have a single Transcode account. You have to pass your SoftLayer account id as a
// parameter.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) CreateTranscodeAccount() (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// CreateTranscodeJob - '''Note'''. This method is obsolete. Please use the
// [[SoftLayer_Network_Media_Transcode_Job::createObject|createObject]] method on
// SoftLayer_Network_Media_Transcode_Job object instead.
// SoftLayer_Network_Media_Transcode_Job::createObject returns an object of a newly created Transcode
// Job. With this method, you can create a transcode job. The very first step of creating a transcode
// job is to upload your media files to the /in directory on your Transcode FTP space. Then, you have
// to pass a [[SoftLayer_Network_Media_Transcode_Job|Transcode job]] object as a parameter for this
// method. There are 4 required properties of SoftLayer_Network_Media_Transcode_Job object:
// transcodePresetName, transcodePresetGuid, inputFile, and outputFile. A transcode preset is a
// configuration that defines a certain media output. You can retrieve all the supported presets with
// the [[SoftLayer_Network_Media_Transcode_Account::getPresets|getPresets]] method. You can also use
// [[SoftLayer_Network_Media_Transcode_Account::getPresetDetail|getPresetDetail]] method to get more
// information on a preset. Use these two methods to determine appropriate values for
// "transcodePresetName" and "transcodePresetGuid" properties. For an "inputFile", you must specify a
// file that exists in the /in directory of your Transcode FTP space. An "outputFile" name will be used
// by the Transcode server for naming a transcoded file. An output file name must be in /out directory.
// If your outputFile name already exists in the /out directory, the Transcode server will append a
// file name with _n (an underscore and the total number of files with the identical name plus 1). The
// "name" property is optional and it can help you keep track of transcode jobs easily.
// "autoDeleteDuration" is another optional property that you can specify. It determines how soon your
// input file will be deleted. If autoDeleteDuration is set to zero, your input file will be removed
// immediately after the last transcode job running on it is completed. A value for autoDeleteDuration
// property is in seconds and the maximum value is 259200 which is 3 days. An example
// SoftLayer_Network_Media_Transcode_Job parameter looks like this: * name: My transcoding *
// transcodePresetName: F4V 896kbps 640x352 16x9 29.97fps * transcodePresetGuid:
// {87E01268-C3E3-4A85-9701-052C9AC42BD4} * inputFile: /in/my_birthday.wmv * outputFile:
// /out/my_birthday_flash Notice that an output file does not have a file extension. The Transcode
// server will append a file extension based on an output format. A newly created transcode job will be
// in "Pending" status and it will be added to the Transcoding queue. You will receive a notification
// email whenever there is a status change on your transcode job. For example, the Transcode server
// starts to process your transcode job, you will be notified via an email. You can add up to 3 pending
// jobs at a time. Transcode jobs with any other status such as "Complete" or "Error" will not be
// counted toward your pending jobs. Once a job is complete, the Transcode server will place the output
// file into the /out directory along with a notification email. The files in the /out directory will
// be removed 3 days after they were created. You will need to use an FTP client to download transcoded
// files.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) CreateTranscodeJob(newJob SoftLayer_Network_Media_Transcode_Job) (bool, error) {
	var returnValue bool
	return returnValue, nil
}

// GetDirectoryInformation - This method returns a collection of
// SoftLayer_Container_Network_Ftp_Directory objects. You can retrieve directory information for /in
// and /out directories. A [[SoftLayer_Container_Network_Directory_Listing|Directory Listing]] object
// contains a type (indicating whether it is a file or a directory), name and file count if it is a
// directory.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetDirectoryInformation(directoryName string, extensionFilter string) ([]*SoftLayer_Container_Network_Directory_Listing, error) {
	var returnValue []*SoftLayer_Container_Network_Directory_Listing
	return returnValue, nil
}

// GetFileDetail - This method returns detailed information of a media file that resides in the
// Transcode FTP server. A [[SoftLayer_Container_Network_Media_Information|media information]] object
// contains media details such as file size, media format, frame rate, aspect ratio and so on. This
// information is merely for reference purposes. You should not rely on this data. Our library grabs
// small pieces of data from a media file to gather media details. This information may not be
// available for some files.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetFileDetail(source string) (*SoftLayer_Container_Network_Media_Information, error) {
	var returnValue *SoftLayer_Container_Network_Media_Information
	return returnValue, nil
}

// GetFtpAttributes - This method returns your Transcode FTP login credentials to the
// transcode.service.softlayer.com server. The Transcode FTP server is available via the SoftLayer
// private network. There is no API method that you can upload a file to Transcode server so you need
// to use an FTP client. You will have /in and /out directories on the Transcode FTP server. You will
// have read-write privileges for /in directory and read-only privilege for /out directory. All the
// files in both /in and /out directories will be deleted after 72 hours from the creation date.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetFtpAttributes() (*SoftLayer_Container_Network_Authentication_Data, error) {
	var returnValue *SoftLayer_Container_Network_Authentication_Data
	return returnValue, nil
}

// GetObject - getObject method retrieves the SoftLayer_Network_Media_Transcode_Account object whose ID
// number corresponds to the ID number of the initial parameter passed to the
// SoftLayer_Network_Media_Transcode_Account service. You can only retrieve a Transcode account
// assigned to your SoftLayer customer account.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetObject() (*SoftLayer_Network_Media_Transcode_Account, error) {
	var returnValue *SoftLayer_Network_Media_Transcode_Account
	return returnValue, nil
}

// GetPresetDetail - This method returns an array of
// [[SoftLayer_Container_Network_Media_Transcode_Preset_Element|preset element]] objects. Each preset
// has its own collection of preset elements such as encoder, frame rate, aspect ratio and so on. Each
// element object has a default value for itself and an array of
// [[SoftLayer_Container_Network_Media_Transcode_Preset_Element_Option|element option]] objects. For
// example, "Frame Rate" element for "Windows Media 9 - Download - 1 Mbps - - Constrained preset has 19
// element options. 15.0 frame rate is selected by default. Currently, you are not able to change the
// default value. Customizing these values may be possible in the future.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetPresetDetail(guid string) ([]*SoftLayer_Container_Network_Media_Transcode_Preset_Element, error) {
	var returnValue []*SoftLayer_Container_Network_Media_Transcode_Preset_Element
	return returnValue, nil
}

// GetPresets - A transcode preset is a configuration that defines a certain media output. This method
// returns an array of transcoding preset objects supported by SoftLayer's Transcode server. Each
// [[SoftLayer_Container_Network_Media_Transcode_Preset|preset object]] contains a property. You will
// need a string when you create a new transcode job.
func (softlayer_network_media_transcode_account *SoftLayer_Network_Media_Transcode_Account) GetPresets() ([]*SoftLayer_Container_Network_Media_Transcode_Preset, error) {
	var returnValue []*SoftLayer_Container_Network_Media_Transcode_Preset
	return returnValue, nil
}
