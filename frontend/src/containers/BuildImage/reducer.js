import * as types from "./constant";
import { toastSuccess, toastError } from "../../helpers/toastHelper";

const initialState = {};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.BUILD_IMAGE: {
      return {
        ...state,
      };
    }
    case types.BUILD_IMAGE_SUCCESS: {
      const { data } = action.payload;
      console.log(data);
      toastSuccess("Build Image thành công");
      return {
        ...state,
      };
    }
    case types.BUILD_IMAGE_FAIL: {
      const { data } = action.payload;
      console.log(data);
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
