import { call, takeLatest, put, delay } from "redux-saga/effects";
import {
  getListLocalImageSuccess,
  getListLocalImageFail,
  buildImageFail,
  buildImageSuccess,
  buildImagePending,
  getImageById,
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

const apiGetListLocalImage = async (data) => {
  if (cancel !== undefined) cancel();
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_LOCAL_IMAGE}`,
    headers: {
      Authorization: `Bear ${token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListLocalImage({ payload }) {
  try {
    const resp = yield call(apiGetListLocalImage, { payload });
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListLocalImageSuccess(data.data));
    }
  } catch (error) {
    yield put(getListLocalImageFail(error));
    yield put(push("/login"));
  }
}

const apiBuildImage = async (image) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let tags = image.tag;

  let formData = new FormData();
  formData.append("file", image.file);

  let urlCall =
    image.file.type === "application/x-tar"
      ? api.API_BUILD_TAR
      : api.API_BUILD_DOCKERFILE;

  let result = await axios({
    method: "POST",
    url: urlCall,
    data: formData,
    headers: {
      Authorization: `Bearer ${token}`,
      "content-type": "multipart/form-data",
    },
    params: {
      tags,
    },
  });
  return result;
};

function* buildImage({ payload }) {
  try {
    let image = payload.data;
    const resp = yield call(apiBuildImage, image);
    yield delay(1000);
    showLoading(false);
    const { data, status } = resp;
    if (status === 200) {
      toastWarning("Build Image is progressing. Please wait");
      yield put(buildImagePending(data));
      yield delay(5000);
      yield put(getImageById(data.data.id));
    }
  } catch (error) {
    yield delay(1000);
    showLoading(false);
    yield put(buildImageFail(error));
  }
}

const apiGetLocalImageById = async (data) => {
  let token = await localStorage.getItem("JWT_TOKEN");
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LOCAL_IMAGE_BY_ID}/${data}`,
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
  return result;
};

function* getLocalImageById({ payload }) {
  try {
    let id = payload.data;
    const abc = yield call(apiGetLocalImageById, id);
    const { data, status } = abc;
    if (status === 200) {
      if (abc.data.data.status === "on progress") {
        yield delay(15000);
        yield put(getImageById(abc.data.data.id));
      } else if (abc.data.data.status === "fail") {
        toastError("Build Image fail");
        yield put(buildImageSuccess(data));
      } else {
        toastSuccess("Build Image success");
        yield put(buildImageSuccess(data));
      }
    }
  } catch (error) {
    yield put(buildImageFail(error));
  }
}

function* onListLocalImageSaga() {
  yield takeLatest(types.BUILD_IMAGE, buildImage);
  yield takeLatest(types.GET_LIST_LOCAL_IMAGE, getListLocalImage);
  yield takeLatest(types.GET_IMAGE_BY_ID, getLocalImageById);
}

export default onListLocalImageSaga();
