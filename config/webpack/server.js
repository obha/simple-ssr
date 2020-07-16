const path = require('path');
module.exports = {
  entry: [
    "@babel/polyfill",
    './src/server'
  ],
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ['babel-loader']
      }
    ]
  },
  resolve: {
    extensions: ['*', '.js', '.jsx']
  },
  output: {
    path: path.resolve(__dirname, '..','..','public'),
    libraryTarget: "umd",
    library: "server",
    filename: 'server.js',
    globalObject: "this",
  },
  
};