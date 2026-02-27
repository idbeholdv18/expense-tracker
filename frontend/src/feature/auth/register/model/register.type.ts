import { RegisterInput } from "./register.schema";

export type RegisterFormState =
  | {
      success: true;
    }
  | {
      success: false;
      errors: Partial<Record<keyof RegisterInput | "form", string>>;
    };
