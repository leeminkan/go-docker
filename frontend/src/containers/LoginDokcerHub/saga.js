import { call, takeLatest, put } from "redux-saga/effects";
import * as types from "./constants";
import * as api from "../../constants/config";
import { loginDockerHubSuccess } from "./action";
import axios from "axios";
//import { push } from "connected-react-router";
import { toastError } from "../../helpers/toastHelper";

const CancelToken = axios.CancelToken;
let cancel;

const apiLoginDockerHub = async (data) => {
  if (cancel !== undefined) cancel();

  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "POST",
    url: `${api.API_LOGIN_DOCKERHUB}`,
    data: {
      username: data.username,
      password: data.password,
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* onLoginDockerHub({ payload }) {
  try {
    const result = payload.data;
    const resp = yield call(apiLoginDockerHub, result);
    const { data, status } = resp;
    if (status === 200) {
      yield put(loginDockerHubSuccess(data));
    }
  } catch (error) {
    toastError(error);
  }
}

function* onLoginDockerHubSuccess({ payload }) {
  const { data } = payload.data;
  console.log(data);
  yield localStorage.setItem("DOCKERHUB", "abcxyz");
  //yield put(push("/"));
}

function* onLoginDockerHubSaga() {
  yield takeLatest(types.LOGIN_DOCKERHUB, onLoginDockerHub);
  yield takeLatest(types.LOGIN_DOCKERHUB_SUCCESS, onLoginDockerHubSuccess);
}

export default onLoginDockerHubSaga();
