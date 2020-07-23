import * as types from "./constant";

const initialState = {
  listLocalImage: [],
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.GET_LIST_LOCAL_IMAGE: {
      return {
        ...state,
        listLocalImage: [],
      };
    }
    case types.GET_LIST_LOCAL_IMAGE_SUCCESS: {
      const { data } = action.payload;
      return {
        ...state,
        listLocalImage: data,
      };
    }
    case types.GET_LIST_LOCAL_IMAGE_FAIL: {
      return {
        ...state,
      };
    }

    default:
      return state;
  }
};

export default reducer;
