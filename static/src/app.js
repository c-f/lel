import React from "react";
import ReactDOM from "react-dom";
import LelApp from "./components/LelApp";

ReactDOM.render(<LelApp />, document.getElementById("app"));

// Development
if (module.hot) {
  // Whenever a new version of App.js is available
  module.hot.accept("./components/LelApp", function() {
    // Require the new version and render it instead
    var LelApp = require("./components/LelApp");
    ReactDOM.render(<LelApp />, document.getElementById("app"));
  });
}
