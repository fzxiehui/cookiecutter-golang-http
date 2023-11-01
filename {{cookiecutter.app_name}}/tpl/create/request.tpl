package request

/*
 * Basic Create
 */
type Create{[.Name]}Request struct {
	Name string `json:"name" example:"name"` // 姓名
	Sort int    `json:"sort" example:"1"`    // 排序
}

/*
 * Basic Update
 */
type Update{[.Name]}Request struct {
	Name string `json:"name" example:"new name"` // 姓名
	Sort int    `json:"sort" example:"1"`        // 排序
}
