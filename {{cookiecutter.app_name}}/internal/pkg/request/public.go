package request

/*
 * columns 查询时关联的字段
 * query 查询关键字
 */
type PublicQueryColumnsRequest struct {
	Field      string `form:"Field" json:"Field" example:"name"`                                                         // 字段名
	Query      string `form:"Query" json:"Query" example:"hello"`                                                        // 查询关键字
	Exp        string `form:"EXP" json:"EXP"  binding:"omitempty,oneof=and or not" example:"and"`                        // 表达式
	Conditions string `form:"Conditions" json:"Conditions" binding:"omitempty,oneof=<> = IN LIKE > < >= <=" example:"="` // 过滤条件: <> = < > <= >= IN LIKE
}

/*
 * 公用列表查询
 * Page 查询第几⻚
 * page_size 每⻚有几条数据
 * sort 按哪个列进行排序 在后面加上 desc 为倒序
 */
type PublicQueryListRequest struct {
	Page     int                         `form:"Page" json:"Page" example:"1"`          // 查询起始⻚
	PageSize int                         `form:"PageSize" json:"PageSize" example:"10"` // 每⻚有几条数据
	Columns  []PublicQueryColumnsRequest `form:"Columns" json:"Columns"`                // 查询条件
	Sort     string                      `form:"Sort" json:"Sort" example:"sort"`       // 按哪个列进行排序 尾部加desc 为倒序
}
