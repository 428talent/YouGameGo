import {Button, Divider, Segment, Statistic} from "semantic-ui-react";
import GameCollection from "../components/GameCollection";
import PropTypes from "prop-types";
import {connect} from "dva";
import React from "react";
import CartGroup from "../components/CartGroup";

const CartPage = ({dispatch,cartListItems,totalPrice,...props}) => {
    return (
        <div>
            <Segment className="page-container" style={{textAlign: "left"}}>
                <h3>购物车</h3>
                <CartGroup cartItems={[...cartListItems]} />
                <Divider/>
                <div style={{textAlign: "right", paddingRight: 20}}>
                    <div>
                        <div>
                            <Statistic size="mini" label='合计' value={`￥${totalPrice}`}/>
                        </div>
                        <div style={{marginTop: 8}}>
                            <Button primary>付款</Button>
                        </div>
                    </div>
                </div>
            </Segment>
        </div>
    )
};
CartPage.propTypes = {
    cartListItems: PropTypes.array.isRequired,
    totalPrice: PropTypes.number.isRequired
};
export default connect(({cartpage}) => ({
    cartListItems: cartpage.cartListItems,
    totalPrice :cartpage.totalPrice
}))(CartPage)