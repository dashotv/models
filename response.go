package models

import "github.com/go-bongo/bongo"

type BaseResponse struct {
	Header *bongo.PaginationInfo `json:"header"`
}

func (b *BaseResponse) Pagination(p *bongo.PaginationInfo) {
	b.Header = p
}
