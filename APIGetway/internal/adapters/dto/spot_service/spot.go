package spotservicedto

type Output struct {
	ID   string
	Name string
}

type Input struct {
	UserID    string
	UserRole  string
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
}
