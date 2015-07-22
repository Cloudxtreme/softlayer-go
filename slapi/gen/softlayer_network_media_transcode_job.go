package sl

// DO NOT EDIT. THIS FILE WAS AUTOMATICALLY GENERATED

import (
	time "time"
)

// SoftLayer_Network_Media_Transcode_Job - The SoftLayer_Network_Media_Transcode_Job contains
// information regarding a transcode job such as input file, output format, user id and so on.
type SoftLayer_Network_Media_Transcode_Job struct {

	// AutoDeleteDuration - The auto-deletion duration in seconds. This value determines how long the input
	// file will be kept on the storage.
	AutoDeleteDuration int `json:"autoDeleteDuration"`

	// ByteIn - no documentation
	ByteIn int `json:"byteIn"`

	// CreateDate - no documentation
	CreateDate *time.Time `json:"createDate"`

	// History - <nil>
	History []*SoftLayer_Network_Media_Transcode_Job_History `json:"history"`

	// HistoryCount - no documentation
	HistoryCount uint64 `json:"historyCount"`

	// Id - no documentation
	Id int `json:"id"`

	// InputFile - no documentation
	InputFile string `json:"inputFile"`

	// ModifyDate - no documentation
	ModifyDate *time.Time `json:"modifyDate"`

	// Name - no documentation
	Name string `json:"name"`

	// OutputFile - no documentation
	OutputFile string `json:"outputFile"`

	// TranscodeAccount - no documentation
	TranscodeAccount *SoftLayer_Network_Media_Transcode_Account `json:"transcodeAccount"`

	// TranscodeAccountId - no documentation
	TranscodeAccountId int `json:"transcodeAccountId"`

	// TranscodeJobGuid - no documentation
	TranscodeJobGuid string `json:"transcodeJobGuid"`

	// TranscodePresetGuid - no documentation
	TranscodePresetGuid string `json:"transcodePresetGuid"`

	// TranscodePresetName - no documentation
	TranscodePresetName string `json:"transcodePresetName"`

	// TranscodeStatus - no documentation
	TranscodeStatus *SoftLayer_Network_Media_Transcode_Job_Status `json:"transcodeStatus"`

	// TranscodeStatusId - no documentation
	TranscodeStatusId int `json:"transcodeStatusId"`

	// TranscodeStatusName - no documentation
	TranscodeStatusName string `json:"transcodeStatusName"`

	// User - no documentation
	User *SoftLayer_User_Customer `json:"user"`

	// UserId - The internal identifier of the user who created a transcode job
	UserId int `json:"userId"`

	// Watermark - no documentation
	Watermark *SoftLayer_Container_Network_Media_Transcode_Job_Watermark `json:"watermark"`
}

// CreateObject - With this method, you can create a transcode job. The very first step of creating a
// transcode job is to upload your media files to the /in directory on your Transcode FTP space. Then,
// you have to pass a [[SoftLayer_Network_Media_Transcode_Job|Transcode job]] object as a parameter for
// this method. There are 4 required properties of SoftLayer_Network_Media_Transcode_Job object:
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
func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job) CreateObject(templateObject SoftLayer_Network_Media_Transcode_Job) (*SoftLayer_Network_Media_Transcode_Job, error) {
	var returnValue *SoftLayer_Network_Media_Transcode_Job
	return returnValue, nil
}

// GetObject - <nil>
func (softlayer_network_media_transcode_job *SoftLayer_Network_Media_Transcode_Job) GetObject() (*SoftLayer_Network_Media_Transcode_Job, error) {
	var returnValue *SoftLayer_Network_Media_Transcode_Job
	return returnValue, nil
}
