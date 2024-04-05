package helpers

type Seo struct {
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
}

func (pagination Pagination) SetSeo() Seo {
	seoDesc := Seo{
		Title:       "title",
		Keyword:     "keyword",
		Description: "description",
	}

	return seoDesc
}
