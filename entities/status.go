package entities

type ServerStatus struct {
	Code  int    `json:"code"`
	Value string `json:"status"`
}
type Statuses ServerStatus
