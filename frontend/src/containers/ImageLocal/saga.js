import { call, takeLatest, put } from "redux-saga/effects";
import { getListLocalImageSuccess, getListLocalImageFail } from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";

const CancelToken = axios.CancelToken;
let cancel;

const apiGetListLocalImage = async (data) => {
  if (cancel !== undefined) cancel();
  let result = await axios({
    method: "GET",
    url: `${api.API_GET_LIST_LOCAL_IMAGE}`,
    headers: {
      Authorization: `Bear ${data.token}`,
    },
    cancelToken: new CancelToken((c) => (cancel = c)),
  });
  return result;
};

function* getListLocalImage({ payload }) {
  try {
    let token = yield localStorage.getItem("JWT_TOKEN");
    const resp = yield call(apiGetListLocalImage, { token, payload });
    const { data, status } = resp;
    if (status === 200) {
      yield put(getListLocalImageSuccess(data.data));
    }
  } catch (error) {
    yield put(getListLocalImageFail(error));
  }
}

function* onListLocalImageSaga() {
  yield takeLatest(types.GET_LIST_LOCAL_IMAGE, getListLocalImage);
}

export default onListLocalImageSaga();
