import React from 'react';
import { Command } from 'gg-editor';
import upperFirst from 'lodash/upperFirst';
import styles from './index.less';
import {Icon} from 'antd';

const MenuItem = (props) => {
  const { command, icon, text } = props;

  return (
    <Command name={command}>
      <div className={styles.item}>
        <Icon type={`${icon || command}`} />
        <span>{text || upperFirst(command)}</span>
      </div>
    </Command>
  );
};

export default MenuItem;
