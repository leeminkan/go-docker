import { call, takeLatest, put, delay } from "redux-saga/effects";
import {
  getListDockerHubImageSuccess,
  getListDockerHubImageFail,
  getDHImageById,
  pushImagePending,
  pushImageFail,
  pushImageSuccess,
} from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";
import { toastWarning } from "../../helpers/toastHelper";

const CancelToken = axios.CancelToken;
let cancel;

const apiGetListDockerHubImage = async (data) => {
  if (cancel !== undefined) cancel();

  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_PUSH}`,
    headers: {
      Authorization: `Bear ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListDockerHubImage({ payload }) {
  try {
    const resp = yield call(apiGetListDockerHubImage, { payload });
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListDockerHubImageSuccess(data.data));
    }
  } catch (error) {
    yield put(getListDockerHubImageFail(error));
  }
}

const apiPushImage = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");

  let result = await axios({
    method: "POST",
    url: `${api.API_PUSH_IMAGE}/${data.id}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* pushImage({ payload }) {
  try {
    let image = payload.data;
    const resp = yield call(apiPushImage, image);
    toastWarning("Push Image is progressing. Please wait");
    const { data, status } = resp;
    if (status === 200) {
      yield put(pushImagePending(data));
      yield delay(15000);
      yield put(getDHImageById(data.data.id));
    }
  } catch (error) {
    yield put(pushImageFail(error));
  }
}

const apiGetDockerHubImageById = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_IMAGE_PUSH_BY_ID}/${data}`,
    headers: {
      Authorization: `Bear ${token}`,
    },
  });
  return result;
};

function* getDockerHubImageById({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetDockerHubImageById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.status === "on progress") {
        yield delay(15000);
        yield put(getDHImageById(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        yield put(pushImageFail("Push image lỗi"));
      } else {
        yield put(pushImageSuccess(data));
      }
    }
  } catch (error) {
    yield put(pushImageFail(error));
  }
}

function* onListDockerHubImageSaga() {
  yield takeLatest(types.PUSH_IMAGE, pushImage);
  yield takeLatest(types.GET_LIST_DOCKER_HUB_IMAGE, getListDockerHubImage);
  yield takeLatest(types.GET_DH_IMAGE_BY_ID, getDockerHubImageById);
}

export default onListDockerHubImageSaga();