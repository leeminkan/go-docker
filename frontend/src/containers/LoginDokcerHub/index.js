import Button from "@material-ui/core/Button";
import { withStyles } from "@material-ui/styles";
import React, { Component } from "react";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../components/FormHelper/TextField";
import styles from "./styles";
import validate from "./validate";
import * as loginDockerHubAction from "./action";
import { Container } from "@material-ui/core";
import logo from "../../assets/img/docker_logo.png";

class LoginDockerHub extends Component {
  onLoginDockerHub = (data) => {
    this.props.loginDockerHubAC.loginDockerHub(data);
  };

  render() {
    const { classes, handleSubmit, submitting } = this.props;
    return (
      <Container component="main" maxWidth="xs">
        <div className={classes.paper}>
          <div className={classes.logoImage}>
            <img src={logo} alt="logo" className={classes.img} />
          </div>
          <form onSubmit={handleSubmit(this.onLoginDockerHub)}>
            <Field
              variant="outlined"
              margin="normal"
              fullWidth
              label="Username"
              name="username"
              autoComplete="username"
              component={TextField}
            />
            <Field
              variant="outlined"
              margin="normal"
              fullWidth
              name="password"
              label="Password"
              type="password"
              component={TextField}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              className={classes.submit}
              disabled={submitting}
            >
              Login Docker Hub
            </Button>
          </form>
        </div>
      </Container>
    );
  }
}

const FORM_NAME = "LOGIN_DOCKER_HUB";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {
  return {
    loginDockerHubAC: bindActionCreators(loginDockerHubAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(
  withStyles(styles),
  withConnect,
  withForm
)(LoginDockerHub);
