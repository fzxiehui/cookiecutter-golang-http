package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/service"
	"github.com/gin-gonic/gin"
)

type FilesHandler interface {
	SaveImage(ctx *gin.Context)

	GetImage(ctx *gin.Context)
}

func NewFilesHandler(handler *Handler, filesService service.FilesService) FilesHandler {
	return &filesHandler{
		Handler:      handler,
		filesService: filesService,
	}
}

type filesHandler struct {
	*Handler
	filesService service.FilesService
}

func (h *filesHandler) GetImage(ctx *gin.Context) {
	bulkName := ctx.Param("bulk")
	if len(bulkName) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "路径错误!"})
		return
	}
	img_user_id, err := strconv.Atoi(ctx.Param("uid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "路径错误!"})
		return
	}
	file_name := ctx.Param("name")
	fileName := fmt.Sprintf("image/%d/%s", img_user_id, file_name)
	data, err := h.filesService.GetImage(ctx, bulkName, fileName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ctx.JSON(http.StatusOK, data)
	ctx.Writer.Write(data)
	return

}

func (h *filesHandler) SaveImage(ctx *gin.Context) {
	// TODO // 如果需要分用户存放 把以下注释打开 并使用 jwt 中间件
	// id := GetUserIdFromCtx(ctx)
	// if id == 0 {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "非法请求!"})
	// 	return
	// }
	var id uint = 1
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName, err := h.filesService.SaveImage(ctx, id, file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"file_name": fileName})
	return
}
