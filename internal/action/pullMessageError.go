package action

// Ответ при ошибке pull message (http 500)
type PullMessageErrorResponse struct {
	Action
	Data pullMessageErrorData `json:"Data"`
}

// Data у action при ошибке pull message
type pullMessageErrorData struct {
	Success bool          `json:"Success"`
	Message string        `json:"Message"`
	Details string        `json:"Details"`
	Data    errorDataData `json:"Data"`
}

type errorDataData struct {
	RawMsg interface{} `json:"RawMsg"`
}
