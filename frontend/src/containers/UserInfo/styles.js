const styles = (theme) => ({
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
    minHeight: "calc(100vh - 142px)",
    borderRadius: "15px",
    boxShadow:
      "0 0.125rem 9.375rem rgba(90,97,105,.1), 0 0.25rem 0.5rem rgba(90,97,105,.12), 0 0.9375rem 1.375rem rgba(90,97,105,.1), 0 0.4375rem 2.1875rem rgba(165,182,201,.1)",
  },
  cardContent: {
    paddingBottom: "0 !important",
  },
  paper: {
    marginTop: theme.spacing(1),
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
  },
  avatar: {
    margin: theme.spacing(1),
    backgroundColor: theme.palette.primary.main,
  },
  form: {
    width: "100%",
    marginTop: theme.spacing(1),
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
    backgroundColor: "#007bff",
    color: "white",
    "&:hover": {
      backgroundColor: "#007bff",
    },
    weight: "200px",
  },
  question: {
    fontSize: "14px",
    marginRight: "5px",
  },
  logoImage: {
    clipPath: "polygon(0% 0, 100% 0, 100% 100%, 0 100%)",
    padding: "5px 10px 10px 10px",
    borderRadius: "5px",
  },
  img: {
    width: "130px",
  },
  button: {
    display: "flex",
    justifyContent: "flex-end",
  },
});

export default styles;
