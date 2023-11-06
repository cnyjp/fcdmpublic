package model

import (
	"encoding/json"
	"time"
)

//ConfigConfig
/*
The single config for plugin, application, and other need config place.
In a group of config, use name to identity.
*/
type ConfigConfig struct {
	Name            string                `json:"name"`                      //Identity for config, which can not be empty.
	Type            string                `json:"type,omitempty"`            //value type, such as string,int,float,bool; the default is string
	Default         string                `json:"default,omitempty"`         //default value in the input controller. default is an empty string.
	InputType       string                `json:"inputType,omitempty"`       //the input controller for the config;such as text,password,select,radio,checkbox,autocomplete, and so on.
	Limits          map[string]string     `json:"limits,omitempty"`          //Optional, limit conditions, use key->name value->value to define the limits.The key can use such as maxvalue,minvalue,maxlength,minlength,pattern,email and so on.
	Options         map[string]string     `json:"options,omitempty"`         //Optional, if the input type is select,radio,checkbox, will need options; key->value,value->text(html option)
	Columns         []ConfigColumn        `json:"columns,omitempty"`         //Optional, if the input type is table, need this config.
	Desc            string                `json:"desc,omitempty"`            //description for config, will display on the option input ui.
	ValidateMessage string                `json:"validateMessage,omitempty"` //Optional, when the config validate error, show this message.
	I18n            map[string]ConfigI18n `json:"i18n,omitempty"`            //i18n information of the config; the key is i18n name, such as zh_CN,en_US.
	JobTypes        []JobType             `json:"jobTypes,omitempty"`        //which job type use the config. if set it to nil or empty, which will used on all job type
	MustJobTypes    []JobType             `json:"mustJobTypes,omitempty"`    //which job type must set the config, if set it to nil or empty, which will not must set on all job type.
}

type ConfigColumn struct {
	Name            string                `json:"name"`
	Type            string                `json:"type,omitempty"`            //column type, such as type of ConfigConfig, string as default.
	InputType       string                `json:"inputType,omitempty"`       //input type for column, such as inputType of ConfigConfig, text is default.
	Default         string                `json:"default,omitempty"`         //same as ConfigConfig Default
	Limits          map[string]string     `json:"limits,omitempty"`          //Optional, limit conditions, use key->name value->value to define the limits.The key can use such as maxvalue,minvalue,maxlength,minlength,pattern,email and so on.
	Options         map[string]string     `json:"options,omitempty"`         //if the type is select, radio, checkbox group or other need multi key-value pairs, use the options to set the select value.
	Desc            string                `json:"desc,omitempty"`            //description for column, sometimes the desc maybe display on the at the column tips
	ValidateMessage string                `json:"validateMessage,omitempty"` //save as ConfigConfig ValidateMessage
	I18n            map[string]ConfigI18n `json:"i18n,omitempty"`            //i18n
}

//ConfigI18n
/*
The config i18n struct. Used in ConfigConfig.
*/
type ConfigI18n struct {
	Name            string            `json:"name"`                      //display i18n for ConfigConfig's Name
	Desc            string            `json:"desc"`                      //display i18n for ConfigConfig's Desc
	ValidateMessage string            `json:"validateMessage,omitempty"` //display i18n for ConfigConfig's ValidateMessage
	Options         map[string]string `json:"options,omitempty"`         //display i18n for ConfigConfig's options.
}

//Application
/*
  Application type struct.
  application in a provider will be identity by secondlyType + Name
  Which identity by providerInstanceId + secondlyType + Name at the serverside.
  When the provider use the application, need not config the id field.
*/
/*
 The fcdm server identify an application with hostid(providerInstanceId) + application.Name + application.SecondlyType.
*/
type Application struct {
	Id                 string            `json:"id"`                     //Identity for the application on server side, useless for provider.
	Name               string            `json:"name,omitempty"`         //provider identity an application by secondlyType + name
	ProviderInstanceId string            `json:"providerInstanceId"`     //the provider instance id for the application, useless for provider.
	ProviderId         string            `json:"providerId"`             //provider id, useless for provider.
	ProviderName       string            `json:"providerName,omitempty"` //provider name, useless for provider
	DisplayName        string            `json:"displayName,omitempty"`  //The display name for the application, if the provider want the application has a friendly name to display
	SecondlyType       string            `json:"secondlyType,omitempty"` //provider identity an application by secondlyType + name
	Size               int64             `json:"size,omitempty"`         //application size, the server side will count it by the info of volumes.
	Volumes            []Volume          `json:"volumes,omitempty"`      //All application use space, provider must set all volume info correct.
	Options            map[string]string `json:"options,omitempty"`      //application options, if the provider need set the application's default value, can use it.
	//if the common information can not include the application's all info, provider can set the extension info in this field.
	//which will be set an env FCDM_EV_APP_EXTENSION when provider execute any command for an application.
	Extensions string `json:"extensions,omitempty"`

	//The display fields info, match the listConfig in secondlyType.
	DisplayFields map[string]string `json:"displayFields,omitempty"`

	StageAutoExpand string `json:"stageAutoExpand,omitempty"` //auto expand
	ExpandPercent   int    `json:"expandPercent,omitempty"`   //percent in auto expand

	Haslog bool `json:"haslog"` //dose the application has a log.

	Remark string `json:"remark,omitempty"` //remark

	//the configs for this application.
	//if the provider does set configs, this will use the config in plugin config.
	Configs []ConfigConfig `json:"configs,omitempty"`

	HostId   string `json:"hostId,omitempty"`
	HostType string `json:"hostType,omitempty"`

	Available bool `json:"available"` //is available

	ParentId string        `json:"parentId,omitempty"`
	Childs   []Application `json:"childs,omitempty"`
	Custom   bool          `json:"custom"` //is the application is created by tenant.
}

//Volume
/*
The volume is the storage space for an application.
In an application, the volume name is the identity.
In provider, need not config id for the volume.
*/
type Volume struct {
	Id            string    `json:"id,omitempty"`            //Identity on the server side, useless on provider.
	TargetId      string    `json:"targetId,omitempty"`      //Target identity on the server side, useless on provider.
	ApplicationId string    `json:"applicationId,omitempty"` //application identity on the server side, useless on provider.
	Name          string    `json:"name"`                    //Name for volume, in an application, it is the identity.
	DisplayName   string    `json:"displayName,omitempty"`   //When the provider need the volume need a friendly name to display, can use it.
	StageType     StageType `json:"stageType"`               //StageType
	Size          int64     `json:"size"`
	BlockSize     int       `json:"blockSize,omitempty"`
	FsType        string    `json:"fsType,omitempty"`   //filesystem type, such ext4,xfs,ntfs and so on.
	FsDetail      string    `json:"fsDetail,omitempty"` //filesystem details, provider can use it to set the filesystem, such blocksize, privilege

	Identity string `json:"identity,omitempty"` //Identity on provider side.

	FsPath string `json:"fsPath,omitempty"` //the volume display in the operate system.

	//Stage properties
	StageAutoExpand string `json:"stageAutoExpand,omitempty"` //auto expand-- true,false|yes,no|inherited

	StageSize int64 `json:"stageSize,omitempty"` //The stagesize provider want for the volume.
}

//PluginConfig
/*
The plugin config struct.
*/
type PluginConfig struct {
	PEName      string            `json:"pename"`            //PE file name，
	SignInfo    SignInfo          `json:"signInfo"`          //signInfo for the peFile.
	AllowCustom bool              `json:"allowCustom"`       //in provider config, this will allow the provider create a custom application.
	Options     map[string]string `json:"options,omitempty"` //The default options for the plugin.

	//if the options is not nil, this will be the option's i18n
	//key for i18n zone, such as zh_CN,en_US, value for the options
	OptionsI18n map[string]ConfigI18n `json:"optionsI18n,omitempty"`

	//configs for the plugin. if the provider does not set the application's configs, it will be the application's config.
	Configs []ConfigConfig `json:"configs,omitempty"`

	ConfigIcon ConfigIcon `json:"configIcon,omitempty"`

	//if the plugin is a host provider, this is used to config all SecondlyType Config.
	SecondlyTypes []SecondlyType `json:"secondlyTypes,omitempty"`

	//search condition for secondlyType: the key is a secondlyType
	//value for a list config.
	AppSearchConditions map[string][]ConfigConfig `json:"appSearchConditions"`

	//the list type supported in the provider. use 	LIST_APP_TYPE_ALL,LIST_APP_TYPE_PAGE,LIST_APP_TYPE_MORE,LIST_APP_TYPE_TREE
	ListAppTypes []ListAppType `json:"listAppTypes"`
}

type ConfigIcon struct {
	Icon32  string `json:"icon32,omitempty"`  //the icon 32*32 encoded base64, use png gif image
	Icon256 string `json:"icon256,omitempty"` //the icon 256*256 encoded base64 use png gif image
}

//SignInfo
/*
The sign info struct for plugin command file.
*/
type SignInfo struct {
	Version string `json:"version"`
	Sign    string `json:"sign"`
}

//PluginCmdInfo
/*
	The plugin info returned struct by the command CMD_PLUGIN_INFO
*/
type PluginCmdInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// TreeNode
// If the  response used tree node mode to return data, use the tree node struct.
type TreeNode struct {
	Identity       string         `json:"identity"`
	Name           string         `json:"name"` //the name of the
	ParentIdentity string         `json:"parentIdentity"`
	IsLeaf         bool           `json:"isLeaf"`                //if the node is a leaf node, false means it is a branch, not a leaf.
	DataFilters    []ConfigConfig `json:"dataFilters,omitempty"` //The data filter rules. just the provider's appSearchCondition
}

// TreeNodeList
// the node list struct
type TreeNodeList struct {
	ParentIdentity string     `json:"parentIdentity"`
	MoreId         string     `json:"moreId,omitempty"`
	Nodes          []TreeNode `json:"nodes"`                //the node list
	HasMore        bool       `json:"hasMore"`              //does the tree node list has more data
	NextMoreId     string     `json:"nextMoreId,omitempty"` //for more data, need use the id to get next data.
}

// TreeNodeData
// The Data for tree node
type TreeNodeData struct {
	NodeId         string              `json:"parentId"`
	Filters        map[string]string   `json:"filters"`          //The filter rules for data list.
	MoreDataId     string              `json:"moreId,omitempty"` //When the data is a response for a more data request, this is the request moreDataId.
	Columns        []ConfigColumn      `json:"columns"`
	Values         []map[string]string `json:"values"`                   //value list, the map key is the column name
	HasMore        bool                `json:"hasMore"`                  //does the data has more elements.
	NextMoreDataId string              `json:"nextMoreDataId,omitempty"` //if the data has more elements, this is the request moreDataId for next request.
}

// Time
/* custom time format, use yyyy-MM-dd HH:mm:ss to transform the datetime.  */
type Time time.Time

const (
	timeFormat = "2006-01-02 15:04:05"
)

// ToLocalTime
/*
	transfer the time to a *time.Time format data with location.
*/
func (t *Time) ToLocalTime(location *time.Location) *time.Time {
	if nil == t {
		return nil
	}
	tt, err := time.ParseInLocation(timeFormat, t.String(), location)

	if nil != err {
		return nil
	}

	return &tt
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
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

// check the time is init value.
func (t *Time) IsZero() bool {
	tt := time.Time(*t)
	return tt.IsZero()
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
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

	//the list item for the secondlyType
	ListConfigs []ConfigConfig `json:"listConfigs,omitempty"`
	//the search item for the secondlyType
	SearchConfigs []ConfigConfig `json:"searchConfigs,omitempty"`
	//the application configs for this SecondlyType
	ApplicationConfigs []ConfigConfig `json:"applicationConfigs,omitempty"`

	ConfigIcon ConfigIcon `json:"configIcon,omitempty"`
}

type ScriptSet struct {
	Step    JobStep  `json:"step"`           //script step,use the JobStep
	Name    string   `json:"name"`           //script name, the filename for the script.
	Timeout int      `json:"timeout"`        //script execute timeout, zero for none timeout.
	Args    []string `json:"args,omitempty"` //script execute args.
}
