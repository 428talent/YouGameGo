import React from 'react';
import {Grid, Image, Segment} from "semantic-ui-react";
import './page.css'
import GameCollection from "./components/GameCollection";

const HomePage = () => {
    return (
        <div>
            <Segment className="page-container">
                <GameCollection/>
            </Segment>
        </div>
    )
}
export default HomePage