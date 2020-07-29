import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import { compose } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../FormHelper/TextField";
import { Typography, Button, Modal } from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Clear";

class RunImage extends Component {
  onRun = (data) => {
    this.props.onSave(data);
  };

  onClose = () => {
    this.props.onCloseModalRI();
    const { reset } = this.props;
    reset();
  };

  render() {
    const { classes, handleSubmit, submitting, openModalRI } = this.props;

    return (
      <Modal open={openModalRI}>
        <div className={classes.modal}>
          <div className={classes.header}>
            <span className={classes.title}>Run Image</span>
            <CloseIcon className={classes.icon} onClick={this.onClose} />
          </div>
          <div className={classes.content}>
            <form onSubmit={handleSubmit(this.onRun)}>
              <Typography
                gutterBottom
                variant="h5"
                component="h2"
                className={classes.nameImage}
              >
                Container name
              </Typography>
              <Field
                className={classes.form}
                variant="outlined"
                margin="normal"
                fullWidth
                name="name"
                component={TextField}
              />
              <div className={classes.button}>
                <Button
                  type="submit"
                  variant="contained"
                  className={classes.submit}
                  disabled={submitting}
                >
                  Run Image
                </Button>
              </div>
            </form>
          </div>
        </div>
      </Modal>
    );
  }
}

const FORM_NAME = "RUN_IMAGE";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

export default compose(withStyles(styles), withForm)(RunImage);
