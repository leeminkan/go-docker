import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose } from "redux";

class ReconfigNode extends Component {
  render() {

    return <div>Reconfig Edge Node</div>;
  }
}

const mapStateToProps = (state) => {};

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(ReconfigNode);
