import React, { Component } from "react";
import TreeView from "./tree.jsx";
import { Drawer, Alert } from "antd";

class LelTagsPanel extends Component {
  render() {
    // maybe arrow function on onSelect

    return (
      <Drawer
        title="Tags"
        placement="left"
        mask={false}
        closable={true}
        onClose={this.props.onClose}
        height="bottom"
        visible={this.props.visible}
        //style={{ paddingLeft: "200px" }}
      >
        <TreeView data={this.props.data} onSelect={this.props.onSelect} />
      </Drawer>
    );
  }
}

class LelNavPanel extends Component {
  render() {
    // maybe arrow function on onSelect
    return (
      <Drawer
        title="Navigation"
        placement="top"
        closable={true}
        onClose={this.props.onClose}
        height="bottom"
        visible={this.props.visible}
      >
        <Alert
          message={"Names can also be search via 'tags:' Command"}
          type="info"
          showIcon
        />
        <TreeView data={this.props.data} onSelect={this.props.onSelect} />
      </Drawer>
    );
  }
}

class LelStickyNavPanel extends Component {
  render() {
    return <TreeView data={this.props.data} onSelect={this.props.onSelect} />;
  }
}

// ant-drawer-content-wrapper
// styles
// eight: auto;
export { LelTagsPanel, LelNavPanel, LelStickyNavPanel };
