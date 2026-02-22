import { T_NavbarState } from "../types/navbar-state.type";
import { E_NavbarAction } from "../types/navbar-action.enum";
import { T_NavbarAction } from "../types/navbar-action.type";
import { E_Navbar } from "../types/navbar.enum";

export const navbarReducer = (
  navbarState: T_NavbarState,
  action: T_NavbarAction,
) => {
  switch (action.type) {
    case E_NavbarAction.Toggle: {
      console.log(navbarState);

      return {
        ...navbarState,
        state:
          navbarState.state === E_Navbar.Open ? E_Navbar.Close : E_Navbar.Open,
      };
    }
    default: {
      return navbarState;
    }
  }
};
