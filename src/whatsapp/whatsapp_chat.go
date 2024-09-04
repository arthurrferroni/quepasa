package whatsapp

type WhatsappChat struct {
	Id           string `json:"id"`
	Title        string `json:"title,omitempty"`
	GroupID      string `json:"groupid,omitempty"`
	GroupName    string `json:"groupName,omitempty"`
	Owner        string `json:"owner,omitempty"`
	Created      string `json:"created,omitempty"`
	Participants int    `json:"participants,omitempty"`
}

var WASYSTEMCHAT = WhatsappChat{Id: "system", Title: "Internal System Message"}
