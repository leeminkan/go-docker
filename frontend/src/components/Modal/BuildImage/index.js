import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import "./styles.css";
import Dropzone from "react-dropzone";
import { compose } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../FormHelper/TextField";
import { CardContent, Typography, Button, Modal } from "@material-ui/core";
import { toastError } from "../../../helpers/toastHelper";
import CloseIcon from "@material-ui/icons/Clear";

class BuildImage extends Component {
  constructor(props) {
    super(props);
    this.state = {
      file: [],
    };
  }

  onDrop = (acceptedFiles) => {
    this.setState({
      file: acceptedFiles[0],
    });
  };

  onBuild = (data) => {
    // eslint-disable-next-line
    if (this.state.file != "") {
      let payload = {
        file: this.state.file,
        tag: data.tag,
      };
      this.props.onSave(payload);
    } else {
      toastError("Vui lòng upload file");
    }
  };

  onClose = () => {
    this.props.onCloseModalBI();
    this.setState({
      file: [],
    });
    const { reset } = this.props;
    reset();
  };

  render() {
    const { classes, handleSubmit, submitting, openModalBI } = this.props;

    let fileUpload =
      // eslint-disable-next-line
      this.state.file != "" ? (
        <div>
          <strong>Files: </strong>
          <span>
            {this.state.file.name} - {this.state.file.size} bytes
          </span>
        </div>
      ) : (
        ""
      );

    return (
      <Modal open={openModalBI}>
        <div className={classes.modal}>
          <div className={classes.header}>
            <span className={classes.title}>Build Image</span>
            <CloseIcon className={classes.icon} onClick={this.onClose} />
          </div>
          <div className={classes.content}>
            <form onSubmit={handleSubmit(this.onBuild)}>
              <Typography
                gutterBottom
                variant="h5"
                component="h2"
                className={classes.nameImage}
              >
                Image tags
              </Typography>
              <Field
                className={classes.form}
                variant="outlined"
                margin="normal"
                fullWidth
                name="tag"
                component={TextField}
              />
              <CardContent className={classes.cardContent}>
                <div>
                  <Dropzone
                    onDrop={this.onDrop}
                    maxSize={3072000}
                    minSize={1}
                    multiple={false}
                  >
                    {({ getRootProps, getInputProps }) => (
                      <div {...getRootProps({ className: "dropzone" })}>
                        <input {...getInputProps()} />
                        <p>Drag and drop files, or click to select files</p>
                      </div>
                    )}
                  </Dropzone>
                </div>
                <div>{fileUpload}</div>
              </CardContent>
              <div className={classes.button}>
                <Button
                  type="submit"
                  variant="contained"
                  className={classes.submit}
                  disabled={submitting}
                >
                  Build Image
                </Button>
              </div>
            </form>
          </div>
        </div>
      </Modal>
    );
  }
}

const FORM_NAME = "BUILD_IMAGE";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

export default compose(withStyles(styles), withForm)(BuildImage);
