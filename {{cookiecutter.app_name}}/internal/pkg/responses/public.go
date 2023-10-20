package responses

/*
 * 公用列表查询响应
 * List interface 数据
 * Total 总条数
 */
type PublicQueryListResponses struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
}
