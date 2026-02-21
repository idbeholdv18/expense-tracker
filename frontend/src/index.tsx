import { createRoot } from "react-dom/client";
import "./app/ui/global.css";
import { App } from "@/app/ui/App";

document.body.innerHTML = '<div id="app"></div>';

const rootNode = document.getElementById("app");
if (!rootNode) {
  throw new Error('node with id="app" doesn\'t exist');
}

const root = createRoot(rootNode);
root.render(<App />);
