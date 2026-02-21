import clsx from "clsx";
import { FC, useActionState, useId } from "react";
import { useFormStatus } from "react-dom";
import { login } from "../model/login";
import { SubmitButton } from "./SubmitButton";

export interface LoginFormProps {}

export const LoginForm: FC<LoginFormProps> = () => {
  const loginId = useId();
  const passwordId = useId();

  const [state, formAction] = useActionState(
    (prevState: any, formData: FormData) => login(formData),
    null,
  );

  return (
    <form
      action={formAction}
      className='flex flex-col gap-4 px-4 py-4 bg-neutral-100 rounded-b-md outline outline-neutral-200'
    >
      <label htmlFor={loginId}>
        login
        <input
          defaultValue={"testuser"}
          required
          name='login'
          type='text'
          id={loginId}
          className={clsx(
            "border border-solid border-neutral-200 w-full bg-white",
          )}
        />
        {state?.errors?.login && (
          <span className='text-red-500'>{state.errors.login}</span>
        )}
      </label>

      <label htmlFor={passwordId}>
        password
        <input
          required
          defaultValue={"123456"}
          name='password'
          type='password'
          id={passwordId}
          className={clsx(
            "border border-solid border-neutral-200 w-full bg-white",
          )}
        />
        {state?.errors?.password && (
          <span className='text-red-500'>{state.errors.password}</span>
        )}
      </label>

      <SubmitButton />

      {state?.errors?.form && (
        <span className='text-red-500'>{state.errors.form}</span>
      )}
    </form>
  );
};
