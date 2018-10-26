import {Container, Grid, Header, Image} from "semantic-ui-react";
import React from "react";
import GameItem from "./GameItem";

const GameCollection = ({colCount,title, gameList, dispatch}) => {
    const style = {
        root: {
            textAlign: "left"
        }
    };
    let gameGrid = [];
    while (gameList.length) gameGrid.push(gameList.splice(0, colCount));
    console.log(gameGrid);
    const content = gameGrid.map(col => {
        const rowContent = col.map(game => {
                return (
                    <Grid.Column key={game.Id}>
                        <GameItem game={game}/>
                    </Grid.Column>
                )
            }
        );
        return (
            <Grid.Row columns={colCount}>
                {rowContent}
            </Grid.Row>
        )

    });
    return (
        <div style={style.root}>
            <Container>
                <Header as='h3'>{title}</Header>
                <Grid>
                    {content}
                </Grid>
            </Container>
        </div>
    )
};

export default GameCollection