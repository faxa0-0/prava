<script setup>
import QuizNavigation from '@/components/QuizNavigation.vue';
import axios from 'axios';
import { computed, onMounted, reactive, ref } from 'vue';

const state = reactive({
  questions: [],
  isLoading: true,
});
const currentQuestionIndex = ref(0);
const currentQuestion = computed(() => state.questions[currentQuestionIndex.value]);

const goToNextQuestion = () => {
  if (currentQuestionIndex.value < state.questions.length - 1) {
    currentQuestionIndex.value++;
  } else if (currentQuestionIndex.value == state.questions.length - 1) {
    currentQuestionIndex.value = 0;
  }
};

const goToPreviousQuestion = () => {
  if (currentQuestionIndex.value > 0) {
    currentQuestionIndex.value--;
  } else if (currentQuestionIndex.value == 0) {
    currentQuestionIndex.value = state.questions.length - 1;
  }
};

const goToQuestion = (index) => {
  if (index >= 0 && index < state.questions.length) {
    currentQuestionIndex.value = index;
  }
};

const answered = ref([]);
const handleAnswer = () => {
  answered.value += 1;
};

onMounted(async () => {
  const savedQuestions = localStorage.getItem('questions');
  if (savedQuestions) {
    state.questions = JSON.parse(savedQuestions);
    state.isLoading = false;
    return;
  }

  axios
    .get('http://localhost:8888/quiz')
    .then((res) => {
      state.questions = res.data.questions;
      localStorage.setItem('questions', JSON.stringify(state.questions));
    })
    .catch((err) => console.error('error fetching questions', err))
    .finally(() => (state.isLoading = false));
});
</script>

<template>
  <div v-if="state.isLoading" class="min-h-screen w-full flex items-center justify-around pb-5">
    <div class="grid gap-3">
      <h2
        class="text-4xl font-manrope font-extrabold text-transparent bg-gradient-to-tr from-indigo-500 to-pink-500 bg-clip-text flex items-center"
      >
        L
        <div
          class="rounded-full flex items-center justify-center w-7 h-7 bg-gradient-to-tr from-indigo-500 to-pink-500 animate-spin"
        >
          <div class="h-5 w-4 rounded-full bg-white"></div>
        </div>
        ading...
      </h2>
    </div>
  </div>
  <div v-else class="min-h-screen p-4 grid-cols-10 grid gap-4 text-white text-sm font-bold leading-6">
    <div
      class="rounded-lg shadow-lg row-span-2 col-span-8 text-gray-700 flex flex-col gap-6 py-6 px-4 overflow-y-auto max-h-[calc(100vh-2rem)] [&::-webkit-scrollbar]:hidden"
    >
      <!-- Question Image -->
      <img
        :src="
          currentQuestion.img_url ||
          'https://images.unsplash.com/photo-1516771317026-14d76f5396e5?q=80&w=2067&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
        "
        alt="Question Image"
        class="self-center h-2/4 w-2/4"
      />

      <!-- Question Text -->
      <div class="rounded-lg shadow-md py-4 px-3 text-xl">
        <p>{{ currentQuestion.text }}</p>
      </div>

      <!-- Question Options -->
      <p
        v-for="(option, optionIndex) in [currentQuestion.answer, ...currentQuestion.options]"
        :key="optionIndex"
        @click="handleAnswer"
        class="rounded-lg border-2 py-4 px-3 transition-colors duration-500 hover:bg-indigo-300 hover:text-white hover:border-indigo-400 cursor-pointer"
      >
        {{ option }}
      </p>

      <!-- Explanation -->
      <div class="hidden rounded-lg shadow-lg shadow-indigo-300 border-2 border-indigo-200 py-4 px-3 text-xl">
        <span>Объяснение</span>
        <p class="font-medium mt-2">{{ currentQuestion.explanation || 'No explanation available' }}</p>
      </div>
    </div>
    <QuizNavigation
      :currentQuestionIndex
      :answered="answered"
      :questions="state.questions"
      @go-to-question="goToQuestion"
      @go-to-next-question="goToNextQuestion"
      @go-to-prev-question="goToPreviousQuestion"
    />
  </div>
</template>
