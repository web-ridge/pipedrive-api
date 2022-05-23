package pipedrive

type PaginationParams struct {
	Limit int `url:"limit,omitempty"`
}

func PaginationOpts(limit int) PaginationParams {
	return PaginationParams{
		Limit: limit,
	}
}
