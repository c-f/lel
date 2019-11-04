import React, { Component } from "react";
import { Alert, Input, Icon } from "antd";

import ImageGallery from "react-image-gallery";
import "react-image-gallery/styles/css/image-gallery.css";

class ImageSlider extends Component {
  state = {
    current: 0,
    link: ""
  };

  render() {
    const link = `![picture](${this.state.link})`;
    return (
      <div>
        <ImageGallery
          thumbnailPosition="left"
          // lazyLoad={true}
          onSlide={e => {
            this.setState({ current: e, link: this.props.images[e].item });
          }}
          onImageLoad={e => {
            console.log("image", e);
          }}
          items={this.props.images}
          sizes={{
            width: 100,
            height: 100
          }}
        />
        <Alert
          message="Copy &amp; paste"
          description={
            <span>
              Image Link:
              <Input
                style={{ width: 300 }}
                suffix={<Icon type="copy"></Icon>}
                value={link}
              />
            </span>
          }
          type="info"
          showIcon
        />
      </div>
    );
  }
}

export default ImageSlider;
