package response

type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Current int         `json:"current"`
	Size    int         `json:"size"`
}
