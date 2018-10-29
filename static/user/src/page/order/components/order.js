import {Button, Divider, Item, Label, Segment, Statistic} from "semantic-ui-react";
import React from "react";
import OrderItem from "./order-item";

const Order = ({orders, ...prop}) => {
    const content = orders.map((order => {
        let totalPrice = 0;
        const orderItems = order.goods.map(good => {
            totalPrice += good.price;
            return (
                <OrderItem good={good} key={good.id}/>
            )
        });
        return (
            <Segment key={order.id}>
                <h4>订单号:{order.id}</h4>
                <Item.Group divided>
                    {orderItems}
                </Item.Group>
                <Divider/>
                <div style={{textAlign: "right", marginRight: 25}}>
                    <div style={{textAlign: "left"}}>
                        <Statistic size='mini' label='合计' value={`￥${totalPrice}`}/>
                        <div>
                            <Button content='付款' primary/>
                        </div>
                    </div>
                </div>
            </Segment>
        )
    }));
    return (
        <div>
            {content}
        </div>
    )
};

export default Order