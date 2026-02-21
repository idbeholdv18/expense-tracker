import { LoginInput } from "./login.schema";

export type LoginFormState =
  | {
      success: true;
    }
  | {
      success: false;
      errors: Partial<Record<keyof LoginInput | "form", string>>;
    };
