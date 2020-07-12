import { combineReducers } from "redux";
import { reducer as formReducer } from "redux-form";
import { connectRouter } from "connected-react-router";
import loginReducer from "../containers/LoginPage/reducer";
import registerReducer from "../containers/RegisterPage/reducer";
import UiReducer from "./ui";

export default (history) =>
  combineReducers({
    login: loginReducer,
    register: registerReducer,
    ui: UiReducer,
    form: formReducer,
    router: connectRouter(history),
  });
