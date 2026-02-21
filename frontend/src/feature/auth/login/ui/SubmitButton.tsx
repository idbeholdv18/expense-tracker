import { FC, memo } from "react";

export interface SubmitButtonProps {
  isSubmitting: boolean;
  isReady: boolean;
}

export const SubmitButton = memo((props: SubmitButtonProps) => {
  return (
    <button
      type='submit'
      disabled={props.isSubmitting || !props.isReady}
      className='h-9 px-4 mt-4 bg-neutral-900 text-white rounded-md'
    >
      {props.isSubmitting ? "Submitting..." : "Login"}
    </button>
  );
});
