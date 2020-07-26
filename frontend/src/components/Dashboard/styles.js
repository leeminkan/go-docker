const styles = (theme) => ({
  wrapper: {
    display: "flex",
    flexDirection: "row",
    //height: "calc(100vh - 64px)",
    backgroundColor: "#5a61",
    marginTop: "64px",
    paddingLeft: "240px",
  },
  wrapperContent: {
    width: "100vw",
    padding: "10px 20px 10px 20px",
    marginLeft: -240,
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.leavingScreen,
    }),
  },
  shiftLeft: {
    marginLeft: 0,
    transition: theme.transitions.create("margin", {
      easing: theme.transitions.easing.easeOut,
      duration: theme.transitions.duration.enteringScreen,
    }),
  },
});

export default styles;
