// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserRoles string

const (
	UserRolesAuthor UserRoles = "author"
	UserRolesReader UserRoles = "reader"
)

func (e *UserRoles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRoles(s)
	case string:
		*e = UserRoles(s)
	default:
		return fmt.Errorf("unsupported scan type for UserRoles: %T", src)
	}
	return nil
}

func AllUserRolesValues() []UserRoles {
	return []UserRoles{
		UserRolesAuthor,
		UserRolesReader,
	}
}

type Blog struct {
	ID           uuid.UUID      `json:"id"`
	Descriptions sql.NullString `json:"descriptions"`
	UserRole     UserRoles      `json:"userRole"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    sql.NullTime   `json:"updatedAt"`
}
