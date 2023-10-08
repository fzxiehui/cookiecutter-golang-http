package request

/*
 * Basic Create
 */
type Create{[.Name]}Request struct {
	Name string `json:"name"`
	Sort int    `json:"sort"`
}


/*
 * Basic Update
 */
type Update{[.Name]}Request struct {
	Name string `json:"name"`
	Sort int    `json:"sort"`
}
