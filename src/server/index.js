import React from 'react';
import ReactDOMServer from 'react-dom/server';
import { StaticRouter } from "react-router-dom";
import App from '../App';

/**
 * @param location {string}
 * @returns {string} - html
 */
export function render(location){

    const html = ReactDOMServer.renderToString(
        <StaticRouter location={location} context={{}}>
          <App/>
        </StaticRouter>
    );
  
    return html
}

