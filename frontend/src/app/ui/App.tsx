import { Navbar } from "@/widget/navbar/ui/Navbar";
import { Navbar as SharedNavbar } from "@/shared/ui/navbar";

import { FC } from "react";
import { Outlet } from "react-router";

export interface AppProps {}

export const App: FC<AppProps> = () => {
  return (
    <div>
      <SharedNavbar.Provider>
        <Navbar />
      </SharedNavbar.Provider>
      <Outlet />
    </div>
  );
};
