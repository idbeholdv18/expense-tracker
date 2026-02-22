import { useContext } from "react";
import { DispatchNavbarContext } from "../context/DispatchNavbar.context";

export const useDispatchNavbar = () => {
  return useContext(DispatchNavbarContext);
};
