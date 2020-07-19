import { call, takeLatest, put } from "redux-saga/effects";
import { buildImageSuccess, buildImageFail } from "./action";
import * as types from "./constant";
import * as api from "../../constants/config";
import axios from "axios";

const apiBuildImage = async (image) => {
  //let token = await localStorage.getItem("JWT_TOKEN");
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
      //Authorization: `Bearer ${token}`,
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
    const { data, status } = resp;
    if (status === 200) {
      yield put(buildImageSuccess(data));
    }
  } catch (error) {
    yield put(buildImageFail(error));
  }
}

function* buildImageSaga() {
  yield takeLatest(types.BUILD_IMAGE, buildImage);
}

export default buildImageSaga();
