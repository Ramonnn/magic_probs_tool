import { useState } from "react";
import axios from "axios";
import ResultTable from "./components/ResultsTable";

function App() {
  const [cardInput, setCardInput] = useState("");
  const [results, setResults] = useState(null);
  const [error, setError] = useState("");

  const submitCards = async () => {
    setError("");
    try {
      const response = await axios.post("/api/calculate", {
        cards: cardInput.split("\n").map((c) => c.trim()).filter(Boolean),
      });

      // No need to sanitize; use as-is
      setResults(response.data);
    } catch (err) {
      setError(err.response?.data?.error || "An error occurred");
    }
  };

  return (
    <div className="max-w-2xl mx-auto p-6 space-y-6">
      <h1 className="text-2xl font-bold">Magic Probability Tool</h1>

      <textarea
        className="w-full p-3 border rounded h-40"
        placeholder="Enter card names, one per line"
        value={cardInput}
        onChange={(e) => setCardInput(e.target.value)}
      />

      <button
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        onClick={submitCards}
      >
        Calculate
      </button>

      {error && <p className="text-red-600">{error}</p>}

      {results && <ResultTable data={results} />}
    </div>
  );
}

export default App;

