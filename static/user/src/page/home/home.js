import React from 'react';
import {Divider, Grid, Image, Segment} from "semantic-ui-react";
import '../page.css'
import GameCollection from "../components/GameCollection";
import {connect} from "dva";
import PropTypes, {object} from "prop-types";

const HomePage = ({dispatch, ...props}) => {
    const {wishGameList} = props;
    return (
        <div>
            <Segment className="page-container">
                <GameCollection title="愿望单" colCount={6} gameList={[...wishGameList]} dispatch={dispatch}/>
                <Divider />
            </Segment>
        </div>
    )
};
HomePage.propTypes = {
    wishlist: PropTypes.array.isRequired
};
export default connect(({homepage}) => ({
    wishlist: homepage.wishlist,
    wishGameList: homepage.wishGameList
}))(HomePage)