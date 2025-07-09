<template>
  <div
    class="modal rounded-3xl w-full max-w-5xl max-h-[90vh] overflow-y-auto relative shadow-2xl bg-[linear-gradient(135deg,_#667eea_0%,_#764ba2_100%)]"
    ref="recommendedStocksModal"
  >
    <div
      class="sticky top-0 z-10 bg-[linear-gradient(135deg,_#667eea_0%,_#764ba2_100%)] px-5 py-7 rounded-2xl flex items-center justify-between"
    >
      <div class="text-left text-white">
        <h1 class="text-3xl font-bold mb-0 flex items-center gap-3">
          <span class="text-2xl">📈</span>
          Recommended Stocks
        </h1>
        <p class="text-lg opacity-90 m-0 mt-1">
          Stocks are recommended based on target change percent, rating change and current action.
        </p>
      </div>
      <button
        class="bg-[#f0f0f0] text-center lowercase border-[none] rounded-full w-[35px] h-[35px] cursor-pointer text-[1.2rem] flex items-center justify-center [transition:all_0.3s_ease] hover:bg-[#e0e0e0] hover:rotate-90"
        @click="recommendedStockStore.setRecommendedStockModalOpen()"
      >
        <svg
          class="w-5 h-5 text-gray-800"
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
          width="24"
          height="24"
          fill="none"
          viewBox="0 0 24 24"
        >
          <path
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="3"
            d="M6 18 17.94 6M18 18 6.06 6"
          />
        </svg>
      </button>
    </div>

    <div class="p-7 pt-0">
      <div class="max-w-none m-0">
        <div class="grid grid-cols-2 gap-5 mb-3">
          <div v-for="stock in stocks" :key="stock.ticker + stock.time">
            <RecommendedStockCard :stock="stock" />
          </div>
        </div>

        <LoadingSpinner v-if="loading" />
        <NoStocksFound textColor="text-white" v-if="stocks.length === 0 && !loading" />
      </div>
    </div>
  </div>
</template>

<style lang="css" scoped>
.modal::-webkit-scrollbar {
  width: 8px;
}

.modal::-webkit-scrollbar-track {
  background: rgba(255, 255, 255, 0.1);
  border-radius: 10px;
}

.modal::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.3);
  border-radius: 10px;
}

.modal::-webkit-scrollbar-thumb:hover {
  background: rgba(255, 255, 255, 0.5);
}
</style>

<script setup lang="ts">
import { useRecommendedStockStore } from '@/stores/recommendedStockStore'
import RecommendedStockCard from './RecommendedStockCard.vue'
import { onMounted, onUnmounted, useTemplateRef } from 'vue'
import { storeToRefs } from 'pinia'
import LoadingSpinner from './LoadingSpinner.vue'
import NoStocksFound from './NoStocksFound.vue'
import { useInfiniteScroll } from '@vueuse/core'

const recommendedStockStore = useRecommendedStockStore()
const { loading, stocks } = storeToRefs(recommendedStockStore)

onMounted(() => {
  recommendedStockStore.fetchStocks(true)
})

onUnmounted(() => {
  recommendedStockStore.resetParams()
})

const recommendedStocksModal = useTemplateRef<HTMLElement>('recommendedStocksModal')

useInfiniteScroll(
  recommendedStocksModal,
  () => {
    recommendedStockStore.fetchStocks()
  },
  {
    distance: 10,
    canLoadMore: () => {
      return recommendedStockStore.hasMore && stocks.value.length > 0
    },
  },
)
</script>
