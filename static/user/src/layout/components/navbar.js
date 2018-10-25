import React, {Component} from 'react'
import {Menu} from 'semantic-ui-react'
import {Link} from "react-router-dom";

export default class MainNavBar extends Component {
    state = {};

    handleItemClick = (e, {name}) => this.setState({activeItem: name});

    render() {
        const {activeItem} = this.state;
        return (
            <Menu stackable inverted>
                <Menu.Item header>YouGame</Menu.Item>

                    <Menu.Item
                        name='Home'
                        active={activeItem === 'features'}
                        onClick={this.handleItemClick}
                    >
                        Home
                    </Menu.Item>


                    <Menu.Item
                        name='testimonials'
                        active={activeItem === 'testimonials'}
                        onClick={this.handleItemClick}
                    >
                        Game

                    </Menu.Item>


                <Menu.Item name='sign-in' active={activeItem === 'sign-in'} onClick={this.handleItemClick}>
                    Sign-in
                </Menu.Item>
            </Menu>
        )
    }
}