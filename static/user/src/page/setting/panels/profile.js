import {Checkbox, Divider, Form, Image, Segment} from "semantic-ui-react";
import {message, Button, Icon, Upload} from "antd";
import {WebServer} from "../../../config/api";
import DropAvatarDialog from "../components/dropavatardialog";
import React from "react";
import Cookies from "js-cookie";
import * as PropTypes from "prop-types";
import ChangeProfileForm from "./forms/change-profile";

const ProfileSettingPanel =
    ({
         user, profile,
         onRefreshUserProfile,
         onChangeUserProfile,
     }) => {
        const style = {
            content: {
                textAlign: "center"
            }
        };
        const uploadProps = {
            name: 'avatar',
            action: "",
            headers: {
                "Authorization": Cookies.get("yougame_token"),
                'Access-Control-Allow-Headers': "*"
            },
            onChange(info) {
                if (info.file.status !== 'uploading') {
                    console.log(info.file, info.fileList);
                }
                if (info.file.status === 'done') {
                    console.log(`${info.file.name} file uploaded successfully`);
                    message.success('更改头像成功', 5);
                    onRefreshUserProfile()
                } else if (info.file.status === 'error') {
                    console.log(`${info.file.name} file upload failed.`);
                }
            },
            showUploadList: false
        };
        if (user != null) {
            uploadProps.action = `${WebServer}/api/user/${user.id}/avatar/upload`
        }
        const submit = values => {
            onChangeUserProfile(values)
        };
        return (
            <div>
                <Segment className="page-container" style={style.content}>
                    <h3 style={{textAlign: "left"}}>头像设置</h3>
                    <Divider/>
                    <Upload {...uploadProps}>
                        {function () {
                            if (user) {
                                return (
                                    <div>
                                        <Image src={`${WebServer}/${user.profile.avatar}`} size='small'/>
                                    </div>
                                )
                            }

                        }()}
                        <Button style={{marginTop: 6}}>
                            <Icon type="upload"/> 点击上传新头像
                        </Button>
                    </Upload>,
                    <DropAvatarDialog {...profile.avatar.dialog} />
                </Segment>
                <Segment className="page-container" style={style.content}>
                    <h3 style={{textAlign: "left"}}>个人信息</h3>
                    <Divider/>
                    {function () {
                        if (user) {
                            return (
                                <div style={{textAlign: "left"}}>
                                    <ChangeProfileForm onSubmit={submit}/>
                                </div>
                            )
                        }
                    }()}
                </Segment>

            </div>
        )
    };
ProfileSettingPanel.propTypes = {
    user: PropTypes.object.isRequired,
    profile: PropTypes.object.isRequired,
    onRefreshUserProfile: PropTypes.func.isRequired,
    onChangeUserProfile: PropTypes.func.isRequired,
    onInputProfileTextChange: PropTypes.func.isRequired
};
export default ProfileSettingPanel