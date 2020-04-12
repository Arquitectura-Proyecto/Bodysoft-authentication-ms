package models

// AuthToken ..
type AuthToken struct {
	Token string
}

// ChangePass ..
type ChangePass struct {
	ID          uint
	Password    string
	NewPassword string
}

// SessionData ..
type SessionData struct {
	ID     float64
	TypeID float64
}
