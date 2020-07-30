import React, { Component, Fragment } from "react";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as DockerHubImageAction from "./action";
import * as LocalImageAction from "../ImageLocal/action";
import { Typography, Button, Box, Grid } from "@material-ui/core";
import { CardContent, Card } from "@material-ui/core";
import ConvertTime from "../../helpers/convertTime";
import pending from "../../assets/img/pending.gif";
import PushImage from "../../components/Modal/PushImage";
import { showLoading } from "../../helpers/loading";

class DockerHubImage extends Component {
  onCloseModalPushImage = () => {
    this.props.DHImageActionCreators.closeModalPushImage();
  };

  openModalPushImage = () => {
    this.props.DHImageActionCreators.openModalPushImage();
  };

  componentDidMount() {
    if (this.props.localImage.length === 0) {
      this.props.LocalImageActionCreators.getListLocalImage();
    }
  }

  onSubmit = (data) => {
    this.props.DHImageActionCreators.pushImage(data);
    this.props.DHImageActionCreators.closeModalPushImage();
    showLoading(true);
  };

  render() {
    const {
      classes,
      dockerHubImage,
      openModalPushImage,
      localImage,
    } = this.props;

    let numberOfPages = dockerHubImage
      ? Math.floor(dockerHubImage.length / 10) + 1
      : 1;

    let columns = [
      {
        key: "id",
        Header: "ID",
        id: "id",
        accessor: "id",
        width: 80,
      },
      {
        key: "repo_name",
        Header: "Repository Name",
        id: "repo_name",
        accessor: "repo_name",
        Cell: (data) => {
          if (data.original.TagDockerHub) {
            return data.original.TagDockerHub.RepoDockerHub.repo_name;
          }
        },
      },
      {
        key: "tag",
        Header: "Tag",
        id: "tag",
        accessor: "tag",
        width: 120,
        Cell: (data) => {
          if (data.original.TagDockerHub) {
            return (
              <div className={classes.status}>
                {data.original.TagDockerHub.tag}
              </div>
            );
          }
        },
      },
      {
        key: "status",
        Header: "Status",
        id: "status",
        accessor: "status",
        width: 120,
        Cell: (data) => {
          if (data.value === "on progress") {
            return (
              <Fragment>
                <div className={classes.logoImage}>
                  <img src={pending} alt="logo" className={classes.img} />
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
        key: "username",
        Header: "Created User",
        id: "username",
        accessor: "username",
        width: 120,
        Cell: (data) => {
          if (data.original.User) {
            return (
              <div className={classes.status}>
                {data.original.User.username}
              </div>
            );
          }
        },
      },
      {
        key: "created_on",
        Header: "Created Time",
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
        <PushImage
          openModalPI={openModalPushImage}
          onCloseModalPI={this.onCloseModalPushImage}
          onSave={this.onSubmit}
          localImage={localImage}
        />
        <Grid container>
          <Grid item xs={10}>
            <Typography
              gutterBottom
              variant="h5"
              component="h2"
              className={classes.titlePage}
            >
              list docker hub image
            </Typography>
          </Grid>
          <Grid item xs={2}>
            <Box flexDirection="row-reverse" display="flex">
              <Button
                variant="outlined"
                size="small"
                color="primary"
                className={classes.button}
                onClick={this.openModalPushImage}
              >
                {/* <AddIcon className={classes.leftIcon} /> */}
                Push Image
              </Button>
            </Box>
          </Grid>
        </Grid>
        <Card>
          <CardContent className={classes.center}>
            <ReactTable
              className="-highlight -striped"
              defaultPageSize={10}
              data={dockerHubImage}
              columns={columns}
              pages={numberOfPages}
              manual
              onFetchData={(state) => {
                this.props.DHImageActionCreators.getListDockerHubImage();
              }}
              getTdProps={(state, rowInfo, column) => {
                return {
                  style: {
                    padding: "12px 9px 9px 9px",
                    background:
                      rowInfo && rowInfo.original.status === "on progress"
                        ? "#fdde53"
                        : "",
                    color:
                      rowInfo && rowInfo.original.status === "fail"
                        ? "red"
                        : "",
                  },
                };
              }}
            />
          </CardContent>
        </Card>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    dockerHubImage: state.DHImage.listDHImage,
    openModalPushImage: state.DHImage.openModalPushImage,
    localImage: state.localImage.listLocalImage,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    DHImageActionCreators: bindActionCreators(DockerHubImageAction, dispatch),
    LocalImageActionCreators: bindActionCreators(LocalImageAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(DockerHubImage);
