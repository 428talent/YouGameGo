import React from 'react'
import {Field, reduxForm} from 'redux-form'
import {Button, Form} from "semantic-ui-react";

let ChangeProfileForm = props => {
    const {handleSubmit} = props;
    return <form onSubmit={handleSubmit}>
        <Form>
            <Form.Field>
                <label htmlFor="nickname">昵称</label>
                <Field name="nickname" type="text" component="input"/>
            </Form.Field>
            <Button type="submit">保存</Button>
        </Form>
    </form>
};

ChangeProfileForm = reduxForm({
    form: 'changeProfileForm'
})(ChangeProfileForm);

export default ChangeProfileForm