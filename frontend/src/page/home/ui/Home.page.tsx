import { Container } from "@/shared/ui/container";
import { Section } from "@/shared/ui/section";
import { FC } from "react";

export interface HomePageProps {}

const HomePage: FC<HomePageProps> = () => {
  return (
    <Section className=''>
      <Container className='min-h-screen flex flex-col'>
        <h1 className='pt-24'>Home Page</h1>
      </Container>
    </Section>
  );
};

export default HomePage;
