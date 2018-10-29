import {Link} from "react-router-dom";
import React, {Component} from "react";
import {Menu} from "semantic-ui-react";
import PropTypes from 'prop-types';
import {connect} from 'dva';

const PageNav = ({dispatch, ...props}) => {
    let {activeTab} = props;

    const handleItemClick = (activeTab) => {
        dispatch({
            type: "app/changeTab",
            activeTab,
        })
    };
    return (
        <Menu secondary>
            <Link exect to="/">
                <Menu.Item name='home' active={activeTab === 'home'}
                           onClick={() => handleItemClick('home')}/>
            </Link>
            <Link exect to="/games">
                <Menu.Item
                    name='games'
                    active={activeTab === 'games'}
                    onClick={() => handleItemClick('games')}
                >
                </Menu.Item>
            </Link>
            <Link exect to="/orders">
                <Menu.Item
                    name='订单'
                    active={activeTab === 'orders'}
                    onClick={() => handleItemClick('orders')}
                />
            </Link>
            <Link exect to="/cart">
                <Menu.Item
                    name='购物车'
                    active={activeTab === 'cart'}
                    onClick={() => handleItemClick('cart')}
                />
            </Link>
        </Menu>
    )


};

PageNav.propTypes = {
    activeTab: PropTypes.string.isRequired
};
export default connect(({app}) => ({
    activeTab: app.activeTab
}))(PageNav);