const validate = (values) => {
  const errors = {};
  if (!values.image) {
    errors.image = "Image không được bỏ trống";
  }

  return errors;
};

export default validate;
