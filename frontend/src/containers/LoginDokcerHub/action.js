import * as types from "./constants";

export const loginDockerHub = (data) => {
  return {
    type: types.LOGIN_DOCKERHUB,
    payload: {
      data,
    },
  };
};

export const loginDockerHubSuccess = (data) => {
  return {
    type: types.LOGIN_DOCKERHUB_SUCCESS,
    payload: {
      data,
    },
  };
};
