<script setup lang="ts">
import { onMounted } from 'vue'
import StockTable from '@/components/StockTable.vue'
import StockFilters from '@/components/StockFilters.vue'
import Stats from '@/components/Stats.vue'
import { useStockStore } from '@/stores/stockStore'
import StockModal from '@/components/StockModal.vue'
import RecommendedStocksModal from '@/components/RecommendedStocksModal.vue'
import { useRecommendedStockStore } from '@/stores/recommendedStockStore'

const stocksStore = useStockStore()
const recommnededStocksStore = useRecommendedStockStore()

onMounted(async () => {
  stocksStore.fetchStocks(true)
})
</script>

<template>
  <div class="flex flex-col rounded-2xl bg-white p-5 gap-4" style="height: 92vh">
    <header class="w-full">
      <h1 class="text-2xl font-bold mb-10">Stocks App</h1>
      <StockFilters />
    </header>

    <Stats />

    <div class="max-h-screen overflow-y-hidden">
      <StockTable :stocks="stocksStore.stocks" />
    </div>

    <div
      v-if="stocksStore.selectedStock"
      class="fixed top-0 left-0 w-full h-full flex justify-center items-center z-[1000] bg-[rgba(0,_0,_0,_0.5)] backdrop-blur-xs animate-[fadeIn_0.3s_ease]"
    >
      <StockModal :stock="stocksStore.selectedStock" />
    </div>

    <div
      v-if="recommnededStocksStore.recommendedStockModalOpen"
      class="fixed left-0 top-0 w-full h-full bg-[rgba(0,_0,_0,_0.5)] p-5 flex items-center justify-center z-[1000]"
    >
      <RecommendedStocksModal />
    </div>
  </div>
</template>
