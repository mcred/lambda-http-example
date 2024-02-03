package events

type LogoutUser struct{}

func (lu *LogoutUser) Call() (string, error) {
	return "user logged out", nil
}

func getLogoutUser() *LogoutUser {
	return &LogoutUser{}
}
