const validate = (values) => {
  const errors = {};
  if (!values.tag) {
    errors.tag = "Tag không được bỏ trống";
  } else if (values.tag && !/^[a-z0-9_.-]{0,}$/i.test(values.tag)) {
    errors.tag = "Tag không hợp lệ. Chỉ chấp nhận kí tự, chữ số, _ và -.";
  }

  return errors;
};

export default validate;
