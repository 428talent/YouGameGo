import {Component} from "react";
import React from "react";
import {Dimmer, Divider, Grid, Header, Image, Loader, Segment} from "semantic-ui-react";
import PageNav from "./pagenav/pagenav";
import {WebServer} from "../../config/api";

const UserCard = ({user,history}) => {
    const style = {
        seg: {
            minHeight: 100
        },
        loader: {
            minHeight: 100
        },
        username: {
            fontSize: 14,
            marginTop: 12
        },
        nickname: {
            fontSize: 24
        }
    };
    if (user === null) {
        return (
            <Segment textAlign='left' style={style.seg}>
                <Dimmer active style={style.loader}>
                    <Loader/>
                </Dimmer>
            </Segment>
        )
    }
    return (

        <Segment textAlign='left' style={style.seg}>
            <div>
                <Grid columns='equal'>
                    <Grid.Row columns={2}>
                        <Grid.Column width={2}>
                            <Image src={`${WebServer}/${user.profile.avatar}`} size='tiny' rounded/>
                        </Grid.Column>
                        <Grid.Column>
                            <div style={style.nickname}>{user.profile.nickname}</div>
                            <div style={style.username}>{user.username}</div>
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
                <Divider/>
                <PageNav history={history}/>
            </div>
        </Segment>

    )

};
export default UserCard