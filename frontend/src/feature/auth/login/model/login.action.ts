import { LoginInput } from "./login.schema";

export async function loginApi(data: LoginInput) {
  const response = await fetch("https://localhost:8080/api/v1/auth/login", {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  if (!response.ok) {
    const error = await response.json().catch(() => null);

    throw new Error(error?.Message || "Login failed");
  }

  return response.json();
}
