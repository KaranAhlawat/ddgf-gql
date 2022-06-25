package repo

import u "ddgf-new/internal/util"

func (r *PSQLRepository) GetAdvices() ([]*Advice, error) {
	var advices []*Advice
	err := r.db.Preload("Tags").Find(&advices).Limit(50).Error
	u.LogError(err)
	return advices, err
}

func (r *PSQLRepository) GetAdvice(id string) (*Advice, error) {
	var advice *Advice
	err := r.db.Preload("Tags").First(&advice, "id = ?", id).Error
	u.LogError(err)
	return advice, err
}

func (r *PSQLRepository) CreateAdvice(content string) (*Advice, error) {
	var advice *Advice = &Advice{
		Content: content,
	}
	err := r.db.Save(advice).Error
	u.LogError(err)
	return advice, err
}

func (r *PSQLRepository) DeleteAdvice(id string) error {
	err := r.db.Delete(&Advice{}, "id = ?", id).Error
	u.LogError(err)
	return err
}

func (r *PSQLRepository) TagAdvice(tid string, aid string) (*Advice, error) {
	advice, err := r.GetAdvice(aid)
	if err != nil {
		u.LogError(err)
		return nil, err
	}

	// Check if the tag is associated
	for _, tag := range advice.Tags {
		if tag.ID.String() == tid {
			return advice, nil
		}
	}

	// If not assiciated, then add the tag
	tag, err := r.GetTag(tid)
	if err != nil {
		u.LogError(err)
		return nil, err
	}

	r.db.Model(advice).Association("Tags").Append(tag)
	return advice, nil
}

func (r *PSQLRepository) UntagAdvice(tid string, aid string) (*Advice, error) {
	tag, err := r.GetTag(tid)
	if err != nil {
		u.LogError(err)
		return nil, err
	}
	advice, err := r.GetAdvice(aid)
	if err != nil {
		u.LogError(err)
		return nil, err
	}
	r.db.Model(advice).Association("Tags").Delete(tag)
	return advice, nil
}
