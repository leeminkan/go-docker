const validate = (values) => {
  const errors = {};
  if (!values.id) {
    errors.id = "Image không được bỏ trống";
  }

  return errors;
};

export default validate;
