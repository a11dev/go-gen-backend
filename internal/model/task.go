package model

type Task struct {
	Key          int64
	Number       int64
	Status       string
	Typecategory string
	Refid        string
}

type TaskDetail struct {
	Key  int64
	Body []byte
}

type ResultTaskDetail struct {
	Key  int64
	Body interface{}
}
