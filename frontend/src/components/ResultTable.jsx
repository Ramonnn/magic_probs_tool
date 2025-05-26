function ResultTable({ data }) {
  if (!data?.probabilities || !data?.cardData) return <p>No data.</p>;

  const cards = data.cardData;

  return (
    <div className="overflow-x-auto animate-fade-in transition-opacity duration-700 ease-out">
      <table className="w-full mt-4 table-auto border-collapse rounded-lg shadow-md">
        <thead className="bg-blue-100 text-left text-sm font-semibold text-blue-800">
          <tr>
            <th className="p-3 border-b">Booster</th>
            <th className="p-3 border-b">Card UUID</th>
            <th className="p-3 border-b">Set</th>
            <th className="p-3 border-b">Foil</th>
            <th className="p-3 border-b">Probability</th>
          </tr>
        </thead>
        <tbody className="text-sm text-gray-800">
          {Object.entries(data.probabilities).map(([cardUuid, cardsMap]) => {
            if (Array.isArray(cardsMap)) {
              return cardsMap.map((entry, index) => (
                <tr key={`${cardUuid}-unknown-${index}`} className="hover:bg-blue-50 transition-colors">
                  <td className="p-3 border-b text-gray-500 italic">{entry.boosterName}</td>
                  <td className="p-3 border-b">{cardUuid}</td>
                  <td className="p-3 border-b">{entry.SetCode || entry.setCode}</td>
                  <td className="p-3 border-b">{entry.IsFoil || entry.isFoil ? "Yes" : "No"}</td>
                  <td className="p-3 border-b">
                    {((entry.Probability || entry.probability) * 100).toFixed(4)}%
                  </td>
                </tr>
              ));
            }

            if (typeof cardsMap !== "object" || cardsMap === null) return null;

            return Object.entries(cardsMap).map(([cardId, entries]) => {
              if (!Array.isArray(entries)) return null;

              const cardName = cards[cardId]?.name || cardId;

              return entries.map((entry, index) => (
                <tr key={`${boosterName}-${cardId}-${index}`} className="hover:bg-blue-50 transition-colors">
                  <td className="p-3 border-b">{boosterName}</td>
                  <td
                    className="p-3 border-b max-w-[200px] truncate"
                    title={cardName}
                  >
                    {cardName}
                  </td>
                  <td className="p-3 border-b">{entry.SetCode || entry.setCode}</td>
                  <td className="p-3 border-b">{entry.IsFoil || entry.isFoil ? "Yes" : "No"}</td>
                  <td className="p-3 border-b">
                    {((entry.Probability || entry.probability) * 100).toFixed(4)}%
                  </td>
                </tr>
              ));
            });
          })}
        </tbody>
      </table>
    </div>
  );
}

export default ResultTable;


