import React, { Component } from "react";
import { Upload, Modal, Input, Form, Icon } from "antd";
import { API } from "../api.js";
const { Dragger } = Upload;

class ImageUploader extends Component {
  state = {
    title: "",
    picture: null,
    visible: false
  };

  // todo information for normal data
  HandlePaste = e => {
    console.log(e.clipboardData);
    var items = (e.clipboardData || e.originalEvent.clipboardData).items;
    console.log(JSON.stringify(items)); // will give you the mime types
    for (const index in items) {
      var item = items[index];

      if (item.kind === undefined) {
        continue;
      }
      console.log("You are trying to upload", item.kind, item.type);
      if (
        item.kind === "file" &&
        ["image/png", "image/jpeg", "image/jpg"].indexOf(item.type) !== -1
      ) {
        console.log("FILLLE", item);
        var blob = item.getAsFile();
        var reader = new FileReader();
        reader.onload = e => {
          this.setState({
            picture: e.target.result,
            visible: true
          });
        };
        reader.readAsDataURL(blob);
      } else {
        console.log("Sorry not supporting", item.type, item.kind);
      }
    }
  };

  empty() {}
  // displayModal, asks the use for the image filename and
  // uploads the data if not empty
  renderModal() {
    return (
      <Modal
        title="Image Name (without .png)"
        visible={this.state.visible}
        onOk={this.HandleUpload}
        //onCancel={this.empty}
        cancelButtonProps={{ disabled: true }}
      >
        <Input
          prefix={<Icon type="safety" style={{ color: "rgba(0,0,0,.25)" }} />}
          placeholder="important_image"
          value={this.state.title}
          onChange={e => {
            this.setState({
              title: e.target.value
            });
          }}
        />
      </Modal>
    );
  }

  HandleUpload = e => {
    this.upload({
      picture: this.state.picture,
      title: this.state.title
    });
  };

  // upload uploads the image to the server
  upload(data) {
    new API()
      .uploadImg(data)
      .then(data => {
        if (this.props.onUpload) {
          this.props.onUpload({
            message: "Upload Successful",
            data: data,
            success: true
          });
        }
      })
      .catch(e => {
        console.log("error", e);
        if (this.props.onError) {
          this.props.onError({
            message: "Upload not Successful",
            data: JSON.stringify(e),
            success: false
          });
        }
      })
      .finally(e => {
        this.setState({
          visible: false
        });
      });
  }

  render() {
    return (
      <div
        style={{
          border: "1px dashed #07a",
          borderRadius: "4px"
        }}
        onPaste={e => {
          console.log("!");
          this.HandlePaste(e);
        }}
      >
        Upload Image
        {this.renderModal()}
      </div>
    );
  }
}

//
export default Form.create({ name: "normal_login" })(ImageUploader);
/*

        {this.renderModal()}{" "}

*/
