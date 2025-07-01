<template>
  <div class="h-screen overflow-y-auto rounded-md" ref="tableContainer" style="height: 65vh">
    <table class="min-w-full divide-y divide-gray-200 max-h-screen overflow-y-auto text-md">
      <thead class="sticky top-0 z-10">
        <tr class="bg-linear-to-t from-sky-500 to-indigo-500">
          <th
            v-for="col in columns"
            :key="col.key"
            class="px-4 py-2 text-left text-white font-bold text-md cursor-pointer select-none"
            @click="handleHeaderClick(col.key)"
          >
            <div class="flex items-center gap-1">
              {{ col.label }}
              <span v-if="stocksStore.sortBy === col.key">
                <svg
                  v-if="stocksStore.sortDir === 'asc'"
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M5 15l7-7 7 7"
                  />
                </svg>
                <svg
                  v-else
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 9l-7 7-7-7"
                  />
                </svg>
              </span>
            </div>
          </th>
        </tr>
      </thead>
      <tbody class="divide-y divide-gray-200 h-full">
        <tr
          v-for="stock in stocks"
          :key="stock.ticker + stock.time"
          class="hover:bg-indigo-50 cursor-pointer transition-all duration-300 hover:shadow-lg animate-fade-in"
          @click="stocksStore.selectStock(stock)"
        >
          <td class="px-6 py-4">
            <span class="font-bold text-indigo-400 text-lg">
              {{ stock.ticker }}
            </span>
          </td>
          <td class="px-4 py-2">{{ stock.company }}</td>
          <td
            class="px-4 py-2 capitalize"
            :class="
              calculateChange(stock.target_from, stock.target_to) >= 0
                ? 'text-green-600'
                : 'text-red-500'
            "
          >
            {{ calculateChange(stock.target_from, stock.target_to) >= 0 ? '+' : '' }}${{
              calculateChange(stock.target_from, stock.target_to).toFixed(2)
            }}
          </td>
          <td>
            <span
              class="px-2 py-2 capitalize text-md rounded-2xl font-bold"
              :class="
                calculateChangePercent(stock.target_from, stock.target_to) >= 0
                  ? 'bg-green-100 text-green-600'
                  : 'bg-red-100 text-red-500'
              "
            >
              {{ calculateChangePercent(stock.target_from, stock.target_to) >= 0 ? '+' : '' }}
              {{ calculateChangePercent(stock.target_from, stock.target_to).toFixed(2) }}%
            </span>
          </td>
          <td class="px-4 py-2">
            <span class="line-through text-gray-400">{{ stock.rating_from }}</span>
            →
            <span class="font-semibold">{{ stock.rating_to }}</span>
          </td>
          <td class="px-4 py-2 text-sm text-gray-400">{{ formatDate(stock.time) }}</td>
        </tr>
      </tbody>
    </table>

    <NoStocksFound v-if="stocks.length === 0 && !loading" />
    <LoadingSpinner v-if="loading" />
  </div>
</template>

<script setup lang="ts">
import { useDateTime } from '@/composables/useDateTime'
import { useStock } from '@/composables/useStock'
import { useStockStore } from '@/stores/stockStore'
import type { Stock } from '@/types/Stock'
import { useInfiniteScroll, useScroll } from '@vueuse/core'
import { useTemplateRef, type PropType } from 'vue'
import NoStocksFound from './NoStocksFound.vue'
import { storeToRefs } from 'pinia'
import LoadingSpinner from './LoadingSpinner.vue'

const tableContainer = useTemplateRef<HTMLElement>('tableContainer')
const stocksStore = useStockStore()

const { loading, stocks } = storeToRefs(stocksStore)

const columns = [
  { label: 'Ticker', key: 'ticker' },
  { label: 'Company', key: 'company' },
  { label: 'Change', key: 'change' },
  { label: 'Change %', key: 'change_percent' },
  { label: 'Rating', key: 'rating_from' },
  { label: 'Time', key: 'time' },
]

const { formatDate } = useDateTime()
const { calculateChange, calculateChangePercent } = useStock()

const { y } = useScroll(tableContainer)

useInfiniteScroll(
  tableContainer,
  () => {
    stocksStore.fetchStocks()
  },
  {
    distance: 10,
    canLoadMore: () => {
      return stocksStore.hasMore && !stocksStore.searchingByTerm && stocks.value.length > 0
    },
  },
)

const handleHeaderClick = (key: string) => {
  y.value = 0
  stocksStore.setSort(key)
}

defineProps({ stocks: { type: Array as PropType<Stock[]>, required: true } })
</script>
