import React from "react";
import {Label} from "semantic-ui-react";

const FilterTag = ({onFilterClick, active, name, ...props}) => {
    return (
        <Label {...props} color={active ? "blue" : undefined} onClick={() => onFilterClick(!!active, name)} as='a'>
            {name}
        </Label>
    )
};

export default FilterTag