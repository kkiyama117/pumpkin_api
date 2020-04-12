package entities

type ServerStatus struct {
	Code  int    `json:"code"`
	Version  string    `json:"version"`
	Value string `json:"status"`
}
type Statuses ServerStatus
