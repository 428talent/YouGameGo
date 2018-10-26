import {Image} from "semantic-ui-react";
import React from "react";
import {ServerUrl} from "../../config/api";

const GameItem = ({dispatch, game}) => {
    return (
        <div>
            <Image src={`${ServerUrl}/${game.Band.Path}`}
                   size='small'/>
        </div>
    )
};
export default GameItem