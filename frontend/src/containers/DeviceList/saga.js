import { call, takeLatest, put } from "redux-saga/effects";
import { getListDeviceSuccess, getListDeviceFail } from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";
import { push } from "connected-react-router";

const CancelToken = axios.CancelToken;
let cancel;

const apiGetListDevice = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  if (cancel !== undefined) cancel();

  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_DEVICE}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListDevice({ payload }) {
  try {
    const resp = yield call(apiGetListDevice);
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListDeviceSuccess(data.data));
    }
  } catch (error) {
    yield put(getListDeviceFail(error));
    yield put(push("/login"));
  }
}

function* onListDeviceSaga() {
  yield takeLatest(types.GET_LIST_DEVICE, getListDevice);
}

export default onListDeviceSaga();
