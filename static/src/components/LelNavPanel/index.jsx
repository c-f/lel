import React, { Component } from "react";
import TreeView from "./tree.jsx";
import { Drawer, Alert, Row, Col } from "antd";

class LelTagsPanel extends Component {
  state = {
    matches: [],
    infoEnable: false
  };

  render() {
    // maybe arrow function on onSelect
    let sty = {};
    if (this.state.infoEnable) {
      sty.boxShadow =
        "0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19)";
    }

    return (
      <Drawer
        title="Tags"
        placement="left"
        mask={false}
        closable={true}
        onClose={this.props.onClose}
        height="bottom"
        visible={this.props.visible}
        width={this.state.infoEnable ? 800 : 400}
      >
        <Row>
          <Col span={12} style={sty}>
            <TreeView
              data={this.props.data}
              onSelect={e => {
                this.setState({ matches: e, infoEnable: true });
              }}
            />
          </Col>
          {this.state.infoEnable && (
            <Col span={12}>
              <ul>
                {this.state.matches.map((match, index) => {
                  match = match.split("__")[1];
                  return (
                    <li key={index}>
                      <a
                        href="#"
                        onClick={e => {
                          this.props.onSelect(match);
                        }}
                      >
                        {match}
                      </a>
                    </li>
                  );
                })}
              </ul>
            </Col>
          )}
        </Row>
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
