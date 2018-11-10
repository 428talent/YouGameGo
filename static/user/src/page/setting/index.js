import {Tab} from "semantic-ui-react";
import '../page.css'
import React from "react";
import {connect} from "dva";
import PropTypes from 'prop-types';
import ProfileSettingPanel from "./panels/profile";

const SettingPage = ({profile, user, dispatch}) => {
    console.log(profile);
    console.log(user);
    const style = {
        content: {
            textAlign: "left"
        }
    };
    const refreshUserProfile = () => {
        dispatch({
            type: "app/refreshUserProfile"
        })
    };
    const changeUserProfile = ({nickname}) => {
        console.log(nickname)
        dispatch({
            type: "settingpage/changeUserProfile",
            payload: {
                nickname
            }
        })
    };

    const panes = [
        {
            menuItem: '个人资料',
            render: () => <ProfileSettingPanel onChangeUserProfile={changeUserProfile} onRefreshUserProfile={refreshUserProfile} {...{user, profile}}/>
        },
       
    ];

    return (
        <div>
            <Tab menu={{fluid: true, vertical: true}} menuPosition='left' panes={panes}/>
        </div>
    )
};
SettingPage.propTypes = {
    profile: PropTypes.object.isRequired
};
export default connect(({settingpage, app}) => ({...settingpage, ...app}))(SettingPage)