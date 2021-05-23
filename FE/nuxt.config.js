const envPath = `../.env`
require("dotenv").config({ path: envPath })
const { NODE_ENV, API_BASE_URL } = process.env
export default {
  // const nuxtConfig: Configuration = {
  buildModules: ["@nuxt/typescript-build", "@nuxtjs/vuetify"],
  /*
   ** Headers of the page
   */
  head: {
    title: process.env.npm_package_name || "",
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      {
        hid: "description",
        name: "description",
        content: process.env.npm_package_description || ""
      }
    ],
    link: [{ rel: "icon", type: "image/x-icon", href: "/favicon.ico" }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: "#fff" },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],
  typescript: {
    typeCheck: true,
    ignoreNotFoundWarnings: true
  },
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    "@nuxtjs/axios",
    "@nuxtjs/eslint-module",
    "@nuxtjs/style-resources"
  ],
  axios: { proxy: true },
  proxy: {
    "/api": {
      target: API_BASE_URL,
      pathRewrite: {
        "^/api/v1": "/api/v1/"
      }
    }
  },
  styleResources: {
    scss: ["~/assets/scss/_index.scss"]
  },
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  /*
   ** Build configuration
   */
  build: {
    extend(config, ctx) {
      // Run ESLint on save
      if (ctx.isDev && process.isClient) {
        // <= `ctx.isClient` を `process.isClient` に変更！
        config.module.rules.push({
          enforce: "pre",
          test: /\.(js|vue)$/,
          loader: "eslint-loader",
          exclude: /(node_modules)/
        })
      }
    }
  },
  srcDir: "src/",
  env: {
    IS_ENV_DEVELOPMENT: NODE_ENV === "development",
    API_BASE_URL
  }
}
