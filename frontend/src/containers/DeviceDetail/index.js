import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { Card, CardContent, Typography, Divider } from "@material-ui/core";
import { connect } from "react-redux";
import { compose } from "redux";
import ReactTable from "react-table-6";

class DeviceDetail extends Component {
  render() {
    const { classes } = this.props;
    const device = [
      {
        name: "Tanner Linsley",
        age: 26,
        friend: {
          name: "Jason Maurer",
          age: 23,
        },
      },
    ];

    const columns = [
      {
        Header: "Name",
        accessor: "name",
      },
      {
        Header: "Age",
        accessor: "age",
        Cell: (props) => <span className="number">{props.value}</span>,
      },
      {
        id: "friendName",
        Header: "Friend Name",
        accessor: (d) => d.friend.name,
      },
      {
        Header: (props) => <span>Friend Age</span>,
        accessor: "friend.age",
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
            List image in device
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={device}
                columns={columns}
                minRows={5}
                //pages={numberOfPages}
                manual
                onFetchData={(state) => {
                  //this.props.DeviceActionCreators.getListDevice();
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
            list container in device
          </Typography>
          <Card>
            <CardContent>
              <ReactTable
                className="-striped -highlight"
                defaultPageSize={10}
                data={device}
                columns={columns}
                minRows={5}
                //pages={numberOfPages}
                manual
                onFetchData={(state) => {
                  //this.props.DeviceActionCreators.getListDevice();
                }}
              />
            </CardContent>
          </Card>
        </div>
      </div>
    );
  }
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {
  return {};
};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(DeviceDetail);
