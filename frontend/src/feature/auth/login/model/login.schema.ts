import z from "zod";

export const LoginSchema = z.object({
  login: z.string("Login must be a string").trim().min(1, "Login is required"),

  password: z
    .string("Password must be a string")
    .min(1, "Password is required")
    .trim(),
});

export type LoginInput = z.infer<typeof LoginSchema>;
