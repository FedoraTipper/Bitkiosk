import Home from "../views/Home.vue";
import Login from "../views/Login.vue";
import Signup from "../views/Signup.vue";

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
  }
};
