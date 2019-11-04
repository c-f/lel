import React, { Component } from "react";
import { Layout, Menu, Breadcrumb, Icon, Input, Tree } from "antd";
const { TreeNode } = Tree;
const { Search } = Input;

import { loop } from "./loop.jsx";

class TreeView extends Component {
  state = {
    expandedKeys: [],
    searchValue: "",
    autoExpandParent: true
  };

  isReady() {
    return this.props.data.length > 0;
  }
  onExpand = expandedKeys => {
    this.setState({
      autoExpandParent: true
    });
  };
  // sucht
  onChange = e => {
    const { value } = e.target;
    console.log("[tree]", "[search]", value);

    this.setState({
      searchValue: value,
      autoExpandParent: true,
      defaultExpandAll: true
    });
  };
  render() {
    const { searchValue, expandedKeys, autoExpandParent } = this.state;
    const { onSelect, data } = this.props;

    return (
      <div>
        <Search
          style={{
            margin: 20,
            position: "flex",
            marginRight: 8,
            zIndex: "1",
            maxWidth: 300
          }}
          placeholder="Search"
          onChange={this.onChange}
        />
        {this.isReady() && (
          <div>
            <Tree
              onSelect={onSelect}
              defaultExpandAll={true}
              autoExpandParent={true}
              onRightClick={e => {
                console.log(e);
              }}
              showLine={true}
              switcherIcon={<Icon type="down" />}
              className="hide-file-icon"
              style={{
                overflowY: "auto",
                height: "90vh",
                padding: "0px 20px",
                zIndex: -100
              }}
            >
              {loop(data, searchValue)}
            </Tree>
          </div>
        )}
      </div>
    );
  }
}

export default TreeView;

// todo
/*

class SearchTree extends React.Component {
 
  // updated state when probs change
  componentWillReceiveProps(nextProps) {
    if (this.props != nextProps) {
      console.log("updated", nextProps.data);
      this.setState({
        data: nextProps.data
      });
    }
  }

  // sucht
  onChange = e => {
    const { value } = e.target;
    console.log(value);
    const expandedKeys = Object.keys(this.state.data)
      .map(i => {
        const item = this.state.data[i];

        if (item.title.indexOf(value) > -1) {
          console.log("match", item.title);
          return this.getParentKey(item.key, this.state.data);
        }
        return null;
      })
      .filter((item, j, self) => item && self.indexOf(item) === j);
    console.log("expanded", expandedKeys);
    this.setState({
      expandedKeys,
      searchValue: value,
      autoExpandParent: true
    });
  };

  getParentKey(key, tree) {
    let parentKey;
    for (let i = 0; i < tree.length; i++) {
      const node = tree[i];
      if (node.children) {
        if (node.children.some(item => item.key === key)) {
          parentKey = node.key;
        } else if (this.getParentKey(key, node.children)) {
          parentKey = this.getParentKey(key, node.children);
        }
      }
    }
    return parentKey;
  }

  // todo
  // bind
  onSelect(a) {
    console.log("selected", a);
  }

  render() {
    const { searchValue, expandedKeys, autoExpandParent } = this.state;

    return (
      <div>
        <Search
          style={{ marginBottom: 8 }}
          placeholder="Search"
          onChange={this.onChange}
        />
        {this.isReady() && (
          <Tree
            onExpand={this.onExpand}
            expandedKeys={expandedKeys}
            autoExpandParent={autoExpandParent}
            showLine={true}
            showIcon={false}
            className="hide-file-icon"
            onSelect={a => this.onSelect(a)}
          >
            {loop(this.state.data)}
          </Tree>
        )}
      </div>
    );
  }
}
*/
