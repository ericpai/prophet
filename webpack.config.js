const path = require('path')
const webpack = require('webpack')
const UglifyJsPlugin = require('uglifyjs-webpack-plugin')

module.exports = {
    entry: path.resolve(__dirname, 'views/app.jsx'),
    output: {
        path: path.resolve(__dirname, 'static/js'),
        filename: 'prophet.js',
    },
    module: {
        loaders: [{
            test: /\.js|jsx$/,
            loader: 'babel-loader',
            exclude: /node_modules/,
            query: {
                presets: ['react', 'env']
            }
        }, {
            test: /\.js|jsx$/,
            loader: 'eslint-loader',
            exclude: /node_modules/
        }, {
            test: /\.js|jsx$/,
            loader: 'ify-loader'
        }, {
            test: /\.css$/,
            loader: 'style!css'
        }, {
            test: /\.json$/,
            loader: 'json-loader'
        }, {
            test: /\.less$/,
            loader: 'style!css!less'
        }, {
            test: /\.(png|jpg)$/,
            loader: 'url?limit=25000'
        }]
    },
    resolve: {
        modules: [path.join(__dirname, 'node_modules')],
        extensions: ['.js', '.jsx']
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: JSON.stringify('dev')
            }
        }),
    ]
};
