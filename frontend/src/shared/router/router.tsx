import { App } from "@/app/ui/App";
import { ExpensesPage } from "@/page/expenses";
import { HomePage } from "@/page/home";
import { LoginPage } from "@/page/login";
import { ProfilePage } from "@/page/profile";
import { RegisterPage } from "@/page/register";
import { Suspense } from "react";
import { createBrowserRouter, Outlet } from "react-router";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      {
        path: "/login",
        element: (
          <Suspense fallback={"loading"}>
            <LoginPage />
          </Suspense>
        ),
      },
      {
        path: "/register",
        element: (
          <Suspense fallback={"loading"}>
            <RegisterPage />
          </Suspense>
        ),
      },
      {
        path: "home",
        element: (
          <Suspense fallback={"loading"}>
            <HomePage />,
          </Suspense>
        ),
      },
      {
        path: "profile",
        element: (
          <Suspense fallback={"loading"}>
            <ProfilePage />,
          </Suspense>
        ),
      },
      {
        path: "expenses",
        element: (
          <Suspense fallback={"loading"}>
            <ExpensesPage />,
          </Suspense>
        ),
      },
    ],
  },
]);
