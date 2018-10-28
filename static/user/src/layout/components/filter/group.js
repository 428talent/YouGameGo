import React from "react";
import {Label} from "semantic-ui-react";
import FilterTag from "./tag";

const FilterGroup = ({onItemClick,filters, ...props}) => {
    let content = filters.map((filter, idx) => {
        return (
            <FilterTag name={filter.name}  active={filter.active} key={idx} onFilterClick={(active, name) =>onItemClick(active,name)}/>
        )
    });
    return (
        <Label.Group>
            {content}
        </Label.Group>
    )
};

export default FilterGroup