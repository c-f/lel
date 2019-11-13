import React from "react";
import {
  message,
  Layout,
  Menu,
  Icon,
  Input,
  Row,
  Col,
  Button,
  Tag,
  BackTop,
  Anchor,
  Tooltip,
  Switch,
  Modal,
  AutoComplete,
  Avatar,
  Descriptions,
  Badge,
  PageHeader,
  Slider,
  Statistic
} from "antd";
const { Search } = Input;
const { Header, Content, Footer, Sider } = Layout;
const { SubMenu } = Menu;

const MODY = {
  home: "home",
  show: "show",
  chronic: "chronic",
  tags: "tags",
  checks: "checks",
  images: "images",
  commands: "commands",
  activities: "activities"
};

import {
  LelNavPanel,
  LelTagsPanel,
  LelStickyNavPanel
} from "../LelNavPanel/index.jsx";
import LelGraphPanel from "../LelGraphPanel/index.jsx";
import ImageUploader from "../LelImage/index.jsx";
import ImageSlider from "../LelImage/slider.jsx";
import SearchPanel from "../LelSearchPanel/index.jsx";
import MarkdownPanel from "./markdownpanel.jsx";
import RecorderPanel from "./recorder.jsx";
import Observer from "../LelCommunication/index.js";
import ChronicPanel from "../LelChronicPanel/index.jsx";
import MetaPanel from "../LelMetaPanel/index.jsx";
import MisatoPanel from "../LelMisatoPanel/index.jsx";
import HomePanel from "../LeLHomePanel/index.jsx";
import LoginPanel from "./loginpanel.jsx";

import { API } from "../api.js";
import { blue } from "@ant-design/colors";
// maybe BrowserRouter
// import { HashRouter as Router, Route, Link, Switch } from "react-router-dom";

// -----------------------------

var config = require("Config");

class LelApp extends React.Component {
  constructor(probs) {
    super(probs);
    this.state = {
      collapsed: true,

      // content display
      content: "",
      contentPath: "",
      meta: {},

      // meta information
      isLoading: true,
      error: null,
      connected: false,

      // content discovery
      nav: [],
      showNavBar: false,
      showStickNavbar: false,

      tags: [],
      showTagsBar: false,

      // Images
      imgs: [],
      showImageBar: false,

      // Misato logger
      misatoResults: [],
      showMisatoBar: false,

      // Graph
      showGraphBar: false,
      graphs: [],

      // globalSearchBar
      searchResults: [],
      showSearchBar: false,

      // Chronic
      chronics: [],
      chronicsMap: {},
      videos: [],
      showChronicBar: true,

      // CheckBar
      metas: [],
      showCheckBar: false,

      mode: MODY.home,
      external: false,
      markdownOnly: false,
      lighttheme: true,

      //
      showNewPath: false,
      newpath: "",

      //
      activities: [],
      showActivityBar: false,

      //
      folder: [],

      // recording Settings
      isRecording: false,

      //
      isLoggedIn: true
    };
    this.api = new API();
  }

  HandleRecording = isrecording => {
    this.setState({ isRecording: isrecording });
  };

  componentDidMount() {
    console.log("Welcome to LEL - debug javascript consoles");

    if (this.api.isLoggedIn()) {
      console.log("Using previously used auth data");
      this.setState({ isLoggedIn: true });
      this.refresh();
      Observer(config.core.channel, this.HandleUpdates);
    } else {
      this.setState({ isLoggedIn: false });
    }
  }

  refresh() {
    this.getNavigationData();
    this.getFolderData();
    this.getTagsData();
    this.getMarkdownData();

    // TODO later remove
    this.getMileStoneData();
    this.getVideoData();
    this.getImagesListData();
    this.getGraphsData();
    this.getMetasData();

    message.success("Refreshed data", 0.5);
  }

  // reconnect reconnects to the update socket
  reconnect = e => {
    Observer(config.core.channel, this.HandleUpdates);
  };

  // HandleUpdates listen for update events from the server
  HandleUpdates = (evnt, msg) => {
    if (msg === "opened") {
      message.success("Connected to server", 1);
      this.setState({
        connected: true
      });
    }
    if (msg === "message") {
      const data = JSON.parse(evnt.data);
      // rerender stuff
      // new image
      // new path
      // mod path
      console.log("Trigger refresh", evnt);
      message.info("Refreshed: " + data.message, 0.5);
      this.getNavigationData();
      this.getTagsData();
      //this.getMarkdownData();
      // TODO optimize rerendering

      this.getMileStoneData();
      this.getVideoData();
      this.getGraphsData();
      this.getImagesListData();
      this.getMetasData();
    }
    if (msg == "closed") {
      message.error("Disconnected to server");
      this.setState({
        connected: false
      });
    }
  };

  // getMetasData returns all metainformation
  getMetasData() {
    this.setState({ isLoading: true });
    this.api
      .metas()
      .then(data => {
        console.log("[Fetch]", "[metas]", data.length);
        this.setState({ metas: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  // getImages returns all the images
  getImagesListData() {
    this.setState({ isLoading: true });
    this.api
      .images()
      .then(data => {
        console.log("[Fetch]", "[imgs]", data.length);
        let imgs = data.map((item, key) => {
          return {
            original: `${config.image.get}/${item}`,
            thumbnail: `${config.image.get}/${item}`,
            item: item
          };
        });
        this.setState({ imgs: imgs, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  // getNavigationData fetches the navigation from the config.nav endpoint
  getNavigationData() {
    this.setState({ isLoading: true });

    this.api
      .nav()
      .then(data => {
        console.log("[Fetch]", "[nav]", data.length);
        this.setState({ nav: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  getFolderData() {
    this.api
      .folder()
      .then(data => {
        console.log("[Fetch]", "[folder]", data.length);
        this.setState({ folder: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  getMileStoneData() {
    this.setState({ isLoading: true });

    this.api
      .milestones()
      .then(data => {
        console.log("[Fetch]", "[milestone]", data.length);

        let dataMap = {};
        const options = { year: "numeric", month: "2-digit", day: "2-digit" };

        data.map(item => {
          //   console.log(item);
          //   return;
          let d = new Date(item.time * 1000);

          let index = d.toLocaleDateString("de-de", options);

          dataMap[index] = dataMap[index] || [];
          dataMap[index].push(item);

          return;
        });

        this.setState({
          chronics: data,
          chronicsMap: dataMap,
          isLoading: false
        });
        return false;
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  // get a list of all the video data available
  getVideoData() {
    this.setState({ isLoading: true });

    this.api
      .videos()
      .then(data => {
        console.log("[Fetch]", "[videos]", data.length);
        this.setState({ videos: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  // get a list of all the graphs available
  getGraphsData() {
    this.setState({ isLoading: true });

    this.api
      .graphs()
      .then(data => {
        console.log("[Fetch]", "[graphs]", data.length);
        this.setState({ graphs: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  // getTagsData fetches the tags data from the config.
  getTagsData() {
    this.setState({ isLoading: true });

    this.api
      .tags()
      .then(data => {
        console.log("[Fetch]", "[tags]", data.length);
        this.setState({ tags: data, isLoading: false });
      })
      .catch(error => this.setState({ error, isLoading: false }));
  }

  //getMarkdownData fetches the markdowndata and renders it
  getMarkdownData() {
    if (this.state.contentPath === "") {
      return;
    }

    this.setState({ isLoading: true });

    this.api
      .get(this.state.contentPath)
      .then(data => {
        console.log("[Fetch]", "[content]", data.length);
        this.setState({ content: data, isLoading: false });
      })
      .catch(error => this.setState({ isLoading: false }));
  }

  onCollapse = collapsed => {
    this.setState({ collapsed });
  };

  // getMetaData returns the metadata of the requested path
  getMetaData() {
    if (this.state.contentPath === "") {
      return;
    }
    this.setState({ isLoading: true });

    this.api
      .meta(this.state.contentPath)
      .then(data => {
        console.log("[Fetch]", "[meta]", data.length);
        this.setState({ meta: data, isLoading: false });
      })
      .catch(error => this.setState({ isLoading: false }));
  }

  HandleNavOnEdit = (key, title) => {
    console.log("lel");
    if (this.state.external) {
      this.HandleOpenDocument(null, key);
    }
    this.HandleNavOnSelect(key, title);
  };

  // Navigation
  HandleNavOnSelect = (key, title) => {
    let contentPath = key;

    console.log("[nav] [select] ", key, title);

    if (!contentPath.endsWith(".md")) {
      message.error("Cannot open folder", 2.5);
      return;
    }

    this.setState(
      {
        contentPath: contentPath,
        mode: MODY.show
      },
      () => {
        this.getMarkdownData();
        this.getMetaData();
      }
    );
    this.handleNavBar(false);
  };

  HandleAuth = (username, token) => {
    this.api.auth(username, token).then(ok => {
      if (ok) {
        message.success("loggedIn", 1.5);
        this.setState({ isLoggedIn: true });
        this.refresh();

        Observer(config.core.channel, this.HandleUpdates);
      } else {
        message.error("wrong login ", 2.5);
      }
    });
  };

  // Tags
  HandleTagsOnSelect = (key, title) => {
    this.HandleNavOnSelect(key, title);
    this.setState({ showTagsBar: false });
  };

  // t handler
  // todo
  handleKeyPress = event => {
    if (
      this.state.showNavBar ||
      this.state.showTagsBar ||
      this.state.showImageBar ||
      this.state.showMisatoBar ||
      this.state.showGraphBar ||
      this.state.showSearchBar ||
      this.state.showStickNavbar
    ) {
      // console.log("this.state.showNavBar", this.state.showNavBar);
      // console.log("this.state.showTagsBar", this.state.showTagsBar);
      // console.log("this.state.showImageBar", this.state.showImageBar);
      // console.log("this.state.showMisatoBar", this.state.showMisatoBar);
      // console.log("this.state.showGraphBar", this.state.showGraphBar);
      // console.log("this.state.showSearchBar", this.state.showSearchBar);
      // console.log("this.state.showStickNavbar", this.state.showStickNavbar);

      return;
    }
    /*
    if (event.key === "t") {
      this.handleNavBar(true);
    }
    if (event.key === "g") {
      this.handleGraphBar(true);
    }
    if (event.key === "s") {
      this.handleSearchBar(true);
    }
    if (event.key === "c") {
      this.handleMisatoBar(true);
    }
    if (event.key === "q") {
      this.handleTagsBar(true);
    }
    */
  };

  renderRecordStatus = () => {
    if (this.state.isRecording) {
      return <Badge status="processing" />;
    }
  };

  render() {
    //  <Route path="/" exact component={}
    let sty = {
      background: "#16314a"
      //backgroundColor: "rgb(0, 80, 100)"
    };
    if (this.state.lighttheme) {
      sty.background = "#f0f2f5";
      //sty.backgroundColor = "rgb(0, 80, 100)";
    }

    return (
      <Layout
        style={sty}
        // style={{
        //   minHeight: "100vh",
        //   marginLeft: this.state.collapsed ? "80px" : "200px"
        // }}
        // onKeyPress={this.handleKeyPress}
      >
        {/*  <input
          style={{ display: "hidden", position: "absolute", zIndex: "-1" }}
          type="text"
          id="one"
          autoFocus={true}
          onKeyPress={this.handleKeyPress}
        />
        */}
        <Header className="header" style={{ backgroundColor: "rgba(0,0,0,0)" }}>
          <Row>
            <Col span={20}>
              {/*<Sider
          collapsible
          collapsed={this.state.collapsed}
          onCollapse={this.onCollapse}
          style={{
            overflow: "auto",
            height: "100vh",
            position: "fixed",
            zIndex: 5,
            left: 0
          }}
        >
        */}
              <div className="logo" />
              <Menu
                theme="light"
                defaultSelectedKeys={["1"]}
                mode="horizontal"
                // openKeys={["navigations,search", "view-nav"]}
              >
                <Menu.Item key="view-imgs">
                  <span
                    onClick={e => {
                      this.setState({ mode: MODY.home });
                    }}
                  >
                    {this.renderStatus()}
                    <span>LeL Viewer {this.renderRecordStatus()}</span>
                  </span>
                </Menu.Item>
                <SubMenu
                  key="menu-logo"
                  title={<Icon type="setting" />}
                  subMenuCloseDelay={10}
                >
                  <Menu.ItemGroup
                    key="settings-editor"
                    title={
                      <span>
                        <Icon type="font-size" /> Editor
                      </span>
                    }
                  >
                    <Menu.Item>
                      external:{" "}
                      <span>
                        <Tooltip title="External editor">
                          <Switch
                            style={{ width: "35px" }}
                            // export
                            checkedChildren={<Icon type="block" />}
                            //unCheckedChildren="lel"
                            defaultChecked={this.state.external}
                            onChange={this.toggleExternal}
                          />
                        </Tooltip>
                      </span>
                    </Menu.Item>
                    <Menu.Item
                      onClick={e => {
                        e.preventDefault;
                      }}
                    >
                      markdown:{" "}
                      <span>
                        <Tooltip title="Markdown only">
                          <Switch
                            style={{ width: "35px" }}
                            // export
                            checkedChildren={<Icon type="fullscreen" />}
                            //unCheckedChildren="lel"
                            defaultChecked={this.state.markdownOnly}
                            onChange={this.toggleMarkdownOnly}
                          />
                        </Tooltip>
                      </span>
                    </Menu.Item>
                    <Menu.Item>
                      light:{" "}
                      <span>
                        <Tooltip title="Switch between light and dark theme">
                          <Switch
                            style={{ width: "35px" }}
                            // export
                            checkedChildren={<Icon type="bulb" />}
                            //unCheckedChildren="lel"
                            defaultChecked={this.state.lighttheme}
                            onChange={this.toggleTheme}
                          />
                        </Tooltip>
                      </span>
                    </Menu.Item>
                  </Menu.ItemGroup>
                  <Menu.ItemGroup
                    key="settings-api"
                    title={
                      <span>
                        <Icon type="api" /> API
                      </span>
                    }
                  >
                    <Menu.Item>
                      <Button
                        icon="sync"
                        onClick={e => {
                          this.refresh();
                        }}
                      >
                        {" "}
                        Refresh
                      </Button>
                    </Menu.Item>
                  </Menu.ItemGroup>
                  <Menu.ItemGroup
                    key="settings-record"
                    title={
                      <span>
                        <Icon type="audio" /> Record
                      </span>
                    }
                  >
                    <RecorderPanel onRecord={this.HandleRecording} />
                  </Menu.ItemGroup>
                  <Menu.ItemGroup
                    key="settings-login"
                    title={
                      <span>
                        <Icon type="user" /> Login
                      </span>
                    }
                  >
                    <Menu.Item>
                      <Button
                        type="primary"
                        onClick={e => {
                          e.preventDefault();
                          message.info(
                            "Currently not fully deployed feature :)",
                            2.5
                          );
                          //this.setState({ isLoggedIn: false });
                        }}
                      >
                        Change login
                      </Button>
                    </Menu.Item>
                  </Menu.ItemGroup>
                  {/*<Menu.ItemGroup key="a" title="title">
                    <Menu.Item>Option</Menu.Item>
                </Menu.ItemGroup> */}
                </SubMenu>
                <SubMenu
                  title={
                    <span className="submenu-title-wrapper">
                      <Icon type="rocket" />
                    </span>
                  }
                >
                  <Menu.Item
                    key="view-imgs"
                    onClick={() => this.handleImageBar(true)}
                  >
                    <Icon type="picture" />
                    <span>Images </span>
                  </Menu.Item>
                  <Menu.Item
                    key="view-chronic"
                    onClick={() => this.handleChronicBar(true)}
                  >
                    <Icon type="fire" />
                    <span>Chronic</span>
                  </Menu.Item>
                  <Menu.Item
                    key="view-logs"
                    onClick={() => this.handleMisatoBar(true)}
                  >
                    <Icon type="experiment" />
                    <span>CommandLog </span>
                  </Menu.Item>
                  <Menu.Item
                    key="view-checks"
                    onClick={() => this.handleCheckBar(true)}
                  >
                    <Icon type="check-square" />
                    <span>Todos</span>
                  </Menu.Item>
                  {/* <Menu.Divider></Menu.Divider> */}

                  {/* <Menu.Divider></Menu.Divider> */}
                  <Menu.Item
                    key="view-activities"
                    disabled={true}
                    onClick={() => {
                      e.preventDefault();
                      message.info(
                        "Currently not fully deployed feature :)",
                        2.5
                      );
                      // this.handleActivityBar(true)
                    }}
                  >
                    <Icon type="alert" />
                    <span>Activities</span>
                  </Menu.Item>
                </SubMenu>

                {/* <Menu.Divider></Menu.Divider> */}
                <Menu.Item
                  key="view-nav"
                  onClick={() => this.handleNavBar(true)}
                >
                  <Icon type="menu" />
                  <span>Navigation</span>
                </Menu.Item>
                <Menu.Item
                  key="view-graph"
                  onClick={() => this.handleGraphBar(true)}
                >
                  <Icon type="branches" />
                  <span>Graph </span>
                </Menu.Item>
                <Menu.Item
                  key="view-tags"
                  onClick={() => this.handleTagsBar(true)}
                >
                  <Icon type="tags" />
                  <span>Tags </span>
                </Menu.Item>

                <Menu.Item
                  key="view-search"
                  onClick={() => this.handleSearchBar(true)}
                >
                  <Icon type="search" />
                  <span>Search </span>
                </Menu.Item>

                {/* <Menu.Divider></Menu.Divider> */}
              </Menu>
            </Col>
            <Col span={4}>
              <Menu
                theme="light"
                defaultSelectedKeys={["1"]}
                mode="horizontal"
                // openKeys={["navigations,search", "view-nav"]}
              >
                <Menu.Item
                  key="view-new-file"
                  onClick={e => {
                    this.setState({
                      showNewPath: true
                    });
                  }}
                >
                  <Icon type="file-add" />
                  <span>New File</span>
                </Menu.Item>
                <Menu.Item
                  key="new-snippet"
                  onClick={e => {
                    this.handleNewDocument();
                  }}
                >
                  <Icon type="snippets" />
                  <span>Snippet</span>
                </Menu.Item>
              </Menu>
            </Col>
          </Row>
        </Header>
        {/* </Sider> */}
        <Layout style={sty}>
          <Sider
            theme="light"
            collapsible
            width={600}
            collapsedWidth={0}
            onCollapse={e => {
              this.setState({ showStickNavbar: !e });
            }}
            defaultCollapsed={!this.state.showStickNavbar}
            style={{
              //overflow: "auto",
              height: "100vh",
              position: "fixed",
              zIndex: 5,
              left: 0
            }}
          >
            <LelStickyNavPanel
              data={this.state.nav}
              onSelect={this.HandleNavOnSelect}
              onClickSpecial={this.HandleNavOnEdit}
              //visible={this.state.showNavBar}
              //onClose={() => this.handleNavBar(false)}
            />
          </Sider>

          <LelGraphPanel
            visible={this.state.showGraphBar}
            graphs={this.state.graphs}
            onSelect={this.HandleGraphOnSelect}
            onOpen={this.HandleOpenDocument}
            onShow={this.HandleNavOnSelect}
            onClose={() => this.handleGraphBar(false)}
          />
          <LelTagsPanel
            data={this.state.tags}
            onSelect={this.HandleTagsOnSelect}
            visible={this.state.showTagsBar}
            onClose={() => this.handleTagsBar(false)}
          />
          <SearchPanel
            title="Content Search"
            dataSource={this.state.searchResults}
            placement="left"
            height="100%"
            width="100%"
            columns={[
              {
                title: "Path",
                dataIndex: "path",
                key: "path",
                render: text => {
                  return (
                    <span>
                      <a href="#" onClick={e => this.handleSearchClick(text)}>
                        {text}
                      </a>
                      <Icon
                        style={{ marginLeft: 5 }}
                        type="edit"
                        onClick={e => this.HandleOpenDocument(null, text)}
                      />
                    </span>
                  );
                }
                // fixed: "left"
              },
              {
                title: "Match",
                dataIndex: "match",
                key: "match"
                // width: 150
              }
            ]}
            onClose={() => this.handleSearchBar(false)}
            onSearch={this.HandleSearch}
            visible={this.state.showSearchBar}
          />

          <LelNavPanel
            data={this.state.nav}
            onSelect={this.HandleNavOnSelect}
            onClickSpecial={this.HandleNavOnEdit}
            visible={this.state.showNavBar}
            onClose={() => this.handleNavBar(false)}
          />
          <LoginPanel
            handleAuth={this.HandleAuth}
            isLoggedIn={this.state.isLoggedIn}
          />

          <Content
            style={{
              margin: "0px 16px 10px",
              overflow: "initial",
              marginLeft: this.state.showStickNavbar ? "600px" : "10px",
              padding: "20px",
              width: "100%",
              minHeight: "calc(100vh - 170px)"
            }}
          >
            {this.renderContent()}
          </Content>

          <Modal
            title="New Path"
            visible={this.state.showNewPath}
            onOk={e => {
              let newpath = this.state.newpath;
              if (!newpath.endsWith(".md")) {
                newpath += ".md";
              }
              this.handleNewDocument(newpath);
              this.setState({
                showNewPath: false
              });
            }}
            onCancel={e => {
              this.setState({ showNewPath: false });
            }}
          >
            <AutoComplete
              value={this.state.newpath}
              prefix={<Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />}
              dataSource={this.state.folder}
              style={{ width: 200 }}
              onSelect={this.onSelect}
              onSearch={this.onSearch}
              onChange={e => {
                this.setState({
                  newpath: e
                });
              }}
            />
          </Modal>
        </Layout>
        <Footer
          style={{
            textAlign: "center",
            display: "block",
            background: sty.background
          }}
        >
          <Tooltip title="Copy your image and simple paste it into this space (maybe click first)">
            <div
              style={{
                textAlign: "center",
                width: "200px",

                float: "right"
              }}
            >
              <ImageUploader
                onUpload={this.HandleImageUpload}
                onError={this.HandleError}
              />
            </div>
          </Tooltip>
          <Button
            type="dashed"
            target="_blank"
            href="https://editor.l3l.lol/user/"
            icon="question-circle"
          >
            Help
          </Button>
        </Footer>
        <BackTop />
      </Layout>
    );
  }

  onSelect = searchText => {
    console.log("Select", searchText);
  };

  // fnc
  onSearch = searchText => {
    console.log("Searches", this.state.nav);
  };

  /*
  // TODO 
       <AutoComplete
          value={this.state.newpath}
          prefix={<Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />}
          dataSource={this.state.nav}
          style={{ width: 200 }}
          onSelect={onSelect}
          onSearch={this.onSearch}
          onChange={e => {
                this.setState({
                  newpath: e.target.value
                });
              }}
          
        />

        <Input
              value={this.state.newpath}
              prefix={<Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />}
              onChange={e => {
                this.setState({
                  newpath: e.target.value
                });
              }}
  */

  // renderContent renders content
  renderContent() {
    switch (this.state.mode) {
      case MODY.chronic:
        return (
          <ChronicPanel
            style={{ background: "white" }}
            milestones={this.state.chronics}
            milestonesMap={this.state.chronicsMap}
            videos={this.state.videos}
          ></ChronicPanel>
        );
      case MODY.show:
        return (
          <MarkdownPanel
            content={this.state.content}
            meta={this.state.meta}
            contentPath={this.state.contentPath}
            onImageLoad={this.HandleImageOnLoad}
            onHelpClick={this.HandleHelpClick}
            onBack={this.HandleBackClick}
            HandleOpenDocument={this.HandleOpenDocument}
            markdownOnly={this.state.markdownOnly}
            // onClick
            //isLoading={this.state.isLoading}
          />
        );
      case MODY.show:
        return <div>Damn currently not implemented :/</div>;
      case MODY.images:
        return <ImageSlider images={this.state.imgs} />;
      case MODY.checks:
        return <MetaPanel data={this.state.metas} />;
      case MODY.home:
        return <HomePanel></HomePanel>;
      case MODY.commands:
        return (
          <MisatoPanel
            dataSource={this.state.misatoResults}
            onSearch={this.HandleMisatoSearch}
          />
        );
      case MODY.activities:
        return <div>Damn currently not implemented</div>;
      default:
        return "not sure what you mean ....";
    }
  }

  HandleHelpClick = event => {
    window.open("https://editor.l3l.lol/user", "_blank");
  };

  HandleBackClick = () => {
    this.setState({
      mode: MODY.home
    });
  };

  handleSearchClick = path => {
    this.setState(
      {
        contentPath: path,
        mode: MODY.show
      },
      this.getMarkdownData
    );
    this.handleSearchBar(false);
  };

  HandleSearch = search => {
    this.api.searchContent(search).then(data => {
      this.setState({
        searchResults: data
      });
    });
  };

  // HandleMisatoSearch searches for commands
  HandleMisatoSearch = search => {
    this.api.misatoSearch(search).then(data => {
      data.map(element => {
        let line = JSON.parse(element.match);
        element["command"] = line["command"];

        return element;
      });
      this.setState({
        misatoResults: data
      });
    });
  };

  // HandleImageOnLoad -->
  HandleImageOnLoad = uri => {
    console.log("Image URL", uri);
    uri = uri.replace(/^([\/.]{1,2})(images\/)?/, "");
    return `${config.image.get}/${uri}`;
  };
  // HandleOpenDocument fetches triggers a http request to open the document via LEL
  HandleOpenDocument = (data, key) => {
    let path = key;
    if (key == undefined) {
      path = this.state.contentPath;
    }
    console.log("Open Document", data, path);
    fetch(config.notes.open + "?path=" + path).then(response => {
      message.success("Opening: " + path, 2.5);
    });
  };

  // HandleError is the handler for all messenges
  HandleError(data) {
    message.error(JSON.stringify(data), 2.5);
  }

  // renderMessages returns all Messages
  renderMessages() {
    //this.state.e;
    message.success(
      "This is a prompt message for success, and it will disappear in 10 seconds",
      10
    );
    message.loading("Action in progress..", 2.5);
  }

  // HandleImageUpload gets called when the upload is finished
  HandleImageUpload(data) {
    message.info(JSON.stringify(data), 2);
  }

  renderStatus() {
    let status = <Icon type="deployment-unit" style={{ color: "green" }} />;
    if (!this.state.connected) {
      status = (
        <Icon
          type="warning"
          onClick={this.reconnect}
          theme="twoTone"
          twoToneColor="#eb2f96"
        />
      );
    }
    let mode = MODE;
    if (mode !== "prod") {
      status = (
        <span>
          <Tag>{mode}</Tag>
          {status}
        </span>
      );
    }
    return status;
  }

  toggleMarkdownOnly = e => {
    this.setState({
      markdownOnly: !this.state.markdownOnly
    });
  };

  toggleTheme = e => {
    this.setState({
      lighttheme: !this.state.lighttheme
    });
  };

  toggleExternal = e => {
    this.setState({
      external: !this.state.external
    });
  };

  handleCheckBar(state) {
    this.setState({
      showCheckBar: state,
      mode: MODY.checks
    });
  }

  handleActivityBar(state) {
    console.log("clicky");
    this.setState({
      showActivityBar: state,
      mode: MODY.activities
    });
  }

  handleChronicBar(state) {
    this.setState({
      showChronicBar: state,
      mode: MODY.chronic
    });
  }
  handleSearchBar(state) {
    this.setState({
      showSearchBar: state
    });
  }
  handleMisatoBar(state) {
    this.setState({
      showMisatoBar: state,
      mode: MODY.commands
    });
  }
  handleImageBar(state) {
    this.setState({
      showImageBar: state,
      mode: MODY.images
    });
  }
  handleTagsBar(state) {
    this.setState({
      showTagsBar: state
    });
  }

  handleGraphBar(state) {
    this.setState({
      showGraphBar: state
    });
  }

  handleNavBar(state) {
    this.setState({
      showNavBar: state
    });
  }
  handleNewDocument(path) {
    if (path == undefined || path == "") {
      path = uuidv4() + ".md";
    }
    if (!this.state.external) {
      this.setState({
        contentPath: path,
        mode: "show",
        content: ""
      });
    }

    // toggleExternal
    fetch(
      config.notes.open + "?editor=" + this.state.external + "&path=" + path
    ).then(e => {
      message.success("Opening: " + path, 1.5);
    });
  }
}

function uuidv4() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function(c) {
    var r = (Math.random() * 16) | 0,
      v = c == "x" ? r : (r & 0x3) | 0x8;
    return v.toString(16);
  });
}

export default LelApp;
