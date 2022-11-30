package modle

type Message struct {
	Id         int    `json:"Id"`
	SendUid    string `json:"SendUid"`
	ReceiveUid string `json:"ReceiveUid"`
	Detail     string `json:"Detail"`
}

var LoginUser string = ""

type Add struct {
	ReceiveUid string `json:"ReceiveUid"`
	Detail     string `json:"Detail"`
}
