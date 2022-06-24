package repo

import (
	"ddgf-new/internal/model"
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
	if b.ID != uuid.Nil {
		return nil
	}
	b.ID = uuid.New()
	return nil
}

type Page struct {
	Content string `gorm:"not null"`
	Base
}

func (p *Page) ToModel() *model.Page {
	return &model.Page{
		ID:      p.ID.String(),
		Time:    p.CreatedAt,
		Content: p.Content,
	}
}

type Advice struct {
	Content string `gorm:"not null"`
	Tags    []*Tag `gorm:"many2many:advice_tag;constraint:OnDelete:CASCADE"`
	Base
}

func (a *Advice) ToModel() *model.Advice {
	tags := make([]*model.Tag, 0)
	for _, tag := range a.Tags {
		tags = append(tags, tag.ToModel())
	}
	return &model.Advice{
		ID:      a.ID.String(),
		Content: a.Content,
		Tags:    tags,
	}
}

type Tag struct {
	Tag     string    `gorm:"not null;index;unique"`
	Advices []*Advice `gorm:"many2many:advice_tag;constraint:OnDelete:CASCADE"`
	Base
}

func (t *Tag) ToModel() *model.Tag {
	return &model.Tag{
		ID:  t.ID.String(),
		Tag: t.Tag,
	}
}
