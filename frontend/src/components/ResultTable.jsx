import DrilldownView from "./DrilldownView";

function ResultTable({ data }) {
  if (!data?.probabilities) return <p>No data.</p>;

  const aggregatedByBooster = data.aggregatedByBooster || {};
  const aggregatedByBoosterFoil = data.aggregatedByFoil || {};
  const aggregatedByBoosterSet = data.aggregatedBySet || {};
  const renderList = (list) => Array.isArray(list) ? list.join(", ") : "";

  const renderAggregates = () => {
    return (
      <div className="mt-6">
        <h3 className="text-md font-bold text-blue-800 mb-2">Aggregated Probabilities</h3>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <h4 className="text-sm font-semibold text-blue-700 mb-1">By Booster</h4>
            <table className="w-full text-sm border-collapse">
              <thead className="bg-blue-50">
                <tr>
                  <th className="p-2 text-left border-b">Set + Booster</th>
                  <th className="p-2 text-left border-b">Total Probability</th>
                </tr>
              </thead>
              <tbody>
                {Object.entries(aggregatedByBoosterSet).map(([booster, prob]) => (
                  <tr key={booster}>
                    <td className="p-2 border-b">{booster}</td>
                    <td className="p-2 border-b">{(prob * 100).toFixed(4)}%</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          <div>
            <h4 className="text-sm font-semibold text-blue-700 mb-1">By Booster</h4>
            <table className="w-full text-sm border-collapse">
              <thead className="bg-blue-50">
                <tr>
                  <th className="p-2 text-left border-b">Booster</th>
                  <th className="p-2 text-left border-b">Total Probability</th>
                </tr>
              </thead>
              <tbody>
                {Object.entries(aggregatedByBooster).map(([booster, prob]) => (
                  <tr key={booster}>
                    <td className="p-2 border-b">{booster}</td>
                    <td className="p-2 border-b">{(prob * 100).toFixed(4)}%</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>

          <div>
            <h4 className="text-sm font-semibold text-blue-700 mb-1">By Booster + Foil</h4>
            <table className="w-full text-sm border-collapse">
              <thead className="bg-blue-50">
                <tr>
                  <th className="p-2 text-left border-b">Booster + Foil</th>
                  <th className="p-2 text-left border-b">Total Probability</th>
                </tr>
              </thead>
              <tbody>
                {Object.entries(aggregatedByBoosterFoil).map(([key, prob]) => (
                  <tr key={key}>
                    <td className="p-2 border-b">{key}</td>
                    <td className="p-2 border-b">{(prob * 100).toFixed(4)}%</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    );
  };

  return (
    <div className="overflow-x-auto animate-fade-in transition-opacity duration-700 ease-out">
      {renderAggregates()}

      {data.drilldown && (
        <div className="mt-6">
          <h3 className="text-md font-bold text-blue-800 mb-2">Drilldown by Set & Booster</h3>
          <DrilldownView drilldown={data.drilldown} />
        </div>
      )}

      <table className="w-full mt-4 table-auto border-collapse rounded-lg shadow-md">
        <thead className="bg-blue-100 text-left text-sm font-semibold text-blue-800">
          <tr>
            <th className="p-3 border-b">Booster</th>
            <th className="p-3 border-b">Booster Variant</th>
            <th className="p-3 border-b">UUID</th>
            <th className="p-3 border-b">Sheet</th>
            <th className="p-3 border-b">Sheet Picks</th>
            <th className="p-3 border-b">Set</th>
            <th className="p-3 border-b">Foil</th>
            <th className="p-3 border-b">Promo Types</th>
            <th className="p-3 border-b">Frame Effects</th>
            <th className="p-3 border-b">Probability</th>
          </tr>
        </thead>
        <tbody className="text-sm text-gray-800">
          {Object.entries(data.probabilities).map(([cardUuid, entries]) => {
            return entries.map((entry, index) => (
              <tr key={`${cardUuid}-${index}`} className="hover:bg-blue-50 transition-colors">
                <td className="p-3 border-b">{entry.boosterName}</td>
                <td className="p-3 border-b">{entry.boosterVariant}</td>
                <td
                  className="p-3 border-b max-w-[200px] break-words"
                  title={cardUuid}
                >
                  {cardUuid}
                </td>
                <td className="p-3 border-b">{entry.sheetName}</td>
                <td className="p-3 border-b">{entry.sheetPicks}</td>
                <td className="p-3 border-b">{entry.setCode}</td>
                <td className="p-3 border-b">{entry.isFoil ? "Yes" : "No"}</td>
                <td className="p-3 border-b">{renderList(entry.promoTypes)}</td>
                <td className="p-3 border-b">{renderList(entry.frameEffects)}</td>
                <td className="p-3 border-b">
                  {(entry.probability * 100).toFixed(4)}%
                </td>
              </tr>
            ));
          })}
        </tbody>
      </table>
    </div>
  );
}

export default ResultTable;




