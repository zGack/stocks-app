<script setup lang="ts">
import { onMounted, useTemplateRef } from 'vue'
import StockTable from '@/components/StockTable.vue'
import StockFilters from '@/components/StockFilters.vue'
import Stats from '@/components/Stats.vue'
import { useStockStore } from '@/stores/stockStore'
import { useInfiniteScroll } from '@vueuse/core'
import StockModal from '@/components/StockModal.vue'

const stocksStore = useStockStore()

const cardsContainer = useTemplateRef<HTMLElement>('cardsContainer')

onMounted(async () => {
  stocksStore.fetchStocks(true)
})

useInfiniteScroll(
  cardsContainer,
  () => {
    stocksStore.fetchStocks()
  },
  {
    distance: 10,
    canLoadMore: () => {
      // inidicate when there is no more content to load so onLoadMore stops triggering
      // if (noMoreContent) return false
      return true // for demo purposes
    },
  },
)

</script>

<template>
  <div class=" flex flex-col rounded-2xl bg-white p-5 gap-4" style="height: 90vh">
    <header clas>
      <h1 class="text-2xl font-bold mb-4">Stock Ratings</h1>
      <StockFilters />
    </header>

    <Stats/>

    <div class="max-h-screen overflow-y-hidden">
      <StockTable :stocks="stocksStore.stocks" />
    </div>

    <div v-if="stocksStore.selectedStock" class="fixed top-0 left-0 w-full h-full flex justify-center items-center z-[1000] bg-[rgba(0,_0,_0,_0.5)] backdrop-blur-xs animate-[fadeIn_0.3s_ease]">
      <StockModal :stock="stocksStore.selectedStock"/>
    </div>
  </div>
</template>
