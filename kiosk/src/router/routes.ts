import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Signup from "../views/Signup.vue";
import Logout from "../views/Logout.vue";

export default {
  home: {
    path: "/",
    name: "home",
    component: Home
  },
  login: {
    path: "/login",
    name: "login",
    component: Login
  },
  signup: {
    path: "/signup",
    name: "signup",
    component: Signup
  },
  logout: {
    path: "/logout",
    name: "logout",
    component: Logout
  },
};
