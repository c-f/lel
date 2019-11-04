// api contains every apiCall

var config = require("Config");

export class API {
  constructor(endpoint = "") {
    if (!!API.instance) {
      return API.instance;
    }

    API.instance = this;

    this.endpoint = endpoint;
    let token = window.localStorage.getItem("token");
    let username = window.localStorage.getItem("username");
    window.localStorage.setItem("auth", false);
    this.auth(username, token);

    return this;
  }

  isLoggedIn = () => {
    return window.localStorage.getItem("auth") || false;
  };

  auth = (username, token) => {
    return this.ok(token).then(ok => {
      if (ok) {
        this.username = username;
        this.token = token;
        window.localStorage.setItem("username", username);
        window.localStorage.setItem("token", username);
      }
      return ok;
    });
  };

  ok = token => {
    return fetch(config.core.ok, {
      headers: {
        Authorization: `Token ${token}`
      }
    }).then(res => {
      if (res.ok) {
        window.localStorage.setItem("auth", true);
        return true;
      }
      return false;
    });
  };

  searchContent = search => {
    return fetch(config.core.search + "?search=" + encodeURI(search), {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  metas = () => {
    return fetch(config.core.metas, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  images = () => {
    return fetch(config.core.images, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  nav = () => {
    return fetch(config.core.nav, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  folder = () => {
    return fetch(config.core.folder, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  milestones = () => {
    return fetch(config.core.milestone, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  videos = () => {
    return fetch(config.core.video + "?structure=fast", {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  tags = () => {
    return fetch(config.core.tags, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  get = contentPath => {
    return fetch(
      config.notes.get + "?path=" + encodeURIComponent(contentPath),
      {
        headers: {
          Authorization: `Token ${this.token}`
        }
      }
    ).then(response => {
      if (response.ok) {
        return response.text();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  meta = contentPath => {
    return fetch(
      config.notes.meta + "?path=" + encodeURIComponent(contentPath),
      {
        headers: {
          Authorization: `Token ${this.token}`
        }
      }
    ).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  misatoSearch = search => {
    return fetch(config.misato.search + "?search=" + encodeURI(search), {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      } else {
        throw new Error("Something went wrong ...");
      }
    });
  };

  uploadImg = data => {
    return fetch(config.image.upload, {
      method: "PUT",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      }

      throw new Error("Something went wrong");
    });
  };

  graph = () => {
    return fetch(config.graph.get, {
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      }

      throw new Error("Something went wrong");
    });
  };

  misatoByTime = (start, end) => {
    return fetch(
      config.misato.section +
        `?start=${encodeURI(parseInt(start))}&stop=${encodeURI(parseInt(end))}`,
      {
        headers: {
          Authorization: `Token ${this.token}`
        }
      }
    ).then(response => {
      if (response.ok) {
        return response.json();
      }

      throw new Error("Something went wrong");
    });
  };

  videoByTime = (start, end) => {
    return fetch(
      config.core.video +
        `?start=${encodeURI(parseInt(start))}&stop=${encodeURI(parseInt(end))}`,
      {
        headers: {
          Authorization: `Token ${this.token}`
        }
      }
    ).then(response => {
      if (response.ok) {
        return response.json();
      }

      throw new Error("Something went wrong");
    });
  };

  uploadVideo = body => {
    return fetch(config.video.upload, {
      method: "post",
      body: body,
      headers: {
        Authorization: `Token ${this.token}`
      }
    }).then(response => {
      if (response.ok) {
        return response.json();
      }

      throw new Error("Something went wrong");
    });
  };

  uploadMarkdown = (body, contentPath) => {
    return fetch(config.notes.upload + "?path=" + encodeURI(contentPath), {
      method: "post",
      body: body,
      headers: {
        Authorization: `Token ${this.token}`
      }
    });
  };
}
