<template>
  <div
    class="flex flex-row items-center justify-around p-3 rounded-lg text-white bg-linear-to-bl from-violet-500 to-fuchsia-500"
  >
    <div class="flex text-center flex-col">
      <span class="text-2xl font-bold block">{{ stocksStore.stocks.length }}</span>
      <span class="text-md">Loaded Stocks</span>
    </div>
    <div class="flex text-center flex-col">
      <span class="text-2xl font-bold block">{{ stocksStore.gainersCount }}</span>
      <span class="text-lg">Gainers</span>
    </div>
    <div class="flex text-center flex-col">
      <span class="text-2xl font-bold block">{{ stocksStore.losersCount }}</span>
      <span class="text-lg">Loosers</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useStockStore } from '@/stores/stockStore'

const stocksStore = useStockStore()
</script>
