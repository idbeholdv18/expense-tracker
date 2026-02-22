import { useNavbar } from "../hooks/useNavbar";
import clsx from "clsx";
import { createPortal } from "react-dom";
import { MobileMenuSection } from "../mobile-menu/MobileMenuSection";
import { Link } from "./types";
import { E_Navbar } from "../types/navbar.enum";
import { FC } from "react";

interface I_NavbarMobileMenuProps {
  className?: string;
}
export const NavbarMobileMenu: FC<I_NavbarMobileMenuProps> = (props) => {
  const { state } = useNavbar();
  return (
    state === E_Navbar.Open &&
    createPortal(
      <div
        className={clsx(
          "flex bg-amber-200 gap-2 fixed top-0 left-0 w-full h-full pt-16 px-2 z-40",
          props.className,
        )}
      >
        <h1>test</h1>
      </div>,
      document.body,
    )
  );
};
