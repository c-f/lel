import React from "react";
import EditorMinimap from "../EditorMinimap";
import { KoniContextMenu } from "../EditorContextMenu";
import { KoniToolbar } from "../EditorToolbar";
import { KoniItemPanel } from "../EditorItemPanel";
import LelDetailPanel from "../LelEditorDetailPanel/index.jsx";
import { HandleLayout, FilterNodes } from "./layout.js";
import RegisterLelNodes from "./node.jsx";
import GGEditor, { Koni, withPropsAPI, Item, ItemPanel } from "gg-editor";

import { Row, Col } from "antd";

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
      ok: false
    };
  }
  componentDidMount() {
    this.refreshData();
    this.setState({
      ok: true
    });
  }
  show() {
    console.log(this.state);
  }
  nodeTrack(e) {
    console.log(e);
    //console.log(this.editor.propsAPI.save())
    //var old = this.HandleLayout(data);
    //this.editor.propsAPI.read(old);
  }

  componentWillReceiveProps(nextProps) {
    console.log("updateing");
    if (this.props.filter !== nextProps.filter) {
      this.setState(
        {
          filter: nextProps.filter
        },
        this.HandleData
      );
    }
  }
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

  filter = () => {
    // console.log("lel")

    filter;
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
            <a href="#" onClick={e => this.refreshData(e)}>
              Refresh
            </a>
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
