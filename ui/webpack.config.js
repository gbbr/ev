var path = require('path');

module.exports = {
  entry: './src/index.js',
  devtool: 'inline-source-map',

  devServer: {
    contentBase: './dist'
  },

  resolve: {
    modules: ['src', 'node_modules'],
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules)/,
        use: {
          loader: 'babel-loader',
          options: {
              presets: ['env', 'react'],
              plugins: ['transform-object-rest-spread']
          }
        }
      },
      { 
        test: /\.css$/,
        use: [
          { loader: "style-loader" },
          { loader: "css-loader" }
        ]
      },
      {
        test: /\.(png|woff|jpg|svg|gif|eot|ttf)$/,
        exclude: /\/icon\/.*.svg/,
        loader: 'url-loader',
        options: { limit: 10000 },
      },
      {
        test: /\.scss$/,
        use: [
          { loader: "style-loader" },
          { loader: "css-loader" },
          { loader: "sass-loader" }
        ]
      }
    ]
  },

  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist'),
    publicPath: '/dist/'
  }
};
