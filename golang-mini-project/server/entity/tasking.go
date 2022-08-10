package entity

//Tasking object for REST(CRUD)
type Tasking struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Assigne  string `json:"Assigne"`
	Deadline string `json:"deadline"`
}
