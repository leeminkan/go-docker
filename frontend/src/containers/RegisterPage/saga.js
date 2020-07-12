import { call, takeLatest, put } from "redux-saga/effects";
import * as types from "./constants";
import * as api from "../../constants/config";
import { registerSuccess } from "./action";
import axios from "axios";
import { push } from "connected-react-router";
import { toastSuccess } from "../../helpers/toastHelper";

const CancelToken = axios.CancelToken;
let cancel;

const apiRegister = async (data) => {
  if (cancel !== undefined) cancel();

  let formData = new FormData();
  formData.append("username", data.username);
  formData.append("password", data.password);
  formData.append("confirm_password", data.rePassword);
  formData.append("is_admin", data.isAdmin);

  let result = await axios({
    method: "POST",
    url: `${api.API_REGISTER}`,
    data: formData,
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* onRegister({ payload }) {
  try {
    const result = payload.data;
    const resp = yield call(apiRegister, result);
    const { data, status } = resp;
    if (status === 200) {
      yield put(registerSuccess(data));
    }
  } catch (error) {
    toastSuccess(error);
  }
}

function* onRegisterSuccess({ payload }) {
  const { data } = payload.data;
  console.log(data);
  yield put(push("/login"));
}

function* onRegisterSaga() {
  yield takeLatest(types.REGISTER, onRegister);
  yield takeLatest(types.REGISTER_SUCCESS, onRegisterSuccess);
}

export default onRegisterSaga();
