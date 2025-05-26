import { useState, useMemo } from "react";
import axios from "axios";
import debounce from "lodash.debounce";

export default function useCalculate() {
	const [results, setResults] = useState(null);
	const [error, setError] = useState("");
	const [loading, setLoading] = useState(false);

	const _calculate = async (cardInput) => {
		setError("");
		setLoading(true);
		try {
			const response = await axios.post("/api/calculate", {
				cards: cardInput
					.split("\n")
					.map((c) => c.trim())
					.filter(Boolean),
			});
			setResults(response.data);
		} catch (err) {
			setError(err.response?.data?.error || "An error occurred");
		} finally {
			setLoading(false);
		}
	};

	// Only recreate debounce once per instance
	const debouncedCalculate = useMemo(
		() => debounce(_calculate, 500),
		[] // don't recreate unless the component remounts
	);

	return { results, error, loading, calculate: debouncedCalculate };
}

