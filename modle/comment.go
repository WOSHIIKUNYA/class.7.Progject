package modle

type Comment struct {
	Message   string `json:"Message"`
	Detail    string `json:"Detail"`
	Commenter string `json:"Commenter"`
}

var CommentBook []Comment

type Comment1 struct {
	Message string `json:"comment"`
	Detail  string `json:"Detail"`
}
type Comment3 struct {
	Message string `json:"comment"`
}

var Comment2 bool = false
var Comment4 bool = false

type Comment5 struct {
	Id        int    `json:"Id"`
	Message   string `json:"Message"`
	Detail    string `json:"Detail"`
	Commenter string `json:"Commenter"`
}
