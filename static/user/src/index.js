import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import dva from "dva";

const app = dva({
    initialState: {
        pagenav:{
            activeTab:"home"
        }
    },
});
app.router(({history}) => <App/>);
app.model(require('./layout/components/pagenav/model').default);
app.model(require('./model/app').default);
app.model(require('./page/home/model').default);

app.start('#root');
// ReactDOM.render(<App />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
