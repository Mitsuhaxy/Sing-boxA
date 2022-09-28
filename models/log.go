package models

// RS: RequestSet
type Log struct {
	Disabled  bool   `json:"disabled"`
	Level     string `json:"level"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}
