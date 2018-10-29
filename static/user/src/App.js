import React from 'react';
import './App.css';
import MainNavBar from "./layout/components/navbar";
import 'semantic-ui-css/semantic.min.css';
import UserCard from "./layout/components/usercard";
import {Container} from "semantic-ui-react";
import {Route} from "react-router-dom";
import HomePage from "./page/home/home";
import GamePage from "./page/game";
import OrderPage from "./page/order/index";
import CartPage from "./page/cart/cart";
import {connect} from "dva";
import PropTypes from "prop-types";
import HashRouter from "react-router-dom/es/HashRouter";

const App = ({history, ...props}) => {
    let {user} = props;
    return (
        <div className="App">
            <HashRouter history={history}>
                <div>
                    <MainNavBar/>
                    <Container>
                        <UserCard user={user}/>
                        <div>
                            <Route exact path="/" component={HomePage}/>
                            <Route path="/games" component={GamePage}/>
                            <Route path="/orders" component={OrderPage}/>
                            <Route path="/cart" component={CartPage}/>
                        </div>
                    </Container>
                </div>
            </HashRouter>
        </div>
    );
};

App.propTypes = {
    user: PropTypes.object
};
export default connect(({app}) => ({
    user: app.user
}))(App);