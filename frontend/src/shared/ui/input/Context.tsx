import { createContext, useContext } from "react";

interface InputContextValue {
  id: string;
  error?: string;
}

export const InputContext = createContext<InputContextValue | null>(null);

export const useInputContext = () => {
  const ctx = useContext(InputContext);
  if (!ctx) {
    throw new Error("Input subcomponents must be used inside <Input>");
  }
  return ctx;
};
