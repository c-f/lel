import React from "react";
import { Tree, Tag, Icon } from "antd";
const { TreeNode } = Tree;

// todo maybe with slashes test/yees

export const loop = (
  data,
  searchValue,
  selectHandler = () => {},
  clickHandler = null,
  previous = null
) => {
  return data
    .map((item, mapIndex) => {
      if (item == null) {
        return null;
      }
      let tagFound = false;

      const index = item.title.indexOf(searchValue);
      const beforeStr = item.title.substr(0, index);
      const afterStr = item.title.substr(index + searchValue.length);
      let title =
        index > -1 ? (
          <span
            onClick={e => {
              e.preventDefault();
              selectHandler(item.key, item.title);
            }}
          >
            {beforeStr}
            <span style={{ color: "#f50" }}>{searchValue}</span>
            {afterStr}
          </span>
        ) : (
          <span
            onClick={e => {
              e.preventDefault();
              selectHandler(item.key, item.title);
            }}
          >
            {item.title}
          </span>
        );

      if (searchValue.startsWith("tags:")) {
        if (item.other !== undefined && item.other.meta !== undefined) {
          tagFound = item.other.meta.names.some(i => {
            return (
              i.indexOf(
                searchValue.substr("tags:".length, searchValue.length)
              ) >= 0
            );
          });
        }
      }

      // display all children if folder was found
      const prevFound = previous || index > -1 || tagFound;

      if (item.children) {
        const children = loop(
          item.children,
          searchValue,
          selectHandler,
          clickHandler,
          prevFound
        ).filter(item => {
          return item != null;
        });

        if (children.length > 0) {
          return (
            <TreeNode key={mapIndex + "__" + item.key} title={title}>
              {children}
            </TreeNode>
          );
        }
      }
      if (index > -1 || searchValue == "" || prevFound || tagFound) {
        let names = "";
        let attrs = "";
        if (item.other !== undefined && item.other.meta !== undefined) {
          names = item.other.meta.names.map(i => {
            return <Tag key={`${item.key}-${i}`}>{i} </Tag>;
          });
        }

        if (clickHandler != null) {
          attrs = (
            <Icon
              type="edit"
              onClick={e => {
                e.preventDefault();
                clickHandler(item.key, title);
              }}
            ></Icon>
          );
        }

        title = (
          <span>
            {title} {names} {attrs}
          </span>
        );

        return <TreeNode key={mapIndex + "__" + item.key} title={title} />;
      } else {
        return null;
      }
    })
    .filter(item => {
      return item != null;
    });
};
