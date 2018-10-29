import {Image, Item} from "semantic-ui-react";
import React from "react";

const CartItem = ({cartItem}) => {
    console.log(cartItem)
    return (
        <Item>
            <Item.Image size='small' src={cartItem.game.band}/>

            <Item.Content>
                <Item.Header as='a'>{cartItem.game.name}</Item.Header>
                <Item.Meta>{cartItem.good.name}</Item.Meta>
                <Item.Description>

                </Item.Description>
                <Item.Extra>￥{cartItem.good.price}</Item.Extra>
            </Item.Content>
        </Item>
    )
}

export default CartItem