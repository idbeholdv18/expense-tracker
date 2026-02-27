import { Container } from "@/shared/ui/container";
import { Input } from "@/shared/ui/input/Input";
import { LinkButton } from "@/shared/ui/link-button/LinkButton";
import { Logo } from "@/shared/ui/logo/Logo";
import { zodResolver } from "@hookform/resolvers/zod";
import { FC } from "react";
import { useForm } from "react-hook-form";
import { SubmitButton } from "../../login/ui/SubmitButton";
import { registerApi } from "../model/register.action";
import { RegisterInput, RegisterSchema } from "../model/register.schema";

export interface RegisterFormProps {}

export const RegisterForm: FC<RegisterFormProps> = () => {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting, isValid },
    setError,
    clearErrors,
  } = useForm<RegisterInput>({
    resolver: zodResolver(RegisterSchema),
    mode: "all",
  });

  const clearRootErrorIfExists = () => {
    if (errors.root) {
      clearErrors("root");
    }
  };

  const onSubmit = handleSubmit(async (data) => {
    try {
      await registerApi({
        email: data.email,
        password: data.password,
        username: data.username,
      });
      console.log("Register success");
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
        <Input error={errors.email?.message} className='mt-8'>
          <Input.Label className='text-fg-secondary'>Email</Input.Label>
          <Input.Field
            {...register("email", {
              onChange: clearRootErrorIfExists,
            })}
          />
          <Input.Error className='text-red-500 text-sm' />
        </Input>

        {/* TODO: check if username is taken */}
        <Input error={errors.username?.message}>
          <Input.Label className='text-fg-secondary'>Username</Input.Label>
          <Input.Field
            {...register("username", {
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

        <Input error={errors.confirmPassword?.message}>
          <Input.Label className='text-fg-secondary'>
            Confirm Password
          </Input.Label>
          <Input.Field
            type='password'
            {...register("confirmPassword", {
              onChange: clearRootErrorIfExists,
            })}
          />
          <Input.Error className='text-red-500 text-sm' />
        </Input>

        <SubmitButton
          isSubmitting={isSubmitting}
          isReady={isButtonReady}
          className='w-full bg-fg-accent text-bg-accent font-semibold mt-4'
        >
          Sign Up
        </SubmitButton>
        {errors.root && <p className='text-red-500'>{errors.root.message}</p>}

        <div className='mt-6 flex flex-col items-center text-fg-secondary gap-2'>
          <p>Already have an account?</p>

          <LinkButton to={"/login"} className='font-semibold text-fg-primary'>
            SIGN IN
          </LinkButton>
        </div>
      </Container>
    </form>
  );
};
