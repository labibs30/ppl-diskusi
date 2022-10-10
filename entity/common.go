package entity

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type UUIDModel struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;not null" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

func (u *UUIDModel) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	if u.ID == uuid.Nil {
		return errors.New("can't generate uuid")
	}

	return nil
}

type CommonModel struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// SliceOfString defined []string data type in Go, need to implements driver.Valuer and sql.Scanner interface
// because gorm doesn't know what type of []string in database
type SliceOfString []string

// Scan will scan value into []string, implements sql.Scanner interface
// From DB to us
func (s *SliceOfString) Scan(src any) error {
	if str, ok := src.(string); ok {
		*s = strings.Split(str, ",")
		return nil
	}

	return errors.New("failed to scan SliceOfString field - source is not a string")
}

// Value will return []string value, implements driver.Valuer interface
// From us to DB
func (s SliceOfString) Value() (driver.Value, error) {
	if s == nil || len(s) <= 0 {
		println("slice of string value nya kosong")
		return nil, nil
	}

	return strings.Join(s, ","), nil
}

// GormDataType is the gorm common data type
func (SliceOfString) GormDataType() string {
	return "text"
}

// GormDBDataType is the gorm db data type
func (SliceOfString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "postgres":
		return "text"
	}

	return ""
}
