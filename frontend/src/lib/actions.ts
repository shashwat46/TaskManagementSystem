"use server";

import { redirect } from 'next/navigation';

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export async function loginAction(formData: FormData) {
  const email = formData.get("email");
  const password = formData.get("password");

  try {
    console.log("Making login request to:", API_URL); // Debug log
    
    const response = await fetch(`${API_URL}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });

    if (!response.ok) {
      const error = await response.text();
      console.error("Login response not ok:", error); // Debug log
      return { error: "Invalid credentials" };
    }

    const data = await response.json();
    return { token: data.token, user: data.user };
  } catch (error) {
    console.error("Login action error:", error); // Debug log
    return { error: "Something went wrong" };
  }
}

export async function registerAction(formData: FormData) {
  const name = formData.get("name");
  const email = formData.get("email");
  const password = formData.get("password");

  try {
    const response = await fetch(`${API_URL}/auth/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, email, password }),
    });

    if (!response.ok) {
      return { error: "Registration failed" };
    }

    redirect("/login"); // Now redirect is properly imported
  } catch (error) {
    return { error: "Something went wrong" };
  }
}
