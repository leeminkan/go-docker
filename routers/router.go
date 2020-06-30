package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "go-docker/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"go-docker/middleware/jwt"
	"go-docker/pkg/export"
	"go-docker/pkg/qrcode"
	"go-docker/pkg/upload"
	"go-docker/routers/api"
	v1 "go-docker/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	////////////////////////////////////////////////////////////////////
	///////						Start							////////
	////////////////////////////////////////////////////////////////////

	//Create example image
	apiv1.GET("/images/create-example", v1.CreateImages)
	//Get list images
	apiv1.GET("/images", v1.GetImages)
	//Run example container
	apiv1.GET("/containers/run-example", v1.RunContainers)
	//Get list container
	apiv1.GET("/containers", v1.GetContainers)

	////////////////////////////////////////////////////////////////////
	///////						End								////////
	////////////////////////////////////////////////////////////////////
	apiv1.Use(jwt.JWT())
	{
		////////////////////////////////////////////////////////////////////
		///////						Start							////////
		////////////////////////////////////////////////////////////////////
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
		////////////////////////////////////////////////////////////////////
		///////							End							////////
		////////////////////////////////////////////////////////////////////
	}

	return r
}
