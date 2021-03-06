package repo

type Repository interface {
	// Read operations
	GetPages() ([]*Page, error)
	GetPage(id string) (*Page, error)
	GetAdvices() ([]*Advice, error)
	GetAdvice(id string) (*Advice, error)
	GetTags() ([]*Tag, error)
	GetTag(id string) (*Tag, error)
	AdvicesForTag(id string) ([]*Advice, error)

	// Write operations
	CreatePage(content string) (*Page, error)
	CreateAdvice(content string) (*Advice, error)
	CreateTag(content string) (*Tag, error)
	DeletePage(id string) error
	DeleteAdvice(id string) error
	DeleteTag(id string) error
	TagAdvice(tid string, aid string) (*Advice, error)
	UntagAdvice(tid string, aid string) (*Advice, error)
}
