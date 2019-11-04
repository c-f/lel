import React from "react";
import { RegisterNode } from "gg-editor";

import svg_database_cube from "./shapes/database_cube.svg";
import svg_generic_node from "./shapes/generic_node.svg";
import svg_internet_cloud from "./shapes/internet_cloud.svg";
import svg_sql_server from "./shapes/sql_server.svg";
import svg_sub_site from "./shapes/sub_site.svg";
import svg_user from "./shapes/user.svg";
import svg_vista_terminal from "./shapes/vista_terminal.svg";
import svg_windows_server_2 from "./shapes/windows_server_2.svg";
import svg_windows_server from "./shapes/windows_server.svg";

// generateConfig returns a valid config and renders the image if no custom icon is found
let generateConfig = shapeConfig => {
  return {
    draw(item) {
      //item.visible = false;
      //const keyShape = this.drawKeyShape(item);
      //console.log(this.drawKeyShape);
      // draw label
      item.model.labelOffsetY = shapeConfig.labelOffsetY;
      this.drawLabel(item);

      // draw image
      const group = item.getGraphicGroup();
      const model = item.getModel();

      let attrs = {
        x: shapeConfig.x || 0,
        y: shapeConfig.y || 0,
        img: model.icon || shapeConfig.icon,
        height: model.height || shapeConfig.height,
        width: model.width || shapeConfig.width
      };

      let lol = group.addShape("path", {
        attrs: {
          path: "M 0,0",
          stroke: "blue",
          lineWidth: 0
        }
      });

      /*
      if (model.height || shapeConfig.height) {
        attrs.height = model.height || shapeConfig.height;
      }
      if (model.width || shapeConfig.width) {
        attrs.width = model.width || shapeConfig.width;
      }
      */
      // test = group.addShape("path", {});
      group.addShape("image", {
        attrs: attrs
      });

      //console.log("item", item);
      // console.log("keyshape", keyShape);
      //console.log("group", group);

      return lol;
    }
  };
};

// default supported shapes
// credits for drawio - you guys rock !
// https://github.com/jgraph/drawio
const shapeConfig = [
  //{ name: "database", icon: svg_database_cube },
  {
    name: "database",
    icon: "/static/icons/database_cube.svg",
    height: 60,
    width: 54,
    labelOffsetY: 45,
    x: -25,
    y: -25
  },
  {
    name: "node",
    icon: "/static/icons/generic_node.svg",
    width: 54,
    height: 60,
    labelOffsetY: 10,
    x: -25,
    y: -60
  },
  {
    name: "globe",
    icon: "/static/icons/internet_globe.svg",
    width: 50,
    height: 50,
    labelOffsetY: 40,
    x: -25,
    y: -25
  },
  {
    name: "cserver",
    icon: "/static/icons/sql_server.svg",
    width: 53,
    height: 70,
    labelOffsetY: 10,
    x: -25,
    y: -70
  },
  {
    name: "site",
    icon: "/static/icons/sub_site.svg",
    width: 50,
    height: 43,
    labelOffsetY: 32,
    x: -25,
    y: -25
  },
  {
    name: "user",
    icon: "/static/icons/user.svg",
    width: 18.5,
    height: 50,
    labelOffsetY: 35,
    x: -10,
    y: -25
  },
  {
    name: "client",
    icon: "/static/icons/vista_terminal.svg",
    width: 38,
    height: 50,
    labelOffsetY: 40,
    x: -15,
    y: -25
  },
  {
    name: "dc",
    icon: "/static/icons/windows_server_2.svg",
    width: 53,
    height: 70,
    labelOffsetY: 10,
    x: -25,
    y: -70
  },
  {
    name: "win",
    icon: "/static/icons/windows_server_2.svg",
    width: 53,
    height: 70,
    labelOffsetY: 10,
    x: -25,
    y: -70
  },
  {
    name: "server",
    icon: "/static/icons/windows_server.svg",
    width: 53,
    height: 70,
    labelOffsetY: 10,
    x: -25,
    y: -70
  }
];

// RegisterLelNode registers the supported nodes
class RegisterLelNodes extends React.Component {
  renderNodes() {
    return shapeConfig.map(shape => {
      const config = generateConfig(shape);
      return (
        <RegisterNode key={shape.name} name={shape.name} config={config} />
      );
    });
  }
  renderTest() {
    const config = generateConfig("/static/icons/svg_generic_node.svg");
    return (
      <React.Fragment>
        <RegisterNode name="database" config={config} />
      </React.Fragment>
    );
  }

  render() {
    return this.renderNodes();
  }
}

export default RegisterLelNodes;
