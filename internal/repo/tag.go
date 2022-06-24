package repo

func (r *PSQLRepository) GetTags() ([]*Tag, error) {
	var tags []*Tag
	err := r.db.Preload("Advices").Find(&tags).Limit(50).Error
	return tags, err
}

func (r *PSQLRepository) GetTag(id string) (*Tag, error) {
	var tag *Tag
	err := r.db.Preload("Advices").First(&tag, "id = ?", id).Error
	return tag, err
}

func (r *PSQLRepository) CreateTag(content string) (*Tag, error) {
	var tag *Tag = &Tag{
		Tag: content,
	}
	err := r.db.Save(&tag).Error
	return tag, err
}

func (r *PSQLRepository) DeleteTag(id string) error {
	err := r.db.Delete(&Tag{}, "id = ?", id).Error
	return err
}
