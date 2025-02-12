package pipedrive

type PaginationParams struct {
	Limit int `url:"limit,omitempty"`
	Start int `url:"start,omitempty"`
}

func PaginationOpts(start int) PaginationParams {
	return PaginationParams{
		Limit: 500,
		Start: start,
	}
}
