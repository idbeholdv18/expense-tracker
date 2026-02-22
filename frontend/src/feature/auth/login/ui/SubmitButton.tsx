import clsx from "clsx";
import { FC, memo } from "react";

export interface SubmitButtonProps {
  isSubmitting: boolean;
  isReady: boolean;
  className?: string;
}

export const SubmitButton = memo((props: SubmitButtonProps) => {
  return (
    <button
      type='submit'
      disabled={props.isSubmitting || !props.isReady}
      className={clsx(
        "h-9 px-4 rounded-md",
        (props.isSubmitting || !props.isReady) &&
          "bg-disabled-pattern text-fg-secondary",
        props.className,
      )}
    >
      {props.isSubmitting ? "Submitting..." : "Login"}
    </button>
  );
});
