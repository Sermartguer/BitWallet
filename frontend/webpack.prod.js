const path = require('path');
const optimize = require('webpack').optimize;
const webpack = require('webpack');
const ExtractTextPlugin = require("extract-text-webpack-plugin");
const HtmlWebpackPlugin = require('html-webpack-plugin');

          
module.exports = {
    context: path.resolve('src'),
    entry: {
        bundle: path.resolve(__dirname, 'src') + './app/index.js',
        vendor: ['react', 'react-dom', 'react-router-dom']
    },
    output: {
        path: path.resolve(__dirname, 'dist') + '/app',
        filename: '[name].[hash].js',
        sourceMapFilename: '[name].map'
    },
    module: {
        rules: [
            {
                test: /.js[x]?$/,
                include: path.resolve(__dirname, 'src'),
                exclude: path.resolve(__dirname, 'node_modules'),
                use: 'babel-loader',
            },
            {
                test: /\.css$/,
                use: ['style-loader', 'css-loader'],
                include: /flexboxgrid/
              },
              {
                test: /\.scss$/,
                use: [{
                    loader: "style-loader" // creates style nodes from JS strings
                }, {
                    loader: "css-loader" // translates CSS into CommonJS
                }, {
                    loader: "sass-loader" // compiles Sass to CSS
                }]
            },
            {
                test: /\.(eot|woff|woff2|ttf|svg|ico|png|jpe?g|gif)$/,
                use: ['file-loader?name=[name].[ext]&outputPath=app/assets/images/',
                      'image-webpack-loader']
            },
            {
                test: /\.(gif|png|jpe?g|svg)$/i,
                use: [
                  'file-loader',
                  {
                    loader: 'image-webpack-loader',
                    options: {
                      bypassOnDebug: true,
                    },
                  },
                ],
              }
        ]
    },
    plugins: [
        /*new optimize.UglifyJsPlugin(),
        new optimize.CommonsChunkPlugin({
            names: ['vendor', 'manifest']
        }),*/
        new HtmlWebpackPlugin({
            template: './src/index.html',
            hash: true
        }),
        new webpack.NamedModulesPlugin(),
        new ExtractTextPlugin('app.css')
    ]
};