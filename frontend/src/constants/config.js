export const HOST = "http://localhost:8001/api/v1";

export const API_REGISTER = `${HOST}/users`;

export const API_LOGIN = `${HOST}/users/login`;

export const API_BUILD_DOCKERFILE = `${HOST}/images/build-from-docker-file`;

export const API_BUILD_TAR = `${HOST}/images/build-from-tar`;

export const API_LOGIN_DOCKERHUB = `${HOST}/docker/login`;

export const API_GET_LIST_DEVICE = `${HOST}/devices`;

export const API_GET_LIST_LOCAL_IMAGE = `${HOST}/images-list-build`;
