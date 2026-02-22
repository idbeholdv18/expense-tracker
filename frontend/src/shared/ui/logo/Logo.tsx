import clsx from "clsx";
import { FC } from "react";

export interface LogoProps {
  className?: string;
}

export const Logo: FC<LogoProps> = (props) => {
  return <div className={clsx("w-6 h-6 rounded-full", props.className)} />;
};
