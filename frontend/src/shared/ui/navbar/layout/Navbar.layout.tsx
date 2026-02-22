import clsx from "clsx";
import { FC, ReactNode } from "react";

interface I_NavbarLayoutProps {
  children: ReactNode;
  className?: string;
}

export const NavbarLayout: FC<I_NavbarLayoutProps> = (props) => {
  return (
    <div className={clsx("flex w-full z-50 fixed", props.className)}>
      {props.children}
    </div>
  );
};
