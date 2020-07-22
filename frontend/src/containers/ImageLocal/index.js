import React, { Component, Fragment } from "react";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose } from "redux";
import { Typography, Button } from "@material-ui/core";
import { CardContent, Card } from "@material-ui/core";
import DeleteIcon from "@material-ui/icons/Delete";

class ImageLocal extends Component {
  render() {
    const { classes } = this.props;

    const data = [
      {
        devicename: "Tanner Linsley",
        os: "Jetson Nano",
        deviceid: 1,
      },
      {
        devicename: "Tanner Harry",
        os: "Jetson Nano",
        deviceid: 2,
      },
    ];

    let columns = [
      {
        key: "deviceid",
        Header: "Device ID",
        accessor: "deviceid",
        //sortable: false,
        //filterable: false,
        width: 100,
      },
      {
        key: "devicename",
        Header: "Device Name",
        id: "devicename",
        accessor: "devicename",
        //sortable: false,
        //filterable: false,
      },
      {
        key: "os",
        Header: "OS",
        id: "os",
        accessor: "os",
        //sortable: false,
        //filterable: false,
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
            list image in local
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={data}
                columns={columns}
                filterable
                //pages={numberOfPages}
                //loading={loading}
                manual
                multiSort={false}
                // onFetchData={(state) => {
                //   this.props.LearnActionCreators.getListLearn(
                //     state.page,
                //     state.pageSize,
                //     state.sorted,
                //     state.filtered
                //   );
                // }}
              />
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(ImageLocal);
