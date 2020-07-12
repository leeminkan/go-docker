import React, { Component } from "react";
import { Redirect, Route } from "react-router-dom";
import Dashboard from "../../components/Dashboard";

class AdminLayoutRoute extends Component {
  render() {
    const { component: YourComponent, ...remainProps } = this.props;
    const token = localStorage.getItem("JWT_TOKEN");
    return (
      <Route
        {...remainProps}
        render={(routeProps) => {
          return token ? (
            <Dashboard>
              <YourComponent {...routeProps} />
            </Dashboard>
          ) : (
            <Redirect to="/login" />
          );
        }}
      />
    );
  }
}

export default AdminLayoutRoute;
