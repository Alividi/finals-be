package helper

import (
	"net/http"
)

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

func NewPagination(page, pageSize, total int) Pagination {
	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}
	return Pagination{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: totalPage,
		TotalData: total,
	}
}

type PaginationRequest struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Keyword   string `json:"keyword"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func (p *PaginationRequest) ReadFromQuery(r *http.Request) {
	p.Page = GetQueryInt(r, "page", 1)
	p.PageSize = GetQueryInt(r, "page_size", 10)
	p.Keyword = GetQueryString(r, "keyword", "")
	p.OrderBy = GetQueryString(r, "order_by", "")
	p.OrderType = GetQueryString(r, "order_type", "")
	p.StartDate = GetQueryString(r, "start_date", "")
	p.EndDate = GetQueryString(r, "end_date", "")
}

func (p *PaginationRequest) Offset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationRequest) Limit() int {
	return p.PageSize
}
