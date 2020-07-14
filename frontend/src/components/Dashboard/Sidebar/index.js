import React, { Component } from "react";
import Drawer from "@material-ui/core/Drawer";
import { List, ListItem, Typography, Toolbar } from "@material-ui/core";
import { withStyles } from "@material-ui/styles";
import { NavLink } from "react-router-dom";
import styles from "./styles";
import { ADMIN_ROUTES } from "../../../constants/route";

class Sidebar extends Component {
  toggleDrawer = (value) => {
    const { onToggleSidebar } = this.props;
    if (onToggleSidebar) {
      onToggleSidebar(value);
    }
  };

  renderList() {
    const { classes } = this.props;
    let xhtml = null;
    xhtml = (
      <List component={classes.list}>
        {ADMIN_ROUTES.map((item, index) => {
          return (
            <NavLink
              key={item.path}
              to={item.path}
              exact={item.exact}
              className={classes.menuLink}
              activeClassName={classes.menuLinkActive}
            >
              <ListItem className={classes.menuItem} button>
                <Toolbar className={classes.tool}>
                  <div className={classes.icon}>
                    <span className="material-icons">{item.icon}</span>
                  </div>
                  <Typography className={classes.menuTitle}>
                    {item.name}
                  </Typography>
                </Toolbar>
              </ListItem>
            </NavLink>
          );
        })}
      </List>
    );
    return xhtml;
  }

  render() {
    const { classes, showSidebar } = this.props;
    return (
      <Drawer
        open={showSidebar}
        onClose={() => this.toggleDrawer(false)}
        variant="persistent"
        classes={{
          paper: classes.drawerPaper,
        }}
      >
        {this.renderList()}
      </Drawer>
    );
  }
}

export default withStyles(styles)(Sidebar);
