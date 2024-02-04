package events

// LogoutUser is an event for logging out a user
type LogoutUser struct{}

// Call logs out a user
func (lu *LogoutUser) Call() (string, error) {
	return "user logged out", nil
}

func newLogoutUser() *LogoutUser {
	return &LogoutUser{}
}
