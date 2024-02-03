package events

type LogoutUser struct {
	name string
}

func (lu *LogoutUser) GetName() string {
	return lu.name

}

func (lu *LogoutUser) Call() (string, error) {
	return "user logged out", nil
}

func getLogoutUser() *LogoutUser {
	return &LogoutUser{name: "logout_user"}
}
