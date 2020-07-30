import * as constRegister from "./constants";
import { toastSuccess } from "../../helpers/toastHelper";

const initialState = {};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case constRegister.REGISTER: {
      return {
        ...state,
      };
    }
    case constRegister.REGISTER_SUCCESS: {
      toastSuccess("Register success");
      return {
        ...state,
      };
    }
    default:
      return state;
  }
};

export default reducer;
