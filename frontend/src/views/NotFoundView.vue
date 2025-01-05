<script setup></script>

<template>
  <div class="h-screen flex flex-col items-center justify-center">
    <h1 class="text-[120px] font-extrabold text-gray-600">404</h1>
    <p class="text-2xl font-medium text-gray-500 mb-6">Хмм, где мы?</p>
    <RouterLink
      to="/"
      class="px-4 py-2 font-medium text-white bg-indigo-500 rounded-md hover:bg-indigo-500 transition-all duration-200 ease-in-out"
    >
      Домой
    </RouterLink>
  </div>
</template>
