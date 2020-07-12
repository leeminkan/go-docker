import { all } from "redux-saga/effects";
import onLoginSaga from "../containers/LoginPage/saga";
import onRegisterSaga from "../containers/RegisterPage/saga";

function* rootSaga() {
  yield all([onLoginSaga, onRegisterSaga]);
}

export default rootSaga;
