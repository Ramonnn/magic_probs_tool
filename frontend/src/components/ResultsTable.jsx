function ResultTable({ data }) {
  if (!data?.probabilities || !data?.cardData) return <p>No data.</p>;

  const cards = data.cardData;

  return (
    <table className="w-full mt-4 border border-gray-300">
      <thead className="bg-gray-100">
        <tr>
          <th className="p-2 border">Booster</th>
          <th className="p-2 border">Card UUID</th>
          <th className="p-2 border">Set</th>
          <th className="p-2 border">Foil</th>
          <th className="p-2 border">Probability</th>
        </tr>
      </thead>
      <tbody>
        {Object.entries(data.probabilities).map(([boosterName, cardsMap]) => {
          // Defensive: if cardsMap is an array (not object), treat differently
          if (Array.isArray(cardsMap)) {
            // This means cardsMap is an array of CardProbability directly (no cardId)
            return cardsMap.map((entry, index) => (
              <tr key={`${boosterName}-unknown-${index}`}>
                <td className="p-2 border">{boosterName}</td>
                <td className="p-2 border">Unknown Card</td>
                <td className="p-2 border">{entry.SetCode || entry.setCode}</td>
                <td className="p-2 border">{entry.IsFoil || entry.isFoil ? "Yes" : "No"}</td>
                <td className="p-2 border">{(entry.Probability || entry.probability) * 100?.toFixed(4)}%</td>
              </tr>
            ));
          }

          if (typeof cardsMap !== "object" || cardsMap === null) return null;

          // cardsMap is an object with cardId keys
          return Object.entries(cardsMap).map(([cardId, entries]) => {
            if (!Array.isArray(entries)) return null;

            const cardName = cards[cardId]?.name || cardId;

            return entries.map((entry, index) => (
              <tr key={`${boosterName}-${cardId}-${index}`}>
                <td className="p-2 border">{boosterName}</td>
                <td className="p-2 border">{cardName}</td>
                <td className="p-2 border">{entry.SetCode || entry.setCode}</td>
                <td className="p-2 border">{entry.IsFoil || entry.isFoil ? "Yes" : "No"}</td>
                <td className="p-2 border">{((entry.Probability || entry.probability) * 100).toFixed(4)}%</td>
              </tr>
            ));
          });
        })}
      </tbody>
    </table>
  );
}

export default ResultTable;



