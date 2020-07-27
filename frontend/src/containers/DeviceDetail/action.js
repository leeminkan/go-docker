import * as types from "./constant";

export const openModalPullImage = () => {
  return {
    type: types.OPEN_MODAL_PULL_IMAGE,
  };
};

export const closeModalPullImage = () => {
  return {
    type: types.CLOSE_MODAL_PULL_IMAGE,
  };
};

export const getListImageInDevice = (data) => {
  return {
    type: types.GET_LIST_IMAGE_IN_DEVICE,
    payload: {
      data,
    },
  };
};

export const getListImageInDeviceSuccess = (data) => {
  return {
    type: types.GET_LIST_IMAGE_IN_DEVICE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const getListImageInDeviceFail = (data) => {
  return {
    type: types.GET_LIST_IMAGE_IN_DEVICE_FAIL,
    payload: {
      data,
    },
  };
};

export const getListContainerInDevice = (data) => {
  return {
    type: types.GET_LIST_CONTAINER_IN_DEVICE,
    payload: {
      data,
    },
  };
};

export const getListContainerInDeviceSuccess = (data) => {
  return {
    type: types.GET_LIST_CONTAINER_IN_DEVICE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const getListContainerInDeviceFail = (data) => {
  return {
    type: types.GET_LIST_CONTAINER_IN_DEVICE_FAIL,
    payload: {
      data,
    },
  };
};

export const pullImage = (data) => {
  return {
    type: types.PULL_IMAGE,
    payload: {
      data,
    },
  };
};

export const pullImageSuccess = (data) => {
  return {
    type: types.PULL_IMAGE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const pullImagePending = (data) => {
  return {
    type: types.PULL_IMAGE_PENDING,
    payload: {
      data,
    },
  };
};

export const pullImageFail = (data) => {
  return {
    type: types.PULL_IMAGE_FAIL,
    payload: {
      data,
    },
  };
};

export const getDeviceImageById = (data) => {
  return {
    type: types.GET_DEVICE_IMAGE_BY_ID,
    payload: {
      data,
    },
  };
};

export const runImageDevice = (data) => {
  return {
    type: types.RUN_IMAGE_DEVICE,
    payload: {
      data,
    },
  };
};

export const runImageDeviceSuccess = (data) => {
  return {
    type: types.RUN_IMAGE_DEVICE_SUCCESS,
    payload: {
      data,
    },
  };
};

export const runImageDevicePending = (data) => {
  return {
    type: types.RUN_IMAGE_DEVICE_PENDING,
    payload: {
      data,
    },
  };
};

export const runImageDeviceFail = (data) => {
  return {
    type: types.RUN_IMAGE_DEVICE_FAIL,
    payload: {
      data,
    },
  };
};

export const getDeviceContainerById = (data) => {
  return {
    type: types.GET_DEVICE_CONTAINER_BY_ID,
    payload: {
      data,
    },
  };
};
