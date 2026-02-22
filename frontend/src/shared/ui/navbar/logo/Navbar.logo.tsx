import clsx from "clsx";
import { FC, ReactNode } from "react";

interface I_NavbarLogoProps {
  className?: string;
  children?: ReactNode;
}

export const NavbarLogo: FC<I_NavbarLogoProps> = (props) => {
  return <div className={clsx(props.className)}>{props.children}</div>;
};
