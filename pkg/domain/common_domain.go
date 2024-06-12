package domain

import "strconv"

type PagingPointer struct {
	Limit  int     `json:"limit"`
	LastID *string `json:"last_id"`
}

func GetPagingDTO(
	queryParams map[string][]string,
	defaultLimit int,
) PagingPointer {
	response := PagingPointer{
		Limit: defaultLimit,
	}
	if queryParams["limit"] != nil {
		limit := queryParams["limit"][0]
		if limit != "" {
			val, err := strconv.Atoi(limit)
			if err != nil {
				response.Limit = val
			}
		}
	}

	if queryParams["last_id"] != nil {
		lastId := queryParams["last_id"][0]
		if lastId != "" {
			response.LastID = &lastId
		}
	}
	return response
}
