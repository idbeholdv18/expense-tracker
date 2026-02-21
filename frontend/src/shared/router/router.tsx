import { LoginPage } from "@/page/login";
import { RegisterPage } from "@/page/register";
import { Suspense } from "react";
import { createBrowserRouter } from "react-router";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <div>Hello World</div>,
  },
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
]);
