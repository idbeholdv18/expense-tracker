import { createContext } from "react";
import { T_NavbarState } from "../types/navbar-state.type";
import { E_Navbar } from "../types/navbar.enum";

export const NavbarContext = createContext<T_NavbarState>({
  state: E_Navbar.Close,
});
