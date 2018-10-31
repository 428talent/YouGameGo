import {Button, Dimmer, Grid, Loader, Modal, Segment} from "semantic-ui-react";
import '../page.css'
import React from "react";
import {connect} from "dva";
import FilterGroup from "../../layout/components/filter/group";
import PropTypes from "prop-types";
import Order from "./components/order";
import OrderModal from "./components/OrderModal";

const OrderPage = ({filters, orders, orderModal, dispatch, ...props}) => {
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
    const onPayOrder = (orderId) => {
        dispatch({
            type: 'orderpage/payOrder',
            payload: {
                orderId
            }
        })
    };
    const onPayButtonClick = (isShow, order) => {
        dispatch({
            type: 'orderpage/setOrderModel',
            isShow, order
        })
    };
    const closeOrderModal = () => {
        dispatch({
            type: 'orderpage/setOrderModel',
            isShow: false,
            order: {}
        })
    };
    return (
        <div>
            <Grid columns='equal' style={style.content}>
                <Grid.Column width={12}>
                    <Segment className="page-container">
                        <Order orders={[...orders]} onPayButtonClick={(order) => {
                            onPayButtonClick(true, order)
                        }}/>
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
            <OrderModal
                isPaying={orderModal.isPaying}
                onCloseModal={() => closeOrderModal()}
                open={orderModal.isShow}
                order={orderModal.order}/>
        </div>
    )
};
OrderPage.propTypes = {
    filters: PropTypes.array.isRequired,
    orders: PropTypes.array.isRequired,
    orderModal: PropTypes.object.isRequired
};
export default connect(({orderpage}) => ({
    filters: orderpage.filters,
    orders: orderpage.orders,
    orderModal: orderpage.orderModal
}))(OrderPage)