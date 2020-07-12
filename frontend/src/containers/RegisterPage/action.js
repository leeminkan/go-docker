import * as types from "./constants";

export const register = (data) => {
  return {
    type: types.REGISTER,
    payload: {
      data,
    },
  };
};

export const registerSuccess = (data) => {
  return {
    type: types.REGISTER_SUCCESS,
    payload: {
      data,
    },
  };
};
