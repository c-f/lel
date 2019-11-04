import React, { Component } from "react";
import {
  message,
  Layout,
  Menu,
  Breadcrumb,
  Icon,
  Input,
  Row,
  Col,
  Button,
  Drawer,
  Table
} from "antd";
const { Search } = Input;

// dataSource
// columns

class SearchPanel extends Component {
  state = {};
  // TODO maybe states over constructor

  render() {
    return (
      <Drawer
        title={this.props.title}
        placement={this.props.placement}
        height="100%"
        width={this.props.width}
        closable={true}
        onClose={this.props.onClose}
        visible={this.props.visible}
      >
        <Search
          placeholder="currently without function"
          enterButton="Search"
          onSearch={this.props.onSearch}
          // onSearch={value => this.search(value)}
        />
        <Table
          style={{ maxHeight: "100%" }}
          dataSource={this.props.dataSource}
          columns={this.props.columns}
          pagination={{ pageSize: 50 }}
          scroll={{ y: "max-content" }}
        />
        ;
      </Drawer>
    );
  }
}

export default SearchPanel;
