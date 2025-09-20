package common

type CursorPageResponse struct {
	Previous *string
	Next     *string
	Total    int
}

type CursorPageRequest struct {
	Size   int
	Before *string
	After  *string
}
