import React, { Component, Fragment } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import {
  Card,
  CardContent,
  Typography,
  Divider,
  Grid,
  Box,
  Button,
} from "@material-ui/core";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import PullImage from "../../components/Modal/PullImage";
import * as DeviceDetailAction from "./action";
import * as DockerHubImageAction from "../ImageDockerHub/action";
import logo from "../../assets/img/pending.gif";
import ConvertTime from "../../helpers/convertTime";

class DeviceDetail extends Component {
  componentDidMount() {
    if (this.props.dockerHubImage.length === 0) {
      this.props.DHImageActionCreators.getListDockerHubImage();
    }
  }

  onCloseModalPullImage = () => {
    this.props.DeviceDetailActionCreators.closeModalPullImage();
  };

  openModalPullImage = () => {
    this.props.DeviceDetailActionCreators.openModalPullImage();
  };

  onSubmit = (data) => {
    this.props.DeviceDetailActionCreators.pullImage(data);
    this.props.DeviceDetailActionCreators.closeModalPullImage();
  };

  onRunContainer = (data) => {
    this.props.DeviceDetailActionCreators.runImageDevice(data);
  };

  render() {
    const {
      classes,
      openModalPullImage,
      dockerHubImage,
      imageInDevice,
      containerInDevice,
    } = this.props;

    let numberOfPagesImage = imageInDevice
      ? Math.floor(imageInDevice.length / 10) + 1
      : 1;

    let numberOfPagesContainer = containerInDevice
      ? Math.floor(containerInDevice.length / 10) + 1
      : 1;

    let columnsImage = [
      {
        key: "id",
        Header: "ID",
        id: "id",
        accessor: "id",
        width: 60,
      },
      {
        key: "repo_name",
        Header: "Repository Name",
        id: "repo_name",
        accessor: "repo_name",
        width: 150,
      },
      {
        key: "image_id",
        Header: "Image ID",
        accessor: "image_id",
        id: "image_id",
      },
      {
        key: "tag",
        Header: "Tag",
        accessor: "tag",
        id: "tag",
        width: 100,
        Cell: (data) => {
          if (data.value) {
            return <div className={classes.status}>{data.value}</div>;
          }
        },
      },
      {
        key: "status",
        Header: "Status",
        id: "status",
        accessor: "status",
        width: 100,
        Cell: (data) => {
          if (data.value === "on progress") {
            return (
              <Fragment>
                <div className={classes.logoImage}>
                  <img src={logo} alt="logo" className={classes.img} />
                </div>
              </Fragment>
            );
          } else {
            return (
              <div className={classes.status}>
                {data.value.charAt(0).toUpperCase() + data.value.slice(1)}
              </div>
            );
          }
        },
      },
      {
        key: "created_on",
        Header: "Create Time",
        id: "created_on",
        accessor: "created_on",
        width: 160,
        Cell: (data) => {
          if (data) {
            return ConvertTime(data.value);
          }
        },
      },
    ];

    let columnsContainer = [
      {
        key: "id",
        Header: "ID",
        id: "id",
        accessor: "id",
        width: 60,
      },
      {
        key: "repo_name",
        Header: "Repository Name",
        id: "repo_name",
        accessor: "repo_name",
        width: 150,
      },
      {
        key: "image_id",
        Header: "Image ID",
        accessor: "image_id",
        id: "image_id",
      },
      {
        key: "tag",
        Header: "Tag",
        accessor: "tag",
        id: "tag",
        width: 100,
        Cell: (data) => {
          if (data.value) {
            return <div className={classes.status}>{data.value}</div>;
          }
        },
      },
      {
        key: "status",
        Header: "Status",
        id: "status",
        accessor: "status",
        width: 100,
        Cell: (data) => {
          if (data.value === "on progress") {
            return (
              <Fragment>
                <div className={classes.logoImage}>
                  <img src={logo} alt="logo" className={classes.img} />
                </div>
              </Fragment>
            );
          } else {
            return (
              <div className={classes.status}>
                {data.value.charAt(0).toUpperCase() + data.value.slice(1)}
              </div>
            );
          }
        },
      },
      {
        key: "created_on",
        Header: "Create Time",
        id: "created_on",
        accessor: "created_on",
        width: 160,
        Cell: (data) => {
          if (data) {
            return ConvertTime(data.value);
          }
        },
      },
    ];

    return (
      <div className={classes.root}>
        <div>
          <PullImage
            openModalPullImage={openModalPullImage}
            onCloseModalPullImage={this.onCloseModalPullImage}
            onSave={this.onSubmit}
            dockerHubImage={dockerHubImage}
          />
          <Grid container>
            <Grid item xs={10}>
              <Typography
                gutterBottom
                variant="h5"
                component="h2"
                className={classes.titlePage}
              >
                list image
              </Typography>
            </Grid>
            <Grid item xs={2}>
              <Box flexDirection="row-reverse" display="flex">
                <Button
                  variant="outlined"
                  size="small"
                  color="primary"
                  className={classes.button}
                  onClick={this.openModalPullImage}
                >
                  {/* <AddIcon className={classes.leftIcon} /> */}
                  Pull Image
                </Button>
              </Box>
            </Grid>
          </Grid>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={imageInDevice}
                columns={columnsImage}
                minRows={5}
                pages={numberOfPagesImage}
                manual
                onFetchData={(state) => {
                  this.props.DeviceDetailActionCreators.getListImageInDevice();
                }}
              />
            </CardContent>
          </Card>
          <Divider className={classes.divider} variant="middle" />
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            list container
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={containerInDevice}
                columns={columnsContainer}
                minRows={5}
                pages={numberOfPagesContainer}
                manual
                onFetchData={(state) => {
                  this.props.DeviceDetailActionCreators.getListContainerInDevice();
                }}
              />
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    dockerHubImage: state.DHImage.listDHImage,
    openModalPullImage: state.deviceDetail.openModalPullImage,
    imageInDevice: state.deviceDetail.imageInDevice,
    containerInDevice: state.deviceDetail.containerInDevice,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    DHImageActionCreators: bindActionCreators(DockerHubImageAction, dispatch),
    DeviceDetailActionCreators: bindActionCreators(
      DeviceDetailAction,
      dispatch
    ),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(DeviceDetail);
