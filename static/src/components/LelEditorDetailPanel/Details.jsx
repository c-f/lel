import React, { Fragment } from "react";
import { Card, Form, Input, Select } from "antd";
import { withPropsAPI } from "gg-editor";
import upperFirst from "lodash/upperFirst";

const { Item } = Form;
const { Option } = Select;

const inlineFormItemLayout = {
  labelCol: {
    sm: { span: 8 }
  },
  wrapperCol: {
    sm: { span: 16 }
  }
};

class Details extends React.Component {
  get item() {
    const { propsAPI } = this.props;

    return propsAPI.getSelected()[0];
  }

  handleSubmit = e => {
    if (e && e.preventDefault) {
      e.preventDefault();
    }

    const { form, propsAPI } = this.props;
    const { getSelected, executeCommand, update } = propsAPI;

    setTimeout(() => {
      form.validateFieldsAndScroll((err, values) => {
        if (err) {
          return;
        }

        const item = getSelected()[0];

        if (!item) {
          return;
        }

        executeCommand(() => {
          update(item, {
            ...values
          });
        });
      });
    }, 0);
  };

  renderEdgeShapeSelect = () => {
    return (
      <Select onChange={this.handleSubmit}>
        <Option value="flow-smooth">Smooth</Option>
        <Option value="flow-polyline">Polyline</Option>
        <Option value="flow-polyline-round">Polyline Round</Option>
      </Select>
    );
  };

  openDocument = (e, id) => {
    e.preventDefault();
    var url = `${config.open}?path=${id}`;
    console.log("start open");
    fetch(url);
  };

  renderNodeDetail = () => {
    const { form } = this.props;
    const { id, references, tags, label } = this.item.getModel();

    // todo open should not be a link :)
    console.log(this.item);
    return (
      <div>
        <h4>{id}</h4>
        <ul>
          <li>
            <a
              onClick={e => {
                this.props.onOpen(e, id);
              }}
            >
              Open Document
            </a>
          </li>
          <li>
            <a
              target="_blank"
              onClick={e => {
                console.log("SSS", id);
                this.props.onShow(id, label);
              }}
            >
              Show Document
            </a>
          </li>
        </ul>
        References:
        <ul>
          {references !== undefined &&
            references.map((value, index) => {
              return (
                <li key={index}>
                  <a onClick={e => this.props.onOpen(e, value)}>{value}</a>
                </li>
              );
            })}
        </ul>
        Tags:
        <ul>
          {tags.map((value, index) => {
            return <li key={index}>{value}</li>;
          })}
        </ul>
      </div>
    );
  };

  renderEdgeDetail = () => {
    const { form } = this.props;
    const { label = "", shape = "flow-smooth" } = this.item.getModel();

    return (
      <Fragment>
        <Item label="Label" {...inlineFormItemLayout}>
          {form.getFieldDecorator("label", {
            initialValue: label
          })(<Input onBlur={this.handleSubmit} />)}
        </Item>
        <Item label="Shape" {...inlineFormItemLayout}>
          {form.getFieldDecorator("shape", {
            initialValue: shape
          })(this.renderEdgeShapeSelect())}
        </Item>
      </Fragment>
    );
  };

  renderGroupDetail = () => {
    const { form } = this.props;
    const { label = "新建分组" } = this.item.getModel();

    return (
      <Item label="Label" {...inlineFormItemLayout}>
        {form.getFieldDecorator("label", {
          initialValue: label
        })(<Input onBlur={this.handleSubmit} />)}
      </Item>
    );
  };

  render() {
    const { type } = this.props;

    if (!this.item) {
      return null;
    }

    return (
      <Card type="inner" size="small" title={upperFirst(type)} bordered={false}>
        <Form onSubmit={this.handleSubmit}>
          {type === "node" && this.renderNodeDetail()}
          {type === "edge" && this.renderEdgeDetail()}
          {type === "group" && this.renderGroupDetail()}
        </Form>
      </Card>
    );
  }
}

export default Form.create()(withPropsAPI(Details));
