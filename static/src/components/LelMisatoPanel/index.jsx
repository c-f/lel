import React, { Component } from "react";

import { Table, Input, Icon, Tag, Tooltip } from "antd";
const { Search } = Input;

import Moment from "react-moment";

class MisatoPanel extends React.Component {
  onChange = e => {
    const { value } = e.target;
    this.props.onSearch(value);
  };

  render() {
    return (
      <div
        style={{
          background: "white",
          padding: 10,
          borderRadius: 3,
          width: "100%"
        }}
      >
        <Search onSearch={this.props.onSearch} />
        <Table
          style={{ maxWidth: "100%" }}
          size="small"
          expandedRowRender={record => (
            <p style={{ margin: 0 }}>{record.match}</p>
          )}
          columns={[
            {
              title: "Path",
              dataIndex: "path",
              key: "path",
              render: text => {
                return (
                  <span>
                    {text}
                    <Icon
                      style={{ marginLeft: 5 }}
                      type="edit"
                      onClick={e => console.log("Clicked", e)}
                    />
                  </span>
                );
              },
              width: 150
            },
            {
              title: "Time",
              dataIndex: "ctx.time",
              key: "ctx.time",
              render: opt => {
                return (
                  <Tooltip title={<Moment unix>{opt}</Moment>}>
                    <Moment fromNow unix>
                      {opt}
                    </Moment>
                  </Tooltip>
                );
              },
              width: 150
              // fixed: "left"
            },
            {
              title: "Command",
              dataIndex: "command",
              key: "command"
              // fixed: "left"
            },
            {
              title: "User",
              dataIndex: "ctx.user",
              key: "ctx.user",
              render: opt => {
                return <Tag color="gold">{opt}</Tag>;
              }
            }
          ]}
          dataSource={this.props.dataSource}
          scroll={{ x: 1500, y: "100%" }}
          rowKey={record => record.key}
        />
      </div>
    );
  }
}

export default MisatoPanel;

/*


            

*/
