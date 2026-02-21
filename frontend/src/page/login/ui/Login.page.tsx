import clsx from "clsx";
import { ChangeEvent, FC, useId, useState } from "react";

export interface LoginPageProps {}

export const LoginPage: FC<LoginPageProps> = () => {
  const loginId = useId();
  const passwordId = useId();

  const login = async (formData: FormData) => {
    const login = formData.get("login");
    const password = formData.get("password");

    try {
      const response = await fetch("https://localhost:8080/api/v1/login", {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          login,
          password,
        }),
      });
      if (!response.ok) {
        throw new Error("login failed");
      }
      const data = await response.json();

      console.log(data);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className={clsx("px-4 pt-8")}>
      <form
        action={login}
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
        </label>

        <input
          type='submit'
          value='Sign In'
          className='h-9 px-4 mt-4 bg-neutral-900 text-white rounded-md'
        />
      </form>
    </div>
  );
};
