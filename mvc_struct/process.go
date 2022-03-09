package mvc_struct

type ActionHistoryItem struct {
	ManagerId   string `json:"manager_id"`
	Action      string `json:"action"`
	Description string `json:"description"`
	EditAt      string `json:"edit_at"`
}

type BaseProcess struct {
	History []ActionHistoryItem `json:"history"`
	Status  int                 `json:"status"`
}
