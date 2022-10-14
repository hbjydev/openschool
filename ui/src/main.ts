import { createApp } from "vue";
import App from "./App.vue";
import { router } from "./router";

import "@fontsource/inter/400.css";
import "@fontsource/inter/700.css";
import "./style.css";

const app = createApp(App);

app.use(router);

app.mount("#app");
