import { NavbarContext } from "../context/Navbar.context";
import { FC, ReactNode, useEffect, useReducer } from "react";
import { navbarReducer } from "../reducer/Navbar.reducer";
import { DispatchNavbarContext } from "../context/DispatchNavbar.context";
import { E_Navbar } from "../types/navbar.enum";

interface I_NavbarProviderProps {
  defaultState?: E_Navbar;
  children?: ReactNode;
}

export const NavbarProvider: FC<I_NavbarProviderProps> = (props) => {
  const [state, dispatch] = useReducer(navbarReducer, {
    state: props.defaultState || E_Navbar.Close,
  });

  useEffect(() => {
    document.body.style.overflowY =
      state.state === E_Navbar.Open ? "hidden" : "auto";
  }, [state.state]);

  return (
    <NavbarContext.Provider value={state}>
      <DispatchNavbarContext.Provider value={dispatch}>
        {props.children}
      </DispatchNavbarContext.Provider>
    </NavbarContext.Provider>
  );
};
