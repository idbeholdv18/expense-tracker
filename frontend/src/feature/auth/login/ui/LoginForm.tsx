import { Input } from "@/shared/ui/input/Input";
import { Logo } from "@/shared/ui/logo/Logo";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC, useId } from "react";
import { useForm } from "react-hook-form";
import { loginApi } from "../model/login.action";
import { LoginInput, LoginSchema } from "../model/login.schema";
import { SubmitButton } from "./SubmitButton";
import { LinkButton } from "@/shared/ui/link-button/LinkButton";
import { Container } from "@/shared/ui/container";

export interface LoginFormProps {}

export const LoginForm: FC<LoginFormProps> = () => {
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

  const clearRootErrorIfExists = () => {
    if (errors.root) {
      clearErrors("root");
    }
  };

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

  const hasRootError = !!errors.root;
  const isButtonReady = isValid || hasRootError;

  return (
    <form onSubmit={onSubmit}>
      <Container
        borderB
        borderT
        borderX
        className='flex flex-col items-center gap-4 py-8 rounded-md bg-bg-accent border-stroke-primary mt-24'
      >
        <div className='flex items-center gap-2'>
          <Logo className='bg-fg-accent' />
          <h1 className='font-bold text-2xl text-fg-accent'>MONETA</h1>
        </div>
        <Input error={errors.login?.message} className='mt-8'>
          <Input.Label className='text-fg-secondary'>Login</Input.Label>
          <Input.Field
            {...register("login", {
              onChange: clearRootErrorIfExists,
            })}
          />
          <Input.Error className='text-red-500 text-sm' />
        </Input>

        <Input error={errors.password?.message}>
          <Input.Label className='text-fg-secondary'>Password</Input.Label>
          <Input.Field
            type='password'
            {...register("password", {
              onChange: clearRootErrorIfExists,
            })}
          />
          <Input.Error className='text-red-500 text-sm' />
        </Input>

        <SubmitButton
          isSubmitting={isSubmitting}
          isReady={isButtonReady}
          className='w-full bg-fg-accent text-bg-accent font-semibold mt-4'
        />
        {errors.root && <p className='text-red-500'>{errors.root.message}</p>}

        <div className='mt-6 flex flex-col items-center text-fg-secondary gap-2'>
          <p>Don't have an account?</p>

          <LinkButton
            to={"/register"}
            className='font-semibold text-fg-primary'
          >
            SIGN UP
          </LinkButton>
        </div>
      </Container>
    </form>
  );
};
