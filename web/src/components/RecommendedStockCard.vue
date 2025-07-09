<template>
  <div
    class="bg-white rounded-2xl p-8 backdrop-blur-sm shadow-lg border-1 border-gray-200 [transition:all_0.3s_ease] hover:translate-[-5px]"
  >
    <div class="flex justify-between items-start mb-5">
      <div class="text-2xl font-bold text-indigo-400">{{ stock?.ticker }}</div>
      <div class="flex flex-col items-center">
        <h3>Score</h3>
        <span class="font-semibold">{{ stock?.stock_score.toFixed(2) }}</span>
      </div>
    </div>
    <div class="text-lg text-[#333] pb-1 font-medium">{{ stock?.company }}</div>
    <div class="flex items-center justify-between mb-5">
      <div class="text-2xl font-bold text-[#333]">{{ stock?.target_to }}</div>
      <div
        class="text-md font-semibold text-green-600"
        :class="
          calculateChange(stock!.target_from, stock!.target_to) >= 0
            ? 'text-green-600'
            : 'text-red-500'
        "
      >
        {{ calculateChange(stock!.target_from, stock!.target_to) >= 0 ? '+' : '' }}${{
          calculateChange(stock!.target_from, stock!.target_to).toFixed(2)
        }}
        ({{ calculateChangePercent(stock!.target_from, stock!.target_to) >= 0 ? '+' : '' }}
        {{ calculateChangePercent(stock!.target_from, stock!.target_to).toFixed(2) }}%)
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStock } from '@/composables/useStock'
import type { Stock } from '@/types/Stock'
import type { PropType } from 'vue'

const { calculateChange, calculateChangePercent } = useStock()

defineProps({ stock: { type: Object as PropType<Stock | null>, required: true } })
</script>
