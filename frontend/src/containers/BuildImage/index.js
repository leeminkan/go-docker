import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import "./styles.css";
import Dropzone from "react-dropzone";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import { Field, reduxForm } from "redux-form";
import TextField from "../../components/FormHelper/TextField";
import * as buildImageAction from "./action";
import { Grid, Card, CardContent, Typography, Button } from "@material-ui/core";
import { toastError } from "../../helpers/toastHelper";

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
      this.props.buildImageActionCreators.buildImage(payload);
    } else {
      toastError("Vui lòng upload file");
    }
  };

  render() {
    const { classes, handleSubmit, submitting } = this.props;

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
      <div className={classes.root}>
        <div>
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            build image
          </Typography>
        </div>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Card className={classes.card}>
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
              {/* <Typography
                gutterBottom
                variant="h5"
                component="h2"
                className={classes.nameImage}
              >
                Result
              </Typography> */}
              {/* <CardContent>
                <div className={classes.result}>
                  <Typography
                    gutterBottom
                    variant="h5"
                    component="body1"
                    className={classes.text}
                  ></Typography>
                </div>
              </CardContent> */}
            </Card>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const FORM_NAME = "BUILD_IMAGE";
const withForm = reduxForm({
  form: FORM_NAME,
  validate,
});

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {
  return {
    buildImageActionCreators: bindActionCreators(buildImageAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect, withForm)(BuildImage);
