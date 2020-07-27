const styles = () => ({
  root: {
    flexGrow: 1,
  },
  titlePage: {
    fontSize: "25px",
    fontWeight: "500",
    color: "#3d5170",
    padding: "10px 0",
    textTransform: "uppercase",
  },
  card: {
    minHeight: "calc(100vh)",
    borderRadius: "15px",
    boxShadow:
      "0 0.125rem 9.375rem rgba(90,97,105,.1), 0 0.25rem 0.5rem rgba(90,97,105,.12), 0 0.9375rem 1.375rem rgba(90,97,105,.1), 0 0.4375rem 2.1875rem rgba(165,182,201,.1)",
  },
  cardContent: {
    paddingBottom: "0 !important",
  },
  divider: {
    margin: "30px 0 0 0",
  },
  button: {
    margin: "10px 10px 0 0",
    fontSize: "18px",
  },
});

export default styles;
