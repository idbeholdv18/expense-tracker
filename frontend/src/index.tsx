import { createRoot } from "react-dom/client";
import React from "react";
import "./app/ui/global.css";

document.body.innerHTML = '<div id="app"></div>';

const root = createRoot(document.getElementById("app"));
root.render(<h1 className='bg-red-200'>Hello, world</h1>);
