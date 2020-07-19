import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { connect } from "react-redux";
import { compose } from "redux";
import { Grid, Card, Typography, TablePagination } from "@material-ui/core";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import Paper from "@material-ui/core/Paper";

class ImageDockerHub extends Component {
  handleChangePage = (event, newPage) => {};

  handleChangeRowsPerPage = (event) => {};

  render() {
    const { classes } = this.props;

    return (
      <div className={classes.root}>
        <div>
          <Typography
            gutterBottom
            variant="h5"
            component="h2"
            className={classes.titlePage}
          >
            list image
          </Typography>
        </div>
        <Grid container spacing={2}>
          <Grid item xs={12}>
            <Card className={classes.card}>
              <TableContainer component={Paper}>
                <Table stickyHeader aria-label="sticky table">
                  <TableHead>
                    <TableRow>
                      <TableCell>Id</TableCell>
                      <TableCell align="right">Name</TableCell>
                      <TableCell align="right">Age</TableCell>
                      <TableCell align="right">Address</TableCell>
                      <TableCell align="right">City</TableCell>
                      <TableCell align="right">ContactNum</TableCell>
                      <TableCell align="right">Salary</TableCell>
                      <TableCell style={{ paddingRight: "60px" }} align="right">
                        Department
                      </TableCell>
                    </TableRow>
                  </TableHead>

                  <TableBody>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                    <TableRow>
                      <TableCell component="th" scope="row">
                        1
                      </TableCell>
                      <TableCell align="right">A</TableCell>
                      <TableCell align="right">35</TableCell>
                      <TableCell align="right">C</TableCell>
                      <TableCell align="right">HCM</TableCell>
                      <TableCell align="right">S</TableCell>
                      <TableCell align="right">FSSS</TableCell>
                      <TableCell style={{ paddingRight: "64px" }} align="right">
                        A
                      </TableCell>
                    </TableRow>
                  </TableBody>
                </Table>
              </TableContainer>
              <TablePagination
                rowsPerPageOptions={[5, 10, 15]}
                component="div"
                count="100"
                rowsPerPage="10"
                page="0"
                onChangePage={this.handleChangePage}
                onChangeRowsPerPage={this.handleChangeRowsPerPage}
              />
            </Card>
          </Grid>
        </Grid>
      </div>
    );
  }
}

const mapStateToProps = (state) => ({});

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(ImageDockerHub);
