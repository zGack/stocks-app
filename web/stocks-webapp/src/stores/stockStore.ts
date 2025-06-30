// src/stores/stockStore.ts
import { defineStore } from 'pinia'
import type { Stock } from '@/types/Stock'
import { useStockService } from '@/service/stockService'
import { useStock } from '@/composables/useStock'

const { sortStocks } = useStock()

export const useStockStore = defineStore('stock', {
  state: () => ({
    stocks: [] as Stock[],
    loading: false,
    error: null as string | null,
    offset: 0,
    limit: 20,
    hasMore: true,
    sortBy: 'time',
    sortDir: 'desc' as 'asc' | 'desc',
    selectedStock: null as Stock | null,
    searchTerm: '',
    searchingByTerm: false,
  }),

  actions: {
    async fetchStocks(initial = false) {
      if (this.loading || (!this.hasMore && !initial)) return

      this.loading = true
      this.error = null

      try {
        const { getStocks } = useStockService()

        const stocks = await getStocks(this.offset, this.sortBy, this.sortDir, this.searchTerm)

        if (!stocks || stocks.length === 0) {
          this.hasMore = false
          this.stocks = []
          return
        }

        const items: Stock[] = stocks

        if (initial) {
          this.stocks = items
          this.offset = items.length
        } else {
          this.stocks.push(...items)
          this.offset += items.length
        }

        if (['change', 'change_percent'].includes(this.sortBy)) {
          this.stocks = sortStocks(this.stocks, this.sortBy, this.sortDir)
        }
      } catch (err: any) {
        console.error('Error fetching stocks:', err)
        this.error = err.message || 'Failed to fetch stocks'
        this.hasMore = false
      } finally {
        this.loading = false
      }
    },
    setSort(column: string) {
      if (this.sortBy === column) {
        this.sortDir = this.sortDir === 'asc' ? 'desc' : 'asc'
      } else {
        this.sortBy = column
        this.sortDir = 'asc'
      }

      if (['change', 'change_percent'].includes(this.sortBy)) {
        this.stocks = sortStocks(this.stocks, this.sortBy, this.sortDir)
        return
      }

      this.offset = 0
      this.hasMore = true
      this.fetchStocks(true) // re-fetch sorted data
    },
    selectStock(stock: Stock | null) {
      this.selectedStock = stock
    },
    setSearchTerm(term: string) {
      this.offset = 0
      this.searchTerm = term
      this.searchingByTerm = term.length > 0
      this.hasMore = !(term.length > 0) // reset hasMore to enable fetching with scroll again
      this.fetchStocks(true) // re-fetch data
    },
  },
})
