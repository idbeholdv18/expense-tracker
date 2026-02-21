import { FC } from "react";
import { LoginPage } from "@/page/login";

export interface AppProps {}

export const App: FC<AppProps> = () => {
  return (
    <div>
      <LoginPage />
    </div>
  );
};
