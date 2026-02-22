import { E_NavbarAction } from "./navbar-action.enum";
import { E_Navbar } from "./navbar.enum";

export type T_NavbarPayload = {
  [E_NavbarAction.Toggle]: undefined,
  [E_NavbarAction.Set]: E_Navbar
}