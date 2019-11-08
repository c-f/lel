import React, { Component } from "react";
import { Tag, Descriptions, Badge, Alert, Icon, List } from "antd";
import { API } from "../api";
class Homepanel extends Component {
  state = {};

  componentDidMount() {
    this.getStatsData();
  }

  getStatsData = () => {
    new API().stats().then(data => {
      console.log(data);
      this.setState({ stats: data });
    });
  };

  render() {
    let infos = (
      <h1>
        LeL Project{" "}
        <Tag
          color="blue"
          onClick={e => {
            this.getStatsData();
          }}
        >
          {VERSION}
        </Tag>{" "}
      </h1>
    );
    let description = <Icon type="loading" />;
    if (this.state.stats !== undefined) {
      description = (
        <div style={{ marginTop: 40 }}>
          <Descriptions>
            <Descriptions.Item label="Notes">
              <Badge
                count={this.state.stats.notes}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="Images">
              <Badge
                count={this.state.stats.images}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="Videos">
              <Badge
                count={this.state.stats.videos}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="Misato">
              <Badge
                count={this.state.stats.misato}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="Milestones">
              <Badge
                count={this.state.stats.milestones}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
            <Descriptions.Item label="Graphs">
              <Badge
                count={this.state.stats.graphs}
                style={{ backgroundColor: "#52c41a" }}
              />
            </Descriptions.Item>
          </Descriptions>

          {this.state.stats.errors && (
            <List
              style={{ marginTop: 20 }}
              header={<h4>Errors</h4>}
              //footer={<div>Footer</div>}
              //bordered
              dataSource={this.state.stats.errors}
              renderItem={item => (
                <List.Item>
                  <Alert
                    message={item}
                    type="error"
                    showIcon
                    style={{ marginBottom: 10 }}
                  />
                </List.Item>
              )}
            />
          )}
          <div></div>
        </div>
      );
    }
    return (
      <div
        style={{
          backgroundColor: "white",
          padding: "30px",
          borderRadius: "3px"
        }}
      >
        {infos}
        <div style={{ textAlign: "center" }}>
          <img
            id="Vektor-Smartobjekt"
            height="150"
            src="data:img/png;base64,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"
          />
          <p>Explode your Notes!</p>
          {description}
        </div>
      </div>
    );
  }
}

export default Homepanel;
