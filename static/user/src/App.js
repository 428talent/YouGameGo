import React, {Component} from 'react';
import './App.css';
import MainNavBar from "./layout/components/navbar";
import 'semantic-ui-css/semantic.min.css';
import UserCard from "./layout/components/usercard";
import {Container} from "semantic-ui-react";
import {BrowserRouter as Router, Route} from "react-router-dom";
import HomePage from "./page/home";
import GamePage from "./page/game";
import PageNav from "./layout/components/pagenav/pagenav";
import {connect} from "dva";
import PropTypes from "prop-types";

const App = ({history, ...props}) => {
    let {user} = props;
    return (
        <div className="App">
            <Router history={history}>
                <div>
                    <MainNavBar/>
                    <Container>
                        <UserCard user={user}/>
                        <div>
                            <Route exact path="/" component={HomePage}/>
                            <Route path="/games" component={GamePage}/>
                        </div>
                    </Container>
                </div>
            </Router>
        </div>
    );
};

App.propTypes = {
    user: PropTypes.object.isRequired
};
export default connect(({app}) => ({
    user: app.user
}))(App);
