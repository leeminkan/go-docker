import HomePage from "../containers/HomePage";
import LoginPage from "../containers/LoginPage";
import RegisterPage from "../containers/RegisterPage";
import ImageLocal from "../containers/ImageLocal";
// import BuildImage from "../containers/BuildImage";
import ReconfigNode from "../containers/ReconfigNode";
import ImageDockerHub from "../containers/ImageDockerHub";
import LoginDokcerHub from "../containers/LoginDokcerHub";
import DeviceList from "../containers/DeviceList";
import UserInfo from "../containers/UserInfo";
import DeviceDetail from "../containers/DeviceDetail";

// eslint-disable-next-line
const AuthDockerHub = localStorage.getItem("DOCKERHUB");

export const ADMIN_ROUTES = [
  {
    path: "/",
    name: "Trang Chủ",
    exact: true,
    component: HomePage,
    icon: "home",
    hidden: false,
  },

  {
    path: "/imagelist",
    name: "List Image",
    component: ImageLocal,
    icon: "dynamic_feed",
    hidden: false,
  },
  {
    path: "/image-dockerhub",
    name: "List Repository Docker Hub",
    component: ImageDockerHub,
    icon: "content_copy",
    hidden: false,
  },
  {
    path: "/device-list",
    name: "List Edge Node Device",
    component: DeviceList,
    icon: "developer_board",
    hidden: false,
  },
  {
    path: "/reconfig-node",
    name: "Reconfig Edge Node",
    component: ReconfigNode,
    icon: "cached",
    hidden: false,
  },
  {
    path: "/user-info",
    name: "User Info",
    component: UserInfo,
    icon: "perm_identity",
    hidden: false,
  },
  {
    path: "/device-detail/:id",
    component: DeviceDetail,
    hidden: true,
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
  {
    path: "/login-docker-hub",
    name: "Login Docker Hub",
    component: LoginDokcerHub,
  },
];
