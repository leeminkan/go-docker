import * as types from "./constant";
import { toastError } from "../../helpers/toastHelper";

const initialState = {
  listLocalImage: [],
  openModalBuildImage: false,
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.GET_LIST_LOCAL_IMAGE: {
      return {
        ...state,
        listLocalImage: [],
      };
    }
    case types.GET_LIST_LOCAL_IMAGE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        listLocalImage: data,
      };
    }
    case types.GET_LIST_LOCAL_IMAGE_FAIL: {
      return {
        ...state,
      };
    }
    case types.OPEN_MODAL_BUILD_IMAGE: {
      return {
        ...state,
        openModalBuildImage: true,
      };
    }
    case types.CLOSE_MODAL_BUILD_IMAGE: {
      return {
        ...state,
        openModalBuildImage: false,
      };
    }
    case types.BUILD_IMAGE: {
      return {
        ...state,
      };
    }
    case types.BUILD_IMAGE_SUCCESS: {
      let newImage = action.payload.data.data;
      const { listLocalImage } = state;
      const index = listLocalImage.findIndex((item) => item.id === newImage.id);
      if (index !== -1) {
        const newList = [
          ...listLocalImage.slice(0, index),
          newImage,
          ...listLocalImage.slice(index + 1),
        ];
        return {
          ...state,
          listLocalImage: newList,
        };
      }
      return {
        ...state,
      };
    }
    case types.BUILD_IMAGE_PENDING: {
      let newImage = action.payload.data.data;
      let dataNewImage = [newImage].concat(state.listLocalImage);
      return {
        ...state,
        listLocalImage: dataNewImage,
      };
    }
    case types.BUILD_IMAGE_FAIL: {
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
