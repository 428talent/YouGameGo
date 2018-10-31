import {Button, Divider, Icon, Modal, Item, Statistic, TransitionablePortal} from "semantic-ui-react";
import React from "react";
import OrderItem from "./order-item";
import * as PropTypes from "prop-types";

const OrderModal = ({order, open, isPaying, onCloseModal}) => {
    if (!open) {
        return (
            <div/>
        )
    }
    let totalPrice = 0.0;
    const orderItems = order.goods.map(good => {
        totalPrice += good.price;
        return (
            <OrderItem good={good} key={good.id}/>
        )
    });
    const closeModal = () =>{
        if (!isPaying){
            onCloseModal()
        }
    }

    return (

            <Modal open={open} onClose={() => closeModal()}>
                <Modal.Header>订单</Modal.Header>
                <Modal.Content scrolling>
                    <Modal.Description>
                        <h4>订单号:{order.id}</h4>
                        <Item.Group divided>
                            {orderItems}
                        </Item.Group>
                        <Divider/>
                        <div style={{textAlign: "right", marginRight: 25}}>
                            <div style={{textAlign: "left"}}>
                                <Statistic size='mini' label='合计' value={`￥${totalPrice}`}/>
                            </div>
                        </div>
                    </Modal.Description>
                </Modal.Content>
                <Modal.Actions>
                    <Button primary loading={isPaying} disabled={isPaying}>
                        确认
                    </Button>
                </Modal.Actions>
            </Modal>

    )
};
OrderModal.propTypes = {
    order: PropTypes.object.isRequired,
    open: PropTypes.bool,
    isPaying: PropTypes.bool,
    onCloseModal: PropTypes.func
};
OrderModal.defaultProps = {
    open: false,
    isPaying: false,
    onCloseModal: () => {
    }
};
export default OrderModal