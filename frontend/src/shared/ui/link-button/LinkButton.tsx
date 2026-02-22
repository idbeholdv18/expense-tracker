import clsx from "clsx";
import { FC } from "react";
import { Link, LinkProps } from "react-router";

export interface LinkButtonProps extends LinkProps {}

export const LinkButton: FC<LinkButtonProps> = ({
  className,
  children,
  ...props
}) => {
  return (
    <Link
      {...props}
      className={clsx(
        "px-4 h-9 flex items-center justify-center rounded-sm",
        className,
      )}
    >
      {children}
    </Link>
  );
};
