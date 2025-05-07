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
	Name               string                `json:"name"`                         //Identity for config, which can not be empty.
	Type               string                `json:"type,omitempty"`               //Optional, value type, such as string,int,float,bool; the default is string
	Default            string                `json:"default,omitempty"`            //Optional, default value in the input controller. default is an empty string.
	InputType          string                `json:"inputType,omitempty"`          //Optional, the input controller for the config;such as text,password,select,radio,checkbox,autocomplete, and so on.
	Limits             map[string]string     `json:"limits,omitempty"`             //Optional, limit conditions, use key->name value->value to define the limits.The key can use such as maxvalue,minvalue,maxlength,minlength,pattern,email and so on.
	Options            map[string]string     `json:"options,omitempty"`            //Optional, if the input type is select,radio,checkbox, will need options; key->value,value->text(html option)
	Columns            []ConfigColumn        `json:"columns,omitempty"`            //Optional, if the input type is table, need this config.
	Desc               string                `json:"desc,omitempty"`               //description for config, will display on the option input ui.
	ValidateMessage    string                `json:"validateMessage,omitempty"`    //Optional, when the config validate error, show this message.
	LimitsErrorMessage map[string]string     `json:"limitsErrorMessage,omitempty"` //Optional, custom the limits validate error message.
	I18n               map[string]ConfigI18n `json:"i18n,omitempty"`               //Optional, i18n information of the config; the key is i18n name, such as zh_CN,en_US.
	JobTypes           []JobType             `json:"jobTypes,omitempty"`           //Optional, which job type use the config. if set it to nil or empty, which will used on all job type
	MustJobTypes       []JobType             `json:"mustJobTypes,omitempty"`       //Optional, which job type must set the config, if set it to nil or empty, which will not must set on all job type.
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
	Name               string            `json:"name"`                         //display i18n for ConfigConfig's Name
	Desc               string            `json:"desc"`                         //display i18n for ConfigConfig's Desc
	ValidateMessage    string            `json:"validateMessage,omitempty"`    //display i18n for ConfigConfig's ValidateMessage
	Options            map[string]string `json:"options,omitempty"`            //display i18n for ConfigConfig's options.
	LimitsErrorMessage map[string]string `json:"limitsErrorMessage,omitempty"` //display i18n for ConfigConfig's Limits error message
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

	StageSize int64 `json:"stageSize,omitempty"` //The stageSize provider want for the volume.

	CopyVersion int `json:"copyVersion,omitempty"` //same as image

	ImageId             string `json:"imageId,omitempty"`             //image id for volume
	StoragePoolId       string `json:"storagePoolId,omitempty"`       //pool id for volume
	StoragePoolEngineId string `json:"storagePoolEngineId,omitempty"` //engine if for pool
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
	if nil == location {
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

type StorageInterface struct {
	Id                 string            `json:"id,omitempty"`                 //need not set when discover
	HostId             string            `json:"hostId,omitempty"`             //need not set when discover
	ProviderInstanceId string            `json:"providerInstanceId,omitempty"` //need not set when discover
	Protocol           StorageProtocol   `json:"protocol,omitempty"`           //storage protocol, iscsi,iser,fc,ib,nfs and so on.
	AdapterType        AdapterType       `json:"adapterType,omitempty"`        //target or initiator
	Address            string            `json:"address,omitempty"`            //target address or initiator name.
	TargetAddress      string            `json:"targetAddress"`                //access address for target interface
	Options            map[string]string `json:"options"`                      //for some protocol, need options. such as：iscsi use chap,usename,password
	Configs            []ConfigConfig    `json:"configs"`                      //configs for interface，
}

/*
StorageEngine
*/
type StorageEngine struct {
	Id             string         `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	EngineTypeName string         `json:"engineTypeName,omitempty"` //the name of the engine type. if blank, means no storage engine type.
	PoolConfigs    []ConfigConfig `json:"poolConfigs,omitempty"`    //the configs for storage create pool. if the configs is nil, then need not input any param when create pool.

	StageTypes    []StageType `json:"stageTypes,omitempty"`    //which stageType the engine support
	PoolNeedSpace bool        `json:"poolNeedSpace,omitempty"` //if the storage engine need space when create or modify a storage pool.

	Level int `json:"level"` //Level for storage engine, 0 - full, 1 - archive
}

type StorageEngineType struct {
	Name            string         `json:"name,omitempty"`          //name for storage engine type
	EngineConfigs   []ConfigConfig `json:"engineConfigs,omitempty"` //configs for storage engine
	PoolConfigs     []ConfigConfig `json:"poolConfigs,omitempty"`   //configs for storage pool
	StageTypes      []StageType    `json:"stageTypes,omitempty"`    //stage types supported, if the engine type level is 0, need this.
	PoolNeedSpace   bool           `json:"poolNeedSpace"`           //when create/modify storage pool, is need point the space.
	CanCreateEngine bool           `json:"canCreateEngine"`         //if the storage engine type can create a new storage engine.
	Level           int            `json:"level"`                   //storage engine level -- 0 full, 1 archive

}

type StoragePool struct {
	Id       string            `json:"id,omitempty"`
	Name     string            `json:"name,omitempty"`
	EngineId string            `json:"engineId,omitempty"`
	Options  map[string]string `json:"options,omitempty"`

	Size         int64          `json:"size"`
	Used         int64          `json:"used"`
	Free         int64          `json:"free"`
	UsedSpaces   []StorageSpace `json:"usedSpaces,omitempty"`   //the space pool using
	AddSpaces    []StorageSpace `json:"addSpaces,omitempty"`    //the space will add to pool, used on create or modify
	RemoveSpaces []StorageSpace `json:"removeSpaces,omitempty"` //the space need remove from used space, used on modify
	CacheSpaces  []StorageSpace `json:"cacheSpaces,omitempty"`  //the space use for cache, when the pool support cache
	LogSpaces    []StorageSpace `json:"logSpaces,omitempty"`    //the log space, when the pool support log

	ValidSpaces []StorageSpace `json:"validSpaces"` //the spaces can add to the pool
}

type StorageEngineStatus struct {
	MaxSize      int64             `json:"maxSize"`  //max size of storage engine
	UsedSize     int64             `json:"usedSize"` //used size of storage engine
	Options      map[string]string `json:"options,omitempty"`
	AllSpaces    []StorageSpace    `json:"allSpaces"`    //all spaces
	ValidSpaces  []StorageSpace    `json:"validSpaces"`  //spaces can add
	UsedSpaces   []StorageSpace    `json:"usedSpaces"`   //used spaces
	StoragePools []StoragePool     `json:"storagePools"` //已建立的存储池
	StatusTime   time.Time         `json:"statusTime"`   //time for status get
}

type StorageSpace struct {
	Name      string            `json:"name"`              //space name
	Identity  string            `json:"identity"`          //space identity
	Path      string            `json:"path"`              //space path
	Username  string            `json:"username"`          //usename
	Password  string            `json:"password"`          //password
	Size      int64             `json:"size"`              //space total size
	BlockSize int32             `json:"blockSize"`         //space block size, optional
	Options   map[string]string `json:"options,omitempty"` //space options
	IsMounted bool              `json:"-"`                 //if the space is mounted to another pool.
}

// VolumeStorageSummary
// the summary for a volume on storage
type VolumeStorageSummary struct {
	Status int    `json:"status"`
	Size   int    `json:"size"` //the volume size defined, maybe there is no size defined on the volume.
	Used   int    `json:"used"` //the volume used space -- real used size on the storage device.
	Data   string `json:"data,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// CoupleVolumeMetadata
// @Description: 单一一个卷在copy任务中的metadata，包含了Source和Target
type CoupleVolumeMetadata struct {
	Source VolumeMetadata `json:"source,omitempty"`
	Target VolumeMetadata `json:"target,omitempty"`
}

// 用于描述 复制 volume 的MetaData的数据结构
// 异构引擎如何实现差异化复制？
// 如果是一个归档类型的存储引擎，其自身并不实现存储的格式化，而是直接保留源存储引擎的数据格式，保留和记录源数据引擎的所有metadata，并只是简单进行的数据归档保存
// 那么，理论上，它是可以实现所有类型存储引擎的差异化复制的。那么，此时提取其metadata，其返回的metadata的数据格式就与原始的引擎完全一致
// 那么，此时需要对应的metadata提交给存储引擎，由存储引擎进行相应的处理（保存或丢弃）
// 问题在于，源存储引擎如何判断应该使用自身的数据处理方式进行数据处理呢？
// 数据引擎判断能否自己处理的方式为：返回的targetMetaData符合自身的数据格式要求，自身解析下能够实现差异处理
// 但是归档引擎并不具备该处理能力，所以归档引擎应该返回一个特殊的标记，表明自己是一个归档引擎。
// 同时，归档引擎并不能准确判断不同的存储卷之间的依赖关系，因而无法在进一步的复制中体现卷的依赖关系
// 如果将一个增量卷向一个全存储引擎进行复制，那么想要全存储引擎能够正常识别和运行该镜像，需要发送一串卷的数据
// 那么copyproxy中发送的数据结构中就必须要能够表现这些卷的前后顺序以及每个包所属的卷信息
type VolumeMetadata struct {
	Id            string `json:"id"`                 //使用ID和poolId，引擎应该很容易找到对应的volume
	SourceId      string `json:"sourceId,omitempty"` //卷的来源id, 对应在来源的卷中进行数据的比对，看是否具备差异化复制的能力 -- clone的卷需要明确表明自身复制的卷id
	OriginId      string `json:"originId"`           //卷所对应的原始卷的id
	ImageId       string `json:"imageId,omitempty"`  //镜像的id
	StoragePoolId string `json:"storagePoolId"`      //卷所在的存储池的id
	StageSize     int64  `json:"stageSize"`          //为卷的接收准备的空间大小，由存储引擎进行实际的大小计算
	BlockSize     int    `json:"blockSize"`          //BlockSize与StageSize为BlockDevice模式的卷需要的参数，在BlockDevice与filesystem转换的复制过程中，如何保证两者的信息不丢失？

	StoragePoolOptions map[string]string `json:"storagePoolOptions,omitempty"` //the pool config options

	StorageStageType StageType `json:"storageStageType"`          //存储端的stageType，这个需要在进行数据获取和处理的时候进行处理和赋值
	StorageIdentity  string    `json:"storageIdentity,omitempty"` //存储端的实际地址，需要存储引擎在target端的时候给出并返回

	JobId     string `json:"jobId,omitempty"`     //任务id
	MainjobId string `json:"mainjobId,omitempty"` //主任务id

	EngineName     string `json:"engineName"` //存储引擎的名字，引擎名用于在存储服务器端找到对应的存储引擎
	EngineId       string `json:"engineId"`   //存储引擎的id，用于在存储服务器端找到对应的存储引擎
	EngineTypeName string `json:"engineTypeName"`
	//StorageEngineType runmodel.StorageEngineType `json:"storageEngineType"` //存储引擎type相同的认为是同构引擎，同构引擎可能会生成不同名字的引擎
	EngineMeta string `json:"engineMeta"` //存储引擎对数据的描述结构，由存储引擎给出，不同的存储引擎给出的数据结构不同
	//IsArchive         bool              `json:"isArchive,omitempty"` //是否为归档引擎，如果是归档引擎， 是否为归档引擎可以从StorageEngineType的level中判定 level == 1 为归档引擎
	DependVolumeMetadata *VolumeMetadata `json:"dependVolumeMetadata,omitempty"` //该卷的数据依赖的卷的metadata -- 当卷的数据是 差异 数据的时候，指定其所依赖的卷id
	//直到最后该数值为nil，表示卷可以独立存在，不依赖于其它的卷
	MaxSequence uint64 `json:"maxSequence,omitempty"` //最大的sequence,需要进行记录
}

/*
系统参数的结构及处理函数
*/
type SysArgs struct {
	Argmap  map[string]string `json:"argmap,omitempty"` //经过命令参数解析后的命令-->参数值的键值对
	PeName  string            `json:"peName,omitempty"` //命令行命令本身的值，对应os.Args[0]
	Command string            `json:"command"`          //命令行中指定的命令
}

func (sysargs *SysArgs) HasHelp() bool {
	return sysargs.HasCommand("help") || sysargs.HasCommand("h")
}

func (sysargs *SysArgs) HasVersion() bool {
	return sysargs.HasCommand("version") || sysargs.HasCommand("v")
}

// 参数列表中是否有某个命令
func (sysargs *SysArgs) HasCommand(command string) bool {
	_, hasvalue := sysargs.Argmap[command]
	return hasvalue
}

// 获取参数列表中的某个参数的数据
func (sysargs *SysArgs) GetCommandValue(command string) (string, bool) {
	value, hasvalue := sysargs.Argmap[command]
	return value, hasvalue
}
