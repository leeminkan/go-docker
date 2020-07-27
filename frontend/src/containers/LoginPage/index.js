import Button from "@material-ui/core/Button";
import { withStyles } from "@material-ui/styles";
import React, { Component } from "react";
import { connect } from "react-redux";
import { Link as RouterLink } from "react-router-dom";
import { compose, bindActionCreators } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../components/FormHelper/TextField";
import styles from "./styles";
import validate from "./validate";
import * as loginAction from "./action";
import { Container, Grid, Link } from "@material-ui/core";
import { showLoading } from "../../helpers/loading";

class LoginPage extends Component {
  onLogin = (data) => {
    showLoading(true);
    this.props.loginActionCreators.login(data);
  };

  render() {
    const { classes, handleSubmit, submitting } = this.props;
    return (
      <Container component="main" maxWidth="xs">
        <div className={classes.paper}>
          {/* <div className={classes.logoImage}>
            <img src={logo} alt="logo" className={classes.img} />
          </div> */}
          <form onSubmit={handleSubmit(this.onLogin)}>
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
              Login
            </Button>
          </form>
          <Grid container>
            <Grid item xs>
              <Link to="/register" component={RouterLink} variant="body2">
                Create new account ?
              </Link>
            </Grid>
          </Grid>
        </div>
      </Container>
    );
  }
}

const FORM_NAME = "LOGIN";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {
  return {
    loginActionCreators: bindActionCreators(loginAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect, withForm)(LoginPage);
