<template>
  <div class="pt-5 flex items-center  justify-between w-full">
    <div class="flex relative max-w-xs w-full">
      <span class="h-full absolute left-0 flex items-center pl-2 text-black">
        <svg viewBox="0 0 24 24" class="h-5 w-5 fill-current text-gray-500">
          <path
            d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z"
          ></path>
        </svg>
      </span>
      <input
        placeholder="Search Stocks by Company"
        class="rounded-lg border-2 border-gray-400 block pl-8 pr-6 py-2 w-full bg-white text-md placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700"
        @input="onSearchInput"
      />
    </div>
    <button
      type="button"
      class="text-white bg-blue-500 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 text-md font-medium rounded-lg px-5 py-2.5 me-2 mb-2 focus:outline-none cursor-pointer"
      @click="recommnededStocksStore.setRecommendedStockModalOpen()"
    >
      Recommended Stocks
    </button>
  </div>
</template>

<script setup lang="ts">
import { useDebounceFn } from '@vueuse/core'
import { useStockStore } from '@/stores/stockStore'
import { useRecommendedStockStore } from '@/stores/recommendedStockStore'

const stocksStore = useStockStore()
const recommnededStocksStore = useRecommendedStockStore()

const onSearchInput = useDebounceFn((e: Event) => {
  const input = e.target as HTMLInputElement

  stocksStore.setSearchTerm(input.value.trim())
}, 400)
</script>
