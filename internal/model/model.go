package model

type Message struct {
	ChatId  int64
	UserId  int64
	Message string
}

type AnalysisJob struct {
	Target  Message
	Context []Message
}
