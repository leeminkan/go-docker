import HomePage from "../containers/HomePage";
import LoginPage from "../containers/LoginPage";
import RegisterPage from "../containers/RegisterPage";
import ImageList from "../containers/ImageList";
import BuildImage from "../containers/BuildImage";
import ReconfigNode from "../containers/ReconfigNode";

export const ADMIN_ROUTES = [
  {
    path: "/",
    name: "Trang Chủ",
    exact: true,
    component: HomePage,
    icon: "home",
  },
  {
    path: "/buildimage",
    name: "Build Image",
    component: BuildImage,
    icon: "construction",
  },
  {
    path: "/imagelist",
    name: "List Image",
    component: ImageList,
    icon: "dynamic_feed",
  },
  {
    path: "/reconfig-node",
    name: "Reconfig Edge Node",
    component: ReconfigNode,
    icon: "cached",
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
