import React, { Component } from "react";
import { Button, Slider, Menu, message } from "antd";
import { RecordRTCPromisesHandler, getSeekableBlob } from "recordrtc";

var config = require("Config");
import * as EBML from "ts-ebml";
import { API } from "../api";
// var ffmpeg = require("ffmpeg.js");

class RecorderPanel extends Component {
  constructor(props) {
    super(props);

    //this.blockSite(true);
    this.state = {
      // record interval in minutes
      isRecording: false,
      interval: 5,
      framerate: 30
    };
    this.api = new API();
  }

  //  Upload the video
  HandleUploadVideo = (blob, start, converted) => {
    var fd = new FormData();
    fd.append("uploadFile", blob, "blobby.txt");
    fd.append("timestamp-end", new Date().getTime());
    fd.append("timestamp-start", start);
    fd.append("converted", converted);

    this.api.uploadVideo(fd);
  };

  getSeekableBlob = (inputBlob, callback) => {
    // EBML.js copyrights goes to: https://github.com/legokichi/ts-ebml
    if (typeof EBML === "undefined") {
      throw new Error("Please link: https://cdn.webrtc-experiment.com/EBML.js");
    }
    var reader = new EBML.Reader();
    var decoder = new EBML.Decoder();
    var tools = EBML.tools;
    var fileReader = new FileReader();
    fileReader.onload = function(e) {
      var ebmlElms = decoder.decode(this.result);
      ebmlElms.forEach(function(element) {
        reader.read(element);
      });
      reader.stop();
      var refinedMetadataBuf = tools.makeMetadataSeekable(
        reader.metadatas,
        reader.duration,
        reader.cues
      );
      var body = this.result.slice(reader.metadataSize);
      var newBlob = new Blob([refinedMetadataBuf, body], {
        type: "video/webm"
      });
      callback(newBlob);
    };
    fileReader.readAsArrayBuffer(inputBlob);
  };

  // todo save record :)
  stop = e => {
    this.state.recorder.stopRecording();

    if (this.state.stream) {
      this.state.stream.getTracks().forEach(track => track.stop());
    }

    this.setState({ stream: null, isRecording: false });
    this.props.onRecord(false);
    this.blockSite(false);
  };

  start = e => {
    navigator.mediaDevices
      .getDisplayMedia({
        video: {
          cursor: "always",
          frameRate: this.state.framerate

          //width: 1920, //  max: 1920 },
          //height: 1080 // { min: 776, ideal: 720, max: 1080 }
        },
        audio: false
      })
      .then(stream => {
        // enable record
        this.setState({ isRecording: true, stream: stream });
        this.props.onRecord(true);
        this.blockSite(true);

        let track = stream.getVideoTracks()[0];
        let settings = track.getSettings();

        stream.addEventListener(
          "inactive",
          function() {
            console.log("MHM TRIGGERED");
          },
          false
        );
        stream.getTracks().forEach(track => {
          track.addEventListener(
            "ended",
            e => {
              this.stop();
            },
            false
          );
          track.addEventListener(
            "inactive",
            e => {
              this.stop();
            },
            false
          );
        });

        message.info(
          `Start recording  width:${settings.width} height:${settings.height}`,
          3
        );
        (async () => {
          let recorder = new RecordRTCPromisesHandler([stream], {
            type: "video",
            video: {
              width: settings.width,
              height: settings.height,
              frameRate: this.state.framerate
            },
            mimeType: "video/webm;codecs=vp8"
          });
          this.setState({ recorder: recorder });

          // sleep or maybe stop
          const sleep = m => new Promise(r => setTimeout(r, m));

          let start = null;
          let tmp = null;
          let recordtime = 1 * 60 * 1000;

          for (; this.state.isRecording; ) {
            start = new Date().getTime();
            recordtime = parseInt(this.state.interval) * 60 * 1000;
            // await sleep(1000);

            console.log("[record]", "[start]", start, recordtime);
            recorder.startRecording();

            await sleep(recordtime);
            await recorder.stopRecording();

            let blob = await recorder.getBlob();
            tmp = start;
            this.getSeekableBlob(blob, rawblob => {
              console.log("[record]", "[send]", rawblob);
              this.HandleUploadVideo(rawblob, tmp, false);
            });

            if (!this.state.stream.active) {
              console.log("INTERRUPT BY USER");
              this.setState({ isRecording: false });
            }

            blob = null;
            await recorder.reset();
          }
        })();
      });
  };

  //
  render() {
    let btn = (
      <Button icon="play-circle" onClick={this.start}>
        Record
      </Button>
    );
    if (this.state.isRecording) {
      btn = (
        <Button icon="stop" onClick={this.stop}>
          Stop
        </Button>
      );
    }
    return (
      <div style={{ padding: "0px 16px 0px 24px" }}>
        <div>{btn}</div>
        <div style={{ marginTop: "25px" }}>
          <span>Record time</span>
          <Slider
            max={30}
            min={1}
            key="record-interval"
            onChange={val => {
              this.setState({ interval: val });
            }}
            value={this.state.interval}
          />
          <span>Framerate</span>
          <Slider
            max={60}
            min={1}
            key="frame-rate"
            onChange={val => {
              this.setState({ framerate: val });
            }}
            value={this.state.framerate}
          />
        </div>
      </div>
    );
  }

  blockSite = state => {
    let fnc = null;
    if (state) {
      fnc = () => {
        return true;
      };
    }
    //window.onbeforeunload = state ? () => { return true} : null
    window.onbeforeunload = fnc;
  };
}

export default RecorderPanel;
