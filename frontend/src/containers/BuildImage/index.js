import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import validate from "./validate";
import "./styles.css";
import Dropzone from "react-dropzone";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as buildImageAction from "./action";
import {
  Grid,
  Card,
  CardContent,
  Typography,
  TextField,
  Button,
} from "@material-ui/core";
import { Field, reduxForm } from "redux-form";

class BuildImage extends Component {
  constructor(props) {
    super(props);
    this.state = { file: [] };
  }

  onDrop = (acceptedFiles) => {
    this.setState({
      file: acceptedFiles[0],
    });
  };

  onSubmit = (data) => {
    console.log(this.state.file);
    let payload = this.state.file;
    this.props.buildImageActionCreators.buildImage(payload);
  };

  render() {
    const { classes, handleSubmit, submitting } = this.props;

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
              <form onSubmit={handleSubmit(this.onSubmit)}>
                <Typography
                  gutterBottom
                  variant="h5"
                  component="h2"
                  className={classes.nameImage}
                >
                  Image name
                </Typography>
                <Field
                  className={classes.form}
                  variant="outlined"
                  margin="normal"
                  fullWidth
                  name="username"
                  autoComplete="username"
                  component={TextField}
                />
                <CardContent className={classes.cardContent}>
                  <div>
                    <Dropzone
                      onDrop={this.onDrop}
                      maxSize={3072000}
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
                  <div>
                    <strong>Files:</strong>
                    {/* <ul>{onDrop}</ul> */}
                  </div>
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
                </CardContent>
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
                  >
                    Contrary to popular belie f, Lorem Ipsum is not simply
                    random text. It has roots in a piece of classical Latin
                    literature from 45 BC, making it over 2000 years old.
                    Richard McClintock, a Latin professor at Hampden-Sydney
                    College in Virginia, looked up one of the more o bscure
                    Latin words, consectetur, from a Lorem Ipsum passage, and
                    going through the cites of the word in classical literature,
                    discovered the undoubtable source. Lorem Ipsum comes from
                    sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et
                    Malorum" (The Extremes of Good and Evil) by Cicero, written
                    in 45 BC. This book is a treatise on the theory of ethics,
                    very popular during the Renaissance. The first line of Lorem
                    Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in
                    section 1.10.32. The standard chunk of Lorem Ipsum used
                    since the 1500s is reproduced below for those interested.
                    Sections 1.10.3 2 and 1.10.33 from "de Finibus Bonorum et
                    Malorum" by Cicero are also reproduced in their exact
                    original form, accompanied by English Contrary to popular
                    belie f, Lorem Ipsum is not simply random text. It has roots
                    in a piece of classical Latin literature from 45 BC, making
                    it over 2000 years old. 
                  </Typography>
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

const mapStateToProps = (state) => {};

const mapDispatchToProps = (dispatch) => {
  return {
    buildImageActionCreators: bindActionCreators(buildImageAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect, withForm)(BuildImage);
