package models

// AuthToken ..
type AuthToken struct {
	Token string
}

// ChangePass ..
type ChangePass struct {
	Token       string
	Password    string
	NewPassword string
}

// SessionData ..
type SessionData struct {
	ID     uint
	TypeID uint
}
