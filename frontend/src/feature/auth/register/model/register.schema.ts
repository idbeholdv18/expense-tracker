import z from "zod";

export const RegisterSchema = z
  .object({
    email: z.email(),
    username: z
      .string("username must be a string")
      .min(3, "username too short")
      .max(48, "username too long")
      .trim(),

    password: z
      .string("Password must be a string")
      .min(6, "Password too short")
      .trim(),

    confirmPassword: z
      .string("Confirm password must be a string")
      .min(1, "Please confirm your password"),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords don't match",
    path: ["confirmPassword"],
  });

export type RegisterInput = z.infer<typeof RegisterSchema>;
