package image

type OptionsBuildImage struct {
	Tags []string `form:"tags"`
}

type OptionsPushImage struct {
	RegistryAuth string `form:"registryAuth"` // RegistryAuth is the base64 encoded credentials for the registry
}

type InputPushImage struct {
	Image string `form:"image"`
}

type InputGetImageBuild struct {
	ImageID string `form:"imageId"`
	Tag     string `form:"tag"`
}

type InputTagImage struct {
	Image string `form:"image"`
	Tag   string `form:"tag"`
}
