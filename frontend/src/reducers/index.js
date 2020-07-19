import { combineReducers } from "redux";
import { reducer as formReducer } from "redux-form";
import { connectRouter } from "connected-react-router";
import loginReducer from "../containers/LoginPage/reducer";
import registerReducer from "../containers/RegisterPage/reducer";
import buildImageReducer from "../containers/BuildImage/reducer";
import loginDockerHubReducer from "../containers/LoginDokcerHub/reducer";
import UiReducer from "./ui";

export default (history) =>
  combineReducers({
    login: loginReducer,
    register: registerReducer,
    buildImage: buildImageReducer,
    loginDockerHub: loginDockerHubReducer,
    ui: UiReducer,
    form: formReducer,
    router: connectRouter(history),
  });
