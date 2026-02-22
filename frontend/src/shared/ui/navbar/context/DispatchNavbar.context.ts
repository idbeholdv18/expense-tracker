import { createContext, Dispatch } from "react";
import { T_NavbarAction } from "../types/navbar-action.type";

export const DispatchNavbarContext = createContext<Dispatch<T_NavbarAction>>(
  () => {},
);
