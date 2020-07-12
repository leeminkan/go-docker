import * as types from "./constants";

export const login = (data) => {
  return {
    type: types.LOGIN,
    payload: {
      data,
    },
  };
};

export const loginSuccess = (data) => {
  return {
    type: types.LOGIN_SUCCESS,
    payload: {
      data,
    },
  };
};
