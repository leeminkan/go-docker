import * as types from "./constant";
import { toastSuccess, toastError } from "../../helpers/toastHelper";

const initialState = {
  listDHImage: [],
  openModalPushImage: false,
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.GET_LIST_DOCKER_HUB_IMAGE: {
      return {
        ...state,
        listDHImage: [],
      };
    }
    case types.GET_LIST_DOCKER_HUB_IMAGE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        listDHImage: data,
      };
    }
    case types.GET_LIST_DOCKER_HUB_IMAGE_FAIL: {
      return {
        ...state,
      };
    }
    case types.OPEN_MODAL_PUSH_IMAGE: {
      return {
        ...state,
        openModalPushImage: true,
      };
    }
    case types.CLOSE_MODAL_PUSH_IMAGE: {
      return {
        ...state,
        openModalPushImage: false,
      };
    }
    case types.PUSH_IMAGE: {
      return {
        ...state,
      };
    }
    case types.PUSH_IMAGE_SUCCESS: {
      let newImage = action.payload.data.data;
      const { listDHImage } = state;
      const index = listDHImage.findIndex((item) => item.id === newImage.id);
      if (index !== -1) {
        const newList = [
          ...listDHImage.slice(0, index),
          newImage,
          ...listDHImage.slice(index + 1),
        ];
        return {
          ...state,
          listDHImage: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.PUSH_IMAGE_PENDING: {
      let newImage = action.payload.data.data;
      let dataNewImage = [newImage].concat(state.listDHImage);
      return {
        ...state,
        listDHImage: dataNewImage,
      };
    }
    case types.PUSH_IMAGE_FAIL: {
      const { data } = action.payload;
      toastError(data);
      return {
        ...state,
      };
    }
    default:
      return state;
  }
};

export default reducer;
