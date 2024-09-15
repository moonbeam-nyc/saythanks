"use client";

import { useState } from "react";

export default function Home() {
  const [address, setAddress] = useState("");
  const [city, setCity] = useState("");
  const [state, setState] = useState("");
  const [responseJson, setResponseJson] = useState(null);
  const [error, setError] = useState("");

  const handleValidateAddress = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");
    setResponseJson(null);

    try {
      const response = await fetch(
        "http://localhost:8080/api/address/validate",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ address, city, state }),
        }
      );

      if (!response.ok) {
        throw new Error("Failed to validate address");
      }

      const data = await response.json();
      setResponseJson(data);
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Address Validation</h1>
      <form onSubmit={handleValidateAddress} className="mb-4">
        <div className="mb-2">
          <label className="block text-sm font-medium text-gray-700">
            Address
          </label>
          <input
            type="text"
            value={address}
            onChange={(e) => setAddress(e.target.value)}
            className="mt-1 block w-full p-2 border border-gray-300 rounded-md text-black"
            placeholder="Address"
            required
          />
        </div>
        <div className="mb-2">
          <label className="block text-sm font-medium text-gray-700">
            City
          </label>
          <input
            type="text"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            className="mt-1 block w-full p-2 border border-gray-300 rounded-md text-black"
            placeholder="City"
            required
          />
        </div>
        <div className="mb-2">
          <label className="block text-sm font-medium text-gray-700">
            State
          </label>
          <input
            type="text"
            value={state}
            onChange={(e) => setState(e.target.value)}
            className="mt-1 block w-full p-2 border border-gray-300 rounded-md text-black"
            placeholder="State"
            required
          />
        </div>
        <button
          type="submit"
          className="mt-2 px-4 py-2 bg-blue-500 text-white rounded-md"
        >
          Validate Address
        </button>
      </form>

      {error && <p className="text-red-500">{error}</p>}

      {responseJson && (
        <div className="mt-4 text-black">
          <h2 className="text-xl font-bold mb-2">Response</h2>
          <pre className="bg-gray-100 p-4 rounded-md">
            {JSON.stringify(responseJson, null, 2)}
          </pre>
        </div>
      )}
    </div>
  );
}
