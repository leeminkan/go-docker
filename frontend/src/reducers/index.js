import { combineReducers } from "redux";
import { reducer as formReducer } from "redux-form";
import { connectRouter } from "connected-react-router";
import loginReducer from "../containers/LoginPage/reducer";
import registerReducer from "../containers/RegisterPage/reducer";
import loginDockerHubReducer from "../containers/LoginDokcerHub/reducer";
import ListDevice from "../containers/DeviceList/reducer";
import ListLocalImage from "../containers/ImageLocal/reducer";
import UiReducer from "./ui";

export default (history) =>
  combineReducers({
    login: loginReducer,
    register: registerReducer,
    loginDockerHub: loginDockerHubReducer,
    device: ListDevice,
    localImage: ListLocalImage,
    ui: UiReducer,
    form: formReducer,
    router: connectRouter(history),
  });
