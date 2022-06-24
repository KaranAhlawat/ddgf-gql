package repo

func (r *PSQLRepository) GetPages() ([]*Page, error) {
	var pages []*Page
	err := r.db.Find(&pages).Limit(50).Error
	return pages, err
}

func (r *PSQLRepository) GetPage(id string) (*Page, error) {
	var page *Page
	err := r.db.First(&page, "id = ?", id).Error
	return page, err
}

func (r *PSQLRepository) CreatePage(content string) (*Page, error) {
	var page *Page = &Page{
		Content: content,
	}
	err := r.db.Save(&page).Error
	return page, err
}

func (r *PSQLRepository) DeletePage(id string) error {
	err := r.db.Delete(&Page{}, "id = ?", id).Error
	return err
}
