import { LoginForm } from "@/feature/auth/login";
import { Container } from "@/shared/ui/container";
import { Section } from "@/shared/ui/section";
import { FC } from "react";

export interface LoginPageProps {}

const LoginPage: FC<LoginPageProps> = () => {
  return (
    <Section className='bg-auth-pattern'>
      <Container className='min-h-screen flex flex-col'>
        <LoginForm />
      </Container>
    </Section>
  );
};

export default LoginPage;
