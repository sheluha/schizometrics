package model

type Message struct {
	UserId  int64
	Message string
}

type AnalysisJob struct {
	Target    Message
	Context   []Message
	ChatId    int64
	MessageId int
}
