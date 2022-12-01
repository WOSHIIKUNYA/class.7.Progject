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
type Change struct {
	SendUid    string `json:"SendUid"`
	ReceiveUid string `json:"ReceiveUid"`
	Detail     string `json:"Detail"`
}
type Add1 struct {
	SendUid string `json:"SendUid"`
	Detail  string `json:"Detail"`
}
type Add2 struct {
	Id         int    `json:"Id"`
	SendUid    string `json:"SendUid"`
	ReceiveUid string `json:"ReceiveUid"`
}
