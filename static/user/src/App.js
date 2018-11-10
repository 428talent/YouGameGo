import React from 'react';
import './App.css';
import MainNavBar from "./layout/components/navbar";
import 'semantic-ui-css/semantic.min.css';
import UserCard from "./layout/components/usercard";
import {Container} from "semantic-ui-react";
import {Route} from "react-router-dom";
import HomePage from "./page/home/home";
import SettingPage from "./page/setting/index";
import GamePage from "./page/game";
import OrderPage from "./page/order/index";
import CartPage from "./page/cart/cart";
import {connect} from "dva";
import PropTypes from "prop-types";
import HashRouter from "react-router-dom/es/HashRouter";
import LoadingModal from "./layout/components/LoadingModal";
import "antd/dist/antd.css";
const App = ({history, ...props}) => {
    let {user,isLoadingModalShow} = props;
    console.log(history)
    return (
        <div className="App">
            <HashRouter history={history}>
                <div>
                    <MainNavBar/>
                    <Container>
                        <UserCard user={user} history={history}/>
                        <div>
                            <Route exact path="/" component={HomePage}/>
                            <Route path="/games" component={GamePage}/>
                            <Route path="/orders" component={OrderPage}/>
                            <Route path="/cart" component={CartPage}/>
                            <Route path="/setting" component={SettingPage}/>
                        </div>
                    </Container>
                    <LoadingModal open={isLoadingModalShow}/>
                </div>
            </HashRouter>
        </div>
    );
};

App.propTypes = {
    user: PropTypes.object,
    isLoadingModalShow:PropTypes.bool.isRequired
};
export default connect(({app}) => ({
    user: app.user,
    isLoadingModalShow:app.isLoadingModalShow
}))(App);
