import * as types from "./constant";

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

export const buildImageFail = (data) => {
  return {
    type: types.BUILD_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};
