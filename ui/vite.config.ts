import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  base: process.env.IS_DEV !== "true" ? "./" : "/",
  build: {
    outDir: "app/build",
  },
  server: {
    port: 3000,
  },
  plugins: [vue()],
});
