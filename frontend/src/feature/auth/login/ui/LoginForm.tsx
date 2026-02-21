import { zodResolver } from "@hookform/resolvers/zod";
import clsx from "clsx";
import { FC, useId } from "react";
import { useForm } from "react-hook-form";
import { loginApi } from "../model/login.action";
import { LoginInput, LoginSchema } from "../model/login.schema";
import { SubmitButton } from "./SubmitButton";

export interface LoginFormProps {}

export const LoginForm: FC<LoginFormProps> = () => {
  const loginId = useId();
  const passwordId = useId();

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting, isValid },
    setError,
    clearErrors,
  } = useForm<LoginInput>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      login: "testuser",
      password: "123456",
    },
    mode: "onChange",
  });

  const onSubmit = handleSubmit(async (data) => {
    try {
      await loginApi(data);
      console.log("Login success");
    } catch (err: any) {
      setError("root", {
        type: "server",
        message: err.message,
      });
    }
  });

  return (
    <form
      onSubmit={onSubmit}
      className='flex flex-col gap-4 px-4 py-4 bg-neutral-100 rounded-b-md outline outline-neutral-200'
    >
      <label htmlFor={loginId}>login</label>
      <input
        {...register("login", {
          onChange: () => clearErrors("root"),
        })}
        id={loginId}
        className={clsx(
          "border border-solid border-neutral-200 w-full bg-white",
        )}
      />
      {errors.login && <p className='text-red-500'>{errors.login.message}</p>}

      <label htmlFor={passwordId}>password</label>
      <input
        {...register("password", {
          onChange: () => clearErrors("root"),
        })}
        type='password'
        id={passwordId}
        className={clsx(
          "border border-solid border-neutral-200 w-full bg-white",
        )}
      />
      {errors.password && (
        <p className='text-red-500'>{errors.password.message}</p>
      )}

      <SubmitButton isSubmitting={isSubmitting} isReady={isValid} />

      {errors.root && <p className='text-red-500'>{errors.root.message}</p>}
    </form>
  );
};
