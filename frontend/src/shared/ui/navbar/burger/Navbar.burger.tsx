import clsx from "clsx";
import { useDispatchNavbar } from "../hooks/useDispatchNavbar";
import { FC } from "react";
import { E_NavbarAction } from "../types/navbar-action.enum";
import { useNavbar } from "../hooks/useNavbar";

interface I_NavbarBurgerProps {
  className?: string;
}

export const NavbarBurger: FC<I_NavbarBurgerProps> = (props) => {
  const dispatchNavbar = useDispatchNavbar();
  const { state } = useNavbar();

  const toggleNavbar = () => {
    dispatchNavbar({ type: E_NavbarAction.Toggle });
  };

  return (
    <button className={clsx(props.className)} onClick={toggleNavbar}>
      <svg
        xmlns='http://www.w3.org/2000/svg'
        width='24'
        height='24'
        viewBox='0 0 24 24'
        fill='none'
        stroke='currentColor'
        strokeWidth='3'
        strokeLinecap='round'
        strokeLinejoin='round'
      >
        <path d='M4 5h16' />
        <path d='M4 12h16' />
        <path d='M4 19h16' />
      </svg>
    </button>
  );
};
