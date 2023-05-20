package helpers

type Meta struct {
	Pagination Pagination `json:"pagination"`
	// Seo        Seo        `json:"seo"`
}

func (meta Meta) SetMeta(pagination Pagination, seo Seo) Meta {
	meta = Meta{
		Pagination: pagination,
		// Seo:        seo,
	}
	return meta
}
