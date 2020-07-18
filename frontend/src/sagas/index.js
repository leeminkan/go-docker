import { all } from "redux-saga/effects";
import onLoginSaga from "../containers/LoginPage/saga";
import onRegisterSaga from "../containers/RegisterPage/saga";
import buildImageSaga from "../containers/BuildImage/saga";

function* rootSaga() {
  yield all([onLoginSaga, onRegisterSaga, buildImageSaga]);
}

export default rootSaga;
