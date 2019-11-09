const path = require("path");
let webpack = require("webpack");
const pkg = require("./package.json");
let master = require("./webpack.config");

master.output = {
  path: path.resolve(__dirname, "dist"),
  filename: "lel-dev.js"
};
master.plugins = [
  new webpack.DefinePlugin({
    VERSION: JSON.stringify(pkg.version),
    HOMEPAGE: JSON.stringify(pkg.homepage),
    MODE: JSON.stringify("dev")
  })
];

master.devServer = {
  watchContentBase: true,
  publicPath: "/static/",
  contentBase: path.join(__dirname, "dist"),
  proxy: {
    "/static/icons/*": {
      target: "https://127.0.0.1:8080",
      pathRewrite: { "^/static/": "" }
    },
    "/static/vendor/*": {
      target: "https://127.0.0.1:8080",
      pathRewrite: { "^/static/": "" }
    }
  },
  port: 8080,
  host: "127.0.0.1",
  https: true,
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
    "Access-Control-Allow-Headers":
      "X-Requested-With, content-type, Authorization"
  },
  filename: "lel-dev.js"
};

module.exports = master;
