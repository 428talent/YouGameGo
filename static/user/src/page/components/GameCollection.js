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
    const content = gameGrid.map((col,cidx) => {
        const rowContent = col.map((game,ridx) => {
                return (
                    <Grid.Column key={cidx}>
                        <GameItem game={game} key={ridx * cidx} />
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