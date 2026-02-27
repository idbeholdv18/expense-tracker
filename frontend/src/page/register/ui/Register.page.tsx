import { RegisterForm } from "@/feature/auth/register";
import clsx from "clsx";
import { FC } from "react";

export interface RegisterPageProps {}

const RegisterPage: FC<RegisterPageProps> = () => {
  return (
    <div className={clsx("px-4 pt-8")}>
      <RegisterForm />
    </div>
  );
};

export default RegisterPage;
