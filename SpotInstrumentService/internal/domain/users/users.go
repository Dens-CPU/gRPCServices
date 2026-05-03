package domainusers

type UserRole int32

const (
	USER_ROLE_UNSPECIFIED  UserRole = 0
	USER_ROLE_BASIC_USER   UserRole = 1
	USER_ROLE_PREMIUM_USER UserRole = 2
)

type User struct {
	Role UserRole
}

func NewUser(role UserRole) *User {
	return &User{Role: role}
}

type Input struct {
	UserRole  UserRole
	UserId    string
	PageSize  int
	PageToken string
}
