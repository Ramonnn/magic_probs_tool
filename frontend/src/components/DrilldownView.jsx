import React, { useState } from "react";

export default function DrilldownView({ drilldown }) {
  const [expandedSet, setExpandedSet] = useState(null);
  const [expandedBooster, setExpandedBooster] = useState(null);

  return (
    <div className="space-y-4">
      {drilldown.map((set) => (
        <div key={set.setCode} className="border p-4 rounded-xl shadow">
          <div
            className="cursor-pointer font-bold text-xl"
            onClick={() =>
              setExpandedSet((prev) => (prev === set.setCode ? null : set.setCode))
            }
          >
            Set: {set.setCode} ({(set.totalProbability * 100).toFixed(2)}%)
          </div>

          {expandedSet === set.setCode && (
            <div className="pl-4 mt-2 space-y-2">
              {set.boosters.map((booster) => (
                <div key={booster.boosterName} className="border p-2 rounded">
                  <div
                    className="cursor-pointer text-lg font-medium"
                    onClick={() =>
                      setExpandedBooster((prev) =>
                        prev === booster.boosterName ? null : booster.boosterName
                      )
                    }
                  >
                    Booster: {booster.boosterName} (
                    {(booster.totalProbability * 100).toFixed(2)}%)
                  </div>

                  {expandedBooster === booster.boosterName && (
                    <div className="mt-2 pl-4">
                      <table className="w-full text-sm border">
                        <thead>
                          <tr>
                            <th className="text-left border px-2 py-1">Card UUID</th>
                            <th className="text-left border px-2 py-1">Foil</th>
                            <th className="text-left border px-2 py-1">Promo Types</th>
                            <th className="text-left border px-2 py-1">Frame Effects</th>
                            <th className="text-left border px-2 py-1">Probability</th>
                          </tr>
                        </thead>
                        <tbody>
                          {booster.cards.map((card) => (
                            <tr key={card.uuid}>
                              <td className="border px-2 py-1">{card.uuid}</td>
                              <td className="border px-2 py-1">{card.foil ? "Yes" : "No"}</td>
                              <td className="border px-2 py-1">
                                {card.promoTypes?.join(", ") || "-"}
                              </td>
                              <td className="border px-2 py-1">
                                {card.frameEffects?.join(", ") || "-"}
                              </td>
                              <td className="border px-2 py-1">
                                {(card.probability * 100).toFixed(4)}%
                              </td>
                            </tr>
                          ))}
                        </tbody>
                      </table>
                    </div>
                  )}
                </div>
              ))}
            </div>
          )}
        </div>
      ))}
    </div>
  );
}

