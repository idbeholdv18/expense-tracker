import { LoginForm } from "@/feature/auth/login";
import clsx from "clsx";
import { ChangeEvent, FC, useId, useState } from "react";
import { useFormStatus } from "react-dom";

export interface LoginPageProps {}

const LoginPage: FC<LoginPageProps> = () => {
  return (
    <div className={clsx("px-4 pt-8")}>
      <LoginForm />
    </div>
  );
};

export default LoginPage;
