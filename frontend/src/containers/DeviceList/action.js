import * as types from "./constant";

export const getListDevice = (data) => {
  return {
    type: types.GET_LIST_DEVICE,
    payload: {
      data,
    },
  };
};

export const getListDeviceSuccess = (data) => {
  return {
    type: types.GET_LIST_DEVICE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const getListDeviceFail = (data) => {
  return {
    type: types.GET_LIST_DEVICE_FAIL,
    payload: {
      data,
    },
  };
};
