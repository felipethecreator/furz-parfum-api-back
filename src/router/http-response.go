package router

type ERRORMessage struct {
	Error   string `json: "error" bson: "error"`
	Message string `json: "message" bson: "message"`
	Reason  string `json: "reason" bson: "reason"`
}
