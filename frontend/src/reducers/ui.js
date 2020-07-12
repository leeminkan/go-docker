import * as types from "../constants/ui";

const initialState = {
  showSidebar: true,
};

const reducer = (state = initialState, action) => {
  switch (action.type) {
    case types.SHOW_SIDEBAR: {
      return {
        ...state,
        showSidebar: true,
      };
    }
    case types.HIDE_SIDEBAR: {
      return {
        ...state,
        showSidebar: false,
      };
    }
    default:
      return state;
  }
};

export default reducer;
