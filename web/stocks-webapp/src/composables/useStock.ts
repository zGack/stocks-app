import type { Stock } from "@/types/Stock"

export function useStock() {
  function calculateChange(from: string, to: string): number {
    const f = parseFloat(from.replace('$', ''))
    const t = parseFloat(to.replace('$', ''))

    if (isNaN(f) || isNaN(t)) return 0

    return t - f
  }

  function calculateChangePercent(from: string, to: string): number {
    const f = parseFloat(from.replace('$', ''))
    const t = parseFloat(to.replace('$', ''))
    if (isNaN(f) || isNaN(t) || f === 0) return 0
    return ((t - f) / f) * 100
  }

  // Sort stocks by change on the client side
  function sortStocks(stocks: Stock[], sortBy: string, sortDir: string): Stock[] {
    return stocks.sort((a, b) => {
        const aVal = sortBy === 'change'
          ? calculateChange(a.target_from, a.target_to)
          : calculateChangePercent(a.target_from, a.target_to)

        const bVal = sortBy === 'change'
          ? calculateChange(b.target_from, b.target_to)
          : calculateChangePercent(b.target_from, b.target_to)

        return sortDir === 'asc' ? aVal - bVal : bVal - aVal
      })
  }

  return {
    calculateChange,
    calculateChangePercent,
    sortStocks,
  }
}
