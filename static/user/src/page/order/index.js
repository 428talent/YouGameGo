import {Button, Divider, Grid, Icon, Item, Label, Segment, Statistic} from "semantic-ui-react";
import GameCollection from "../components/GameCollection";
import '../page.css'
import React from "react";
import {connect} from "dva";
import FilterGroup from "../../layout/components/filter/group";
import PropTypes from "prop-types";

const OrderPage = ({filters, dispatch}) => {
    const style = {
        content: {
            textAlign: "left"
        }
    };
    let onFilterClick = (name, active) => {
        dispatch({
            type: 'orderpage/setFilter',
            name: name,
            active: !active

        })
    };
    return (
        <div>
            <Grid columns='equal' style={style.content}>
                <Grid.Column width={12}>
                    <Segment className="page-container">
                        <Segment>
                            <h4>订单号:23333</h4>
                            <Item.Group divided>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                            </Item.Group>
                            <Divider/>
                            <div style={{textAlign: "right", marginRight: 25}}>
                                <div style={{textAlign: "left"}}>
                                    <Statistic size='mini' label='合计' value='￥33'/>
                                    <div>
                                        <Button content='付款' primary/>
                                    </div>
                                </div>
                            </div>
                        </Segment>
                        <Segment>
                            <Item.Group divided>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                                <Item>
                                    <Item.Image
                                        src='https://cdn.steamstatic.com.8686c.com/steam/apps/391220/header.jpg?t=1540399490'/>

                                    <Item.Content>
                                        <Item.Header as='a'>12 Years a Slave</Item.Header>
                                        <Item.Meta>
                                            <span className='cinema'>Union Square 14</span>
                                        </Item.Meta>
                                        <Item.Description>Description</Item.Description>
                                        <Item.Extra>
                                            <Label>IMAX</Label>
                                            <Label icon='globe' content='Additional Languages'/>
                                        </Item.Extra>
                                    </Item.Content>
                                </Item>
                            </Item.Group>
                        </Segment>
                    </Segment>
                </Grid.Column>
                <Grid.Column width={4}>
                    <Segment>
                        <h4>过滤器</h4>
                        <FilterGroup filters={filters} onItemClick={(active, name) => onFilterClick(name, active)}
                        />
                    </Segment>
                </Grid.Column>
            </Grid>

        </div>
    )
};
OrderPage.propTypes = {
    filters: PropTypes.array.isRequired
};
export default connect(({orderpage}) => ({filters: orderpage.filters}))(OrderPage)