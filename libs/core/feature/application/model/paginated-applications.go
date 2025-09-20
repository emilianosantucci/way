package model

import "libs/core/common"

type PaginatedApplications struct {
	Data []Application             `json:"data"`
	Page common.CursorPageResponse `json:"page"`
}
