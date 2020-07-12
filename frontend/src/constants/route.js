import HomePage from "../containers/HomePage";
import LoginPage from "../containers/LoginPage";
import RegisterPage from "../containers/RegisterPage";
import ListImage from "../containers/ListImage";

export const ADMIN_ROUTES = [
  {
    path: "/",
    name: "Trang Chủ",
    number: 1,
    exact: true,
    component: HomePage,
  },
  {
    path: "/imagelist",
    name: "Image",
    number: 2,
    component: ListImage,
  },
];

export const ROUTES = [
  {
    path: "/login",
    name: "Đăng Nhập",
    component: LoginPage,
  },
  {
    path: "/register",
    name: "Đăng Ký",
    component: RegisterPage,
  },
];
