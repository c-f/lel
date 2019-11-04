import React, { Component } from "react";
import { Drawer, Input, Icon, Form, Button, message } from "antd";

class LoginPanelForm extends Component {
  handleSubmit = e => {
    e.preventDefault();

    const { getFieldsValue } = this.props.form;

    this.props.form.validateFields((err, values) => {
      if (err) {
        return;
      }
      this.props.handleAuth(values["username"], values["token"]);
    });
  };

  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Drawer
        title="Login"
        placement="right"
        closable={false}
        visible={!this.props.isLoggedIn}
        maskStyle={{
          backgroundColor: "white",
          backgroundImage: `url("/static/icons/login.r.svg")`,
          opacity: 1
        }}
      >
        <Form onSubmit={this.handleSubmit} className="login-form">
          <Form.Item>
            {getFieldDecorator("username", {
              rules: [
                { required: true, message: "Please input your operatorname" }
              ]
            })(
              <Input
                placeholder="Username"
                name="user"
                prefix={
                  <Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />
                }
              />
            )}
          </Form.Item>
          <Form.Item>
            {getFieldDecorator("token", {
              rules: [
                { required: true, message: "Please input your LEL Token!" }
              ]
            })(
              <Input
                placeholder="Token"
                name="token"
                prefix={
                  <Icon type="api" style={{ color: "rgba(0,0,0,.25)" }} />
                }
              />
            )}
          </Form.Item>
          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              className="login-form-button"
            >
              Log in
            </Button>
          </Form.Item>
        </Form>
      </Drawer>
    );
  }
}

const LoginPanel = Form.create({ name: "normal_login" })(LoginPanelForm);
export default LoginPanel;
