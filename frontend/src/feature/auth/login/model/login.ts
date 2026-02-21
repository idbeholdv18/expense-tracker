export const login = async (formData: FormData) => {
  const login = formData.get("login");
  const password = formData.get("password");

  const errors: Record<string, string> = {};

  if (!login) {
    errors.login = "Login is required";
  }

  if (login.toString().length < 3) {
    errors.login = "Login length should be gt 2";
  }

  if (!password) {
    errors.password = "Password is required";
  }

  if (Object.keys(errors).length > 0) {
    return { errors };
  }

  const response = await fetch("https://localhost:8080/api/v1/login", {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      login,
      password,
    }),
  });
  if (!response.ok) {
    return {
      errors: {
        form: "invalid credentials",
      },
    };
  }
  return {
    success: true,
  };
};
