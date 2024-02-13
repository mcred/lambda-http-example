package events

type logoutUser struct{}

func (lu *logoutUser) call() (string, error) {
	return "user logged out", nil
}

func newLogoutUser() *logoutUser {
	return &logoutUser{}
}
