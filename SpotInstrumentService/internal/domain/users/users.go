package domainusers

type UserType int32

const (
	USER_ROLE_UNSPECIFIED UserType = 0
	USER_ROLE_USER        UserType = 1
	USER_ROLE_ADMIN       UserType = 2
)

type User struct {
	Role UserType
}

func NewUser(role UserType) *User {
	return &User{Role: role}
}
