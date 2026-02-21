import { FC } from "react";
import { useFormStatus } from "react-dom";

export interface SubmitButtonProps {}

export const SubmitButton: FC<SubmitButtonProps> = () => {
  const { pending } = useFormStatus();

  return (
    <button
      type='submit'
      disabled={pending}
      className='h-9 px-4 mt-4 bg-neutral-900 text-white rounded-md'
    >
      {pending ? "Submitting..." : "Login"}
    </button>
  );
};
