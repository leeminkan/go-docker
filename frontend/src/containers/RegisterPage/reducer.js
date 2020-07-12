import * as constRegister from "./constants";
import { toastSuccess } from "../../helpers/toastHelper";

const initialState = {
  username: "",
  password: "",
  rePassword: "",
  isAdmin: false,
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case constRegister.REGISTER: {
      return {
        ...state,
      };
    }
    case constRegister.REGISTER_SUCCESS: {
      toastSuccess("Đăng ký thành công");
      return {
        ...state,
      };
    }
    default:
      return state;
  }
};

export default reducer;
