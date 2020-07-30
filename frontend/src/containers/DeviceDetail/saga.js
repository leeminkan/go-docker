import { call, takeLatest, put, delay } from "redux-saga/effects";
import {
  getListImageInDeviceSuccess,
  getListImageInDeviceFail,
  getListContainerInDeviceSuccess,
  getListContainerInDeviceFail,
  getDeviceImageById,
  getDeviceContainerById2,
  pullImagePending,
  pullImageFail,
  pullImageSuccess,
  getDeviceContainerById1,
  getDeviceContainerById3,
  runImageDevicePending,
  runImageDeviceFail,
  runImageDeviceSuccess,
  stopContainerPending,
  stopContainerFail,
  stopContainerSuccess,
  startContainerPending,
  startContainerFail,
  startContainerSuccess,
} from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";
import {
  toastWarning,
  toastSuccess,
  toastError,
} from "../../helpers/toastHelper";
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
    const { data, status } = resp;
    if (status === 200) {
      toastWarning("Pull Image is progressing. Please wait");
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
        toastError("Pull Image fail");
        yield put(pullImageSuccess(data));
      } else {
        toastSuccess("Pull Image success");
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
    const { data, status } = resp;
    if (status === 200) {
      toastWarning("Run Image is progressing. Please wait");
      yield put(runImageDevicePending(data));
      yield delay(15000);
      yield put(getDeviceContainerById1(data.data.id));
    }
  } catch (error) {
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

function* getDeviceContainerById1Saga({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDeviceContainerById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.status === "on progress") {
        yield delay(15000);
        yield put(getDeviceContainerById1(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        toastError("Run Image fail");
        yield put(runImageDeviceSuccess(data));
      } else {
        toastSuccess("Run Image thành công");
        yield put(runImageDeviceSuccess(data));
      }
    }
  } catch (error) {
    yield put(runImageDeviceFail(error));
  }
}

const apiStopContainerDevice = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "POST",
    url: `${api.API_STOP_CONTAINER_IN_DEVICE}`,
    data: {
      containerID: data,
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* stopContainerDevice({ payload }) {
  try {
    let id = payload.data;
    const resp = yield call(apiStopContainerDevice, id);
    const { data, status } = resp;
    if (status === 200) {
      toastWarning("Stop Container is progressing. Please wait");
      yield put(stopContainerPending(data));
      yield delay(15000);
      yield put(getDeviceContainerById2(data.data.id));
    }
  } catch (error) {
    yield put(stopContainerFail(error));
  }
}

function* getDeviceContainerById2Saga({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDeviceContainerById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.active === "stopping") {
        yield delay(15000);
        yield put(getDeviceContainerById2(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        toastError("Stop container fail");
        yield put(stopContainerSuccess(data));
      } else {
        toastSuccess("Stop container success");
        yield put(stopContainerSuccess(data));
      }
    }
  } catch (error) {
    yield put(stopContainerFail(error));
  }
}

const apiStartContainerDevice = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "POST",
    url: `${api.API_START_CONTAINER_IN_DEVICE}`,
    data: {
      containerID: data,
    },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* startContainerDevice({ payload }) {
  try {
    let id = payload.data;
    const resp = yield call(apiStartContainerDevice, id);
    const { data, status } = resp;
    if (status === 200) {
      toastWarning("Start Container is progressing. Please wait");
      yield put(startContainerPending(data));
      yield delay(15000);
      yield put(getDeviceContainerById3(data.data.id));
    }
  } catch (error) {
    yield put(startContainerFail(error));
  }
}

function* getDeviceContainerById3Saga({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDeviceContainerById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.active === "starting") {
        yield delay(15000);
        yield put(getDeviceContainerById3(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        toastError("Start container fail");
        yield put(startContainerSuccess(data));
      } else {
        toastSuccess("Start container success");
        yield put(startContainerSuccess(data));
      }
    }
  } catch (error) {
    yield put(stopContainerFail(error));
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
    types.GET_DEVICE_CONTAINER_BY_ID_1,
    getDeviceContainerById1Saga
  );
  yield takeLatest(
    types.GET_DEVICE_CONTAINER_BY_ID_2,
    getDeviceContainerById2Saga
  );
  yield takeLatest(
    types.GET_DEVICE_CONTAINER_BY_ID_3,
    getDeviceContainerById3Saga
  );
  yield takeLatest(types.STOP_CONTAINER, stopContainerDevice);
  yield takeLatest(types.START_CONTAINER, startContainerDevice);
}

export default onDeviceDetail();
