import {Image, Item} from "semantic-ui-react";
import React from "react";

const CartItem = () => {
    return (
        <Item>
            <Item.Image size='small' src='https://cdn.steamstatic.com.8686c.com/steam/apps/931500/header_schinese.jpg?t=1540732300'/>

            <Item.Content>
                <Item.Header as='a'>Header</Item.Header>
                <Item.Meta>Description</Item.Meta>
                <Item.Description>

                </Item.Description>
                <Item.Extra>Additional Details</Item.Extra>
            </Item.Content>
        </Item>
    )
}

export default CartItem