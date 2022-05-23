package pipedrive

type PaginationParams struct {
	Limit int
}

func PaginationOpts(limit int) PaginationParams {
	return PaginationParams{
		Limit: limit,
	}
}
