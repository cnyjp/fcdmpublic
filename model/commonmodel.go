package model

import (
	"encoding/json"
	"time"
)

/**
  The single config for plugin, applicaiton, and other need config place.
  In a group of config, use name to identity.
*/
type ConfigConfig struct {
	Name         string                `json:"name"`                   //Identity for config, which can not be empty.
	Type         string                `json:"type,omitempty"`         //value type, such as string,int,float,bool; the default is string
	Default      string                `json:"default,omitempty"`      //default value in the input controller. default is a empty string.
	InputType    string                `json:"inputType,omitempty"`    //the input controller for the config;such as text,password,select,radio,checkbox,autocomplete, and so on.
	Limits       map[string]string     `json:"limits,omitempty"`       //Optional, limit conditions, use key->name value->value to define the limits.The key can use such as maxvalue,minvalue,maxlength,minlength,pattern,email and so on.
	Options      map[string]string     `json:"options,omitempty"`      //Optional, if the input type is select,radio,checkbox, will need options; key->value,value->text(html option)
	Desc         string                `json:"desc,omitempty"`         //description for config, will display on the option input ui.
	I18n         map[string]ConfigI18n `json:"i18n,omitempty"`         //i18n infomation of the config; the key is i18n name, such as zh_CN,en_US.
	JobTypes     []JobType             `json:"jobTypes,omitempty"`     //which job type use the config. if set it to nil or empty, which will used on all job type
	MustJobTypes []JobType             `json:"mustJobTypes,omitempty"` //which job type must set the config, if set it to nil or empty, which will not must set on all job type.
}

/**
  The config i18n struct. Used in ConfigConfig.
*/
type ConfigI18n struct {
	Name    string            `json:"name"`              //display i18n for configconfig's Name
	Desc    string            `json:"desc"`              //display i18n for configconfig's Desc
	Options map[string]string `json:"options,omitempty"` //display i18n for configconfig's options.
}

//application的唯一性认定：主机id+name+secondlytype来唯一认定一个application
/**
  Application type struct.
  Applicaiton in a provider will be identity by secondlyType + Name
  Which identity by providerInstanceId + sedondlyType + Name at the serverside.
  When the provider use the application, need not config the id field.
*/
type Application struct {
	Id                 string            `json:"id"`                     //Identity for the appliation on server side, useless for provider.
	Name               string            `json:"name,omitempty"`         //provider identity a applicaiton by secondlytype + name
	ProviderInstanceId string            `json:"providerInstanceId"`     //the provider instance id for the applicaiton, useless for provider.
	ProviderId         string            `json:"providerId"`             //provider id, useless for provider.
	ProviderName       string            `json:"providerName,omitempty"` //provider name, useless for provider
	DisplayName        string            `json:"displayName,omitempty"`  //The display name for the applicaiton, if the provider want the application has a friendly name to display
	SecondlyType       string            `json:"secondlyType,omitempty"` //provider identity a applicaiton by secondlytype + name
	Size               int64             `json:"size,omitempty"`         //application size, the server side will count it by the volumes info.
	Volumes            []Volume          `json:"volumes,omitempty"`      //All application use space, provider must set all volume info correct.
	Options            map[string]string `json:"options,omitempty"`      //applicaiton options, if the provider nned set the application's default value, can use it.
	//if the common information can not include the application's all info, provider can set the extend info in this field.
	//which will be set a env FCDM_EV_APP_EXTENSION when provider execute any command for a application.
	Extensions string `json:"extensions,omitempty"`

	//The display fields info, match the listconfig in secondlytype.
	DisplayFields map[string]string `json:"displayFields,omitempty"`

	StageAutoExpand string `json:"stageAutoExpand,omitempty"` //auto expand
	ExpandPercent   int    `json:"expandPercent,omitempty"`   //percent in auto expand

	Haslog bool `json:"haslog"` //dose the applicaiton has a log.

	Remark string `json:"remark,omitempty"` //remark

	//the configs for this applicaiton.
	//if the provider does set configs, this will use the config in plugin config.
	Configs []ConfigConfig `json:"configs,omitempty"`

	HostId   string `json:"hostId,omitempty"`
	HostType string `json:"hostType,omitempty"`

	Avaiable bool `json:"avaiable"` //is avaiable

	ParentId string        `json:"parentId,omitempty"`
	Childs   []Application `json:"childs,omitempty"`
}

/**
  The volume is the storage space for a applicaiton.
  In an applicaiton, the volume name is the identity.
  In provider, need not config the Id field for the volume.
*/
type Volume struct {
	Id            string    `json:"id"`            //Identity on the server side, useless on provider.
	TargetId      string    `json:"targetId"`      //Target identity on the server side, useless on provider.
	ApplicationId string    `json:"applicationId"` //Applicaiton identity on the server side, useless on provider.
	Name          string    `json:"name"`          //Name for volume, in a application, it is the identity.
	DisplayName   string    `json:"displayName"`   //When the provider need the volume need a friendly name to display, can use it.
	StageType     StageType `json:"stageType"`     //StageType
	Size          int64     `json:"size"`
	BlockSize     int       `json:"blockSize,omitempty"`
	FsType        string    `json:"fsType"`             //filesystem type, such ext4,xfs,ntfs and so on.
	FsDetail      string    `json:"fsDetail,omitempty"` //filesystem details, provider can use it to set the filesystem, such blocksize, privilege

	FsPath string `json:"fsPath,omitempty"` //the volume display in the operate system.

	//Stage properties
	StageAutoExpand string `json:"stageAutoExpand,omitempty"` //auto expand-- true,false|yes,no|inherited

	StageSize int64 `json:"stageSize,omitempty"` //The stagesize provider want for the volume.
}

/**
  The plugin config struct.
*/
type PluginConfig struct {
	PEName   string            `json:"pename"`            //PE file name，
	SignInfo SignInfo          `json:"signInfo"`          //signinfo for the pefile.
	Options  map[string]string `json:"options,omitempty"` //The default options for the plugin.

	//if the options is not nil, this will be the option's i18n
	//key for i18n zone, such as zh_CN,en_US, value for the options
	OptionsI18n map[string]ConfigI18n `json:"optionsI18n,omitempty"`

	//configs for the plugin. if the provider does not set the application's configs, it will be the applicaion's config.
	Configs []ConfigConfig `json:"configs,omitempty"`

	//search condition for secondlytype: the key for secondly_type
	//value for a list config.
	AppSearchConditions map[string][]ConfigConfig `json:"appSearchConditions"`

	//the list type supported in the provider.
	ListAppTypes []string `json:"listAppTypes"`
}

/**
  The sign info struct for plugin command file.
*/
type SignInfo struct {
	Version string `json:"version"`
	Sign    string `json:"sign"`
}

/**
  The plugin info returned struct by the command CMD_PLUGIN_INFO
*/
type PluginCmdInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

//If the  response used tree node mode to return data, use the tree node struct.
type TreeNode struct {
	Id                    string `json:"id"`
	Name                  string `json:"name"`
	ParentId              string `json:"parentId"`
	IsLeaf                bool   `json:"isLeaf"`
	TargetApplicationName string `json:"targetApplicationName"` //The Applicaiton Name which the tree node point.
}

/** custom time format, use yyyy-MM-dd HH:mm:ss to transform the datetime.  */
type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	if nil != err {
		//use the normal format to transform
		tt := time.Time{}
		err = json.Unmarshal(data, &tt)
		*t = Time(tt)
	} else {
		*t = Time(now)
	}
	return nil
}

//check the time is init value.
func (t *Time) IsZero() bool {
	tt := time.Time(*t)
	return tt.IsZero()
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}
func Now() Time {
	return Time(time.Now())
}
func (t *Time) Format(format string) string {
	return time.Time(*t).Format(format)
}

type SecondlyType struct {
	Name string                `json:"name"` //
	I18n map[string]ConfigI18n `json:"i18n"` //i18n for name

	//the list item for the secondlytype
	ListConfigs []ConfigConfig `json:"listConfigs"`
	//the search item for the secondlytype
	SearchConfigs []ConfigConfig `json:"searchConfigs"`
}

type ScriptSet struct {
	Step    JobStep  `json:"step"`           //script step,use the JobStep
	Name    string   `json:"name"`           //script name, the filename for the script.
	Timeout int      `json:"timeout"`        //script execute timeout, zero for none timeout.
	Args    []string `json:"args,omitempty"` //script execute args.
}
