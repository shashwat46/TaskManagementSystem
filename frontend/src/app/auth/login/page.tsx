"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { loginAction } from "@/lib/actions";

export default function LoginPage() {
  const router = useRouter();

  async function handleSubmit(formData: FormData) {
    try {
      const result = await loginAction(formData);
      console.log("Login response:", result); // Debug log
      
      if (result.error) {
        console.error("Login error:", result.error);
        return;
      }

      if (result.token) {
        console.log("Token received, storing..."); // Debug log
        localStorage.setItem("token", result.token);
        console.log("Redirecting to dashboard..."); // Debug log
        router.push("/dashboard");
      }
    } catch (error) {
      console.error("Login failed:", error);
    }
  }

  return (
    <div className="w-full max-w-md mx-auto mt-8">
      <form action={handleSubmit}>
        {/* Existing form JSX */}
        <div className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <div className="mb-4">
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="email">
              Email
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              id="email"
              name="email"
              type="email"
              required
              placeholder="Enter your email"
            />
          </div>
          
          <div className="mb-6">
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor="password">
              Password
            </label>
            <input
              className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
              id="password"
              name="password"
              type="password"
              required
              placeholder="Enter your password"
            />
          </div>
          
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline w-full"
            type="submit"
          >
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
}
