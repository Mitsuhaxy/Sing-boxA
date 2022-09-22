package models

// RS: RequestSet
type Log struct {
	Disabled  bool   `json:"disabled"`
	Leavel    string `json:"leavel"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}
