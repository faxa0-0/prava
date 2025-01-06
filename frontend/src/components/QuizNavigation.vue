<script setup>
import Timer from '@/components/Timer.vue';
import router from '@/router';
import { ref } from 'vue';

const props = defineProps({
  questions: {
    type: Array,
    required: true,
  },
  answered: {
    type: Object,
    required: true,
  },
  currentQuestionIndex: {
    type: Number,
    required: true,
  },
});

defineEmits(['go-to-question'], ['go-to-next-question'], ['go-to-prev-question']);

const isTimeUp = ref(false);

const handleEndTest = () => {
  props.questions.value = [];
  localStorage.removeItem('questions');
  localStorage.removeItem('timerState');
  isTimeUp.value = true;
  router.push('/');
};
</script>

<template>
  <div class="rounded-lg shadow-md row-span-2 col-span-2 flex flex-col gap-8 p-5">
    <Timer :initial-minutes="20" v-model:isTimeUp="isTimeUp" />
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
        Пройдено <span>{{ Object.entries(answered).length }}</span
        >/<span>{{ questions.length }}</span>
      </div>
      <div class="grid grid-cols-4 gap-3">
        <button
          v-for="n in questions.length"
          :key="n"
          @click="$emit('go-to-question', n - 1)"
          :class="[
            'w-11 h-11 flex items-center justify-center rounded-full font-semibold transition-all duration-500',
            n - 1 === currentQuestionIndex ? 'ring-4 ring-indigo-500' : '',
            answered[n - 1] === 'correct' ? 'bg-green-500 text-white' : '',
            answered[n - 1] === 'incorrect' ? 'bg-red-500 text-white' : '',
            !answered[n - 1] ? 'bg-indigo-50 text-gray-700 hover:text-white hover:bg-indigo-500' : '',
          ]"
        >
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
