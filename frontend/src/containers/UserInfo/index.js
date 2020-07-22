import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose } from "redux";

class UserInfo extends Component {
  render() {
    return <div>User Info</div>;
  }
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(UserInfo);
