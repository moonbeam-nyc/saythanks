"use client";

import { useState } from "react";

export default function Home() {
  const [address, setAddress] = useState("");
  const [zipCode, setZipCode] = useState("");
  const [recipients, setRecipients] = useState([]);

  const handleValidateAddress = async () => {
    const response = await fetch("http://localhost:8080/validate-address", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ address, zip_code: zipCode }),
    });
    const data = await response.json();
    console.log(data);
  };

  const handleGetRecipients = async () => {
    const response = await fetch("http://localhost:8080/get-recipients", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ address, zip_code: zipCode }),
    });
    const data = await response.json();
    setRecipients(data.recipients);
  };

  return (
    <div>
      <input
        type="text"
        value={address}
        onChange={(e) => setAddress(e.target.value)}
        placeholder="Address"
      />
      <input
        type="text"
        value={zipCode}
        onChange={(e) => setZipCode(e.target.value)}
        placeholder="Zip Code"
      />
      <button onClick={handleValidateAddress}>Validate Address</button>
      <button onClick={handleGetRecipients}>Get Recipients</button>
      <ul>
        {recipients.map((recipient, index) => (
          <li key={index}>{recipient}</li>
        ))}
      </ul>
    </div>
  );
}
