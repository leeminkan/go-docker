import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import configureStore, { history } from "./redux";
import { ConnectedRouter } from "connected-react-router";
import App from "./App";
import Loading from "./components/Loading";

import * as serviceWorker from "./serviceWorker";

const store = configureStore();

ReactDOM.render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
      <Loading>
        <App />
      </Loading>
    </ConnectedRouter>
  </Provider>,
  document.getElementById("root")
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
