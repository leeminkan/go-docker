import * as types from "./constant";
import { toastError } from "../../helpers/toastHelper";

const initialState = {
  imageInDevice: [],
  openModalPullImage: false,
  openModalRunImage: false,
  containerInDevice: [],
  runID: "",
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.OPEN_MODAL_PULL_IMAGE: {
      return {
        ...state,
        openModalPullImage: true,
      };
    }
    case types.CLOSE_MODAL_PULL_IMAGE: {
      return {
        ...state,
        openModalPullImage: false,
      };
    }
    case types.OPEN_MODAL_RUN_IMAGE: {
      const { data } = action.payload;
      return {
        ...state,
        openModalRunImage: true,
        runID: data,
      };
    }
    case types.CLOSE_MODAL_RUN_IMAGE: {
      return {
        ...state,
        openModalRunImage: false,
        runID: "",
      };
    }
    case types.GET_LIST_IMAGE_IN_DEVICE: {
      return {
        ...state,
        imageInDevice: [],
      };
    }
    case types.GET_LIST_IMAGE_IN_DEVICE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        imageInDevice: data,
      };
    }
    case types.GET_LIST_IMAGE_IN_DEVICE_FAIL: {
      return {
        ...state,
      };
    }
    case types.GET_LIST_CONTAINER_IN_DEVICE: {
      return {
        ...state,
        containerInDevice: [],
      };
    }
    case types.GET_LIST_CONTAINER_IN_DEVICE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        containerInDevice: data,
      };
    }
    case types.GET_LIST_CONTAINER_IN_DEVICE_FAIL: {
      return {
        ...state,
      };
    }
    case types.PULL_IMAGE: {
      return {
        ...state,
      };
    }
    case types.PULL_IMAGE_SUCCESS: {
      let newImage = action.payload.data.data;
      const { imageInDevice } = state;
      const index = imageInDevice.findIndex((item) => item.id === newImage.id);
      if (index !== -1) {
        const newList = [
          ...imageInDevice.slice(0, index),
          newImage,
          ...imageInDevice.slice(index + 1),
        ];
        return {
          ...state,
          imageInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.PULL_IMAGE_PENDING: {
      let newImage = action.payload.data.data;
      let dataNewImage = [newImage].concat(state.imageInDevice);
      return {
        ...state,
        imageInDevice: dataNewImage,
      };
    }
    case types.PULL_IMAGE_FAIL: {
      const { data } = action.payload;
      toastError(data);
      return {
        ...state,
      };
    }
    case types.RUN_IMAGE_DEVICE: {
      return {
        ...state,
      };
    }
    case types.RUN_IMAGE_DEVICE_SUCCESS: {
      let newContainer = action.payload.data.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === newContainer.id
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          newContainer,
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.RUN_IMAGE_DEVICE_PENDING: {
      let newContainer = action.payload.data.data;
      let dataNewContainer = [newContainer].concat(state.containerInDevice);
      return {
        ...state,
        containerInDevice: dataNewContainer,
      };
    }
    case types.RUN_IMAGE_DEVICE_FAIL: {
      const { data } = action.payload;
      console.log(data);
      toastError(data);
      return {
        ...state,
      };
    }

    case types.START_CONTAINER: {
      return {
        ...state,
      };
    }
    case types.START_CONTAINER_SUCCESS: {
      let startContainer = action.payload.data.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === startContainer.id
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          startContainer,
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.START_CONTAINER_PENDING: {
      let startContainer = action.payload.data.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === startContainer.id
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          startContainer,
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.START_CONTAINER_FAIL: {
      const { data } = action.payload;
      toastError(data);
      return {
        ...state,
      };
    }

    case types.STOP_CONTAINER: {
      return {
        ...state,
      };
    }
    case types.STOP_CONTAINER_SUCCESS: {
      let stopContainer = action.payload.data.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === stopContainer.id
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          stopContainer,
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.STOP_CONTAINER_PENDING: {
      let stopContainer = action.payload.data.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === stopContainer.id
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          stopContainer,
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.STOP_CONTAINER_FAIL: {
      const { data } = action.payload;
      toastError(data);
      return {
        ...state,
      };
    }

    case types.DELETE_CONTAINER: {
      return {
        ...state,
      };
    }
    case types.DELETE_CONTAINER_SUCCESS: {
      let deleteContainer = action.payload.data;
      const { containerInDevice } = state;
      const index = containerInDevice.findIndex(
        (item) => item.id === deleteContainer
      );
      if (index !== -1) {
        const newList = [
          ...containerInDevice.slice(0, index),
          ...containerInDevice.slice(index + 1),
        ];
        return {
          ...state,
          containerInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }

    case types.DELETE_CONTAINER_FAIL: {
      const { data } = action.payload;
      toastError(data);
      return {
        ...state,
      };
    }

    case types.DELETE_IMAGE: {
      return {
        ...state,
      };
    }
    case types.DELETE_IMAGE_SUCCESS: {
      let deleteImage = action.payload.data;
      const { imageInDevice } = state;
      const index = imageInDevice.findIndex((item) => item.id === deleteImage);
      if (index !== -1) {
        const newList = [
          ...imageInDevice.slice(0, index),
          ...imageInDevice.slice(index + 1),
        ];
        return {
          ...state,
          imageInDevice: newList,
        };
      }
      return {
        ...state,
      };
    }

    case types.DELETE_IMAGE_FAIL: {
      //const { data } = action.payload;
      toastError(
        "Please stop and delete container built by image before delete image"
      );
      return {
        ...state,
      };
    }

    default:
      return state;
  }
};

export default reducer;
