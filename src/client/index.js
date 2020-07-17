import { loadableReady } from '@loadable/component';
import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import App from '../App';


loadableReady(() =>{
    ReactDOM.hydrate(
        <BrowserRouter>
            <App/>
        </BrowserRouter>,
        document.getElementById('root')
    )
})
