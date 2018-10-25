import {Grid, Header, Image} from "semantic-ui-react";
import React from "react";

const GameCollection = ({}) => {
    const style = {
        root: {
            textAlign: "left"
        }
    };
    return (
        <div style={style.root}>
            <Header as='h2'>Third Header</Header>
            <Grid>
                <Grid.Row columns={3}>
                    <Grid.Column>
                        <Image src='https://cdn.steamstatic.com.8686c.com/steam/apps/359550/header.jpg?t=1538561874'
                               size='medium'/>
                    </Grid.Column>
                    <Grid.Column>
                        <Image src='https://cdn.steamstatic.com.8686c.com/steam/apps/359550/header.jpg?t=1538561874'
                               size='medium'/>
                    </Grid.Column>
                    <Grid.Column>
                        <Image src='https://cdn.steamstatic.com.8686c.com/steam/apps/359550/header.jpg?t=1538561874'
                               size='medium'/>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </div>
    )
};
export default GameCollection