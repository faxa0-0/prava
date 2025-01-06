<script setup>
import Timer from '@/components/Timer.vue';
import router from '@/router';

const props = defineProps({
  questions: {
    type: Array,
    required: true,
  },
  answered: {
    type: Array,
    required: true,
  },
  currentQuestionIndex: {
    type: Number,
    required: true,
  },
});

defineEmits(['go-to-question'], ['go-to-next-question'], ['go-to-prev-question']);

const handleEndTest = () => {
  props.questions.value = [];
  localStorage.removeItem('questions');
  router.push('/');
};
</script>

<template>
  <div class="rounded-lg shadow-md row-span-2 col-span-2 flex flex-col gap-8 p-5">
    <Timer :initial-minutes="1" />
    <div class="flex flex-col gap-2">
      <button
        @click="$emit('go-to-next-question')"
        type="button"
        class="py-2.5 px-6 text-sm bg-indigo-50 text-indigo-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-indigo-200"
      >
        Следующий вопрос
      </button>
      <button
        @click="$emit('go-to-prev-question')"
        type="button"
        class="py-2.5 px-6 text-sm bg-indigo-50 text-indigo-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-indigo-200"
      >
        Предыдущий вопрос
      </button>
    </div>

    <div>
      <div class="col-span-5 text-gray-700 text-lg text-center font-bold mb-5">
        Пройдено <span>{{ answeredQuestions }}</span
        >/<span>{{ questions.length }}</span>
      </div>
      <div class="grid grid-cols-4 gap-3">
        <button
          v-for="n in questions.length"
          :key="n"
          class="w-11 h-11 items-center rounded-full font-semibold transition-all duration-500"
          @click="$emit('go-to-question', n - 1)"
          :class="[
            n === currentQuestionIndex + 1
              ? 'bg-indigo-500 text-white'
              : 'bg-indigo-50 text-gray-700 hover:text-white hover:bg-indigo-500',
          ]"
        >
          {{ console.log(currentQuestionIndex) }}
          {{ n }}
        </button>
      </div>
    </div>

    <button
      @click="handleEndTest"
      type="button"
      class="mt-auto py-2.5 w-full bg-red-100 text-red-500 rounded-lg cursor-pointer font-semibold text-center shadow-xs transition-all duration-500 hover:bg-red-500 hover:text-white"
    >
      Закончить тест
    </button>
  </div>
</template>
