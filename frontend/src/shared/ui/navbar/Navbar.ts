import { NavbarLayout } from "./layout/Navbar.layout";
import { NavbarLogo } from "./logo/Navbar.logo";
import { NavbarBurger } from "./burger/Navbar.burger";
import { NavbarContainer } from "./container/Navbar.container";
import { NavbarProvider } from "./provider/Navbar.provider";
import { NavbarMobileMenu } from "./mobile-menu/Navbar.mobile-menu";

type T_Navbar = {
  Layout: typeof NavbarLayout;
  Logo: typeof NavbarLogo;
  Burger: typeof NavbarBurger;
  Container: typeof NavbarContainer;
  Provider: typeof NavbarProvider;
  MobileMenu: typeof NavbarMobileMenu;
};

export const Navbar: T_Navbar = {
  Layout: NavbarLayout,
  Logo: NavbarLogo,
  Burger: NavbarBurger,
  Container: NavbarContainer,
  Provider: NavbarProvider,
  MobileMenu: NavbarMobileMenu,
};
