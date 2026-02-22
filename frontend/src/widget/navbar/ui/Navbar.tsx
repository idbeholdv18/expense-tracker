import { Container } from "@/shared/ui/container";
import { LinkButton } from "@/shared/ui/link-button/LinkButton";
import { Logo } from "@/shared/ui/logo/Logo";
import { Navbar as SharedNavbar } from "@/shared/ui/navbar";
import { Section } from "@/shared/ui/section";
import { Link } from "react-router";

export const Navbar = () => {
  return (
    <>
      <SharedNavbar.MobileMenu />
      <SharedNavbar.Layout className='h-16 border-b border-solid border-b-stroke-primary bg-bg-primary'>
        <Section>
          <Container className='flex items-center justify-between h-full p-0'>
            <SharedNavbar.Container className='flex gap-8'>
              <SharedNavbar.Logo className='flex items-center gap-2'>
                <Logo className="bg-fg-accent"/>
                <h1 className='font-bold text-2xl text-fg-accent'>MONETA</h1>
              </SharedNavbar.Logo>
              <SharedNavbar.Container className='hidden gap-2 md:flex items-center'></SharedNavbar.Container>
            </SharedNavbar.Container>
            <SharedNavbar.Burger className='block md:hidden' />

            <SharedNavbar.Container className='hidden gap-1 md:flex'>
              <LinkButton to='login' className='bg-fg-accent text-bg-accent'>
                Sign In
              </LinkButton>
              <LinkButton
                to='login'
                className='border border-solid border-bg-fg-accent'
              >
                Sign Up
              </LinkButton>
            </SharedNavbar.Container>
          </Container>
        </Section>
      </SharedNavbar.Layout>
    </>
  );
};
