import React, { Component, Fragment } from "react";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as LocalImageAction from "./action";
import { Typography, Button } from "@material-ui/core";
import { CardContent, Card } from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";

class DeviceList extends Component {
  render() {
    const { classes, localImage } = this.props;

    let columns = [
      {
        key: "id",
        Header: "ID",
        accessor: "id",
        //sortable: false,
        //filterable: false,
        width: 60,
      },
      {
        key: "repo_name",
        Header: "Repository Name",
        id: "repo_name",
        accessor: "repo_name",
        //sortable: false,
        //filterable: false,
        width: 180,
      },
      {
        key: "image_id",
        Header: "Image ID",
        accessor: "image_id",
        //sortable: false,
        //filterable: false,
      },

      {
        key: "status",
        Header: "Status",
        id: "status",
        accessor: "status",
        //sortable: false,
        //filterable: false,
        width: 70,
      },
      {
        key: "action",
        Header: "Action",
        accessor: "action",
        width: 210,
        align: "left",
        sortable: false,
        filterable: false,
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
        <div>
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            list local image
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={localImage}
                columns={columns}
                filterable
                //pages={numberOfPages}
                //loading={true}
                manual
                multiSort={false}
                onFetchData={(state) => {
                  this.props.LocalImageActionCreators.getListLocalImage();
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
    localImage: state.localImage.listLocalImage,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    LocalImageActionCreators: bindActionCreators(LocalImageAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(DeviceList);
