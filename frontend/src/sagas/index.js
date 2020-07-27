import { all } from "redux-saga/effects";
import onLoginSaga from "../containers/LoginPage/saga";
import onRegisterSaga from "../containers/RegisterPage/saga";
import onLoginDockerHubSaga from "../containers/LoginDokcerHub/saga";
import onListDevice from "../containers/DeviceList/saga";
import onListLocalImage from "../containers/ImageLocal/saga";
import onListDockerHubImage from "../containers/ImageDockerHub/saga";
import onDeviceDetail from "../containers/DeviceDetail/saga";

function* rootSaga() {
  yield all([
    onLoginSaga,
    onRegisterSaga,
    onLoginDockerHubSaga,
    onListDevice,
    onListLocalImage,
    onListDockerHubImage,
    onDeviceDetail,
  ]);
}

export default rootSaga;
