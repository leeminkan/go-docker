import { call, takeLatest, put, delay } from "redux-saga/effects";
import {
  getListImageInDeviceSuccess,
  getListImageInDeviceFail,
  getListContainerInDeviceSuccess,
  getListContainerInDeviceFail,
  getDeviceImageById,
  pullImagePending,
  pullImageFail,
  pullImageSuccess,
  getDeviceContainerById,
  runImageDevicePending,
  runImageDeviceFail,
  runImageDeviceSuccess,
} from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";
import { toastWarning } from "../../helpers/toastHelper";
import { push } from "connected-react-router";
import { showLoading } from "../../helpers/loading";

const CancelToken = axios.CancelToken;
let cancel;

const apiGetListImageInDevice = async (data) => {
  if (cancel !== undefined) cancel();
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_IMAGE_IN_DEVICE}/${data.payload.data}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListImageInDevice({ payload }) {
  try {
    const resp = yield call(apiGetListImageInDevice, { payload });
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListImageInDeviceSuccess(data.data));
    } else if (status === 20002) {
      yield put(push("/login"));
    }
  } catch (error) {
    yield put(getListImageInDeviceFail(error));
  }
}

const apiGetListContainerInDevice = async (data) => {
  if (cancel !== undefined) cancel();
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_CONTAINER_IN_DEVICE}/${data.payload.data}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListContainerInDevice({ payload }) {
  try {
    const resp = yield call(apiGetListContainerInDevice, { payload });
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListContainerInDeviceSuccess(data.data));
    } else if (status === 20002) {
      yield put(push("/login"));
    }
  } catch (error) {
    yield put(getListContainerInDeviceFail(error));
  }
}

const apiPullImage = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");

  let result = await axios({
    method: "POST",
    url: `${api.API_PULL_IMAGE}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
    data: data,
  });
  return result;
};

function* pullImage({ payload }) {
  try {
    let image = payload.data;
    const resp = yield call(apiPullImage, image);
    yield delay(1000);
    showLoading(false);
    toastWarning("Pull Image is progressing. Please wait");
    const { data, status } = resp;
    console.log(resp);
    if (status === 200) {
      yield put(pullImagePending(data));
      yield delay(15000);
      yield put(getDeviceImageById(data.data.id));
    }
  } catch (error) {
    yield delay(1000);
    showLoading(false);
    yield put(pullImageFail(error));
  }
}

const apiGetDeviceImageById = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_IMAGE_DEVICE_BY_ID}/${data}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* getDeviceImageByIdSaga({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDeviceImageById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.status === "on progress") {
        yield delay(15000);
        yield put(getDeviceImageById(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        yield put(pullImageFail("Pull image lỗi"));
      } else {
        yield put(pullImageSuccess(data));
      }
    }
  } catch (error) {
    yield put(pullImageFail(error));
  }
}

const apiRunImageDevice = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "POST",
    url: `${api.API_RUN_IMAGE_IN_DEVICE}`,
    data: data,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* runImageDevice({ payload }) {
  try {
    let image = payload.data;
    const resp = yield call(apiRunImageDevice, image);
    yield delay(1000);
    showLoading(false);
    toastWarning("Run Image is progressing. Please wait");
    const { data, status } = resp;
    if (status === 200) {
      yield put(runImageDevicePending(data));
      yield delay(15000);
      yield put(getDeviceContainerById(data.data.id));
    }
  } catch (error) {
    yield delay(1000);
    showLoading(false);
    yield put(runImageDeviceFail(error));
  }
}

const apiGetDeviceContainerById = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_CONTAINER_IN_DEVICE_BY_ID}/${data}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* getDeviceContainerByIdSaga({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDeviceContainerById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.status === "on progress") {
        yield delay(15000);
        yield put(getDeviceContainerById(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        yield put(runImageDeviceFail("Push image lỗi"));
      } else {
        yield put(runImageDeviceSuccess(data));
      }
    }
  } catch (error) {
    yield put(runImageDeviceFail(error));
  }
}

function* onDeviceDetail() {
  yield takeLatest(types.GET_LIST_IMAGE_IN_DEVICE, getListImageInDevice);
  yield takeLatest(
    types.GET_LIST_CONTAINER_IN_DEVICE,
    getListContainerInDevice
  );
  yield takeLatest(types.PULL_IMAGE, pullImage);
  yield takeLatest(types.GET_DEVICE_IMAGE_BY_ID, getDeviceImageByIdSaga);
  yield takeLatest(types.RUN_IMAGE_DEVICE, runImageDevice);
  yield takeLatest(
    types.GET_DEVICE_CONTAINER_BY_ID,
    getDeviceContainerByIdSaga
  );
}

export default onDeviceDetail();
