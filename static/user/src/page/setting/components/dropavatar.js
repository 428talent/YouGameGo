import React from "react";
import Dropzone from "react-dropzone";
import {Image} from "semantic-ui-react";
import AvatarEditor from "react-avatar-editor";

const DropAvatar = () => {
    const style = {
        avatarContainer:{
            textAlign:"center",
            height:"100%"
        }
    };
    return (
        <div>
            <Dropzone>
                <div style={style.avatarContainer}>
                    <Image src='https://randomuser.me/api/portraits/men/74.jpg' size='small' centered />
                </div>
            </Dropzone>
            <AvatarEditor
                image="https://randomuser.me/api/portraits/men/74.jpg"
                width={250}
                height={250}
                border={50}
                color={[255, 255, 255, 0.6]} // RGBA
                scale={1.2}
                rotate={0}
            />
            
        </div>
    )
}
export default DropAvatar