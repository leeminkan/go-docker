import { observer } from "../observer";
import * as command from "../constants/ui";

const showLoading = (isShow) => {
  observer.update(command.SHOW_LOADING, isShow);
};

export { showLoading };
