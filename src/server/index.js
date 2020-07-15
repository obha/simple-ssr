import React from 'react';
import ReactDOMServer from 'react-dom/server';

/**
 * @returns {string} - html
 */
export function render(){
    const htmlContent = ReactDOMServer.renderToString(<div>HELLO</div>);
  
    return htmlContent
}