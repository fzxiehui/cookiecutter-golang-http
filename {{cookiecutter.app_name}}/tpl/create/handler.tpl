package handler

type {[.Name]}Handler interface {
	/* Basic */
	// c
	Create(ctx *gin.Context)
	// r
	Get(ctx *gin.Context)
	// u
	Update(ctx *gin.Context)
	// d
	Delete(ctx *gin.Context)
	// q
	Query(ctx *gin.Context)
}

func New{[.Name]}Handler(handler *Handler, s service.{[.Name]}Service) {[.Name]}Handler {
	return &{[.LowerName]}Handler{
		Handler:         handler,
		s: s,
	}
}

type {[.LowerName]}Handler struct {
	*Handler
	s service.{[.Name]}Service
}

/* Basic */
// c
func (h *{[.LowerName]}Handler) Create(ctx *gin.Context) {
	req := request.Create{[.Name]}Request{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.s.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}
// r
func (h *{[.LowerName]}Handler) Get(ctx *gin.Context) {
	// get uri id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.s.Get(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}
// u
func (h *{[.LowerName]}Handler) Update(ctx *gin.Context) {
	// get uri id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req := request.Update{[.Name]}Request{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.s.Update(ctx, uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}
// d
func (h *{[.LowerName]}Handler) Delete(ctx *gin.Context) {
	// get uri id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.s.Delete(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "删除成功"})
	return
}
// q
func (h *{[.LowerName]}Handler) Query(ctx *gin.Context) {
	req := &request.PublicQueryListRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := h.s.Query(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, data)
	return
}
