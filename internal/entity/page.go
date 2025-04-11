package entity

type PageReqInfo struct {
	Page  int64       `json:"page" db:"page"`   // 頁碼
	Limit int64       `json:"size" db:"size"`   // 每頁數量
	Sort  []SortOrder `json:"sort" db:"sort"`   // 排序
	Index int64       `json:"index" db:"index"` // 查詢索引
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

type SortOrder struct {
	Property  string        `json:"property"`  // 排序欄位
	Direction SortDirection `json:"direction"` // 排序方向(ASC/DESC)
}

type PageRespInfo struct {
	Page       int64
	Limit      int64
	Total      int64
	TotalPages int64
}
