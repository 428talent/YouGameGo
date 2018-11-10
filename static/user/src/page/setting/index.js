import {Divider, Image, Segment} from "semantic-ui-react";
import '../page.css'
import React from "react";
import {connect} from "dva";
import DropAvatarDialog from "./components/dropavatardialog";
import PropTypes from 'prop-types';
import {Button, Icon, Upload} from "antd";
import {WebServer} from "../../config/api";
import Cookies from "js-cookie";

const SettingPage = ({profile, user}) => {
    console.log(profile);
    console.log(user);
    const style = {
        content: {
            textAlign: "left"
        }
    };
    const uploadProps = {
        name: 'avatar',
        action: "",
        headers: {
            "Authorization": Cookies.get("yougame_token"),
        },
        onChange(info) {
            if (info.file.status !== 'uploading') {
                console.log(info.file, info.fileList);
            }
            if (info.file.status === 'done') {
                console.log(`${info.file.name} file uploaded successfully`);
            } else if (info.file.status === 'error') {
                console.log(`${info.file.name} file upload failed.`);
            }
        },
    };
    if (!user){
        uploadProps.action = `${WebServer}/user/${user}/avatar`
    }
    const onClickSave = () => {
        if (this.editor) {
            // This returns a HTMLCanvasElement, it can be made into a data URL or a blob,
            // drawn on another canvas, or added to the DOM.
            const canvas = this.editor.getImage();

            // If you want the image resized to the canvas size (also a HTMLCanvasElement)
            const canvasScaled = this.editor.getImageScaledToCanvas()
        }
    };
    let avatarEditor = null;
    const setEditorRef = (editor) => avatarEditor = editor;

    return (
        <div>
            <Segment className="page-container" style={style.content}>
                <h3>个人资料</h3>
                <Divider/>
                <h4>头像</h4>
                <Upload {...uploadProps}>
                    {function () {
                        if (user) {
                            return (
                                <div>
                                    <Image src={`${WebServer}/${user.profile.avatar}`} size='small' />
                                </div>
                            )
                        }

                    }()}
                    <Button style={{marginTop:6}}>
                        <Icon type="upload"/> 点击上传新头像
                    </Button>
                </Upload>,
                <DropAvatarDialog {...profile.avatar.dialog} />
            </Segment>
        </div>
    )
};
SettingPage.propTypes = {
    profile: PropTypes.object.isRequired
};
export default connect(({settingpage, app}) => ({...settingpage, ...app}))(SettingPage)