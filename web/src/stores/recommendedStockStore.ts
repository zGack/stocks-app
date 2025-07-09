// src/stores/recommendedStockStore.ts
import { defineStore } from 'pinia'
import type { Stock } from '@/types/Stock'
import { useStockService } from '@/service/stockService'

export const useRecommendedStockStore = defineStore('recommendedStock', {
  state: () => ({
    stocks: [] as Stock[],
    loading: false,
    error: null as string | null,
    offset: 0,
    limit: 10,
    hasMore: true,
    sortBy: 'time',
    sortDir: 'desc' as 'asc' | 'desc',
    recommendedStockModalOpen: false,
  }),
  getters: {
    isLoading: (state) => state.loading,
  },
  actions: {
    async fetchStocks(initial = false) {
      if (this.loading || (!this.hasMore && !initial)) return

      this.loading = true
      this.error = null

      try {
        const { getStocks } = useStockService()
        this.sortBy = 'stock_score'
        this.sortDir = 'desc'

        const stocks = await getStocks(this.offset, this.sortBy, this.sortDir)

        if (!stocks || stocks.length === 0) {
          this.hasMore = false
          this.stocks = []
          this.loading = false
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
      } catch (err: any) {
        console.log('Error fetching recommended stocks:', err)
        this.error = err.message || 'Failed to fetch stocks'
        this.hasMore = false
      } finally {
        this.loading = false
      }
    },
    setRecommendedStockModalOpen() {
      this.recommendedStockModalOpen = !this.recommendedStockModalOpen
    },
    resetParams() {
      this.offset = 0
      this.hasMore = true
    }
  },
})
