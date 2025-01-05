<script setup>
defineProps({
  totalQuestions: {
    type: Number,
    required: true,
  },
  answeredQuestions: {
    type: Number,
    required: true,
  },
});
</script>

<template>
  <div class="rounded-lg shadow-md row-span-2 col-span-2 flex flex-col gap-8 p-5">
    <div class="flex flex-col gap-2">
      <button
        type="button"
        class="py-2.5 px-6 text-sm bg-indigo-50 text-indigo-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-indigo-200"
      >
        Следующий вопрос
      </button>
      <button
        type="button"
        class="py-2.5 px-6 text-sm bg-indigo-50 text-indigo-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-indigo-200"
      >
        Предыдущий вопрос
      </button>
    </div>

    <div>
      <div class="col-span-5 text-gray-700 text-lg text-center font-bold mb-5">
        Пройдено <span>{{ answeredQuestions }}</span
        >/<span>{{ totalQuestions }}</span>
      </div>
      <div class="grid grid-cols-4 gap-3">
        <button
          v-for="n in totalQuestions"
          :key="n"
          class="w-11 h-11 items-center rounded-full font-semibold bg-indigo-50 text-gray-700 hover:text-white hover:bg-indigo-500 transition-all duration-500"
        >
          {{ n }}
        </button>
      </div>
    </div>

    <button
      type="button"
      class="mt-auto py-2.5 w-full bg-red-100 text-red-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-red-500 hover:text-white"
    >
      Закончить тест
    </button>
  </div>
</template>
