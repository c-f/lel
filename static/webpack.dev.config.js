const path = require("path");
let webpack = require("webpack");
const pkg = require("./package.json");
let master = require("./webpack.config");

master.output = {
  path: path.resolve(__dirname, "dist"),
  filename: "lel-dev.js"
};

master.devServer = {
  //	watchContentBase: true,
  publicPath: "/static/",
  contentBase: path.join(__dirname, "dist"),
  proxy: {
    "/static/icons/*": {
      target: "http://127.0.0.1:8080",
      pathRewrite: { "^/static/": "" }
    },
    "/static/vendor/*": {
      target: "http://127.0.0.1:8080",
      pathRewrite: { "^/static/": "" }
    }
  },

  filename: "lel-dev.js"
};

module.exports = master;
