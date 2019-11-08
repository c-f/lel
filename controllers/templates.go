package controllers

import (
	"html/template"
	"log"
	"net/http"
)

func getTemplate(w http.ResponseWriter, tmpl string) *template.Template {
	templates := template.Must(template.ParseGlob(tmpl))
	return templates
}

func getBase(w http.ResponseWriter) *template.Template {
	tmpl, err := template.New("app").Parse(appTemplate)
	if err != nil {
		log.Fatal(err)
	}
	return tmpl
}

const appTemplate = `{{ define "base" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta name="description" content="LEL is a visualization layer and helper for relevant documentation and note taking focused on IT operations." />
    <meta name="author" content="CF <script>alert('&')</script> Tim" />
    <meta name="homepage" content="https://editor.l3l.lol" />

    <title>LeL - explode your notes</title>
    <link rel="shortcut icon" type="image/x-icon" href="https://editor.l3l.lol/favicon.png" />

    <link rel="stylesheet" href="/static/vendor/antd.3.16.1.min.css" />
    
    <style type="text/css">
      #root,
      #root > div {
        display: flex;
        width: 100%;
        height: 100%;
        overflow: hidden;
      }

      .ant-tree.ant-tree-show-line.hide-file-icon
        li
        span.ant-tree-switcher-noop {
        background: transparent;
      }
      .ant-tree.ant-tree-show-line li span.ant-tree-switcher {
        background: none;
      }

      /*override */
      .ant-tree.ant-tree-show-line.hide-file-icon li:before {
        content: " ";
        width: 1px;
        border-left: 1px solid #d9d9d9;
        height: 100%;
        position: absolute;
        left: 12px;
        top: 0;
        margin: 0;
      }

      .ant-tree.ant-tree-show-line.hide-file-icon li:first-child:before {
        top: 6px;
        height: calc(100% - 6px);
      }

      .ant-tree.ant-tree-show-line.hide-file-icon li:last-child:before {
        height: 16px;
      }

      .ant-tree.ant-tree-show-line.hide-file-icon
        li:first-child:last-child:before {
        height: 10px;
      }

      .ant-tree.ant-tree-show-line.hide-file-icon
        li
        .ant-tree-switcher-noop
        > i {
        visibility: hidden;
      }

      .ant-tree.ant-tree-show-line.hide-file-icon
        li
        .ant-tree-switcher-noop:after {
        content: " ";
        height: 1px;
        border-bottom: 1px solid #d9d9d9;
        width: 10px;
        position: absolute;
        left: 12px;
        top: 50%;
        margin: 0;
      }
      div.CodeMirror.cm-s-rubyblue,
      div.lel-md-panel {
        height: auto;
        border-radius: 4px;
        min-height: 40vh;
      }
      div.CodeMirror-scroll {
        overflow-x: hidden !important;
      }
      div.ant-anchor,
      div.ant-anchor-wrapper {
        padding-left: 0px;
      }
      div.ant-anchor-ink {
        display: none;
      }

      div.image-gallery {
        margin-bottom: 20px;
      }

      div.image-gallery-image {
        text-align: center;
        border-radius: 5px;

        min-height: 440px;
        padding: 20px;

        background-color: #ffffff;
        /*

        */
        background-image: url("data:image/svg+xml,%3Csvg width='64' height='64' viewBox='0 0 64 64' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M8 16c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8zm0-2c3.314 0 6-2.686 6-6s-2.686-6-6-6-6 2.686-6 6 2.686 6 6 6zm33.414-6l5.95-5.95L45.95.636 40 6.586 34.05.636 32.636 2.05 38.586 8l-5.95 5.95 1.414 1.414L40 9.414l5.95 5.95 1.414-1.414L41.414 8zM40 48c4.418 0 8-3.582 8-8s-3.582-8-8-8-8 3.582-8 8 3.582 8 8 8zm0-2c3.314 0 6-2.686 6-6s-2.686-6-6-6-6 2.686-6 6 2.686 6 6 6zM9.414 40l5.95-5.95-1.414-1.414L8 38.586l-5.95-5.95L.636 34.05 6.586 40l-5.95 5.95 1.414 1.414L8 41.414l5.95 5.95 1.414-1.414L9.414 40z' fill='%231c7c9f' fill-opacity='0.04' fill-rule='evenodd'/%3E%3C/svg%3E");
      }
      div.image-gallery-image img {
        max-height: 400px;

        max-width: 100%;
        width: auto;
        margin: auto;
      }

      div.ant-fullcalendar-calendar-body {
        background-color: white;
      }
    </style>
    <script>
      const base = "https://127.0.0.1:8888";
      
      const socketBase = "wss://127.0.0.1:8888";

        const lel = {
        base: base + "/",
        graph: {
          current: base + "/api/feat/graph/get",
          upload: base + "/api/feat/upload/graph",
          get: base + "/graphs"
        },
        image: {
          upload: base + "/api/feat/upload/image",
          get: base + "/images"
        },
        video: {
          upload: base + "/api/feat/upload/video",
          get: base + "/videos"
        },
        notes: {
          get: base + "/api/notes/get",
          open: base + "/api/notes/open",
          remove: base + "/api/notes/remove",
          meta: base + "/api/notes/meta",
          upload: base + "/api/feat/upload/md"
        },
        core: {
          // list all the things
          ok: base + "/api/core/ok",
          stats: base + "/api/core/stats",

          nav: base + "/api/core/nav",
          folder: base + "/api/core/folder",
          video: base + "/api/core/videos",
          images: base + "/api/core/images",
          graphs: base + "/api/core/graphs",

          // within
          tags: base + "/api/core/tags",
          milestone: base + "/api/core/milestones",
          metas: base + "/api/core/metas",
          search: base + "/api/core/search",

          channel: socketBase + "/api/core/channel"
        },
        misato: {
          search: base + "/api/feat/misato/search",
          section: base + "/api/feat/misato"
        }
      };
      // TODO dynamic Routes 
    </script>
    <!-- 
    href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.25.2/codemirror.min.css"
    href="https://cdnjs.cloudflare.com/ajax/libs/codemirror/5.25.2/theme/rubyblue.min.css"
    href="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.11.0/styles/github.min.css"
    -->
    
    <link
      rel="stylesheet"
      href="/static/vendor/codemirror.5.25.2.min.css"
    />
    <link
      rel="stylesheet"
      href="/static/vendor/github.9.11.0.min.css"
    />
    <link
      rel="stylesheet"
      href="/static/vendor/rubyblue.5.25.2.min.css"
    />
  </head>

  <body>
    <div id="app"></div>

    <!-- App -->
    <script src="/static/app-prod.js"></script>
  </body>
</html>
{{ end }}
`
