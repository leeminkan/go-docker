const styles = (theme) => ({
  list: {
    padding: "0",
  },
  drawerPaper: {
    position: "relative",
    width: 240,
    maxWidth: 240,
    height: "100%",
    zIndex: 99,
    boxShadow:
      "0 0.125rem 9.375rem rgba(90,97,105,.1), 0 0.25rem 0.5rem rgba(90,97,105,.12), 0 0.9375rem 1.375rem rgba(90,97,105,.1), 0 0.4375rem 2.1875rem rgba(165,182,201,.1)",
  },
  menuLinkActive: {
    "&>div": {
      backgroundColor: theme.color.hover,
      color: "#007bff",
      borderLeft: "2px solid #007bff",
    },
  },
  menuLink: {
    textDecoration: "none",
    color: theme.text.default,
  },
  menuItem: {
    height: "50px",
  },
  menuTitle: {
    fontWeight: "450",
    fontSize: "14px",
  },
});

export default styles;
