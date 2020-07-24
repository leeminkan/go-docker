import * as types from "./constant";

export const getListLocalImage = (data) => {
  return {
    type: types.GET_LIST_LOCAL_IMAGE,
    payload: {
      data,
    },
  };
};

export const getListLocalImageSuccess = (data) => {
  return {
    type: types.GET_LIST_LOCAL_IMAGE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const getListLocalImageFail = (data) => {
  return {
    type: types.GET_LIST_LOCAL_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};

export const openModalBuildImage = () => {
  return {
    type: types.OPEN_MODAL_BUILD_IMAGE,
  };
};

export const closeModalBuildImage = () => {
  return {
    type: types.CLOSE_MODAL_BUILD_IMAGE,
  };
};

export const buildImage = (data) => {
  return {
    type: types.BUILD_IMAGE,
    payload: {
      data,
    },
  };
};

export const buildImageSuccess = (data) => {
  return {
    type: types.BUILD_IMAGE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const buildImagePending = (data) => {
  return {
    type: types.BUILD_IMAGE_PENDING,
    payload: {
      data,
    },
  };
};

export const buildImageFail = (data) => {
  return {
    type: types.BUILD_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};

export const getImageById = (data) => {
  return {
    type: types.GET_IMAGE_BY_ID,
    payload: {
      data,
    },
  };
};
