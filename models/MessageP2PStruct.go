package models

type MessageP2PStruct struct {
	Command string
	Id      int
	To      int
	Data    interface{}
}
