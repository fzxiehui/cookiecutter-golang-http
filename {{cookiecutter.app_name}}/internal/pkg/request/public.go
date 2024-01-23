package request

/*
 * columns 查询时关联的字段
 * query 查询关键字
 */
type PublicQueryColumnsRequest struct {
	Field      string `form:"field" json:"field" example:"name"`                                                         // 字段名
	Query      string `form:"query" json:"query" example:"hello"`                                                        // 查询关键字
	Exp        string `form:"exp" json:"exp"  binding:"omitempty,oneof=and or not" example:"and"`                        // 表达式
	Conditions string `form:"conditions" json:"conditions" binding:"omitempty,oneof=<> = IN LIKE > < >= <=" example:"="` // 过滤条件: <> = < > <= >= IN LIKE
}

/*
 * 公用列表查询
 * Page 查询第几⻚
 * page_size 每⻚有几条数据
 * sort 按哪个列进行排序 在后面加上 desc 为倒序
 */
type PublicQueryListRequest struct {
	Page     int                         `form:"page" json:"page" example:"1"`            // 查询起始⻚
	PageSize int                         `form:"page_size" json:"page_size" example:"10"` // 每⻚有几条数据
	Columns  []PublicQueryColumnsRequest `form:"columns" json:"columns"`                  // 查询条件
	Sort     string                      `form:"sort" json:"sort" example:"sort"`         // 按哪个列进行排序 尾部加desc 为倒序
}
