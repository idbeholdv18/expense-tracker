/// <reference path="node_modules/webpack-dev-server/types/lib/Server.d.ts"/>
import { type Configuration } from "webpack";
import path, { dirname } from "path";
import HtmlWebpackPlugin from "html-webpack-plugin";
import { fileURLToPath } from "url";
import { readFileSync } from "fs";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

export default (): Configuration => {
  return {
    mode: "development",
    entry: path.resolve(__dirname, "src", "index.tsx"),
    output: {
      filename: "[name].[contenthash].js",
      path: path.resolve(__dirname, "dist"),
      clean: true,
    },
    module: {
      rules: [
        {
          test: /\.tsx?$/,
          use: "ts-loader",
          exclude: /node_modules/,
        },
        {
          test: /\.css$/i,
          use: ["style-loader", "css-loader", "postcss-loader"],
        },
      ],
    },
    plugins: [
      new HtmlWebpackPlugin({
        template: path.resolve(__dirname, "public", "index.html"),
      }),
    ],
    resolve: {
      alias: {
        "@": path.resolve(__dirname, "src"),
      },
      extensions: [".tsx", ".ts", ".js", ".jsx"],
    },
    devServer: {
      port: 3000,
      server: {
        type: "https",
        options: {
          key: readFileSync(path.resolve(__dirname, "../cert/key.pem")),
          cert: readFileSync(path.resolve(__dirname, "../cert/cert.pem")),
        },
      },
      hot: true,
      open: true,
      historyApiFallback: true,
    },
  };
};
