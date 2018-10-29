import {Button, Divider, Segment, Statistic} from "semantic-ui-react";
import GameCollection from "../components/GameCollection";
import PropTypes from "prop-types";
import {connect} from "dva";
import React from "react";
import CartGroup from "../components/CartGroup";

const CartPage = ({dispatch, ...props}) => {
    const {wishGameList} = props;
    return (
        <div>
            <Segment className="page-container" style={{textAlign: "left"}}>
                <h3>购物车</h3>
                <CartGroup/>
                <Divider/>
                <div style={{textAlign: "right", paddingRight: 20}}>
                    <div>
                        <div>
                            <Statistic size="mini" label='合计' value='￥5,550'/>
                        </div>
                        <div style={{marginTop:8}}>
                            <Button primary>付款</Button>
                        </div>
                    </div>
                </div>
            </Segment>
        </div>
    )
};
CartPage.propTypes = {
    wishlist: PropTypes.array.isRequired
};
export default connect(({cartpage}) => ({}))(CartPage)