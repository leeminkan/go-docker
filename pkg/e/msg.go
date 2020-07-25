package e

var MsgFlags = map[int]string{
	SUCCESS:                         "Successful!",
	ERROR:                           "Something was wrong!",
	INVALID_PARAMS:                  "Invalid Params",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:            "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_GET_TAGS_FAIL:             "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:            "统计标签失败",
	ERROR_ADD_TAG_FAIL:              "新增标签失败",
	ERROR_EDIT_TAG_FAIL:             "修改标签失败",
	ERROR_DELETE_TAG_FAIL:           "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:           "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:           "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",

	ERROR_DOCKER_LOGIN_FAIL:                "Login fail!",
	ERROR_TYPE_MUST_BE_TAR:                 "File type must be tar!",
	ERROR_GET_LIST_IMAGE:                   "Sorry! Cann't get list image in this time",
	ERROR_GET_LIST_CONTAINER:               "Sorry! Cann't get list container in this time",
	ERROR_CREATE_USER_FAIL:                 "Create user fail!",
	ERROR_EXIST_USER_FAIL:                  "The username is already existed",
	ERROR_MATCH_CONFIRM_PASSWORD_USER_FAIL: "Confirm password must match with password",
	ERROR_USER_LOGIN_FAIL:                  "Login fail!",
	ERROR_CREATE_DEVICE_FAIL:               "Create device fail!",
	ERROR_EXIST_DEVICE_FAIL:                "Check device exists fail!",
	ERROR_NOT_EXIST_DEVICE:                 "Device not exists!",
	ERROR_GET_LIST_DEVICE:                  "Cannot get list image!",
	ERROR_DELETE_DEVICE:                    "Cannot delete image!",
	ERROR_EXIST_DEVICE:                     "Device already exist!",
	ERROR_CONNECT_DEVICE_FAIL:              "Device connect fail!",
	ERROR_EXIST_DEVICE_CONTROL_FAIL:        "Request with Device donot exists",
	ERROR_SET_MESSAGE_MQTT:                 "Error when set value to control device pull image",
	ERROR_NOT_EXIST_REPONAME_CONTROL:       "Get repo name from repo id fail",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
