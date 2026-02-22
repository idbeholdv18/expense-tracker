import { forwardRef, InputHTMLAttributes } from "react";
import { useInputContext } from "./Context";
import clsx from "clsx";

export const Field = forwardRef<
  HTMLInputElement,
  InputHTMLAttributes<HTMLInputElement>
>((props, ref) => {
  const { id, error } = useInputContext();

  return (
    <input
      ref={ref}
      id={id}
      className={clsx(
        "border border-stroke-primary w-full p-2 rounded-md bg-bg-accent",
        "focus:outline-none focus:ring focus:ring-stroke-accent focus:border-stroke-accent",
        error && "border-red-500",
      )}
      {...props}
    />
  );
});

Field.displayName = "Input.Field";
