import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { Grid, Card, CardContent, Typography } from "@material-ui/core";
import { connect } from "react-redux";
import { compose } from "redux";

class ReconfigNode extends Component {
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
            Reconfig Node
          </Typography>
        </div>
        <Grid container>
          <Grid item xs={12}>
            <Card className={classes.card}>
              <CardContent className={classes.cardContent}></CardContent>
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

export default compose(withStyles(styles), withConnect)(ReconfigNode);
