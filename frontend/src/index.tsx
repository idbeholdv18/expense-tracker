import { router } from "@/shared/router/router";
import { createRoot } from "react-dom/client";
import { RouterProvider } from "react-router";
import "./app/ui/global.css";

document.body.innerHTML = '<div id="app"></div>';

const rootNode = document.getElementById("app");
if (!rootNode) {
  throw new Error('node with id="app" doesn\'t exist');
}

const root = createRoot(rootNode);
root.render(<RouterProvider router={router} />);
