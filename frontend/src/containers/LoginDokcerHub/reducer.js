import * as constLogin from "./constants";
import { toastSuccess } from "../../helpers/toastHelper";

const initialState = {
  username: "",
  password: "",
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case constLogin.LOGIN_DOCKERHUB: {
      return {
        ...state,
      };
    }
    case constLogin.LOGIN_DOCKERHUB_SUCCESS: {
      toastSuccess("Đăng nhập thành công");
      return {
        ...state,
      };
    }
    default:
      return state;
  }
};

export default reducer;
