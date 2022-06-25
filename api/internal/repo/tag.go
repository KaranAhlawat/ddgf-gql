package repo

import u "ddgf-new/internal/util"

func (r *PSQLRepository) GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := r.db.Find(&tags).Limit(50).Error
	u.LogError(err)
	return tags, err
}

func (r *PSQLRepository) GetTag(id string) (*Tag, error) {
	var tag *Tag
	err := r.db.First(&tag, "id = ?", id).Error
	u.LogError(err)
	return tag, err
}

func (r *PSQLRepository) CreateTag(content string) (*Tag, error) {
	var tag *Tag = &Tag{
		Tag: content,
	}
	err := r.db.Save(tag).Error
	u.LogError(err)
	return tag, err
}

func (r *PSQLRepository) DeleteTag(id string) error {
	err := r.db.Delete(&Tag{}, "id = ?", id).Error
	u.LogError(err)
	return err
}

func (r *PSQLRepository) AdvicesForTag(id string) ([]*Advice, error) {
	var tag Tag
	err := r.db.Preload("Advices.Tags").First(&tag, "id = ?", id).Error
	u.LogError(err)
	return tag.Advices, err
}
