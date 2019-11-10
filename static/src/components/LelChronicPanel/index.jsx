import React, { Component } from "react";
import {
  Timeline,
  Tag,
  Row,
  Col,
  Slider,
  List,
  Tooltip,
  Collapse,
  Avatar,
  message,
  Drawer,
  Icon,
  Button,
  Badge,
  Calendar,
  Tabs,
  DatePicker
} from "antd";
const { Panel } = Collapse;
const { TabPane } = Tabs;
import ReactPlayer from "react-player";
import Moment from "react-moment";
import moment from "moment";
import { API } from "../api";

var config = require("Config");

const TABMODE = {
  calendar: "calendar",
  captures: "captures",
  details: "details"
};
//TODO todos
//TODO activities

class ChronicPanel extends Component {
  // might come in hany with getDuration() and seekTo
  ref = player => {
    this.player = player;
  };

  // HandleMisatoByTime get all commands by a given timeframe
  HandleMisatoByTime = (start, end) => {
    console.log("[fetch]", "[misato]", start, end);
    new API().misatoByTime(start, end).then(data => {
      console.log("[Fetch]", "[commands/byday]", data.length);
      this.setState({ currCommands: data });
    });
  };

  HandleVideoByTime = (start, end) => {
    console.log("[fetch]", "[videos]", start, end);
    new API().videoByTime(start, end).then(data => {
      console.log("[Fetch]", "[videos/byday]", data.length);
      this.setState({ currCaptures: data });
    });
  };

  // initial State
  state = {
    url: undefined,
    current: undefined,
    play: false,

    showCommands: false,
    showDetails: false,
    // startDate

    currTab: TABMODE.calendar,

    currDay: undefined,
    currMilestones: [],
    currDocs: [],
    currCaptures: [],
    currCommands: []
  };

  dateFormat = "DD.MM.YYYY";

  //date = new Date(unix_timestamp*1000);
  // lel
  dateCellRender = value => {
    let mapkey = value.format(this.dateFormat);

    let list = this.props.milestonesMap[mapkey] || [];

    const colors = [
      "cyan", // milestones
      "green",
      "blue",
      "purple",
      "geekblue",
      "magenta",
      "volcano", // due times :)
      "gold",
      "lime"
    ];
    return (
      <ul className="events" style={{ listStyleType: "none" }}>
        {list.map(item => {
          return (
            <li
              key={item.time}
              style={{ paddingLeft: "2em", textIndent: "-2em" }}
            >
              <Badge color="cyan" text={item.milestone} />
            </li>
          );
        })}
      </ul>
    );
  };

  getMonthData = value => {
    if (value.month() === 8) {
      return 1394;
    }
  };

  monthCellRender = value => {
    const num = this.getMonthData(value);
    return num ? (
      <div className="notes-month">
        <section>{num}</section>
        <span>Backlog number</span>
      </div>
    ) : null;
  };

  render() {
    const arr = this.props.milestones;

    let start = arr[0];
    const length = arr.length;
    let end = 0;
    if (length > 0) {
      end = arr[length - 1];
    }

    let days = (
      <Moment unix diff={<Moment unix>start.time</Moment>} unit="days">
        {end.time}
      </Moment>
    );

    let showBtn = (
      <Button
        type="primary"
        size="small"
        onClick={e =>
          this.setState({
            showDetails: true,
            currMilestones: this.props.milestones
          })
        }
        icon="block"
      >
        Show All
      </Button>
    );

    return (
      <div style={{ background: "white", padding: 10, borderRadius: 3 }}>
        <Row>
          <h2>Chronic {days}</h2>
        </Row>
        <Row>
          <Tabs
            tabPosition="left"
            onTabClick={(k, e) => {
              if (k === TABMODE.details) {
                this.setState({ showDetails: true });
              } else {
                this.setState({ currTab: k });
              }
            }}
            activeKey={this.state.currTab}
          >
            <TabPane
              tab={
                <span>
                  Calendar <Icon type="calendar"></Icon>
                </span>
              }
              key={TABMODE.calendar}
            >
              <Calendar
                dateCellRender={this.dateCellRender}
                onSelect={value => {
                  let mapkey = value.format("DD.MM.YYYY");
                  let list = this.props.milestonesMap[mapkey] || [];
                  this.setState({
                    showDetails: true,
                    currMilestones: list,
                    currDay: value
                  });
                  let start = value.startOf("day").format("X");
                  // X is unix timestamp
                  let stop = value.endOf("day").format("X");
                  //
                  this.HandleVideoByTime(start, stop);
                }}
                //monthCellRender={monthCellRender}
                // headerRender={e => {
                //   console.log("headerRender");
                // }}
              />
            </TabPane>
            <TabPane
              tab={
                <span>
                  Captures <Icon type="desktop"></Icon>
                </span>
              }
              key={TABMODE.captures}
            >
              {this.state.url === undefined && "No URL specified"}
              <ReactPlayer
                ref={this.ref}
                url={this.state.url}
                playing={this.state.play}
                controls={true}
                volume={0}
                width="100%"
                style={{ padding: "20px" }}
                onEnded={this.onEndedHandler}
              />
              <Collapse bordered={false}>
                <Panel header="Videos" key="1">
                  <List>{this.renderVideoList()}</List>
                </Panel>
              </Collapse>
            </TabPane>
            <TabPane
              tab={
                <span>
                  <Button type="primary" icon="project">
                    Details
                  </Button>
                </span>
              }
              key={TABMODE.details}
            ></TabPane>
          </Tabs>
        </Row>

        <Row>
          <Drawer
            title={<span>Details: {showBtn}</span>}
            width={520}
            onClose={e => {
              this.setState({ showDetails: false });
            }}
            visible={this.state.showDetails}
            //visible={this.state.visible}
          >
            {this.state.currDay !== undefined && (
              <div>
                <DatePicker
                  onChange={(v, f) => {
                    console.log(v);
                    // TODO  change milestone stuff
                    this.setState({ currDay: v });
                  }}
                  defaultValue={this.state.currDay}
                  format={this.dateFormat}
                />
                <Button type="primary" icon="reload" ghost>
                  Refresh
                </Button>
              </div>
            )}
            <Tabs defaultActiveKey="1">
              <TabPane
                tab={
                  <span>
                    <Icon type="bulb" />
                    Milestones
                  </span>
                }
                key="1"
              >
                <Timeline style={{ padding: "20px" }}>
                  {this.renderTimelineItem()}
                </Timeline>
              </TabPane>
              <TabPane
                tab={
                  <span>
                    <Icon type="video-camera" />
                    Captures
                  </span>
                }
                key="2"
              >
                {this.renderVideoList()}
              </TabPane>
              <TabPane
                tab={
                  <span>
                    <Icon type="file-text" />
                    Docs
                  </span>
                }
                key="3"
              >
                TODO
              </TabPane>
            </Tabs>
            <Drawer
              title="Commands"
              width={640}
              closable={false}
              onClose={e => {
                this.setState({ showCommands: false });
              }}
              // visible={this.state.childrenDrawer}
              visible={this.state.showCommands}
            >
              <List
                itemLayout="horizontal"
                dataSource={this.state.currCommands}
                renderItem={item => (
                  <List.Item>
                    <List.Item.Meta
                      avatar={
                        <Tooltip title={<Moment unix>{item.time}</Moment>}>
                          <Tag color="gold">{item.user}</Tag>
                        </Tooltip>
                      }
                      title={item.command}
                      //description={item.command}
                    />
                  </List.Item>
                )}
              />
            </Drawer>
            <div
              style={{
                position: "absolute",
                bottom: 0,
                width: "100%",
                borderTop: "1px solid #e8e8e8",
                padding: "10px 16px",
                textAlign: "right",
                left: 0,
                background: "#fff",
                borderRadius: "0 0 4px 4px"
              }}
            ></div>
          </Drawer>
        </Row>
      </div>
    );
  }

  onEndedHandler = e => {
    console.log("Ended", e);
    // check if next video is available -
    //this.setState({url: })
  };

  tipFormat = e => {
    return e;
  };

  renderTimelineItem() {
    //    return this.props.data.map((item, key) => (
    return this.state.currMilestones.map((item, key) => (
      <Timeline.Item key={key} time={item.time} user={item.user}>
        <Tooltip title={<Moment unix>{item.time}</Moment>}>
          <Tag color="gold">{item.user}</Tag>
        </Tooltip>
        <a href="#" onClick={e => this.handleClick(e, item.time)}>
          {item.milestone}
        </a>{" "}
        <Icon
          type="code"
          onClick={e => {
            let start = item.time;
            // X is unix timestamp
            let stop = moment
              .unix(item.time)
              .endOf("day")
              .format("X");
            //
            this.HandleMisatoByTime(start, stop);
            this.setState({ showCommands: true });
          }}
        ></Icon>
      </Timeline.Item>
    ));
  }

  // handleClick for all stuff
  handleClick = (e, time) => {
    e.preventDefault();
    console.log("Was clicked !", time);
    let candidates = this.bestVideoCandidates(time);
    if (candidates.length > 0) {
      let candidate = candidates[0];

      this.setState(
        {
          url: this.getURL(candidate.name),
          play: true,
          currTab: TABMODE.captures,
          showCommands: false,
          showDetails: false
        },
        e => {
          console.log("Seeked to ", candidate.seekto);
          //this.ref.See
          this.player.seekTo(candidate.seekto);
        }
      );
    } else {
      let msg = `No candidates matches on: ${new Date(time * 1000)} `;
      message.info(msg, 3);
    }
  };

  handleClickName = (e, name) => {
    e.preventDefault();

    this.setState({
      url: this.getURL(name),
      currTab: TABMODE.captures,
      showCommands: false,
      showDetails: false
    });
  };

  // bestVideoCandidate searches all videos and returns the best/closes video
  bestVideoCandidates = time => {
    let search = time - 20;

    return this.props.videos
      .map(entry => {
        if (entry.start < search && entry.stop > search) {
          let seekto = search - entry.start;
          return {
            success: true,
            seekto: seekto,
            name: entry.name,
            start: entry.start
          };
        } else {
          return { success: false };
        }
      })
      .filter(entry => {
        return entry.success;
      });
  };

  getURL = filename => {
    return `${config.video.get}/${filename}`;
  };

  renderVideoList() {
    return this.state.currCaptures.map((item, key) => (
      <List.Item key={key}>
        <Tag>{item.duration}</Tag>
        <a href="#" onClick={e => this.handleClickName(e, item.name)}>
          {item.name}
        </a>
      </List.Item>
    ));
  }
}

export default ChronicPanel;
