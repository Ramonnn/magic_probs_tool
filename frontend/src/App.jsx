import { useState } from "react";
import InputArea from "./components/InputArea";
import SubmitButton from "./components/SubmitButton";
import ErrorMessage from "./components/ErrorMessage";
import ResultTable from "./components/ResultTable";
import useCalculate from "./hooks/useCalculate";

function App() {
  const [cardInput, setCardInput] = useState("");
  const { results, error, loading, calculate } = useCalculate();

  const submitCards = () => {
    calculate(cardInput);
  };

  return (
    <div className="max-w-2xl mx-auto p-6 space-y-6">
      <h1 className="text-2xl font-bold">Magic Probability Tool</h1>

      <InputArea
        value={cardInput}
        onChange={(e) => setCardInput(e.target.value)}
        disabled={loading}
      />

      <SubmitButton onClick={submitCards} disabled={loading} loading={loading} />

      {results && (
        <div className="animate-fade-in transition-opacity duration-700 ease-out">
          <ResultTable data={results} />
        </div>
      )}

      {
        error && (
          <p className="text-red-600 animate-bounce-in text-sm">
            {error}
          </p>
        )
      }

      <ErrorMessage message={error} />


    </div>
  );
}

export default App;

