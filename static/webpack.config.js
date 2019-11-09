const path = require("path");
let webpack = require("webpack");
const pkg = require("./package.json");

const cssLoader = {
  loader: "css-loader"
};
cssLoader.options = {
  modules: true,
  //minimize: true,
  camelCase: true,
  importLoaders: 1,
  localIdentName: "[local]--[hash:base64:5]"
};

const rules = [
  {
    test: /\.(js|jsx|ts)$/,
    exclude: /node_modules/,
    use: {
      loader: "babel-loader"
    }
  },
  {
    test: /\.svg$/,
    loader: "svg-inline-loader"
  },
  {
    test: /\.css$/i,
    use: ["style-loader", "css-loader"]
  },

  {
    test: /\.less$/,
    exclude: /node_modules/,
    use: [
      {
        loader: "style-loader" // creates style nodes from JS strings
      },
      cssLoader,
      {
        loader: "less-loader", // compiles Less to CSS
        options: {
          modifyVars: {
            "primary-color": "#1DA57A",
            "link-color": "#1DA57A"
          }
        }
      },
      {
        loader: "postcss-loader",
        options: {
          config: {
            path: path.resolve(__dirname, "./postcss.config.js")
          }
        }
      }
    ]
  }
];

module.exports = {
  entry: {
    app: "./src/app.js"
  },
  module: {
    rules
  },
  output: {
    path: path.resolve(__dirname, "dist"),
    filename: "app-prod.js"
  },
  devtool: "cheap-module-source-map",
  externals: {
    // config is defined in template
    Config: "lel"
  },
  plugins: [
    new webpack.DefinePlugin({
      VERSION: JSON.stringify(pkg.version),
      HOMEPAGE: JSON.stringify(pkg.homepage),
      MODE: "prod"
    })
  ]
};
