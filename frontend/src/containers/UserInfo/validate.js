const validate = (values) => {
  const errors = {};
  if (!values.username) {
    errors.username = "Username không được bỏ trống";
  }

  return errors;
};

export default validate;
