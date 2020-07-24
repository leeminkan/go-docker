package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "go-docker/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"go-docker/middleware/cors"
	"go-docker/middleware/docker"
	"go-docker/middleware/jwt"
	"go-docker/pkg/export"
	"go-docker/pkg/qrcode"
	"go-docker/pkg/upload"
	"go-docker/routers/api"
	v1 "go-docker/routers/api/v1"

	"github.com/gin-gonic/contrib/static"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORSMiddleware())

	r.Use(static.Serve("/", static.LocalFile("./runtime/public", true)))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	////////////////////////////////////////////////////////////////////
	///////						Start							////////
	////////////////////////////////////////////////////////////////////

	//Get list images
	apiv1.GET("/images", v1.GetImages)
	//Get image
	apiv1.GET("/images/:id", v1.GetImage)
	//Get list container
	apiv1.GET("/containers", v1.GetContainers)
	//Get container
	apiv1.GET("/containers/:id", v1.GetContainer)
	//Create a container
	apiv1.POST("/containers", v1.CreateContainer)
	//Remove container
	apiv1.DELETE("/containers/:id", v1.RemoveContainer)
	//Get list devices
	apiv1.GET("/devices", v1.GetListDevices)
	//Create device
	apiv1.POST("/devices", v1.CreateDevice)
	//Remove device
	apiv1.DELETE("/devices/:id", v1.RemoveDevice)
	//Connect device
	apiv1.POST("/devices/connect", v1.ConnectDevice)
	//Control a device pull image from dockerhub
	apiv1.POST("/control/devices/pull", v1.ControlDevicePull)
	//Create user
	apiv1.POST("/users", v1.CreateUser)
	//Login user
	apiv1.POST("/users/login", v1.Login)

	apiv1.Use(jwt.JWTCustom())
	{
		//Login user
		apiv1.GET("/users/mine", v1.GetInfo)
		//Login Docker Hub
		apiv1.POST("/docker/login", v1.LoginDockerHub)
		//Remove image
		apiv1.DELETE("/images/:id", v1.RemoveImage)
		//Build image
		apiv1.POST("/images/build-from-docker-file", v1.BuildImageFromDockerFile)
		//Build image
		apiv1.POST("/images/build-from-tar", v1.BuildImageFromTar)
		//Tag image
		apiv1.POST("/images/change-tag", v1.ChangeTagImage)
		//Get image build by id
		apiv1.GET("/images-build/:id", v1.GetImageBuildByID)
		//Get image build
		apiv1.GET("/images-build", v1.GetImageBuild)
		//Get list image build
		apiv1.GET("/images-list-build", v1.GetListImageBuild)
		//Get list image push
		apiv1.GET("/images-list-push", v1.GetListImagePush)

		apiv1.Use(docker.CheckLoginDockerHub())
		{
			//Push image
			apiv1.POST("/images/push", v1.PushImage)
		}
	}

	////////////////////////////////////////////////////////////////////
	///////						End								////////
	////////////////////////////////////////////////////////////////////

	return r
}
