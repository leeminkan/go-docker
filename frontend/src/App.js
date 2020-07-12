import React, { Component } from "react";
import { MuiThemeProvider } from "@material-ui/core/styles";
import { history } from "./redux";
import { Router, Switch } from "react-router-dom";
import theme from "./commons/styles/theme";
import { CssBaseline } from "@material-ui/core";
import { ADMIN_ROUTES, ROUTES } from "./constants/route";
import AdminLayoutRoute from "./layout/AdminLayout";
import DefaultLayoutRoute from "./layout/DefaultLayout";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

class App extends Component {
  renderAdminRoutes() {
    let xhtml = null;
    xhtml = ADMIN_ROUTES.map((item, index) => {
      return (
        <AdminLayoutRoute
          component={item.component}
          exact={item.exact}
          key={index}
          path={item.path}
        />
      );
    });
    return xhtml;
  }

  renderDefaultRoutes() {
    let xhtml = null;
    xhtml = ROUTES.map((item, index) => {
      return (
        <DefaultLayoutRoute
          component={item.component}
          exact={item.exact}
          key={index}
          path={item.path}
        />
      );
    });
    return xhtml;
  }

  renderRoutes() {
    let xhtml = null;
    xhtml = (
      <Switch>
        {this.renderAdminRoutes()}
        {this.renderDefaultRoutes()}
      </Switch>
    );
    return xhtml;
  }

  render() {
    return (
      <Router history={history}>
        <MuiThemeProvider theme={theme}>
          <div className="App">
            <ToastContainer />
            <CssBaseline />
            {this.renderRoutes()}
          </div>
        </MuiThemeProvider>
      </Router>
    );
  }
}

export default App;
