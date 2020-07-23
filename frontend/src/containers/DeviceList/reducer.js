import * as types from "./constant";

const initialState = {
  listDevice: [],
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.GET_LIST_DEVICE: {
      return {
        ...state,
        listDevice: [],
      };
    }
    case types.GET_LIST_DEVICE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        listDevice: data,
      };
    }
    case types.GET_LIST_DEVICE_FAIL: {
      return {
        ...state,
      };
    }

    default:
      return state;
  }
};

export default reducer;
