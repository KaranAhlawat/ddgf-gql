package repo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New()
	return nil
}

type Page struct {
	Content string `gorm:"not null"`
	Base
}

type Advice struct {
	Content string `gorm:"not null"`
	Tags    []*Tag `gorm:"many2many:advice_tag;constraint:OnDelete:CASCADE"`
	Base
}

type Tag struct {
	Tag     string    `gorm:"not null;index;unique"`
	Advices []*Advice `gorm:"many2many:advice_tag;constraint:OnDelete:CASCADE"`
	Base
}
