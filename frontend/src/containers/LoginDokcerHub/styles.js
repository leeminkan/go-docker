import login from "../../assets/img/register.jpg";

const styles = (theme) => ({
  "@global": {
    body: {
      backgroundImage: `url(${login})`,
      backgroundRepeat: "no-repeat",
      backgroundPosition: "center",
      backgroundAttachment: "fixed",
    },
  },
  title: {
    fontSize: "30px",
  },
  paper: {
    marginTop: theme.spacing(4),
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
    width: "160px",
  },
});

export default styles;
