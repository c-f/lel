import React, { Component } from "react";
import { Card } from "antd";
import {
  NodePanel,
  EdgePanel,
  GroupPanel,
  MultiPanel,
  CanvasPanel,
  DetailPanel
} from "gg-editor";
import Details from "./Details.jsx";
import styles from "./index.less";

class LelDetailPanel extends Component {
  render() {
    return (
      <DetailPanel className={styles.detailPanel}>
        <NodePanel>
          <Details
            type="node"
            onOpen={this.props.onOpen}
            onShow={this.props.onShow}
          />
        </NodePanel>
        <EdgePanel>
          <Details type="edge" onOpen={this.props.onOpen} />
        </EdgePanel>
        <GroupPanel>
          <Details type="group" onOpen={this.props.onOpen} />
        </GroupPanel>
        <MultiPanel>
          <Card
            type="inner"
            size="small"
            title="Multi Select"
            bordered={false}
          />
        </MultiPanel>
        <CanvasPanel>
          <Card type="inner" size="small" title="Canvas" bordered={false} />
        </CanvasPanel>
      </DetailPanel>
    );
  }
}

export default LelDetailPanel;
