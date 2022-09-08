const fs = require("fs");
const path = require("path");

const DIR = './frontend';

module.exports = {
  mode: 'production',
  entry: () => new Promise(resolve => {
    fs.readdir(DIR, (err, files) => {
      let res = {}
      for (const file of files) {
        if (!file.match(/\.js$/)) continue;
        res[path.basename(file, '.js')] = `${DIR}/${file}`;
      }
      resolve(res)
    });
  }),
  output: {
    filename: '[name].js',
    path: __dirname + '/public/assets/pages',
  },
  externals: {
    'react': 'React',
    'react-dom': 'ReactDOM'
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        loader: "babel-loader",
        exclude: /node_modules/,
        options: {
          presets: ["@babel/preset-react"],
        },
      },
    ],
  }
};

