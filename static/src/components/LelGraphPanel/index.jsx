import React, { Component } from "react";
import { Drawer, Input, Icon, Tooltip } from "antd";
import Graph from "./graph.jsx";
const { Search } = Input;

class LelGraphPanel extends Component {
  state = {
    search: ""
  };
  render() {
    // maybe arrow function on onSelect
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
        <Search
          addonBefore={
            <Tooltip title="Currently searching only one tag is possible">
              <Icon type="info-circle" theme="twoTone"></Icon>
            </Tooltip>
          }
          onSearch={this.HandleSearch}
        />
        <Graph
          url={this.props.url}
          onSelect={this.props.onSelect}
          onOpen={this.props.onOpen}
          onShow={this.props.onShow}
          filter={this.state.search}
        />
      </Drawer>
    );
  }
  HandleSearch = search => {
    // todo
    console.log("Searching for ", search);
    this.setState({
      search: search
    });
  };
}
// this.props.OnSelect

// ant-drawer-content-wrapper
// styles
// eight: auto;
export default LelGraphPanel;
