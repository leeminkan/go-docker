const validate = (values) => {
  const errors = {};
  if (!values.username) {
    errors.username = "Username không được bỏ trống";
  }

  if (!values.password) {
    errors.password = "Mật khẩu không được bỏ trống";
  } else if (values.password.trim().length < 6) {
    errors.password = "Mật khẩu phải từ 6 ký tự trở lên";
  }

  if (!values.rePassword) {
    errors.rePassword = "Mật khẩu không được bỏ trống";
  } else if (values.rePassword.trim().length < 6) {
    errors.rePassword = "Mật khẩu phải từ 6 ký tự trở lên";
  }

  if (!values.isAdmin) {
    errors.isAdmin = "Role không được bỏ trống";
  }

  return errors;
};

export default validate;
