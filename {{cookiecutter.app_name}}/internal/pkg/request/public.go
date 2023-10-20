package request

/*
 * 公用列表查询
 * Page 查询第几⻚
 * page_size 每⻚有几条数据
 * columns 查询时关联的字段
 * sort 按哪个列进行排序 在后面加上 desc 为倒序
 * query 查询关键字
 */
type PublicQueryListRequest struct {
	Page     int      `form:"page" json:"page"`
	PageSize int      `form:"page_size" json:"page_size"`
	Columns  []string `form:"columns" json:"columns"`
	Sort     string   `form:"sort" json:"sort"`
	Query    string   `form:"query" json:"query"`
}
