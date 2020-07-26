import { withStyles } from "@material-ui/core";
import AppBar from "@material-ui/core/AppBar";
import IconButton from "@material-ui/core/IconButton";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import MenuIcon from "@material-ui/icons/Menu";
import React, { Component } from "react";
import AccountCircle from "@material-ui/icons/AccountCircle";
import MenuItem from "@material-ui/core/MenuItem";
import Menu from "@material-ui/core/Menu";
import { withRouter } from "react-router";
import styles from "./styles";

const menuId = "primary-search-account-menu";
class Header extends Component {
  state = {
    anchorEl: null,
  };

  handleProfileMenuOpen = (event) => {
    this.setAnchorEl(event.currentTarget);
  };

  setAnchorEl(currentTarget) {
    this.setState({
      anchorEl: currentTarget,
    });
  }

  handleMenuClose = () => {
    this.setAnchorEl(null);
  };

  renderDesktopMenu() {
    const isMenuOpen = Boolean(this.state.anchorEl);
    const { anchorEl } = this.state;
    return (
      <Menu
        anchorEl={anchorEl}
        anchorOrigin={{ vertical: "top", horizontal: "right" }}
        id={menuId}
        keepMounted
        transformOrigin={{ vertical: "top", horizontal: "right" }}
        open={isMenuOpen}
        onClose={this.handleMenuClose}
      >
        <MenuItem onClick={this.handleLogout}>Logout</MenuItem>
      </Menu>
    );
  }

  handleToggleSidebar = () => {
    const { onToggleSidebar, open } = this.props;
    if (onToggleSidebar) {
      onToggleSidebar(!open);
    }
  };

  handleLogout = () => {
    localStorage.removeItem("JWT_TOKEN");
    const { history } = this.props;
    if (history) {
      history.push("/login");
    }
  };

  render() {
    const { classes } = this.props;
    return (
      <>
        <AppBar position="fixed">
          <Toolbar>
            <IconButton
              edge="start"
              className={classes.menuButton}
              color="inherit"
              aria-label="Open drawer"
              onClick={this.handleToggleSidebar}
            >
              <MenuIcon />
            </IconButton>
            <Typography className={classes.title} variant="h6" noWrap>
              Đồ án chuyên ngành
            </Typography>

            <div className={classes.grow} />
            <div className={classes.sectionDesktop}>
              <IconButton
                edge="end"
                aria-label="Account of current user"
                aria-controls={menuId}
                aria-haspopup="true"
                onClick={this.handleProfileMenuOpen}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
            </div>
          </Toolbar>
        </AppBar>
        {this.renderDesktopMenu()}
      </>
    );
  }
}

export default withStyles(styles)(withRouter(Header));
