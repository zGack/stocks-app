import { getStocksFromApi } from '@/api/stock'

export function useStockService() {
  async function getStocks(
    offset: number,
    sortBy: string = 'time',
    sortDir: 'desc' | 'asc' = 'desc',
  ) {
    if (['change', 'change_percent'].includes(sortBy)) {
      // since change and change_percent are not supported by the API, use time as fallback
      sortBy = 'time'
    }
    return await getStocksFromApi(offset, sortBy, sortDir)
  }

  return {
    getStocks,
  }
}
