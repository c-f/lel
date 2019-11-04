import React, { Component } from "react";
import { message, Table, Row } from "antd";

const toDoColumns = [
  {
    title: "done",
    dataIndex: "done"
  },
  {
    title: "priority",
    dataIndex: "priority"
  },
  {
    title: "completion_date",
    dataIndex: "completion_date"
  },
  {
    title: "creation_date",
    dataIndex: "creation_date"
  },
  {
    title: "description",
    dataIndex: "description"
  },
  {
    title: "context",
    dataIndex: "context"
  },
  {
    title: "projects",
    dataIndex: "projects"
  },
  {
    title: "repeat",
    dataIndex: "repeat"
  },
  {
    title: "action",
    dataIndex: "action"
  },
  {
    title: "due",
    dataIndex: "due"
  },
  {
    title: "_raw",
    dataIndex: "_raw"
  }
];

class MetaPanel extends Component {
  render() {
    let data = this.props.data;
    let todos = data
      .map((item, key) => {
        return item.todos;
      })
      .flat();
    console.log(data);
    return <Row>{this.renderTodos(todos)}</Row>;
  }
  renderTodos = todos => {
    //todos = todos.map
    return <Table columns={toDoColumns} dataSource={todos} size="small" />;
  };
}

export default MetaPanel;
