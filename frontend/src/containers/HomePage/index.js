import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { Grid, Card, CardContent, Typography } from "@material-ui/core";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as DeviceAction from "../DeviceList/action";
import * as ImageLocalAction from "../ImageLocal/action";
import * as ImageDHAction from "../ImageDockerHub/action";

class HomePage extends Component {
  componentDidMount() {
    this.props.ImageLocalActionCreators.getListLocalImage();
    this.props.ImageDHActionCreators.getListDockerHubImage();
    this.props.DeviceActionCreators.getListDevice();
  }

  render() {
    const { classes, localImage, device, dockerHubImage } = this.props;

    return (
      <div className={classes.root}>
        <div>
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            statistic
          </Typography>
        </div>
        <Grid container spacing={2}>
          <Grid item xs={6} sm={3}>
            <Card small className={classes.card}>
              <CardContent>
                <div className={classes.cardBody}>
                  <Typography
                    gutterBottom
                    variant="h5"
                    component="h2"
                    className={classes.title}
                  >
                    Image In Server
                  </Typography>
                  <Typography
                    gutterBottom
                    variant="h3"
                    component="h1"
                    className={classes.number}
                  >
                    {localImage ? localImage.length : 0}
                  </Typography>
                </div>
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={6} sm={3}>
            <Card small className={classes.card}>
              <CardContent>
                <div className={classes.cardBody}>
                  <Typography
                    gutterBottom
                    variant="h5"
                    component="h2"
                    className={classes.title}
                  >
                    Repository Docker Hub
                  </Typography>
                  <Typography
                    gutterBottom
                    variant="h3"
                    component="h1"
                    className={classes.number}
                  >
                    {dockerHubImage ? dockerHubImage.length : 0}
                  </Typography>
                </div>
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={6} sm={3}>
            <Card small className={classes.card}>
              <CardContent>
                <div className={classes.cardBody}>
                  <Typography
                    gutterBottom
                    variant="h5"
                    component="h2"
                    className={classes.title}
                  >
                    Edge Node Device
                  </Typography>
                  <Typography
                    gutterBottom
                    variant="h3"
                    component="h1"
                    className={classes.number}
                  >
                    {device ? device.length : 0}
                  </Typography>
                </div>
              </CardContent>
            </Card>
          </Grid>
          <Grid item xs={6} sm={3}>
            <Card small className={classes.card}>
              <CardContent>
                <div className={classes.cardBody}>
                  <Typography
                    gutterBottom
                    variant="h5"
                    component="h2"
                    className={classes.title}
                  >
                    Account
                  </Typography>
                  <Typography
                    gutterBottom
                    variant="h3"
                    component="h1"
                    className={classes.number}
                  >
                    5
                  </Typography>
                </div>
              </CardContent>
            </Card>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    device: state.device.listDevice,
    localImage: state.localImage.listLocalImage,
    dockerHubImage: state.DHImage.listDHImage,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    DeviceActionCreators: bindActionCreators(DeviceAction, dispatch),
    ImageLocalActionCreators: bindActionCreators(ImageLocalAction, dispatch),
    ImageDHActionCreators: bindActionCreators(ImageDHAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(HomePage);
