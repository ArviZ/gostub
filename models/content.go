package models

type ContentList struct {
	Default 	Content		`json:"default"`
	Handlers    []Handler  	`json:"handlers"`
}

type Content struct {
	Body 		string 				`json:"body"`
	Status 		int 				`json:"status"`
	Header 		map[string]string 	`json:"header"`
	Cookie 		map[string]string 	`json:"cookie"`
}

type Handler struct {
	Content 	Content 			`json:"content"`
	Header 		map[string]string 	`json:"header"`
	Path 		map[string]string   `json:"path"`
	Param 		map[string]string 	`json:"param"`
}
