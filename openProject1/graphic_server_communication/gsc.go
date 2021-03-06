package graphic_server_communication

type GSIC_data struct {
	UniqueID   string `json:"ID"`
	Timestamp  string `json:"timestamp"`
	Posx       string `json:"Posx"`
	Posy       string `json:"Posy"`
	HP         string `json:"HP"`
	AP         string `json:"AP"`
	Class      string `json:"Class"`
	ClassPower string `json:"ClassPower"`
	GoToPosx   string `json:"GoToPosx"`
	GoToPosy   string `json:"GoToPosy"`
	ActionID   string `json:"ActionID"`
	Md5        string `json:"md5"`
}

type GSOC_data struct {
	UniqueID   string `json:"ID"`
	Timestamp  string `json:"timestamp"`
	Result     string `json:"res"`
	ActionID   string `json:"aid"`
	HP         string `json:"hp"`
	AP         string `json:"ap"`
	ClassPower string `json:"cp"`
	Md5        string `json:"md5"`
}

