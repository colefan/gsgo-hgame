package bs

type UserBind struct {
	UID     uint64
	Account string
}

func CreateUserBindObj() *UserBind {
	return &UserBind{}
}
