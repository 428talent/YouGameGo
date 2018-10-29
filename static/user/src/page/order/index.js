import {Grid, Segment} from "semantic-ui-react";
import '../page.css'
import React from "react";
import {connect} from "dva";
import FilterGroup from "../../layout/components/filter/group";
import PropTypes from "prop-types";
import Order from "./components/order";

const OrderPage = ({filters,orders, dispatch,...props}) => {
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
    return (
        <div>
            <Grid columns='equal' style={style.content}>
                <Grid.Column width={12}>
                    <Segment className="page-container">
                        <Order orders={[...orders]}/>
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

        </div>
    )
};
OrderPage.propTypes = {
    filters: PropTypes.array.isRequired,
    orders: PropTypes.array.isRequired
};
export default connect(({orderpage}) => ({
    filters: orderpage.filters,
    orders: orderpage.orders
}))(OrderPage)