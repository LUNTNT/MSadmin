package Model

type StandardReply struct {
	ID			string `json:"id" bson:"id"`
	Name 		string `json:"name" bson:"name"`
	Content		[]string `json:"content" bson:"content"`
	Channel 	[]string `json:"channel" bson:"channel"`
	Team 		string `json:"team" bson:"team"`
	Assignee 	[]string `json:"assignee" bson:"assignee"`
}