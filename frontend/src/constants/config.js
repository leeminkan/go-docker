// HOST
export const HOST = "http://localhost:8001/api/v1";

// USER
export const API_REGISTER = `${HOST}/users`;
export const API_LOGIN = `${HOST}/users/login`;
export const API_USER = `${HOST}/users/mine`;

// IMAGE BUILD
export const API_BUILD_DOCKERFILE = `${HOST}/images-build/from-docker-file`;
export const API_BUILD_TAR = `${HOST}/images-build/from-tar`;

export const API_GET_LIST_LOCAL_IMAGE = `${HOST}/images-list-build`;
export const API_GET_LOCAL_IMAGE_BY_ID = `${HOST}/images-build`;

// IMAGE PUSH
export const API_GET_LIST_PUSH = `${HOST}/images-list-push`;
export const API_GET_IMAGE_PUSH_BY_ID = `${HOST}/images-push`;
export const API_PUSH_IMAGE = `${HOST}/images-push/from-build-id`;

// DOCKER HUB
export const API_LOGIN_DOCKERHUB = `${HOST}/docker/login`;

// DEVICE
export const API_GET_LIST_DEVICE = `${HOST}/devices`;

// DEVICE DETAIL
export const API_GET_LIST_IMAGE_IN_DEVICE = `${HOST}/device/images`;
export const API_PULL_IMAGE = `${HOST}/control/devices/pull`;
export const API_GET_IMAGE_DEVICE_BY_ID = `${HOST}/device/image`;

export const API_RUN_IMAGE_IN_DEVICE = `${HOST}/images-list-build`;

export const API_GET_LIST_CONTAINER_IN_DEVICE = `${HOST}/device/containers`;
export const API_STOP_CONTAINER_IN_DEVICE = `${HOST}/images-list-build`;
export const API_GET_CONTAINER_IN_DEVICE_BY_ID = `${HOST}/images-push`;
