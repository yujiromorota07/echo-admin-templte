package entity

import "time"

// AccountID - account id
type AccountID uint32

// Password - account password
type Password string

// Email - account email
type Email string

type Account struct {
	ID        uint32     `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type AccountDetail struct {
	ID        uint32     `db:"id"`
	Name      string     `db:"name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	RoleID    uint32     `db:"role_id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

type AccountRole struct {
	ID        uint32    `db:"id"`
	AccountID uint32    `db:"account_id"`
	RoleID    uint32    `db:"role_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
