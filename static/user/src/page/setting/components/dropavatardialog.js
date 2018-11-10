import {Button, Modal} from "semantic-ui-react";
import React from 'react';
import AvatarEditor from "react-avatar-editor";
import {Slider} from "antd";
import PropTypes from 'prop-types';

const DropAvatarDialog = ({isShow,scale}) =>{
    const style = {
        content:{
            textAlign:"center"
        }
    };
  return (
      <div>
      <Modal
          size="tiny"
          open={isShow}
      >
          <Modal.Header>更改头像</Modal.Header>
          <Modal.Content style={style.content}>
              <AvatarEditor
                  image="http://localhost:8888/static/upload/user/avatar/c3a00efffa92d01ba3e518434140729e.jpg"
                  width={250}
                  height={250}
                  border={50}
                  color={[0, 0, 0, 0.6]} // RGBA
                  scale={scale}
                  rotate={0}
              />
              <Slider defaultValue={30} max={100} min={1}/>
          </Modal.Content>
          <Modal.Actions>
              <Button  negative>
                  取消
              </Button>
              <Button

                  positive
                  labelPosition='right'
                  icon='checkmark'
                  content='确认'
              />
          </Modal.Actions>
      </Modal>
      </div>
  )
};
DropAvatarDialog.propTypes={
    scale:PropTypes.number.isRequired,
    isShow:PropTypes.bool.isRequired
};
DropAvatarDialog.defaultProps = {
    isShow:false,
};

export default DropAvatarDialog