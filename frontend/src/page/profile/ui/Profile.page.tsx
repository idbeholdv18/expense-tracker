import { Container } from "@/shared/ui/container";
import { Section } from "@/shared/ui/section";
import { FC } from "react";

export interface ProfilePageProps {}

const ProfilePage: FC<ProfilePageProps> = () => {
  return (
    <Section className=''>
      <Container className='min-h-screen flex flex-col'>
        <h1 className='pt-24'>Profile Page</h1>
      </Container>
    </Section>
  );
};

export default ProfilePage;
