const validate = (values) => {
  const errors = {};
  if (!values.name) {
    errors.name = "Name is required";
  } else if (values.name && !/^[a-z0-9_.-]{0,}$/i.test(values.name)) {
    errors.name = "Tag is invalid. Only accept [0-9a-z_.-]";
  }

  return errors;
};

export default validate;
