import React,{ Component } from 'react';
import { Menu, Icon, Spin } from 'antd';
import styles from 'components/common/Common.less';

const SubMenu = Menu.SubMenu;
const MenuItemGroup = Menu.ItemGroup;

class Layout extends React.Component {
	render(){
		return (
			<div>
				<Menu mode="horizontal">
					<Menu.Item>主菜单</Menu.Item>
					<SubMenu title="子菜单">
						<Menu.Item>子菜单项</Menu.Item>
					</SubMenu>
				</Menu>
				<div className={styles.content}>
					{this.props.children}
				</div>
			</div>
		)
	}
}

export default Layout;
