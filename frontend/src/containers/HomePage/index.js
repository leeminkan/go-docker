import React, { Component } from "react";
import { withStyles } from "@material-ui/styles";
import styles from "./styles";
import { Grid, Card, CardContent, Typography } from "@material-ui/core";
import { connect } from "react-redux";
import { compose } from "redux";

class HomePage extends Component {
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
                    Edge Node
                  </Typography>
                  <Typography
                    gutterBottom
                    variant="h3"
                    component="h1"
                    className={classes.number}
                  >
                    1
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
                    Image
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
                    Image
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
                    Image
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

const mapStateToProps = (state) => {};

const mapDispatchToProps = (dispatch) => {};

const withConnect = connect(mapStateToProps, mapDispatchToProps);

export default compose(withStyles(styles), withConnect)(HomePage);
