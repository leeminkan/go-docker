import { all } from "redux-saga/effects";
import onLoginSaga from "../containers/LoginPage/saga";
import onRegisterSaga from "../containers/RegisterPage/saga";
import onLoginDockerHubSaga from "../containers/LoginDokcerHub/saga";
import onListDevice from "../containers/DeviceList/saga";
import onListLocalImage from "../containers/ImageLocal/saga";

function* rootSaga() {
  yield all([
    onLoginSaga,
    onRegisterSaga,
    onLoginDockerHubSaga,
    onListDevice,
    onListLocalImage,
  ]);
}

export default rootSaga;
