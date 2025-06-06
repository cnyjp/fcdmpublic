package model

/*
This is the environment variables defined for script and plugins.
*/
const (
	FCDM_EV_COMMAND = "FCDM_EV_COMMAND" //The command need plugin to execute.

	FCDM_EV_AD_PREFIX = "FCDM_EV_AD_" //The options configured in application

	FCDM_EV_IMAGE_AD_PREFIX = "FCDM_EV_IMAGE_AD_" //the options configured from config for image

	//the volume access path for the application's volumes. which used PREFIX + volume's name as a key,
	//and the access path as the value.
	FCDM_EV_VOLUME_PREFIX          = "FCDM_EV_VOLUME_"          //the prefix for volume name
	FCDM_EV_VOLUME_IDENTITY_PREFIX = "FCDM_EV_VOLUME_IDENTITY_" //the prefix for volume identity

	FCDM_EV_CUSTOM_VOLUME_NAME_PREFIX = "FCDM_EV_CUSTOM_VOLUME_NAME_" //the prefix for custom volume name

	//FCDM_EV_STAGE_TYPE = "FCDM_EV_STAGE_TYPE"

	//FCDM_EV_STAGE_PROTOCOL        = "FCDM_EV_STAGE_PROTOCOL"
	//FCDM_EV_STAGE_PROTOCOL_TARGET = "FCDM_EV_STAGE_PROTOCOL_TARGET"

	FCDM_EV_APPGROUP_ID   = "FCDM_EV_APPGROUP_ID"   //for provider, to distinguish if the run job is a group job.
	FCDM_EV_APPGROUP_APPS = "FCDM_EV_APPGROUP_APPS" //for provider, to distinguish if the run job is a group job.

	//The application name used in provider run command.
	FCDM_EV_APPNAME = "FCDM_EV_APPNAME"

	//The application extension used in provider run command.
	FCDM_EV_APP_EXTENSION = "FCDM_EV_APP_EXTENSION"

	FCDM_EV_APP_SECONDLYTYPE = "FCDM_EV_APP_SECONDLYTYPE" //application's secondlytype

	//the job's jobstep, in different jobtype, the step is different.
	//this is set for job script.
	//backup's step is:init, prepare, freeze, thaw, final
	//mount's step is:beforemount, aftermount
	//umount's step is:beforeumount, afterumount
	//restore's step is:beforerestore, afterrestore
	FCDM_EV_JOBSTEP = "FCDM_EV_JOBSTEP"
	FCDM_EV_JOB_ID  = "FCDM_EV_JOB_ID" //job id in the env

	//the job's jobtype, backup, restore, mount, umount,copy, receive
	FCDM_EV_JOB_TYPE = "FCDM_EV_JOB_TYPE" //job's type(JobType) in the env
	//the job's policy type, use JobType:backup,restore,mount,copy
	FCDM_EV_POLICY_TYPE  = "FCDM_EV_POLICY_TYPE"  //policy type (policy's JobType) in the env
	FCDM_EV_POLICY_ID    = "FCDM_EV_POLICY_ID"    //policy id in the env
	FCDM_EV_MAINJOB_ID   = "FCDM_EV_MAINJOB_ID"   //mainjob id in the env, init, normal, umount, final.
	FCDM_EV_MAINJOB_STEP = "FCDM_EV_MAINJOB_STEP" //JobStep in mainjob
	//the umount job match mount job's id.
	FCDM_EV_MOUNT_JOB_ID = "FCDM_EV_MOUNT_JOB_ID"
	//the appgroup init job returned init message.
	FCDM_EV_JOB_INIT_MESSAGE = "FCDM_EV_JOB_INIT_MESSAGE"

	//application's id of job. sometimes script need use it.
	FCDM_EV_JOB_APP_ID   = "FCDM_EV_JOB_APP_ID"
	FCDM_EV_JOB_IMAGE_ID = "FCDM_EV_JOB_IMAGE_ID"
	//when the jobtype is backup, need use the backup type to ensure the which backup type is running.
	FCDM_EV_JOB_BACKUP_TYPE = "FCDM_EV_JOB_BACKUP_TYPE"

	//the stage info for job, used for provider instance supply stage to instance host.
	FCDM_EV_JOB_STAGE_INFO = "FCDM_EV_JOB_STAGE_INFO"

	FCDM_EV_HOST_LISTAPP_PAGE = "FCDM_EV_HOST_LISTAPP_PAGE" //list app page
	FCDM_EV_HOST_LISTAPP_SIZE = "FCDM_EV_HOST_LISTAPP_SIZE" //list app page size
	FCDM_EV_HOST_LISTAPP_TYPE = "FCDM_EV_HOST_LISSAPP_TYPE" //list app type
	FCDM_EV_HOST_LISTAPP_NEXT = "FCDM_EV_HOST_LISSAPP_NEXT" //list app next value

	//The prefix for list app options with secondlytype
	//attach to the listapptypeoptions options.
	FCDM_EV_HOST_LISTAPP_TYPE_OPT_PREFIX = "FCDM_EV_HOST_LISTAPP_TYPE_OPT_"

	FCDM_EV_PROVIDER_INSTANCE_OPT_PREFIX = "FCDM_EV_PROVIDER_INSTANCE_OPT_" //providerinstance mode application the providerinstance's options.
	FCDM_EV_PROVIDER_INSTANCE_ID         = "FCDM_EV_PROVIDER_INSTANCE_ID"   //providerinstance的id

	FCDM_EV_TREELIST_NODE_ID            = "FCDM_EV_TREELIST_NODE_ID"            //the tree node to list it's children, empty means root node
	FCDM_EV_RESTORE_NODE_IDENTITIES     = "FCDM_EV_RESTORE_NODE_IDENTITIES"     //the node identities to restore
	FCDM_EV_TREELIST_DATA_FILTER_PREFIX = "FCDM_EV_TREELIST_DATA_FILTER_PREFIX" //the search condition for tree list data prefix

	//for storage engine
	FCDM_EV_ENGINE_ID                 = "FCDM_EV_ENGINE_ID"           //the storage engine id
	FCDM_EV_ENGINE_OPTIONS_PREFIX     = "FCDM_EV_ENGINE_OPTIONS_"     //the options prefix for engine create or modify
	FCDM_EV_ENGINE_OLD_OPTIONS_PREFIX = "FCDM_EV_ENGINE_OLD_OPTIONS_" //the options prefix for engine create or modify

	FCDM_EV_ENGINE_POOL_ID                  = "FCDM_EV_ENGINE_POOL_ID"           //pool id in env
	FCDM_EV_ENGINE_POOL_OPTIONS_PREFIX      = "FCDM_EV_ENGINE_POOL_OPTIONS_"     //the options prefix for pool create or modify
	FCDM_EV_ENGINE_POOL_OLD_OPTIONS_PREFIX  = "FCDM_EV_ENGINE_OLD_POOL_OPTIONS_" //the options prefix for pool create or modify
	FCDM_EV_ENGINE_POOL_SPACE_USED_PREFIX   = "FCDM_EV_ENGINE_POOL_SPACE_USED_"
	FCDM_EV_ENGINE_POOL_SPACE_ADD_PREFIX    = "FCDM_EV_ENGINE_POOL_SPACE_ADD_"
	FCDM_EV_ENGINE_POOL_SPACE_REMOVE_PREFIX = "FCDM_EV_ENGINE_POOL_SPACE_REMOVE_"
	FCDM_EV_ENGINE_POOL_SPACE_LOG_PREFIX    = "FCDM_EV_ENGINE_POOL_LOG_REMOVE_"
	FCDM_EV_ENGINE_POOL_SPACE_CACHE_PREFIX  = "FCDM_EV_ENGINE_POOL_CACHE_REMOVE_"

	FCDM_EV_ENGINE_POOL_VOLUME_INFO           = "FCDM_EV_ENGINE_POOL_VOLUME_INFO"           //volume full information for storage engine and engine type.
	FCDM_EV_ENGINE_POOL_VOLUME_METADATA       = "FCDM_EV_ENGINE_POOL_VOLUME_METADATA"       //volume full metadata for storage engine and engine type.
	FCDM_EV_ENGINE_POOL_IMAGE_METADATA        = "FCDM_EV_ENGINE_POOL_IMAGE_METADATA"        //volume full metadata for storage engine and engine type.
	FCDM_EV_ENGINE_POOL_VOLUME_METADATA_ACTOR = "FCDM_EV_ENGINE_POOL_VOLUME_METADATA_ACTOR" //volume metadata actor for storage engine and engine type.
)

const (
	FCDM_PROVIDER_PROCESS_PERCENT_PREFIX = "FCDM_PERCENT:"      //percent of current job running.
	FCDM_PROVIDER_PROTECT_DATA_PREFIX    = "FCDM_PROTECT_DATA:" //the backup job need return the
)

const (
	FCDM_PROVIDER_RESTORE_NODE_IDENTITIES_CONFIG_NAME = "FCDM_CONFIG_NODE_IDENTITIES" //specific restore node ids config name in provider
)

// The special code return by plugins
const (
	CODE_RUN_SUCCESS           = 0   //which means the plugin run command succeed.
	CODE_STAGE_NEED_EXPAND     = 201 //which means the stage need expand by default rules.
	CODE_STAGE_EXPAND_TO       = 202 //which means the stage need expand by return volume data.
	CODE_DISCOVER              = 203 //which means the provider found application change, need refresh the application's information
	CODE_STORAGE_VOLUME_LOCKED = 101 //the volume is locked, when check volume locked return this code.
)

// command string
// these command is used for FCDM_EV_COMMAND env, FCDM_EV_COMMAND must be one of the command list.
// plugin commands
const (
	CMD_DISCOVER = "discover" //command discover applications

	CMD_DISCOVER_APP_MORE = "discover_app_more" //command to use an application's options to discover more child or and other details

	CMD_APPLICATION_INFO = "application_info" //command used to retrieve an application's information
	CMD_BACKUP           = "backup"           //command for backup operate
	CMD_MOUNT            = "mount"            //command for mount operate

	CMD_UMOUNT = "umount" //command for umount operate

	CMD_RESTORE = "restore" //command for restore operate

	CMD_RESTORE_RESTORETREE = "restore_restoretree" //command for restore data under a tree node. if not specific a tree node, should restore all data.

	CMD_APP_LIST_TREENODE      = "app_list_treenode"      //command to list a tree node's children node data
	CMD_APP_LIST_TREENODE_DATA = "app_list_treenode_data" //command to list a tree node's list data

	CMD_RESTORE_LIST_TREENODE      = "restore_list_treenode"      //command to list a tree node's children node data on restore mode.
	CMD_RESTORE_LIST_TREENODE_DATA = "restore_list_treenode_data" //command to list a tree node's list data on restore mode

	CMD_REQUEST_STAGE         = "request_stage"         //command request stage  -- 由于第三方主机的加入，需要将请求与准备两个动作分开，请求的动作发生在向控制中心发送请求之前和发送中，而准备动作则发生在控制中心返回了stageinfo信息之后
	CMD_PREPARE_STAGE         = "prepare_stage"         //command prepare stage  -- 与storage的preparestage分开，指向为主机端准备本地stage的过程
	CMD_STORAGE_PREPARE_STAGE = "storage_prepare_stage" //command prepare stage on storage server side.  -- 该命令可以发给服务器，也可以发给clientadapter，storagaengine，serveradapter；每个组件对应的操作不同，但是指令应该是相同的。
	CMD_STORAGE_CREATE_IMAGE  = "storage_create_image"  //command create image on storage server side.创建镜像  -- 该命令可以发给服务器，也可以发给clientadapter，storagaengine，serveradapter；每个组件对应的操作不同，但是指令应该是相同的。
	CMD_STORAGE_CLEAR_STAGE   = "storage_clear_stage"   //command clear stage on storage server side.  -- 指向为存储服务器清理stage

	CMD_CLEAR_STAGE  = "clear_stage"  //command clear stage
	CMD_IMAGE_STAGE  = "image_stage"  //对stage制作镜像  -- 返回镜像的id及相关信息。 -- 这个只有在备份成功完成后才会需要发送，是否考虑直接与备份体系合成？
	CMD_CREATE_IMAGE = "create_image" //创建镜像的指令

	CMD_GETINITIRTORS = "getinitiators" //发现客户端备份接口的指令
	CMD_GETTARGETS    = "gettargets"    //发现存储服务器存储接口的指令

	CMD_ADPTC_GETPORTS = "adptc_getports" //指定客户端发现接口的指令，不同协议的adapter发现自己协议内的接口
	CMD_ADPTC_SCANLUNS = "adptc_scanluns" //指定客户端发现存储供给的指令
	CMD_ADPTC_GETSTAGE = "adptc_getstage" //指定客户端给出存储的位置，改位置可能为设备，也可能为一个挂载好的文件系统

	CMD_ADAPTER_CONNECT_TO_STORAGE = "adapter_conn_to_storage"    //连接到存储服务器的clientadapter的指令
	CMD_ADAPTER_MOUNT_VOLUME       = "adapter_mount_volumn"       //连接到存储服务器的clientadapter的指令
	CMD_ADAPTER_PREPARE_STAGE      = "adapter_prepare_stage"      //客户端的adapter准备stage
	CMD_ADAPTER_CLEAR_STAGE        = "adapter_clear_stage"        //客户端的adapter清理stage
	CMD_ADAPTER_CLEAR_STALE_DEVICE = "adapter_clear_stale_device" //客户端的adapter呆设备(无用挂载)

	//third host commands
	CMD_HOST_LIST_APP              = "host_list_app"
	CMD_HOST_LIST_APP_SECONDLYTYPE = "host_list_app_secondlytype"
	CMD_HOST_LIST_APP_OPTION       = "host_list_app_option"

	CMD_HOST_DISCOVER_STORAGE_INTERFACE = "host_discover_storage_interface" //discover storage interface on provider instance

	//connector端的指令
	CMD_SET_AUTH   = "setauth"   //设置认证信息
	CMD_SET_SERVER = "setserver" //设置服务器信息
	//CMD_SET_HOSTID  = "sethostid"  //设置主机id
	CMD_SHOW_SERVER = "showserver" //显示服务器设置的信息和认证信息
	CMD_COMMIT      = "commit"     //确认信息，提交信息
	CMD_CLEARDATA   = "cleardata"  //清除connector端存储的所有数据，包括认证信息，app信息，job信息，policy信息等等。以便于重新进行主机注册

	CMD_SHOW_AUTH = "showauth" //显示所有的认证相关信息，包括了服务器地址，port，zone，authid，authkey，authcode，以及宿主机的hostid

	CMD_PLUGIN_INFO = "cmd_plugin_info" //传递给plugin的提取对应info的指令，plugin需要返回自身的相应信息。PluginCmdInfo

	CMD_ADAPTER_ADDRESSES = "fcdm_cmd_adapter_address" //传递给adapter获取对应的协议及协议地址的指令

	CMD_ADAPTER_INIT = "fcdm_cmd_adapter_init" //command to init adapter

	CMD_SADAPTER_ENSUER_STAGE = "sadapter_ensure_stage" //adapter的确认stage的指令

	CMD_SADAPTER_CLEAR_STAGE    = "sadapter_clear_stage"    //clear stage command for sadapter
	CMD_SADAPTER_DELETE_VOLUME  = "sadapter_delete_volume"  //delete supply volume command for sadapter
	CMD_SADAPTER_LIST_GROUP_IDS = "sadapter_list_group_ids" //delete supply volume command for sadapter
	CMD_SADAPTER_DELETE_GROUP   = "sadapter_delete_group"   //delete supply volume command for sadapter

	/*******  commands for storage   ********/

)

const (
	//commands for storageEngineType
	CMD_STORAGE_ENGINE_TYPE_INIT = "storage_engine_type_init" //init command for storage engine or engine type. Used when storage server start to init storage engine.
	CMD_STORAGE_ENGINE_TYPE_INFO = "storage_engine_type_info" //the info command for storage engine or engine type.Used when get the engine info.
	//CMD_STORAGE_ENGINE_TYPE_CREATE_ENGINE        = "storage_engine_type_create_engine"        //create storage engine command for storage engine type
	//CMD_STORAGE_ENGINE_TYPE_MODIFY_ENGINE        = "storage_engine_type_modify_engine"        //create storage engine command for storage engine type
	//CMD_STORAGE_ENGINE_TYPE_DELETE_ENGINE        = "storage_engine_type_delete_engine"        //create storage engine command for storage engine type
	//CMD_STORAGE_ENGINE_TYPE_CHECK_ENGINE_OPTIONS = "storage_engine_type_check_engine_options" //check the create engine options for enginetype

	//command for storage engine
	CMD_STORAGE_ENGINE_CREATE             = "storage_engine_create" //init command for storage engine or engine type. Used when storage server create a new storage engine.
	CMD_STORAGE_ENGINE_MODIFY             = "storage_engine_modify" //init command for storage engine or engine type. Used when storage server start to init storage engine.
	CMD_STORAGE_ENGINE_INIT               = "storage_engine_init"   //init command for storage engine or engine type. Used when storage server start to init storage engine.
	CMD_STORAGE_ENGINE_INFO               = "storage_engine_info"   //the info command for storage engine or engine type.Used when get the engine info.
	CMD_STORAGE_ENGINE_STATUS             = "storage_engine_status" //the status command for storage engine or engine type.Used when get the engine status.
	CMD_STORAGE_ENGINE_DELETE             = "storage_engine_delete" //the delete command for storage engine or engine type.Used when delete engine
	CMD_STORAGE_ENGINE_CHECK_OPTIONS      = "storage_engine_check_options"
	CMD_STORAGE_ENGINE_CHECK_POOL_OPTIONS = "storage_engine_check_pool_options" //check the engine options can create a valid engine
	CMD_STORAGE_ENGINE_CREATE_POOL        = "storage_engine_create_pool"        //check the engine options can create a valid engine
	CMD_STORAGE_ENGINE_MODIFY_POOL        = "storage_engine_modify_pool"        //check the engine options can create a valid engine
	CMD_STORAGE_ENGINE_DELETE_POOL        = "storage_engine_delete_pool"        //check the engine options can create a valid engine
	CMD_STORAGE_ENGINE_LIST_POOL          = "storage_engine_list_pool"          //check the engine options can create a valid engine
	CMD_STORAGE_ENGINE_GET_POOL_INFO      = "storage_engine_get_pool_info"      //get pool info, include size, spaces. If the pool not exist, return the valid spaces.

	//command for storage engine pool
	CMD_STORAGE_ENGINE_POOL_ENSURE_VOLUME                   = "storage_engien_pool_ensure_volume"
	CMD_STORAGE_ENGINE_POOL_RETRIEVE_VOLUME                 = "storage_engien_pool_retrieve_volume"
	CMD_STORAGE_ENGINE_POOL_DELETE_VOLUME                   = "storage_engine_pool_delete_volume"
	CMD_STORAGE_ENGINE_POOL_LIST_VOLUME                     = "storage_engine_pool_list_volume"
	CMD_STORAGE_ENGINE_POOL_COUNT_VOLUME_METADATA           = "storage_engine_pool_count_volume_metadata"
	CMD_STORAGE_ENGINE_POOL_GET_VOLUME_METADATA_READER      = "storage_engine_pool_get_volume_metadata_reader"
	CMD_STORAGE_ENGINE_POOL_READ_VOLUME                     = "storage_engine_pool_read_volume"
	CMD_STORAGE_ENGINE_POOL_RELEASE_VOLUME_METADATA_READER  = "storage_engine_pool_release_volume_metadata_reader"
	CMD_STORAGE_ENGINE_POOL_GET_VOLUME_METADATA_WRITER      = "storage_engine_pool_get_volume_metadata_writer"
	CMD_STORAGE_ENGINE_POOL_WRITE_VOLUME                    = "storage_engine_pool_write_volume"
	CMD_STORAGE_ENGINE_POOL_COMMIT_VOLUME_METADATA_WRITER   = "storage_engine_pool_commit_volume_metadata_writer"
	CMD_STORAGE_ENGINE_POOL_ROLLBACK_VOLUME_METADATA_WRITER = "storage_engine_pool_rollback_volume_metadata_writer"
	CMD_STORAGE_ENGINE_POOL_BACKUP_COMMIT_VOLUME            = "storage_engine_pool_backup_commit_volume"
	CMD_STORAGE_ENGINE_POOL_UNLOCK_VOLUME                   = "storage_engine_pool_unlock_volume"
	CMD_STORAGE_ENGINE_POOL_BACKUP_CHECK_AND_LOCK_VOLUME    = "storage_engine_pool_backup_check_and_lock_volume"
	CMD_STORAGE_ENGINE_POOL_BACKUP_ROLLBACK_LOCK_VOLUME     = "storage_engine_pool_backup_rollback_lock_volume"
	CMD_STORAGE_ENGINE_POOL_IS_VOLUME_LOCKED                = "storage_engine_pool_is_volume_locked"
	CMD_STORAGE_ENGINE_POOL_IS_VOLUME_USED                  = "storage_engine_pool_is_volume_used"
	CMD_STORAGE_ENGINE_POOL_GET_IMAGE_METADATA              = "storage_engine_pool_get_image_metadata"
	CMD_STORAGE_ENGINE_POOL_ROLLBACK_ALL_VOLUME             = "storage_engine_pool_rollback_all_volume"
	CMD_STORAGE_ENGINE_POOL_LIST_LOCKED_VOLUMES             = "storage_engine_pool_list_locked_volumes"
	CMD_STORAGE_ENGINE_LIST_ALL_LOCKED_VOLUMES              = "storage_engine_list_all_locked_volumes"

	CMD_STORAGE_ENGINE_POOL_SUMMARY_VOLUME = "storage_engine_pool_summary_volume" //get a volume summary in the pool
)

const (
	CLI_CMD_VERSION        = "version" //used for pefile to show version in cli
	CLI_CMD_VERSION_SHORT  = "v"       //used for pefile to show version in cli
	CLI_CMD_VERSION_MIDDLE = "ver"     //used fo pefile to show version in cli

	CLI_CMD_INSTALL   = "install"
	CLI_CMD_UNINSTALL = "uninstall"
)

// The type of list app, used in the ListAppRequest and ListAppResponse.
type ListAppType = string

const (
	LIST_APP_TYPE_ALL  ListAppType = "all"
	LIST_APP_TYPE_PAGE ListAppType = "page"
	LIST_APP_TYPE_MORE ListAppType = "more"
	LIST_APP_TYPE_TREE ListAppType = "tree"
)

// The storage protocols.
type StorageProtocol = string

const (
	STORAGE_PROTOCOLS = "protocols" //protocls name in config files

	STORAGE_PROTOCOL_ISCSI StorageProtocol = "iscsi"
	STORAGE_PROTOCOL_ISER  StorageProtocol = "iser"
	STORAGE_PROTOCOL_FC    StorageProtocol = "fc"
	STORAGE_PROTOCOL_SRP   StorageProtocol = "srp" //ib
	STORAGE_PROTOCOL_NFS   StorageProtocol = "nfs"
	STORAGE_PROTOCOL_NBD   StorageProtocol = "nbd"

	//针对本地数据存储提供的协议类型，一般只用于模拟器和较特殊的需要使用本地存储进行临时替代的场景
	STORAGE_PROTOCOL_VOLUME StorageProtocol = "volume"
	STORAGE_PROTOCOL_PATH   StorageProtocol = "path"

	LOCAL_STORAGE_NAME_INSIDE = "localstorage" //本地存储的adapter的名字，目录名和协议名

)

type JobStep = string

// a job type-step map, include backup,mount,restore jobtype.
// every job type has multi steps.
var JobTypeStepMap = map[string][]JobStep{
	JOB_TYPE_BACKUP:  {BACKUP_STEP_INIT, BACKUP_STEP_PREPARE, BACKUP_STEP_FREEZE, BACKUP_STEP_THAW, BACKUP_STEP_FINAL, BACKUP_STEP_CANCEL},
	JOB_TYPE_MOUNT:   {MOUNT_STEP_BEFORE, MOUNT_STEP_AFTER, UMOUNT_STEP_BEFORE, UMOUNT_STEP_AFTER},
	JOB_TYPE_RESTORE: {RESTORE_STEP_BEFORE, RESTORE_STEP_AFTER},
}

const (
	//step for scripts
	BACKUP_STEP_INIT    JobStep = "backup_init"
	BACKUP_STEP_PREPARE JobStep = "backup_prepare"
	BACKUP_STEP_FREEZE  JobStep = "backup_freeze"
	BACKUP_STEP_THAW    JobStep = "backup_thaw"
	BACKUP_STEP_FINAL   JobStep = "backup_final"
	BACKUP_STEP_CANCEL  JobStep = "backup_cancel"

	MOUNT_STEP_BEFORE JobStep = "beforemount"
	MOUNT_STEP_AFTER  JobStep = "aftermount"

	UMOUNT_STEP_BEFORE JobStep = "beforeumount"
	UMOUNT_STEP_AFTER  JobStep = "afterumount"

	RESTORE_STEP_BEFORE JobStep = "beforerestore"
	RESTORE_STEP_AFTER  JobStep = "afterrestore"

	//mainjob step
	JOB_STEP_INIT   JobStep = "jobinit"
	JOB_STEP_NORMAL JobStep = "jobnormal"
	JOB_STEP_FINAL  JobStep = "jobfinal"
	JOB_STEP_CANCEL JobStep = "jobcancel"
	JOB_STEP_UMOUNT JobStep = "jobumount"

	JOB_STEP_FINISHED JobStep = "jobfinished" //special job step， means the mainjob is finished.
)

type JobType = string //JobType for subjob and mainjob

const (
	JOB_TYPE_BACKUP  JobType = "backup"
	JOB_TYPE_MOUNT   JobType = "mount"
	JOB_TYPE_RESTORE JobType = "restore"
	JOB_TYPE_COPY    JobType = "copy"

	JOB_TYPE_UMOUNT JobType = "umount" //only for subjob

	JOB_TYPE_RECEIVE JobType = "receive" //only for subjob
)

type StageType = string //type for job stage

const (
	STAGE_TYPE_VOLUME      StageType = "volume"      //volume
	STAGE_TYPE_BLOCKDEVICE StageType = "blockdevice" //one blockdevice
	STAGE_TYPE_FILESYSTEM  StageType = "filesystem"  //filesystem
)

// the type for StorageAdapter
type AdapterType = string

const (
	ADAPTER_TYPE_TARGET    AdapterType = "target"
	ADAPTER_TYPE_INITIATOR AdapterType = "initiator"
)

const (
	PLUGIN_OUTPUT_PREFIX_SUBLOG = "##FCDM_PLUGIN_OUTPUT_SUBLOG::" //if a plugin want output a log message saved on server side, use this prefix

)

const (
	JOB_PERCENT_PREFIX = "##FCDM JOB PERCENT"
)

// BackupType
// Used for policy and image.
// the provider will use it to judge the type of backup action.
type BackupType = int

const (
	BACKUP_TYPE_ALL BackupType = 1 //both DB and LOG
	BACKUP_TYPE_DB  BackupType = 2 //only DB
	BACKUP_TYPE_LOG BackupType = 3 //only LOG
)

// CopyActor
// used for copy policy
// label the copy source or target
type CopyActor = string

const (
	COPY_ACTOR_SOURCE CopyActor = "source"
	COPY_ACTOR_TARGET CopyActor = "target"
)

const (
	COMM_CMD_CANCEL = "comm_cmd_cancel" //echo to plugin to cancel current task
)
