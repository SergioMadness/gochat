package config

import "socialchat/models"

type Session struct {
	currentUser *models.Profile
}

var currentSession *Session

func GetSession() *Session {
	if currentSession == nil {
		currentSession = new(Session)
	}
	return currentSession
}

func (s *Session) SetCurrentUser(profile *models.Profile) {
	s.currentUser = profile
}

func (s *Session) GetCurrentUser() *models.Profile {
	return s.currentUser
}

func (s *Session) IsLoggedIn() bool {
	return s.currentUser != nil
}
