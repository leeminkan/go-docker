const styles = (theme) => ({
  list: {
    padding: "0",
  },
  drawerPaper: {
    position: "inherit",
    width: 240,
    marginRight: 10,
    maxWidth: 240,
    minHeight: "120%",
    zIndex: 99,
    boxShadow:
      "0 0.125rem 9.375rem rgba(90,97,105,.1), 0 0.25rem 0.5rem rgba(90,97,105,.12), 0 0.9375rem 1.375rem rgba(90,97,105,.1), 0 0.4375rem 2.1875rem rgba(165,182,201,.1)",
  },
  menuLinkActive: {
    "&>div": {
      backgroundColor: theme.color.hover,
      color: "#007bff",
      //borderLeft: "2px solid #007bff",
    },
  },
  menuLink: {
    "&:hover": {
      color: "#007bff",
    },
    textDecoration: "none",
    color: "#3d5170",
  },
  menuItem: {
    height: "50px",
  },
  menuTitle: {
    fontWeight: "420",
    fontSize: "13px",
  },
  icon: {
    "&>svg": {
      paddingTop: 2,
    },
    display: "inline-block",
    textAlign: "center",
    paddingRight: 8,
    fontSize: 10,
  },
  tool: {
    padding: "5px 8px",
  },
});

export default styles;
