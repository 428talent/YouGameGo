import {Dimmer, Loader, Modal} from "semantic-ui-react";
import React from "react";

const LoadingModal = ({open, content}) => {
    return (
        <Modal size="tiny" open={open} style={{height: 300}}>
            <Modal.Content>
                <Dimmer active inverted>
                    <Loader inverted content={content}/>
                </Dimmer>
            </Modal.Content>
        </Modal>
    )
}
LoadingModal.defaultProps = {
    content: '处理中',
};
export default LoadingModal