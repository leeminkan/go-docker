import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import {
  Grid,
  Card,
  CardContent,
  Typography,
  TextField,
  Button,
} from "@material-ui/core";
import { connect } from "react-redux";
import { compose } from "redux";
import { Field, reduxForm } from "redux-form";
import validate from "./validate";
import logo from "../../assets/img/avatar_default.png";

class UserInfo extends Component {
  onEdit = (data) => {
    //this.props.loginDockerHubAC.loginDockerHub(data);
  };

  render() {
    const { classes, submitting, handleSubmit } = this.props;

    return (
      <div className={classes.root}>
        <div>
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            User Information
          </Typography>
        </div>
        <Grid container>
          <Grid item xs={12}>
            <Card className={classes.card}>
              <CardContent className={classes.cardContent}>
                <div className={classes.paper}>
                  <div className={classes.logoImage}>
                    <img src={logo} alt="logo" className={classes.img} />
                  </div>
                  <form onSubmit={handleSubmit(this.onEdit)}>
                    <Field
                      variant="outlined"
                      margin="normal"
                      fullWidth
                      name="id"
                      label="ID"
                      component={TextField}
                    />
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
                      name="role"
                      label="Role"
                      component={TextField}
                    />
                    <div className={classes.button}>
                      <Button
                        type="submit"
                        variant="contained"
                        className={classes.submit}
                        disabled={submitting}
                      >
                        Edit
                      </Button>
                    </div>
                  </form>
                </div>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const FORM_NAME = "USER_INFO";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect, withForm)(UserInfo);
