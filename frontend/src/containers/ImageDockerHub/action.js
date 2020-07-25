import * as types from "./constant";

export const getListDockerHubImage = (data) => {
  return {
    type: types.GET_LIST_DOCKER_HUB_IMAGE,
    payload: {
      data,
    },
  };
};

export const getListDockerHubImageSuccess = (data) => {
  return {
    type: types.GET_LIST_DOCKER_HUB_IMAGE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const getListDockerHubImageFail = (data) => {
  return {
    type: types.GET_LIST_DOCKER_HUB_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};

export const openModalPushImage = () => {
  return {
    type: types.OPEN_MODAL_PUSH_IMAGE,
  };
};

export const closeModalPushImage = () => {
  return {
    type: types.CLOSE_MODAL_PUSH_IMAGE,
  };
};

export const pushImage = (data) => {
  return {
    type: types.PUSH_IMAGE,
    payload: {
      data,
    },
  };
};

export const pushImageSuccess = (data) => {
  return {
    type: types.PUSH_IMAGE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const pushImagePending = (data) => {
  return {
    type: types.PUSH_IMAGE_PENDING,
    payload: {
      data,
    },
  };
};

export const pushImageFail = (data) => {
  return {
    type: types.PUSH_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};

export const getDHImageById = (data) => {
  return {
    type: types.GET_DH_IMAGE_BY_ID,
    payload: {
      data,
    },
  };
};
