export const Observer = (url, callbackFn) => {
  let ws = new WebSocket(url);
  ws.onopen = () => {
    // on connecting, do nothing but log it to the console
    console.log("connected");
    callbackFn({}, "opened");
    ws.send("Ping");
  };

  ws.onmessage = evt => {
    console.log("Message");
    callbackFn(evt, "message");
  };

  ws.onclose = () => {
    console.log("Closed");
    callbackFn({}, "closed");
    // automatically try to reconnect on connection loss

    ws = new WebSocket(url);
  };
};

export default Observer;
