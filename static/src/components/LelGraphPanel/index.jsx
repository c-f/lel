import React, { Component } from "react";
import { Drawer } from "antd";
import Graph from "./graph.jsx";

class LelGraphPanel extends Component {
  render() {
    return (
      <Drawer
        title="Graph"
        placement="right"
        width="100%"
        closable={true}
        onClose={this.props.onClose}
        height="100vh"
        visible={this.props.visible}
      >
        <Graph
          onSelect={this.props.onSelect}
          onOpen={this.props.onOpen}
          onShow={this.props.onShow}
          graphs={this.props.graphs}
        />
      </Drawer>
    );
  }
}

export default LelGraphPanel;
