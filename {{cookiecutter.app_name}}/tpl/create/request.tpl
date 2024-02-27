package request

/*
 * Basic Create
 */
type Create{[.Name]}Request struct {
	Name string `json:"Name" example:"name"` // 姓名
	Sort int    `json:"Sort" example:"1"`    // 排序
}

/*
 * Basic Update
 */
type Update{[.Name]}Request struct {
	Name string `json:"Name" example:"new name"` // 姓名
	Sort int    `json:"Sort" example:"1"`        // 排序
}
