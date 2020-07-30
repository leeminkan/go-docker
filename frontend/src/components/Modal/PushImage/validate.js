const validate = (values) => {
  const errors = {};
  if (!values.id) {
    errors.id = "Image is required";
  }

  return errors;
};

export default validate;
