package api

type Response struct {
	Code     string `json:"Code"`
	Message  string `json:"Message"`
	CodeDesc string `json:"codeDesc,omitempty"`
}
