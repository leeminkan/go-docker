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
	"go-docker/middleware/role"
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
	///////						Image							////////
	////////////////////////////////////////////////////////////////////

	apiImage := apiv1.Group("/")
	//Get list images
	apiImage.GET("/images", v1.GetImages)
	//Get image
	apiImage.GET("/images/:id", v1.GetImage)
	apiImage.Use(jwt.JWTCustom())
	{
		//Remove image
		apiImage.DELETE("/images/:id", v1.RemoveImage)
		//Tag image
		apiImage.POST("/images/change-tag", v1.ChangeTagImage)
	}

	////////////////////////////////////////////////////////////////////
	///////						End-Image						////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						ImageBuild						////////
	////////////////////////////////////////////////////////////////////

	apiImageBuild := apiv1.Group("/")
	apiImageBuild.Use(jwt.JWTCustom())
	{
		//Get image build by id
		apiImageBuild.GET("/images-build/:id", v1.GetImageBuildByID)
		//Get image build
		apiImageBuild.GET("/images-build", v1.GetImageBuild)
		//Get list image build
		apiImageBuild.GET("/images-list-build", v1.GetListImageBuild)
		//Build image
		apiImageBuild.POST("/images-build/from-docker-file", v1.BuildImageFromDockerFile)
		//Build image
		apiImageBuild.POST("/images-build/from-tar", v1.BuildImageFromTar)
	}

	////////////////////////////////////////////////////////////////////
	///////						End-ImageBuild					////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Image-Push						////////
	////////////////////////////////////////////////////////////////////

	apiImagePush := apiv1.Group("/")
	apiImagePush.Use(jwt.JWTCustom())
	{
		//Get list image push
		apiImagePush.GET("/images-list-push", v1.GetListImagePush)

		//Get image push by id
		apiImagePush.GET("/images-push/:id", v1.GetImagePushByID)

		apiImagePush.Use(docker.CheckLoginDockerHub())
		{
			//Push image
			apiImagePush.POST("/images/push", v1.PushImage)
			//Push image From ID
			apiImagePush.POST("/images-push/from-build-id/:id", v1.PushImageFromID)
		}
	}

	////////////////////////////////////////////////////////////////////
	///////						End-ImagePush					////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Container						////////
	////////////////////////////////////////////////////////////////////

	apiContainer := apiv1.Group("/")
	//Get list container
	apiContainer.GET("/containers", v1.GetContainers)
	//Get container
	apiContainer.GET("/containers/:id", v1.GetContainer)
	//Create a container
	apiContainer.POST("/containers", v1.CreateContainer)
	//Remove container
	apiContainer.DELETE("/containers/:id", v1.RemoveContainer)

	////////////////////////////////////////////////////////////////////
	///////						End-Container					////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Device							////////
	////////////////////////////////////////////////////////////////////

	apiDevice := apiv1.Group("/")

	//Get list devices
	apiDevice.GET("/devices", v1.GetListDevices)
	//Create device
	apiDevice.POST("/devices", v1.CreateDevice)
	//Remove device
	apiDevice.DELETE("/devices/:id", v1.RemoveDevice)
	//Connect device
	apiDevice.POST("/devices/connect", v1.ConnectDevice)
	//Control a device pull image from dockerhub
	apiDevice.POST("/control/devices/pull", v1.ControlDevicePull)

	////////////////////////////////////////////////////////////////////
	///////						End-Device						////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						User							////////
	////////////////////////////////////////////////////////////////////

	apiUser := apiv1.Group("/")
	//Login user
	apiUser.POST("/users/login", v1.Login)

	apiUser.Use(jwt.JWTCustom())
	{
		apiUser.Use(role.IsAdmin())
		{
			//Create user
			apiUser.POST("/users", v1.CreateUser)
		}
		//Login user
		apiUser.GET("/users/mine", v1.GetInfo)

	}

	////////////////////////////////////////////////////////////////////
	///////						End-User						////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Docker							////////
	////////////////////////////////////////////////////////////////////

	apiDocker := apiv1.Group("/")
	apiDocker.Use(jwt.JWTCustom())
	{
		//Login Docker Hub
		apiDocker.POST("/docker/login", v1.LoginDockerHub)
	}

	////////////////////////////////////////////////////////////////////
	///////						End-Docker						////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Repo							////////
	////////////////////////////////////////////////////////////////////

	apiRepo := apiv1.Group("/")
	apiRepo.Use(jwt.JWTCustom())
	{
		//Get repo by id
		apiRepo.GET("/repos/:id", v1.GetRepoByID)
		//Get list repo
		apiRepo.GET("/repos", v1.GetListRepo)
	}

	////////////////////////////////////////////////////////////////////
	///////						End-Repo						////////
	////////////////////////////////////////////////////////////////////

	////////////////////////////////////////////////////////////////////
	///////						Tag							////////
	////////////////////////////////////////////////////////////////////

	apiTag := apiv1.Group("/")
	apiTag.Use(jwt.JWTCustom())
	{
		//Get tag by id
		apiTag.GET("/tags/repo/:id", v1.GetListTagByRepoID)
	}

	////////////////////////////////////////////////////////////////////
	///////						End-Tag						////////
	////////////////////////////////////////////////////////////////////

	return r
}
