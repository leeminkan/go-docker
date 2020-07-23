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
