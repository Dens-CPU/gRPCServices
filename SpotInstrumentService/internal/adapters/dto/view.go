package viewdto

type Output struct {
	ID   string
	Name string
}

type Input struct {
	UserRole  string
	UserId    string
	PageSize  int
	PageToken string
}
