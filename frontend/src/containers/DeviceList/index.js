import React, { Component } from "react";
import ReactTable from "react-table-6";
import "react-table-6/react-table.css";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose, bindActionCreators } from "redux";
import * as DeviceAction from "./action";
import { Typography } from "@material-ui/core";
import { CardContent, Card } from "@material-ui/core";

class DeviceList extends Component {
  render() {
    const { classes, device } = this.props;
    let numberOfPages = device ? Math.floor(device.length / 10) + 1 : 1;

    let columns = [
      {
        key: "id",
        Header: "ID",
        accessor: "id",
        width: 80,
      },
      {
        key: "machine_id",
        Header: "Machine ID",
        accessor: "machine_id",
      },
      {
        key: "device_name",
        Header: "Device Name",
        id: "device_name",
        accessor: "device_name",
        width: 200,
        Cell: (data) => {
          if (data.original.device_name) {
            return (
              <div className={classes.status}>{data.original.device_name}</div>
            );
          }
        },
      },
      {
        key: "os",
        Header: "OS",
        id: "os",
        accessor: "os",
        width: 100,
        Cell: (data) => {
          if (data.original) {
            return <div className={classes.status}>{data.original.os}</div>;
          }
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
            list edge node
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={device}
                columns={columns}
                pages={numberOfPages}
                manual
                onFetchData={(state) => {
                  this.props.DeviceActionCreators.getListDevice();
                }}
                getTdProps={(state, rowInfo, column) => {
                  if (rowInfo) {
                    return {
                      onClick: (e, handleOriginal) => {
                        this.props.history.push(
                          `/device-detail/${rowInfo.original.id}`
                        );
                      },
                      style: {
                        cursor: "pointer",
                        padding: "12px 9px 9px 9px",
                      },
                    };
                  }
                  return {};
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
    device: state.device.listDevice,
  };
};

const mapDispatchToProps = (dispatch) => {
  return {
    DeviceActionCreators: bindActionCreators(DeviceAction, dispatch),
  };
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(DeviceList);
