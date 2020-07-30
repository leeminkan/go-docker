import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import { compose } from "redux";
import { Field, reduxForm } from "redux-form";
import SelectField from "../../FormHelper/Select";
import { Button, Modal } from "@material-ui/core";
import CloseIcon from "@material-ui/icons/Clear";

class PullImage extends Component {
  onPull = (data) => {
    if (data) {
      this.props.onSave(data);
    }
  };

  onClose = () => {
    const { reset } = this.props;
    reset();
    this.props.onCloseModalPullImage();
  };

  onrenderSelect = () => {
    let xml = null;
    if (this.props.dockerHubImage) {
      xml = this.props.dockerHubImage.map((item, index) => {
        if (item.status === "done") {
          return (
            <option key={index} value={item.id}>
              {item.full_repo_name}
            </option>
          );
        }
        return null;
      });
    }
    return xml;
  };

  render() {
    const {
      classes,
      handleSubmit,
      submitting,
      openModalPullImage,
    } = this.props;

    return (
      <Modal open={openModalPullImage}>
        <div className={classes.modal}>
          <div className={classes.header}>
            <span className={classes.title}>Pull Image From Docker Hub</span>
            <CloseIcon className={classes.icon} onClick={this.onClose} />
          </div>
          <div className={classes.content}>
            <form onSubmit={handleSubmit(this.onPull)}>
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
                  Pull Image
                </Button>
              </div>
            </form>
          </div>
        </div>
      </Modal>
    );
  }
}

const FORM_NAME = "PULL_IMAGE";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

export default compose(withStyles(styles), withForm)(PullImage);
