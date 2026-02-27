import { RegisterInput } from "./register.schema";

export async function registerApi(
  data: Omit<RegisterInput, "confirmPassword">,
) {
  const response = await fetch("https://localhost:8080/api/v1/auth/register", {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    const error = await response.json().catch(() => null);

    throw new Error(error?.Message || "Register failed");
  }

  return response.json();
}
