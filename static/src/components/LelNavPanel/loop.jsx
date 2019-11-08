import React from "react";
import { Tree, Tag } from "antd";
const { TreeNode } = Tree;

// todo maybe with slashes test/yees

export const loop = (data, searchValue, previous) => {
  return data
    .map(item => {
      if (item == null) {
        return null;
      }
      let tagFound = false;

      const index = item.title.indexOf(searchValue);
      const beforeStr = item.title.substr(0, index);
      const afterStr = item.title.substr(index + searchValue.length);
      let title =
        index > -1 ? (
          <span>
            {beforeStr}
            <span style={{ color: "#f50" }}>{searchValue}</span>
            {afterStr}
          </span>
        ) : (
          <span>{item.title}</span>
        );

      if (searchValue.startsWith("tags:")) {
        if (item.other !== undefined && item.other.meta !== undefined) {
          tagFound = item.other.meta.names.some(i => {
            return `tags:${i}`.indexOf(searchValue) >= 0;
          });
        }
      }

      // display all children if folder was found
      const prevFound = previous || index > -1 || tagFound;

      if (item.children) {
        const children = loop(item.children, searchValue, prevFound).filter(
          item => {
            return item != null;
          }
        );

        if (children.length > 0) {
          return (
            <TreeNode key={item.key} title={title}>
              {children}
            </TreeNode>
          );
        }
      }
      if (index > -1 || searchValue == "" || prevFound || tagFound) {
        if (item.other !== undefined && item.other.meta !== undefined) {
          const names = item.other.meta.names.map(i => {
            return <Tag key={`${item.key}-${i}`}>{i} </Tag>;
          });
          title = (
            <span>
              {title} {names}
            </span>
          );
        }

        return <TreeNode key={item.key} title={title} />;
      } else {
        return null;
      }
    })
    .filter(item => {
      return item != null;
    });
};
