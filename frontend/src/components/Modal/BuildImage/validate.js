const validate = (values) => {
  const errors = {};
  if (!values.tag) {
    errors.tag = "Tag is required";
  } else if (values.tag && !/^[a-z0-9_.-]{0,}$/i.test(values.tag)) {
    errors.tag = "Tag is invalid. Only accept [0-9a-z_.-]";
  }

  return errors;
};

export default validate;
