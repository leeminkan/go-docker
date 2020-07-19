import { all } from "redux-saga/effects";
import onLoginSaga from "../containers/LoginPage/saga";
import onRegisterSaga from "../containers/RegisterPage/saga";
import buildImageSaga from "../containers/BuildImage/saga";
import onLoginDockerHubSaga from "../containers/LoginDokcerHub/saga";

function* rootSaga() {
  yield all([
    onLoginSaga,
    onRegisterSaga,
    buildImageSaga,
    onLoginDockerHubSaga,
  ]);
}

export default rootSaga;
