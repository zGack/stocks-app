<template>
  <div class="mb-4 flex items-center gap-2 max-w-xs relative">
    <span class="h-full absolute inset-y-0 left-0 flex items-center pl-2 text-black">
      <svg viewBox="0 0 24 24" class="h-5 w-5 fill-current text-gray-500">
        <path
          d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z"
        ></path>
      </svg>
    </span>
    <input
      placeholder="Search Stocks by Company"
      class="rounded-lg border-2 border-gray-400 block pl-8 pr-6 py-2 w-full bg-white text-md placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700"
    />
  </div>
</template>

<script setup>
import { ref } from 'vue'
const emit = defineEmits(['filter'])

const ticker = ref('')

function emitFilters() {
  emit('filter', { ticker: ticker.value })
}
</script>
