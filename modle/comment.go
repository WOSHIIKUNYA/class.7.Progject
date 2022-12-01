package modle

type Comment struct {
	Message   string `json:"Message"`
	Detail    string `json:"Detail"`
	Commenter string `json:"Commenter"`
}
type Comment1 struct {
	Message string `json:"comment"`
	Detail  string `json:"Detail"`
}
type Comment3 struct {
	Message string `json:"comment"`
}

var Comment2 bool = false
