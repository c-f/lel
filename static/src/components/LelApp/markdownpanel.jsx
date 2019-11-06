import React, { Component } from "react";
import { UnControlled as CodeMirror } from "react-codemirror2";
import {
  Row,
  Col,
  Icon,
  Button,
  message,
  Anchor,
  PageHeader,
  Descriptions,
  Avatar,
  Tooltip,
  Tag
} from "antd";
import Moment from "react-moment";
import { API } from "../api";

require("codemirror/mode/markdown/markdown");

var config = require("Config");
const ReactMarkdown = require("react-markdown");

class MarkdownPanel extends Component {
  ref = editor => {
    this.editor = editor;
  };

  constructor(props) {
    super(props);
    this.state = {
      content: "",
      isLoading: true,
      contentPath: "",
      meta: {},
      changed: 2
    };
    CodeMirror;
  }

  componentWillReceiveProps(nextProps) {
    if (this.state.content != nextProps.content) {
      console.log("[markdown]", "[content]", nextProps.content);
      this.setState({
        content: nextProps.content,
        changed: 2
      });
    }
    if (
      this.state.contentPath != nextProps.contentPath &&
      nextProps.contentPath != undefined
    ) {
      console.log("[markdown]", "[contentPath]", nextProps.contentPath);
      this.setState({
        contentPath: nextProps.contentPath
      });
    }
    if (this.state.meta != nextProps.meta) {
      console.log("[markdown]", "[meta]", nextProps.contentPath);
      this.setState({
        meta: nextProps.meta
      });
    }
  }

  handleMarkdownChanges = (e, data, value) => {
    console.log("handled");
    let changed = this.state.changed;
    this.setState({
      content: value,
      changed: changed < 0 ? changed : changed - 1
    });
  };

  renderMarkdown() {
    return (
      <div
        style={{
          background: "white",
          padding: "30px",
          borderRadius: "4px"
        }}
      >
        <ReactMarkdown
          source={this.state.content}
          transformImageUri={this.props.onImageLoad}
          renderers={{ image: this.renderImages }}
        />
      </div>
    );
  }

  renderImages = props => {
    return <img {...props} style={{ maxWidth: "100%", maxHeight: "500px" }} />;
  };

  onDelete = e => {
    e.preventDefault();
    new API().remove(this.props.contentPath).then(res => {
      message.success("File remove: " + this.props.contentPath, 2.5);
    });
  };

  // do something
  onSave = e => {
    e.preventDefault();

    var fd = new FormData();
    fd.append("content", this.state.content);
    new API()
      .uploadMarkdown(fd, this.props.contentPath)
      .then(res => {
        if (res.ok) {
          message.success("Success saved: " + this.props.contentPath, 2.5);
          this.setState({ changed: 1 });
        } else {
          message.error("Could not save: " + this.props.contentPath, 2.5);
        }
      })
      .catch(error => {
        message.error(
          "Could not save: catch" + this.props.contentPath + error,
          2.5
        );
      });
  };

  /*
  <Button
            disabled={this.state.changed <= 0 ? false : true}
            type="success"
            icon="save"
            shape="round"
            onClick={this.onSave}
          >
            Save{this.state.changed <= 0 ? "*" : ""}
          </Button>
  
  */
  renderMeta() {
    if (this.state.meta.tags !== undefined) {
      let { tags, names, entity } = this.state.meta;
      let tagtags = tags.map((item, key) => {
        return (
          <Tag color="cyan" key={key}>
            {item}
          </Tag>
        );
      });
      let tagnames = names.map((item, key) => {
        return (
          <Tag color="lime" key={key}>
            {item}
          </Tag>
        );
      });
      let tagentity = (
        <Tag color="geekblue" key="entity">
          {entity}
        </Tag>
      );
      return (
        <Row type="flex" justify="end">
          <Col span={24} style={{ textAlign: "right" }}>
            {tagnames}
            {tagentity}
            {tagtags}
          </Col>
        </Row>
      );
    } else {
      console.log("Spinner");
    }
  }

  render() {
    let editbtn = <span></span>;

    if (this.state.contentPath !== "") {
      editbtn = (
        <Button
          icon="edit"
          onClick={this.props.HandleOpenDocument}
          size="small"
        />
      );
    }
    let parts = this.state.contentPath.split(/(\\|\/)/g);
    let filename = parts.pop();
    let folder = parts.join("");

    return (
      <div>
        <Row>
          <Anchor style={{ margin: "10px auto", maxWidth: "1200px" }}>
            <PageHeader
              style={{
                border: "2px solid #eee"
              }}
              onBack={e => {
                window.history.back();
              }}
              title={filename}
              subTitle={folder}
              tags={editbtn}
              extra={[
                <Icon
                  key="help"
                  type="question-circle"
                  onClick={this.props.onHelpClick}
                ></Icon>,
                <Button
                  key="remove"
                  icon="delete"
                  onClick={this.onDelete}
                ></Button>,
                <Button key="change" icon="save" onClick={this.onSave}>
                  {this.state.changed <= 0 ? "*" : ""}
                </Button>
              ]}
            >
              <Descriptions>
                <Descriptions.Item label="Info">
                  <Avatar
                    style={{ backgroundColor: "teal" }}
                    size="small"
                    icon="user"
                  />{" "}
                  <Tooltip title={this.state.meta.modified}>
                    <Moment fromNow>{this.state.meta.modified}</Moment>
                  </Tooltip>
                </Descriptions.Item>
                <Descriptions.Item label="Tags">
                  {this.renderMeta()}
                </Descriptions.Item>
                <Descriptions.Item label="Operations"></Descriptions.Item>
              </Descriptions>
            </PageHeader>
          </Anchor>
        </Row>
        <Row justify="center">
          <Col span={12}>
            <form className="editor pure-form">
              <CodeMirror
                autoCursor={false}
                value={this.state.content}
                defaultValue={this.state.content}
                onChange={this.handleMarkdownChanges}
                editorDidMount={editor => {
                  editor.refresh();
                }}
                options={{
                  mode: "markdown",
                  theme: "rubyblue",
                  lineNumbers: true,
                  spellcheck: true
                  //lineWrapping: false,

                  // keyMap: "sublime"
                }}
              ></CodeMirror>
            </form>
          </Col>
          <Col span={12}>
            <div className="lel-md-panel">{this.renderMarkdown()}</div>
          </Col>
        </Row>
      </div>
    );
  }
}

export default MarkdownPanel;
