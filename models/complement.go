package models

// Auth ..
type Auth struct {
	Email    string
	Password string
}

// TokenData ..
type TokenData struct {
	ID     uint
	TypeID uint
}

// ChangePass ..
type ChangePass struct {
	ID          uint
	Password    string
	NewPassword string
}
