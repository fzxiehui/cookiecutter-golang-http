package handler

import (
	"net/http"
	"strconv"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/pkg/request"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/service"
	"github.com/gin-gonic/gin"
)

/*
 1: 在 cmd migration_wire.go 中添加:
 model.{[.Name]}{},

 2: 在 cmd statr_wire.go 中添加:
 handler.New{[.Name]}Handler,
 service.New{[.Name]}Service,
 repository.New{[.Name]}Repository,

 3: 添加路由

 v1.POST("/{[.LowerName]}", {[.LowerName]}Handler.Create)
 v1.GET("/{[.LowerName]}/:id", {[.LowerName]}Handler.Get)
 v1.PATCH("/{[.LowerName]}/:id", {[.LowerName]}Handler.Update)
 v1.DELETE("/{[.LowerName]}/:id", {[.LowerName]}Handler.Delete)
 v1.POST("/{[.LowerName]}/query", {[.LowerName]}Handler.Query)
 */

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
// @Tags {[.Name]} 
// @Summary 标准方法:创建({[.Name]})
// @Description  标准方法:创建({[.Name]})
// @Produce application/json
// @Param req body request.Create{[.Name]}Request true "object"
// @Success 200 {object} model.{[.Name]} "{[.Name]}"
// @Router /v1/{[.LowerName]} [post]
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
// @Tags {[.Name]}
// @Summary 标准方法: 查询指定ID({[.Name]})
// @Description  标准方法: 查询指定ID({[.Name]})
// @Produce application/json
// @Param id path int true "{[.Name]}.ID"
// @Success 200 {object} model.{[.Name]} "{[.Name]}"
// @Router /v1/{[.LowerName]}/{id} [get]
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
// @Tags {[.Name]}
// @Summary 标准方法:修改({[.Name]})
// @Description  标准方法:修改({[.Name]})
// @Produce application/json
// @Param id path int true "{[.Name]}.ID"
// @Param req body request.Update{[.Name]}Request true "object"
// @Success 200 {object} model.{[.Name]} "{[.Name]}"
// @Router /v1/{[.LowerName]}/{id} [patch]
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
// @Tags {[.Name]}
// @Summary 标准方法:删除({[.Name]})
// @Description  标准方法:删除({[.Name]})
// @Produce application/json
// @Param id path int true "{[.Name]}.ID"
// @Success 200 {string} success "{[.Name]}"
// @Router /v1/{[.LowerName]}/{id} [delete]
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
// @Tags {[.Name]}
// @Summary 标准方法:高级查询({[.Name]})
// @Description 标准方法:高级查询({[.Name]})
// @Produce application/json
// @Param req body request.PublicQueryListRequest true "object"
// @Success 200 {object} responses.PublicQueryListResponses "{[.Name]}"
// @Router /v1/{[.LowerName]}/query [post]
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
