import clsx from "clsx";
import { useInputContext } from "./Context";

export const Error = ({ className }: { className?: string }) => {
  const { error } = useInputContext();

  if (!error) return null;

  return <p className={clsx("", className)}>{error}</p>;
};
