import Button from "@material-ui/core/Button";
import { withStyles } from "@material-ui/styles";
import React, { Component } from "react";
import { connect } from "react-redux";
import { Link as RouterLink } from "react-router-dom";
import { compose, bindActionCreators } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../components/FormHelper/TextField";
import SelectField from "../../components/FormHelper/Select";
import styles from "./styles";
import validate from "./validate";
import * as registerAction from "./action";
import { Grid, Link, Container } from "@material-ui/core";

class RegisterPage extends Component {
  onRegister = (data) => {
    this.props.registerActionCreators.register(data);
  };

  render() {
    const { classes, handleSubmit, submitting } = this.props;
    return (
      <Container component="main" maxWidth="xs">
        <div className={classes.paper}>
          {/* <div className={classes.logoImage}>
            <img src={logo} alt="logo" className={classes.img} />
          </div> */}
          <form onSubmit={handleSubmit(this.onRegister)}>
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
            <Field
              variant="outlined"
              margin="normal"
              fullWidth
              name="rePassword"
              label="rePassword"
              type="password"
              component={TextField}
            />
            <Field
              name="isAdmin"
              id="isAdmin"
              component={SelectField}
              fullWidth
              label="Role"
            >
              <option value={false}></option>
              <option value={true}>Admin</option>
              <option value={false}>User</option>
            </Field>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              className={classes.submit}
              disabled={submitting}
            >
              Register
            </Button>
          </form>
          <Grid container>
            <Grid item xs>
              <Link to="/login" component={RouterLink} variant="body2">
                Login ?
              </Link>
            </Grid>
          </Grid>
        </div>
      </Container>
    );
  }
}

const FORM_NAME = "REGISTER";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {
  return {
    registerActionCreators: bindActionCreators(registerAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect, withForm)(RegisterPage);
