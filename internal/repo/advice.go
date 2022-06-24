package repo

func (r *PSQLRepository) GetAdvices() ([]*Advice, error) {
	var advices []*Advice
	err := r.db.Preload("Tags").Find(&advices).Limit(50).Error
	return advices, err
}

func (r *PSQLRepository) GetAdvice(id string) (*Advice, error) {
	var advice *Advice
	err := r.db.Preload("Tags").First(&advice, "id = ?", id).Error
	return advice, err
}

func (r *PSQLRepository) CreateAdvice(content string) (*Advice, error) {
	var advice *Advice = &Advice{
		Content: content,
	}
	err := r.db.Save(&advice).Error
	return advice, err
}

func (r *PSQLRepository) DeleteAdvice(id string) error {
	err := r.db.Delete(&Advice{}, "id = ?", id).Error
	return err
}
