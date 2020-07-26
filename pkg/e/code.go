package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_NOT_EXIST_ARTICLE        = 10011
	ERROR_CHECK_EXIST_ARTICLE_FAIL = 10012
	ERROR_ADD_ARTICLE_FAIL         = 10013
	ERROR_DELETE_ARTICLE_FAIL      = 10014
	ERROR_EDIT_ARTICLE_FAIL        = 10015
	ERROR_COUNT_ARTICLE_FAIL       = 10016
	ERROR_GET_ARTICLES_FAIL        = 10017
	ERROR_GET_ARTICLE_FAIL         = 10018
	ERROR_GEN_ARTICLE_POSTER_FAIL  = 10019

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003

	ERROR_DOCKER_LOGIN_FAIL                = 1001
	ERROR_TYPE_MUST_BE_TAR                 = 1002
	ERROR_GET_LIST_IMAGE                   = 1100
	ERROR_GET_LIST_CONTAINER               = 1110
	ERROR_CREATE_USER_FAIL                 = 1510
	ERROR_EXIST_USER_FAIL                  = 1511
	ERROR_MATCH_CONFIRM_PASSWORD_USER_FAIL = 1512
	ERROR_USER_LOGIN_FAIL                  = 1612
	ERROR_CREATE_DEVICE_FAIL               = 1710
	ERROR_EXIST_DEVICE_FAIL                = 1711
	ERROR_NOT_EXIST_DEVICE                 = 1712
	ERROR_GET_LIST_DEVICE                  = 1713
	ERROR_DELETE_DEVICE                    = 1714
	ERROR_EXIST_DEVICE                     = 1715
	ERROR_CONNECT_DEVICE_FAIL              = 1716
	ERROR_EXIST_DEVICE_CONTROL_FAIL        = 1717
	ERROR_NOT_EXIST_REPONAME_CONTROL       = 1718

	ERROR_IMAGE_BUILD_NOT_FOUND    = 1810
	ERROR_IMAGE_BUILD_FILE_INVALID = 1811
	ERROR_IMAGE_BUILD_PARSE_QUERY  = 1812
	ERROR_BUILD_IMAGE_FAIL         = 1813
	ERROR_TAG_IMAGE                = 1820
	ERROR_PUSH_IMAGE_FAIL          = 1910
	ERROR_PUSH_IMAGE_EXISTED       = 1911
	ERROR_GET_LIST_IMAGE_PUSH      = 1912
	ERROR_GET_IMAGE_PUSH_BY_ID     = 1913
	ERROR_REPO_CREATE              = 2010
	ERROR_GET_LIST_REPO            = 2011
	ERROR_REPO_NOT_FOUND           = 2012
	ERROR_TAG_CREATE               = 2020
	ERROR_GET_LIST_TAG_BY_REPO_ID  = 2021

	ERROR_SET_MESSAGE_MQTT = 40001
)
