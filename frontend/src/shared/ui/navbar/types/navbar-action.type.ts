import { T_NavbarPayload } from "./navbar-payload.type";
import { T_ActionMap } from "@/shared/types/action-map.type";

export type T_NavbarAction =
  T_ActionMap<T_NavbarPayload>[keyof T_ActionMap<T_NavbarPayload>];
