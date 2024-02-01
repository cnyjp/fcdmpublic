package model

/*
*

	The struct used by provider backup command.
	When the provider run backup command and return a special code, will use the BackupResponse struct to unmarshal
	the last line printed by the provider.
	If the unmarshal error, then do not use the returned value to set the next step message.
*/
type BackupResponse struct {
	InitMessage       string `json:"initMessage,omitempty"`       //the init jobStep job returned message.
	ProtectedDataSize int64  `json:"protectedDataSize,omitempty"` //the backup job protected data size.
	ImageName         string `json:"imageName,omitempty"`         //the provider named image name.
	//sometimes we need the backup image configure some option for the mount or restore action, use the configs return
	//and save then in the image options. which will cover the option of the application's options.
	Configs []ConfigConfig `json:"configs,omitempty"`
}

/*
*

	 The struct used by provider list_app command.
	 Whent the provider run list_app command and return a zero code, will use the struct to unmarshal the last line
	printed by the provider.
*/
type ListAppResponse struct {
	HostId             string            `json:"hostId,omitempty"`
	AgentHostId        string            `json:"agentHostId,omitempty"`
	ProviderInstanceId string            `json:"providerInstanceId,omitempty"`
	Applications       []Application     `json:"applications"`
	Options            map[string]string `json:"options,omitempty"`      //search condition
	SecondlyType       string            `json:"secondlyType,omitempty"` //search secondly type
	Total              int               `json:"total"`                  //total record number
	Page               int               `json:"page"`                   //page number
	Size               int               `json:"size"`                   //page size
	Next               string            `json:"next"`                   //next page url -- for more type
	Type               ListAppType       `json:"type"`                   //list type
	TreeNodes          []TreeNode        `json:"treeNodes"`              //if the tree mode used, this is the all tree nodes.
}
