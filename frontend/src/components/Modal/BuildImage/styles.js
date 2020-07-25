const styles = (theme) => ({
  root: {
    flexGrow: 1,
  },
  cardContent: {
    paddingBottom: "0 !important",
  },
  form: {
    padding: "5px 17px",
    margin: 0,
    "&>div>input": {
      padding: "10px 10px",
    },
  },
  nameImage: {
    fontWeight: "500",
    padding: "10px 10px 0 22px",
    fontSize: "14px",
  },
  submit: {
    margin: theme.spacing(3, 0, 2),
    backgroundColor: "#25944c",
    color: "white",
    "&:hover": {
      backgroundColor: "#074d19",
    },
    weight: "200px",
  },
  button: {
    display: "flex",
    justifyContent: "center",
  },
  modal: {
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    position: "absolute",
    width: "60%",
    height: "auto",
    backgroundColor: "white",
    boxShadow: "#000000",
    outline: "none",
    borderRadius: "5px",
  },
  header: {
    backgroundColor: "#20c997",
    padding: "10px 15px",
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    borderRadius: "5px 5px 0 0",
  },
  title: {
    color: "white",
    fontWeight: 700,
    textTransform: "capitalize",
    fontSize: "18px",
    marginTop: "0px",
  },
  icon: {
    cursor: "pointer",
    fontSize: 30,
    color: "white",
  },
  content: {
    padding: "10px 15px",
    maxHeight: "calc(100vh - 150px)",
    overflowY: "auto",
  },
  name: {
    paddingTop: "5px",
    borderBottom: "0.5px solid red",
  },
  textField: {
    width: "100%",
  },
  select: {
    display: "flex",
    flexWrap: "wrap",
  },
  buttonSave: {
    width: "85px",
    marginLeft: "10px",
    backgroundColor: "#20c997",
  },
  formControl: {
    marginTop: "10px",
  },
});

export default styles;
