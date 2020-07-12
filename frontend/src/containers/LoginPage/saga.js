import { call, takeLatest, put } from "redux-saga/effects";
import * as types from "./constants";
import * as api from "../../constants/config";
import { loginSuccess } from "./action";
import axios from "axios";
import { push } from "connected-react-router";
import { toastError } from "../../helpers/toastHelper";

const CancelToken = axios.CancelToken;
let cancel;

const apiLogin = async (data) => {
  if (cancel !== undefined) cancel();

  let formData = new FormData();
  formData.append("username", data.username);
  formData.append("password", data.password);

  let result = await axios({
    method: "POST",
    url: `${api.API_LOGIN}`,
    data: formData,
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* onLogin({ payload }) {
  try {
    const result = payload.data;
    const resp = yield call(apiLogin, result);
    const { data, status } = resp;
    if (status === 200) {
      yield put(loginSuccess(data));
    }
  } catch (error) {
    toastError(error);
  }
}

function* onLoginSuccess({ payload }) {
  const { data } = payload.data;
  console.log(data);
  yield localStorage.setItem("JWT_TOKEN", data.token);
  yield put(push("/"));
}

function* onLoginSaga() {
  yield takeLatest(types.LOGIN, onLogin);
  yield takeLatest(types.LOGIN_SUCCESS, onLoginSuccess);
}

export default onLoginSaga();
