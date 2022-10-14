import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";

import IndexVue from "./pages/Index.vue";
import LoginVue from "./pages/auth/Login.vue";

const About = { template: "<div>About</div>" };

// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const routes: RouteRecordRaw[] = [
  { path: "/", component: IndexVue },
  { path: "/about", component: About },

  {
    path: "/auth/login",
    component: LoginVue,
  },
];

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
export const router = createRouter({
  // 4. Provide the history implementation to use. We are using the hash history for simplicity here.
  history: createWebHashHistory(),
  routes, // short for `routes: routes`
});
