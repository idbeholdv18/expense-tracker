import z from "zod";

export const LoginSchema = z.object({
  login: z
    .string("Login must be a string")
    .min(3, "Login too short")
    .max(320, "Login too long")
    .trim(),

  password: z
    .string("Password must be a string")
    .min(6, "Password too short")
    .trim(),
});

export type LoginInput = z.infer<typeof LoginSchema>;
