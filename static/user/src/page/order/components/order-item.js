import {Item, Label} from "semantic-ui-react";
import React from "react";

const OrderItem = ({good,...props}) => {
    return (
        <Item>
            <Item.Image
                src={good.band_pic}/>

            <Item.Content>
                <Item.Header as='a'>{good.name}</Item.Header>
                <Item.Meta>
                    <span className='cinema'>{good.good_name}</span>
                </Item.Meta>
                <Item.Description>￥{good.price}</Item.Description>

            </Item.Content>
        </Item>
    )
};
export default OrderItem