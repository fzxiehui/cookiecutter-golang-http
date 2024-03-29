package responses

/*
 * 公用列表查询响应
 * List interface 数据
 * Total 总条数
 */
type PublicQueryListResponses struct {
	List  interface{} `json:"List"`  // 查询列表
	Total int64       `json:"Total"` // 总条数
}
