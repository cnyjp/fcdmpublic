package model

/**
  This is the environment variables defined for script and plugins.
*/
const (
	FCDM_EV_COMMAND = "FCDM_EV_COMMAND" //The command need plugin to execute.

	FCDM_EV_AD_PREFIX = "FCDM_EV_AD_" //The options configured in applicaiton

	//the volume access path for the applicaiton's volumes. which used PREFIX + volumename as a key,
	//and the access path as the value.
	FCDM_EV_VOLUME_PREFIX = "FCDM_EV_VOLUME_"

	//FCDM_EV_STAGE_TYPE = "FCDM_EV_STAGE_TYPE"

	//FCDM_EV_STAGE_PROTOCOL        = "FCDM_EV_STAGE_PROTOCOL"
	//FCDM_EV_STAGE_PROTOCOL_TARGET = "FCDM_EV_STAGE_PROTOCOL_TARGET"

	//The applicaiton name used in provider run command.
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

	FCDM_EV_HOST_LISTAPP_PAGE = "FCDM_EV_HOST_LISTAPP_PAGE" //list app page
	FCDM_EV_HOST_LISTAPP_SIZE = "FCDM_EV_HOST_LISTAPP_SIZE" //list app page size
	FCDM_EV_HOST_LISTAPP_TYPE = "FCDM_EV_HOST_LISSAPP_TYPE" //list app type
	FCDM_EV_HOST_LISTAPP_NEXT = "FCDM_EV_HOST_LISSAPP_NEXT" //list app next value

	//The prefix for list app options with secondlytype
	//attach to the listapptypeoptions options.
	FCDM_EV_HOST_LISTAPP_TYPE_OPT_PREFIX = "FCDM_EV_HOST_LISTAPP_TYPE_OPT_"

	FCDM_EV_PROVIDER_INSTANCE_OPT_PREFIX = "FCDM_EV_PROVIDER_INSTANCE_OPT_" //providerinstance mode applicaiton the providerinstance's options.
	FCDM_EV_PROVIDER_INSTANCE_ID         = "FCDM_EV_PROVIDER_INSTANCE_ID"   //providerinstance的id
)

const (
	FCDM_PROVIDER_PROCESS_PERCENT_PREFIX = "FCDM_PERCENT:"      //percent of current job running.
	FCDM_PROVIDER_PROTECT_DATA_PREFIX    = "FCDM_PROTECT_DATA:" //the backup job need return the
)

//指令的字符串常量
//指令，指的是FCDM_EV_COMMAND环境变量中的数据，这个是不能任意设置的，必须在指令常量的范围内
//Provider Commands
const (
	CMD_DISCOVER         = "discover"         //发现应用的指令
	CMD_APPLICATION_INFO = "applicaiton_info" //单独提取一个应用的信息
	CMD_BACKUP           = "backup"           //备份应用的指令
	CMD_MOUNT            = "mount"            //挂载镜像的指令

	CMD_UMOUNT = "umount" //停止挂载的指令  --  停止挂载是否应该时一个job？？？ -- 似乎也应该是一个job。

	//CMD_APPRESTORE = "apprestore"  //应用数据恢复的指令 -- 数据全部恢复
	CMD_RESTORE               = "restore"               //文件恢复的指令  --  文件恢复的指令应该有一组文件列表的指令跟随  -- 似乎也没有必要？？？
	CMD_RESTORE_MOUNT         = "restore_mount"         // -- 为了恢复而进行的mount操作？？？？？似乎没有必要？？？？
	CMD_RESTORE_TREELIST      = "restore_treelist"      //列出某个tree节点的下级节点，如果没有指定节点则为根节点
	CMD_RESTORE_LISTRECORD    = "restore_listrecode"    //列出某个tree节点下的所有数据记录（非tree节点），如果没有指定节点则为根节点，参数包括页码
	CMD_RESTORE_RESTORERECORD = "restore_restorerecord" //恢复某个记录的数据  -- 数据库可以实现 表空间、表、记录恢复，文件系统可以实现 目录、文件恢复
	CMD_RESTORE_RESTORETREE   = "restore_restoretree"   //恢复某个树节点下的所有数据  -- 如果不指定节点，那就是全部恢复了。

	CMD_REQUEST_STAGE         = "request_stage"         //请求stage  -- 由于第三方主机的加入，需要将请求与准备两个动作分开，请求的动作发生在向控制中心发送请求之前和发送中，而准备动作则发生在控制中心返回了stageinfo信息之后
	CMD_PREPARE_STAGE         = "prepare_stage"         //准备stage  -- 与storage的preparestage分开，指向为主机端准备本地stage的过程
	CMD_STORAGE_PREPARE_STAGE = "storage_prepare_stage" //准备stage  -- 该命令可以发给服务器，也可以发给clientadapter，storagaengine，serveradapter；每个组件对应的操作不同，但是指令应该是相同的。
	CMD_STORAGE_CREATE_IMAGE  = "storage_create_image"  //创建镜像  -- 该命令可以发给服务器，也可以发给clientadapter，storagaengine，serveradapter；每个组件对应的操作不同，但是指令应该是相同的。
	CMD_STORAGE_CLEAR_STAGE   = "storage_clear_stage"   //清理stage  -- 指向为存储服务器清理stage
	CMD_CLEAR_STAGE           = "clear_stage"           //清理stage的相应资源  -- 同上
	CMD_IMAGE_STAGE           = "image_stage"           //对stage制作镜像  -- 返回镜像的id及相关信息。 -- 这个只有在备份成功完成后才会需要发送，是否考虑直接与备份体系合成？
	CMD_CREATE_IMAGE          = "create_image"          //创建镜像的指令

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

	//第三方主机相关的增补指令
	CMD_HOST_LIST_APP              = "host_list_app"
	CMD_HOST_LIST_APP_SECONDLYTYPE = "host_list_app_secondlytype"
	CMD_HOST_LIST_APP_OPTION       = "host_list_app_option"

	//connector与server之间的通讯指令
	CMD_REGISTER = "register" //connector注册到服务器的指令

	CMD_UPDAET_IMAGE_INFO = "update_image_info" //connector备份完成后获取到对应的image的附加信息，将该附加信息返回给服务器端 ---- 似乎不是很有必要，只需要在备份完成的通知中将image的信息加进去就好了。todo:待验证。

	//connector端的指令
	CMD_SET_AUTH   = "setauth"   //设置认证信息
	CMD_SET_SERVER = "setserver" //设置服务器信息
	//CMD_SET_HOSTID  = "sethostid"  //设置主机id
	CMD_SHOW_SERVER = "showserver" //显示服务器设置的信息和认证信息
	CMD_COMMIT      = "commit"     //确认信息，提交信息
	CMD_CLEARDATA   = "cleardata"  //清除connector端存储的所有数据，包括认证信息，app信息，job信息，policy信息等等。以便于重新进行主机注册

	//CMD_INHERIT = "inherit" //继承，继承之前某个已经存在于服务器中的connector的信息；问题在于，继承而来的数据应该如何获取？

	//CMD_SET_CONNECTORINFO = "setconnectorinfo" //作为cli向connector一次性设置所有server相关信息的方法。将来的客户端或cli的想到可以通过这个命令进行参数设置。
	CMD_SHOW_AUTH = "showauth" //显示所有的认证相关信息，包括了服务器地址，port，zone，authid，authkey，authcode，以及宿主机的hostid

	CMD_PLUGIN_INFO = "cmd_plugin_info" //传递给plugin的提取对应info的指令，plugin需要返回自身的相应信息。PluginCmdInfo

	CMD_ADAPTER_ADDRESSES = "fcdm_cmd_adapter_address" //传递给adapter获取对应的协议及协议地址的指令

	/******* 第三方主机相关的特殊指令 **************/
	//CMD_GET_SEARCH_OPTIONS = "get_search_options" //提取所有的检索条件  ---- 问题在于，检索条件本身可能会出现需要 select的option的现象 -- 如果发生了该现象应该从这个接口获取所有的options
	//第三方主机的search选项可以考虑直接在pluginconfig中进行配置，配置不同种类的application使用何种类型的检索条件，并指定检索条件的输入方式；唯一不同的是，某些select类型的options不是配置文件配置出来的，而是通过上述指令获取的。
	//CMD_SEARCH_APPLICATIONS = "search_applications" //根据给定的条件检索应用  ---- 对应的指令应包括检索条件，返回值则为检索出来结果的applications。

	CMD_ADAPTER_INIT = "fcdm_cmd_adapter_init" //传递给adapter进行初始化的指令

	/********  存储部分的相关指令   ********/

	CMD_GET_STORAGE_STATUS = "get_storage_status" //获取storage server中存储的相关信息

	CMD_STORAGE_ADD_POOL        = "storage_add_pool"
	CMD_STORAGE_MODIFY_POOL     = "storage_modify_pool"
	CMD_STORAGE_DELETE_POOL     = "storage_delete_pool"
	CMD_STORAGE_LIST_POOL       = "storage_list_pool"
	CMD_STORAGE_INFO_POOL       = "storage_info_pool"
	CMD_STORAGE_ADD_COPY        = "storage_add_copy"
	CMD_STORAGE_MODIFY_COPY     = "storage_modify_copy"
	CMD_STORAGE_DELETE_COPY     = "storage_delete_copy"
	CMD_STORAGE_CREATE_SNAPSHOT = "storage_create_snapshot"
	CMD_STORAGE_DELETE_SNAPSHOT = "storage_delete_snapshot"
	CMD_STORAGE_CREATE_STAGE    = "storage_create_stage"
	CMD_STORAGE_DELETE_STAGE    = "storage_delete_stage"

	CMD_STORAGE_SERVER_ENSURE_JOB_STAGE = "storage_server_ensure_job_stage" //向存储服务器发送确认job的stage的命令

	CMD_ADD_POLICY    = "add_policy"    //增加policy  server向connector发送的指令
	CMD_SAVE_POLICY   = "save_policy"   //增加policy  server向connector发送的指令
	CMD_MODIFY_POLICY = "modify_policy" //修改policy  server向connector发送的指令
	CMD_DELETE_POLICY = "delete_policy" //删除policy server向connector发送的指令
	CMD_RUN_POLICY    = "run_policy"    //运行policy  server向connector发送的指令

	CMD_POLICY_TRIGGERED  = "policy_triggered"   //policy触发后发送的指令
	CMD_JOB_STATUS_CHANGE = "job_status_change"  //状态变化
	CMD_JOB_RETRY_CHANGE  = "job_retry_change"   //job的retry的状态变化
	CMD_HOST_JOB_START    = "host_job_start"     //从connector发送到server端的jobstart指令，该指令通知server jobstart的同时，从server要求job执行的必要信息；同时，如果job为一个appgroup的应用的话，server需要判断是否需要向不同的主机发送jobstart的指令
	CMD_SERVER_RUN_JOB    = "server_run_job"     //从server端向conenctor发送的jobstart指令，用于在代理host通知server启动job后，server向job需要执行的connector发送jobstart的指令。事实上，这个需要的是
	CMD_SERVER_JOB_CANCEL = "server_job_cancel"  //从server端向conenctor发送的job cancel指令，用于在代理host通知server启动job后，server向job需要执行的connector发送jobstart的指令。事实上，这个需要的是
	CMD_JOB_GET_STATUS    = "cmd_job_get_status" //从server端向conenctor发送的job cancel指令，用于在代理host通知server启动job后，server向job需要执行的connector发送jobstart的指令。事实上，这个需要的是

	CMD_HOST_JOBLIMIT_GET = "cmd_host_joblimit_get" //从server向主机发送的提取主机同时运行任务数的指令
	CMD_HOST_JOBLIMIT_SET = "cmd_host_joblimit_set" //从server向主机发送的设置主机同时运行任务数的指令

	CMD_HOST_DATA_SYNC   = "host_data_sync"   //从主机端向server同步数据的指令
	CMD_SERVER_DATA_SYNC = "server_data_sync" //从server向主机端同步数据的指令
)

/**
  The type of list app, used in the ListAppRequest and ListAppResponse.
*/
type ListAppType = string

const (
	LIST_APP_TYPE_ALL  ListAppType = "all"
	LIST_APP_TYPE_PAGE ListAppType = "page"
	LIST_APP_TYPE_MORE ListAppType = "more"
	LIST_APP_TYPE_TREE ListAppType = "tree"
)

/**
  存储通讯协议的自定义type及其对应的函数和常量。
*/
type StorageProtocol = string

const (
	STORAGE_PROTOCOLS = "protocols" //配置文件中标志多协议的字符串标志

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

//func (step *JobStep) String() string {
//	return string(*step)
//}

//定义一个全部job类型与对应的阶段的统一变量，以便后续引用
var JobTypeStepMap = map[string][]JobStep{
	JOB_TYPE_BACKUP:  {BACKUP_STEP_INIT, BACKUP_STEP_PREPARE, BACKUP_STEP_FREEZE, BACKUP_STEP_THAW, BACKUP_STEP_FINAL, BACKUP_STEP_CANCEL},
	JOB_TYPE_MOUNT:   {MOUNT_STEP_BEFORE, MOUNT_STEP_AFTER, UMOUNT_STEP_BEFORE, UMOUNT_STEP_AFTER},
	JOB_TYPE_RESTORE: {RESTORE_STEP_BEFORE, RESTORE_STEP_AFTER},
}

const (
	//这些是脚本的step
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

	//这个是为在整体job设计的step，为mainjob的step，不是subjob的step，但是可以标志job运行的阶段
	JOB_STEP_INIT   JobStep = "jobinit"
	JOB_STEP_NORMAL JobStep = "jobnormal"
	JOB_STEP_FINAL  JobStep = "jobfinal"
	JOB_STEP_CANCEL JobStep = "jobcancel"
	JOB_STEP_UMOUNT JobStep = "jobumount"

	JOB_STEP_FINISHED JobStep = "jobfinished" //一个特殊的阶段，标志任务已经结束，并不是运行中会产生的实际的阶段
)

type JobType = string //工作类型

const (
	JOB_TYPE_BACKUP  JobType = "backup"
	JOB_TYPE_MOUNT   JobType = "mount"
	JOB_TYPE_RESTORE JobType = "restore"
	JOB_TYPE_COPY    JobType = "copy"

	//umount类型的job，主要用于subjob的umount类型的处理
	JOB_TYPE_UMOUNT JobType = "umount"

	JOB_TYPE_RECEIVE JobType = "receive" //复制策略中的接收任务
)

/**
  使用的存储类型。指向为stage（舞台）的类型
*/
type StageType = string

const (
	STAGE_TYPE_VOLUME      StageType = "volume"      //卷格式，一般来说意味着多个blockdevice
	STAGE_TYPE_BLOCKDEVICE StageType = "blockdevice" //单个的块设备
	STAGE_TYPE_FILESYSTEM  StageType = "filesystem"  //文件系统
)

type AdapterType = string

const (
	ADAPTER_TYPE_TARGET    AdapterType = "target"
	ADAPTER_TYPE_INITIATOR AdapterType = "initiator"
)

/**
主机类型。
  主机对应的宿主机的问题：
*/

//const (
//	HOST_TYPE = "host_type" //配置文件中主机类型的配置字符串key
//
//	HOST_TYPE_STORAGESERVER = "StorageServer" //存储服务器  --  是否要列在主机列表中？？？
//
//	HOST_TYPE_LOCAL      = "local"      //本地主机，如果为nil或空字符串也是表示本地主机，默认的主机类型
//	HOST_TYPE_REMOTE     = "remote"     //远端主机，通过telnet，ssh等协议进行连接的远程主机
//	HOST_TYPE_VSPHERE    = "vsphere"    //esxi主机，通过esxi接口进行访问的主机
//	HOST_TYPE_VCENTER    = "vcenter"    //vcenter主机，通过vcenter接口进行访问的主机
//	HOST_TYPE_KVMHOST    = "kvmhost"    //标准的kvm宿主机，connector通过远程进行连接的kvm宿主机
//	HOST_TYPE_HYPERVHOST = "hypervhost" //标准的hyperv宿主机，connector通过远程连接连接宿主机
//)

const (
	JOB_PERCENT_PREFIX = "##FCDM JOB PERCENT"
)

type PluginConfigOption = string

const (
	PLUGIN_CONFIG_PROVIDER_ALLOW_CUSTOM PluginConfigOption = "FCDM_PROVIDER_AllowCustom" //是否允许进行自定义的选项的选项名，只有设置为true的时候才可以进行应用自定义
)

/**
根据主机类型的分析结果，主机类型对应的配置选项
这是一个事实上的常量，所以，对应的type不使用指针作为成员

*/
//var HOSTTYPE_CONFIGS = map[string][]ConfigConfig{
//	HOST_TYPE_REMOTE: []ConfigConfig{},
//}
//
//var USER_TYPE_MANAGER = 0
//var USER_TYPE_CONSUMER = 1
//
//func UserIsManager(authcode int) bool {
//	n := authcode & USER_TYPE_MANAGER
//	return n == USER_TYPE_MANAGER
//}
//func UserIsConsumer(authcode int) bool {
//	return USER_TYPE_CONSUMER == (authcode & USER_TYPE_CONSUMER)
//}

type BackupType = int

const (
	BACKUP_TYPE_ALL BackupType = 1
	BACKUP_TYPE_DB  BackupType = 2
	BACKUP_TYPE_LOG BackupType = 3
)
