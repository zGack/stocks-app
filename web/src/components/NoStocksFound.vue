<template>
  <div :class="['text-center p-10 text-lg ', textColor || 'text-[#666]']">
      <h3>No Stocks Found</h3>
      <p class="text-sm mt-2 opacity-80">Try adjusting your search criteria and try again.</p>
    </div>
</template>

<script setup lang="ts">

defineProps({ textColor: { type: String } })
</script>
