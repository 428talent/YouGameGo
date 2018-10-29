import {Image, Item} from "semantic-ui-react";
import React from "react";
import CartItem from "./CartItem";

const CartGroup = ({cartItems}) => {
    const content = cartItems.map((cartItem, idx) => {
        return (
            <CartItem cartItem={cartItem} key={idx}/>
        )
    });
    return (
        <Item.Group>
            {content}
        </Item.Group>
    )
}
export default CartGroup