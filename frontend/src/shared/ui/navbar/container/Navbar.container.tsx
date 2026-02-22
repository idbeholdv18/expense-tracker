import clsx from "clsx";
import { FC, ReactNode } from "react";

interface I_NavbarContainerProps {
  className?: string;
  children?: ReactNode;
}

export const NavbarContainer: FC<I_NavbarContainerProps> = (props) => {
  return <div className={clsx(props.className)}>{props.children}</div>;
};
