import { createRoot } from "react-dom/client";
import "./app/ui/global.css";
import { App } from "@/app/ui/App";

document.body.innerHTML = '<div id="app"></div>';

const root = createRoot(document.getElementById("app"));
root.render(<App />);
