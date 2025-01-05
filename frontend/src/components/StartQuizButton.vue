<script setup>
import { RouterLink } from 'vue-router';
</script>

<template>
  <button
    class="relative rounded-lg px-8 py-4 overflow-hidden group bg-indigo-500 hover:bg-gradient-to-r hover:from-indigo-500 hover:to-indigo-400 text-white hover:ring-4 hover:ring-offset-2 hover:ring-indigo-400 transition-all ease-out duration-300"
  >
    <span
      class="absolute right-0 w-12 h-40 -mt-16 transition-all duration-1000 transform translate-x-16 bg-white opacity-10 rotate-12 group-hover:-translate-x-60 ease"
    ></span>
    <RouterLink to="/quiz" class="relative text-xl font-bold">Начать тест</RouterLink>
  </button>
</template>
