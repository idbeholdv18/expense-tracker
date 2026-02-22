import { ReactNode } from "react";
import { useInputContext } from "./Context";
import clsx from "clsx";

export const Label = ({
  children,
  className,
}: {
  children: ReactNode;
  className?: string;
}) => {
  const { id } = useInputContext();

  return (
    <label htmlFor={id} className={clsx("", className)}>
      {children}
    </label>
  );
};
