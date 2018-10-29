import {Image, Item} from "semantic-ui-react";
import React from "react";
import CartItem from "./CartItem";
const CartGroup = () =>{
    return (
        <Item.Group>
            <CartItem/>
            <CartItem/>
            <CartItem/>
            <CartItem/>
        </Item.Group>
    )
}
export default CartGroup