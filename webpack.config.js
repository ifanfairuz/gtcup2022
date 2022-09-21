const fs = require("fs");
const path = require("path");

const v = "0-3-1";
const DIR = "./frontend";

module.exports = {
  mode: "production",
  entry: () =>
    new Promise((resolve) => {
      fs.readdir(DIR, (err, files) => {
        let res = {};
        for (const file of files) {
          if (!file.match(/\.js$/)) continue;
          res[path.basename(file, ".js") + `.${v}`] = `${DIR}/${file}`;
        }
        resolve(res);
      });
    }),
  resolve: {
    extensions: [".json", ".html", ".js", ".jsx"],
  },
  output: {
    filename: "[name].js",
    path: __dirname + "/public/assets/pages",
  },
  externals: {
    luxon: "luxon",
    react: "React",
    "react-dom": "ReactDOM",
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
  },
};
