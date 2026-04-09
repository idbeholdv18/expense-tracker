import { FC, useId } from "react";

export interface LoginInputProps {
  loginId: string;
}

export const LoginInput: FC<LoginInputProps> = (props) => {
  const id = useId();

  return (
    <div className='flex flex-col w-full'>
      <label htmlFor={id} className='text-fg-primary'>
        login
      </label>
      <input
        {...register("login", {
          onChange: () => clearErrors("root"),
        })}
        id={id}
        className={clsx("border border-solid border-stroke-primary w-full p-1")}
      />
      {errors.login && <p className='text-red-500'>{errors.login.message}</p>}
    </div>
  );
};
