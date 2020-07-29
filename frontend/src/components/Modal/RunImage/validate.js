const validate = (values) => {
  const errors = {};
  if (!values.name) {
    errors.name = "Name không được bỏ trống";
  } else if (values.name && !/^[a-z0-9_.-]{0,}$/i.test(values.name)) {
    errors.name = "Tag không hợp lệ. Chỉ chấp nhận kí tự, chữ số, _ và -.";
  }

  return errors;
};

export default validate;
