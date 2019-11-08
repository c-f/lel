import React from "react";
import EditorMinimap from "../EditorMinimap";
import { KoniContextMenu } from "../EditorContextMenu";
import { KoniToolbar } from "../EditorToolbar";
import { KoniItemPanel } from "../EditorItemPanel";
import LelDetailPanel from "../LelEditorDetailPanel/index.jsx";
import { HandleLayout, FilterNodes } from "./layout.js";
import RegisterLelNodes from "./node.jsx";
import GGEditor, { Koni, withPropsAPI, Item, ItemPanel } from "gg-editor";

import {
  Row,
  Col,
  message,
  Select,
  Input,
  Tooltip,
  Icon,
  Button,
  Tag,
  Modal
} from "antd";
const InputGroup = Input.Group;
const { Search } = Input;

import styles from "./graph.less";
import { API } from "../api.js";

GGEditor.setTrackable(false);

class Graph extends React.Component {
  constructor(probs) {
    super(probs);

    this.state = {
      data: { nodes: [], edges: [] },
      raw: { nodes: [], edges: [] },
      filter: "",
      previousFilter: "",
      ok: false,
      name: "",
      newName: "",

      graphs: []
    };
  }

  getGraphsData = () => {
    new API().graphs().then(data => {
      this.setState({ graphs: data });
    });
  };

  componentDidMount() {
    this.refreshData();
    this.getGraphsData();

    this.setState({
      ok: true
    });
  }

  nodeTrack(e) {
    console.log(e);
    //console.log(this.editor.propsAPI.save())
    //var old = this.HandleLayout(data);
    //this.editor.propsAPI.read(old);
  }

  shouldComponentUpdate(nextProps, nextState) {
    if (this.state.newName !== nextState.newName) {
      return false;
    }
    // if (this.state.in === nextState.in) {
    //   if (this.state.inn === nextState.inn) {
    //     return false;
    //   }
    // }
    return true;
  }

  // LoadData by
  LoadData = name => {
    new API().graphByName(name).then(res => {
      // data, filter
      console.log("RES", res);
      if (res.data != undefined && res.filter != undefined) {
        this.setState({
          data: res.data,
          filter: res.filter,
          previousFilter: res.filter
        });
      }
    });
  };

  HandleData = () => {
    let filtered = FilterNodes(this.state.raw, this.state.filter);
    console.log("filtered", filtered);

    const prettyData = HandleLayout(filtered);

    this.setState({
      data: prettyData
    });
  };

  refreshData = e => {
    if (e != null) {
      e.preventDefault();
    }
    new API().graph().then(data => {
      this.setState(
        {
          raw: data
        },
        this.HandleData
      );
    });
  };

  saveData = name => {
    let data = {};

    if (!name.endsWith("graph.json")) {
      name = `${name}.graph.json`;
    }

    new API()
      .uploadGraph(name, {
        data: this.state.data,
        filter: this.state.filter
      })
      .then(res => {
        message.success("Saved Graph", 2.5);
      });
  };

  filter = () => {
    // console.log("lel")

    filter;
  };

  HandleSelectGraph = name => {
    this.setState({ name: name });
    this.LoadData(name);
  };

  HandleSearch = filter => {
    // todo
    this.setState({ filter: filter }, this.HandleData);
  };

  HandleCompass = e => {
    e.preventDefault();
    this.setState({ filter: "", name: "" }, this.refreshData);
  };

  showModal = () => {
    Modal.info({
      title: "New Graph",
      content: (
        <div>
          <Input
            onChange={e => {
              this.setState({ newName: e.target.value });
            }}
          ></Input>
        </div>
      ),
      onOk: () => {
        console.log(this.state.newName);
        this.saveData(this.state.newName);
        this.getGraphsData();
      },
      maskClosable: true
    });
  };

  render() {
    if (this.state.ok) {
      var read = this.editor.propsAPI.read;
      read(this.state.data);
    }

    return (
      <GGEditor className={styles.editor} ref={e => (this.editor = e)}>
        <Row type="flex" className={styles.editorHd}>
          <Col span={24}>
            <InputGroup>
              <Row>
                <Col span={1}>
                  <Button icon="compass" onClick={this.HandleCompass}></Button>
                </Col>
                <Col span={1}>
                  <Button
                    icon="cloud-sync"
                    onClick={e => message.error("Currently not implemented", 1)}
                  ></Button>
                </Col>
                <Col span={1}>
                  <Button
                    icon="plus-circle"
                    onClick={e => {
                      this.showModal();
                    }}
                  ></Button>
                </Col>
                <Col span={1}>
                  <Button
                    icon="save"
                    onClick={e => this.saveData(this.state.name)}
                  ></Button>
                </Col>
                <Col span={8}>
                  <Select
                    defaultValue=""
                    style={{ width: "100%" }}
                    onSelect={this.HandleSelectGraph}
                  >
                    {this.state.graphs.map(e => {
                      return (
                        <Select.Option value={e} key={e}>
                          {e}
                        </Select.Option>
                      );
                    })}
                  </Select>
                </Col>
                <Col span={12}>
                  <Search
                    addonBefore={
                      <Tooltip title="Currently searching only one tag is possible">
                        <Icon type="info-circle" theme="twoTone"></Icon>
                      </Tooltip>
                    }
                    //value={this.state.filter}
                    //onChange={e => this.setState({ filter: e.target.value })}
                    onSearch={this.HandleSearch}
                    style={{ marginBottom: 5 }}
                  />
                  {this.state.previousFilter != "" && (
                    <span>
                      Previous:{" "}
                      <Tag color="blue">{this.state.previousFilter}</Tag>
                    </span>
                  )}
                  {this.state.filter != "" && (
                    <span>
                      Filter: <Tag color="volcano">{this.state.filter}</Tag>
                    </span>
                  )}
                </Col>
              </Row>
            </InputGroup>
            <KoniToolbar />
          </Col>
        </Row>
        <Row type="flex" className={styles.editorBd}>
          <Col span={1} className={styles.editorSidebar} />
          <Col span={19} className={styles.editorContent}>
            <Koni
              className={styles.koni}
              data={this.state.data}
              onNodeClick={e => this.nodeTrack(e)}
            />

            <ItemPanel>
              <RegisterLelNodes />
            </ItemPanel>
          </Col>
          <Col span={4} className={styles.editorSidebar}>
            <LelDetailPanel
              onOpen={this.props.onOpen}
              onShow={this.props.onShow}
            />
            <EditorMinimap />
          </Col>
        </Row>

        <KoniContextMenu />
      </GGEditor>
    );
  }
}
//
//

export default withPropsAPI(Graph);
