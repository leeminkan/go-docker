import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import { compose } from "redux";
import { Field, reduxForm } from "redux-form";
import SelectField from "../../FormHelper/Select";
import { Button, Modal } from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Clear";

class PushImage extends Component {
  onPush = (data) => {
    if (data) {
      this.props.onSave(data);
    }
  };

  onClose = () => {
    this.props.onCloseModalPI();
    const { reset } = this.props;
    reset();
  };

  onrenderSelect = () => {
    let xml = null;
    xml = this.props.localImage.map((item, index) => {
      return (
        <option key={index} value={item.id}>
          {item.repo_name}
        </option>
      );
    });
    return xml;
  };

  render() {
    const { classes, handleSubmit, submitting, openModalPI } = this.props;

    return (
      <Modal open={openModalPI}>
        <div className={classes.modal}>
          <div className={classes.header}>
            <span className={classes.title}>Push Image To Docker Hub</span>
            <CloseIcon className={classes.icon} onClick={this.onClose} />
          </div>
          <div className={classes.content}>
            <form onSubmit={handleSubmit(this.onPush)}>
              <Field
                name="id"
                id="id"
                component={SelectField}
                fullWidth
                label="Image"
              >
                <option value={""}></option>
                {this.onrenderSelect()}
              </Field>
              <div className={classes.button}>
                <Button
                  type="submit"
                  variant="contained"
                  className={classes.submit}
                  disabled={submitting}
                >
                  Push Image
                </Button>
              </div>
            </form>
          </div>
        </div>
      </Modal>
    );
  }
}

const FORM_NAME = "PUSH_IMAGE";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

export default compose(withStyles(styles), withForm)(PushImage);
