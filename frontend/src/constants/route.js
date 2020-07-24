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

// eslint-disable-next-line
const AuthDockerHub = localStorage.getItem("DOCKERHUB");

export const ADMIN_ROUTES = [
  {
    path: "/",
    name: "Trang Chủ",
    exact: true,
    component: HomePage,
    icon: "home",
  },
  // {
  //   path: "/buildimage",
  //   name: "Build Image",
  //   component: BuildImage,
  //   icon: "construction",
  // },
  {
    path: "/login-docker-hub",
    name: "Login Docker Hub",
    component: LoginDokcerHub,
    icon: "lock",
  },
  {
    path: "/imagelist",
    name: "List Image",
    component: ImageLocal,
    icon: "dynamic_feed",
  },
  {
    path: "/image-dockerhub",
    name: "List Image Docker Hub",
    component: ImageDockerHub,
    icon: "content_copy",
  },
  {
    path: "/device-list",
    name: "List Device",
    component: DeviceList,
    icon: "developer_board",
  },
  {
    path: "/reconfig-node",
    name: "Reconfig Edge Node",
    component: ReconfigNode,
    icon: "cached",
  },
  {
    path: "/user-info",
    name: "User Info",
    component: UserInfo,
    icon: "perm_identity",
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
