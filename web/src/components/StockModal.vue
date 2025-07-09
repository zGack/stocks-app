<template>
  <div
    class="bg-white w-[90%] max-h-[90vh] shadow-2xl rounded-2xl p-8 max-w-2xl overflow-y-auto relative animate-[slideUp_0.3s_ease]"
    @click.stop
  >
    <header class="flex items-center justify-between mb-6 pb-4 border-b-2 border-gray-200">
      <div class="flex flex-col items-start">
        <h1 class="text-3xl font-bold text-indigo-400">{{ stock?.ticker }}</h1>
        <h2 class="text-xl text-[#666] font-semibold">{{ stock?.company }}</h2>
      </div>
      <button
        @click="closeModal"
        class="bg-[#f0f0f0] text-center lowercase border-[none] rounded-full w-[35px] h-[35px] cursor-pointer text-[1.2rem] flex items-center justify-center [transition:all_0.3s_ease] hover:bg-[#e0e0e0] hover:rotate-90"
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
    </header>
    <main class="grid gap-5 pt-5">
      <section class="grid grid-cols-2 gap-5">
        <div class="bg-[#f8f9fa] p-5 rounded-xl border-l-5 border-[#4facfe]">
          <div class="text-sm text-[#666] uppercase mb-2 font-semibold">Brokerage</div>
          <div class="text-xl font-bold text-[#333]">{{ stock?.brokerage }}</div>
        </div>
        <div class="bg-[#f8f9fa] p-5 rounded-xl border-l-5 border-[#4facfe]">
          <div class="text-sm text-[#666] uppercase mb-2 font-semibold">Action</div>
          <div class="text-xl font-bold text-[#333]">{{ stock?.action }}</div>
        </div>
        <div class="bg-[#f8f9fa] p-5 rounded-xl border-l-5 border-[#4facfe]">
          <div class="text-sm text-[#666] uppercase mb-2 font-semibold">Target</div>
          <div>
            <span
              class="line-through text-gray-400"
              :class="
                calculateChange(stock!.target_from, stock!.target_to) >= 0
                  ? 'text-green-600'
                  : 'text-red-500'
              "
              >{{ stock?.target_from }}</span
            >
            →
            <span
              class="font-semibold"
              :class="
                calculateChange(stock!.target_from, stock!.target_to) >= 0
                  ? 'text-green-600'
                  : 'text-red-500'
              "
              >{{ stock?.target_to }}</span
            >
          </div>
        </div>
        <div class="bg-[#f8f9fa] p-5 rounded-xl border-l-5 border-[#4facfe]">
          <div class="text-sm text-[#666] uppercase mb-2 font-semibold">Current Rating</div>
          <div class="text-xl font-bold text-[#333]">{{ stock?.rating_to }}</div>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { useStockStore } from '@/stores/stockStore'
import type { Stock } from '@/types/Stock'
import type { PropType } from 'vue'
import { useStock } from '@/composables/useStock'

const stocksStore = useStockStore()
const { calculateChange } = useStock()

const closeModal = () => {
  stocksStore.selectStock(null)
}

defineProps({ stock: { type: Object as PropType<Stock | null>, required: true } })
</script>
