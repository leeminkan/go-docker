import React, { Component, Fragment } from "react";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as LocalImageAction from "./action";
import { Typography, Button, Box, Grid } from "@material-ui/core";
import { CardContent, Card } from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";
import logo from "../../assets/img/pending.gif";
import BuildImage from "../../components/Modal/BuildImage";

class LocalImage extends Component {
  onCloseModalBuildImage = () => {
    this.props.LocalImageActionCreators.closeModalBuildImage();
  };

  openModalBuildImage = () => {
    this.props.LocalImageActionCreators.openModalBuildImage();
  };

  onSubmit = (data) => {
    this.props.LocalImageActionCreators.buildImage(data);
    this.props.LocalImageActionCreators.closeModalBuildImage();
  };

  // shouldComponentUpdate() {
  //   if (this.props.imagePending) {
  //     this.props.LocalImageActionCreators.getImageById(this.props.imagePending);
  //   }
  // }

  render() {
    const { classes, localImage, openModalBuildImage } = this.props;

    let columns = [
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
                {/* <div className={classes.status}>{data.value}</div> */}
              </Fragment>
            );
          } else {
            return <div className={classes.status}>{data.value}</div>;
          }
        },
      },
      {
        key: "action",
        Header: "Action",
        accessor: "action",
        width: 210,
        align: "left",
        Cell: (data) => {
          return (
            <Fragment>
              <Button
                variant="outlined"
                color="primary"
                className={classes.icon}
                startIcon={<DeleteIcon />}
                //onClick={() => this.openModalEditLearn(data.original)}
              >
                Edit
              </Button>
              <Button
                variant="outlined"
                color="secondary"
                className={classes.icon}
                startIcon={<DeleteIcon />}
                //onClick={() => this.openModalEditLearn(data.original)}
              >
                Delete
              </Button>
            </Fragment>
          );
        },
      },
    ];

    return (
      <div className={classes.root}>
        <BuildImage
          openModalBI={openModalBuildImage}
          onCloseModalBI={this.onCloseModalBuildImage}
          onSave={this.onSubmit}
        />
        <Grid container>
          <Grid item xs={10}>
            <Typography
              gutterBottom
              variant="h5"
              component="h2"
              className={classes.titlePage}
            >
              list local image
            </Typography>
          </Grid>
          <Grid item xs={2}>
            <Box flexDirection="row-reverse" display="flex">
              <Button
                variant="outlined"
                size="small"
                color="primary"
                className={classes.button}
                onClick={this.openModalBuildImage}
              >
                {/* <AddIcon className={classes.leftIcon} /> */}
                Build Image
              </Button>
            </Box>
          </Grid>
        </Grid>
        <Card>
          <CardContent className={classes.center}>
            <ReactTable
              className="-highlight"
              defaultPageSize={10}
              data={localImage}
              columns={columns}
              //filterable
              //pages={numberOfPages}
              //loading={true}
              manual
              //multiSort={false}
              onFetchData={(state) => {
                this.props.LocalImageActionCreators.getListLocalImage();
              }}
              getTdProps={(state, rowInfo, column) => {
                return {
                  style: {
                    background:
                      rowInfo && rowInfo.original.status === "on progress"
                        ? "#fdde53"
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
    localImage: state.localImage.listLocalImage,
    openModalBuildImage: state.localImage.openModalBuildImage,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    LocalImageActionCreators: bindActionCreators(LocalImageAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(LocalImage);
