import clsx from "clsx";
import { ReactNode, useId } from "react";
import { InputContext } from "./Context";

interface InputRootProps {
  children: ReactNode;
  error?: string;
  id?: string;
  className?: string;
}

export const Root = ({ children, error, id, className }: InputRootProps) => {
  const generatedId = useId();
  const inputId = id ?? generatedId;

  return (
    <InputContext.Provider value={{ id: inputId, error }}>
      <div className={clsx("flex flex-col gap-1 w-full", className)}>
        {children}
      </div>
    </InputContext.Provider>
  );
};
