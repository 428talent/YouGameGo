import {Link} from "react-router-dom";
import React, {Component} from "react";
import {Menu} from "semantic-ui-react";
import PropTypes from 'prop-types';
import {connect} from 'dva';

const PageNav = ({dispatch, ...props}) => {
    let {activeTab} = props;

    const handleItemClick = (tag) => {
        dispatch({
            type: "pagenav/changePage",
            tag: tag
        })
    };
    console.log(props);
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
                />
            </Link>
            <Link exect to="/orders">
                <Menu.Item
                    name='订单'
                    active={activeTab === 'orders'}
                    onClick={() => handleItemClick('orders')}
                />
            </Link>
        </Menu>
    )


}

PageNav.propTypes = {
    activeTab: PropTypes.number.isRequired
};
export default connect(({pagenav}) => ({
    activeTab: pagenav.activeTab
}))(PageNav);