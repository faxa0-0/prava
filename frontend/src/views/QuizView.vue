<script setup>
import QuizLoader from '@/components/QuizLoader.vue';
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
  } else {
    currentQuestionIndex.value = 0;
  }
  explanation.value = !!answered.value[currentQuestionIndex.value];
};

const goToPreviousQuestion = () => {
  if (currentQuestionIndex.value > 0) {
    currentQuestionIndex.value--;
  } else {
    currentQuestionIndex.value = state.questions.length - 1;
  }
  explanation.value = !!answered.value[currentQuestionIndex.value];
};

const goToQuestion = (index) => {
  if (index >= 0 && index < state.questions.length) {
    currentQuestionIndex.value = index;
    explanation.value = !!answered.value[currentQuestionIndex.value];
  }
};

const shuffleOptions = (question) => {
  const options = [question.answer, ...question.options];
  for (let i = options.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [options[i], options[j]] = [options[j], options[i]];
  }
  return options;
};

const explanation = ref(false);
const answered = ref({});

const handleAnswer = (option) => {
  const isCorrect = option === currentQuestion.value.answer;
  answered.value[currentQuestionIndex.value] = isCorrect ? 'correct' : 'incorrect';
  explanation.value = true;
};

onMounted(async () => {
  const savedQuestions = localStorage.getItem('questions');
  if (savedQuestions) {
    state.questions = JSON.parse(savedQuestions);
    state.questions.forEach((q) => (q.shuffledOptions = shuffleOptions(q)));
    state.isLoading = false;
    return;
  }

  axios
    .get('http://localhost:8888/quiz')
    .then((res) => {
      state.questions = res.data.questions;
      state.questions.forEach((q) => (q.shuffledOptions = shuffleOptions(q)));
      localStorage.setItem('questions', JSON.stringify(state.questions));
    })
    .catch((err) => console.error('error fetching questions', err))
    .finally(() => (state.isLoading = false));
});
</script>

<template>
  <QuizLoader v-if="state.isLoading" />

  <div v-else class="min-h-screen p-4 grid-cols-10 grid gap-4 text-white text-sm font-bold leading-6">
    <div
      class="rounded-lg shadow-lg row-span-2 col-span-8 text-gray-700 flex flex-col gap-6 py-6 px-4 overflow-y-auto max-h-[calc(100vh-2rem)] [&::-webkit-scrollbar]:hidden"
    >
      <img
        :src="
          currentQuestion.img_url ||
          'https://images.unsplash.com/photo-1516771317026-14d76f5396e5?q=80&w=2067&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
        "
        alt="Question Image"
        class="self-center h-2/4 w-2/4"
      />

      <div class="rounded-lg shadow-md py-4 px-3 text-xl">
        <p>{{ currentQuestion.text }}</p>
      </div>

      <p
        v-for="(option, optionIndex) in currentQuestion.shuffledOptions"
        :key="optionIndex"
        @click="handleAnswer(option)"
        :class="[
          'rounded-lg border-2 py-4 px-3 transition-colors duration-500 cursor-pointer',
          selectedOptions.value[currentQuestionIndex.value] === option &&
          answered.value[currentQuestionIndex.value] === 'correct'
            ? 'bg-green-500 text-white'
            : '',
          selectedOptions.value[currentQuestionIndex.value] === option &&
          answered.value[currentQuestionIndex.value] === 'incorrect'
            ? 'bg-red-500 text-white'
            : '',
          !explanation ? 'hover:bg-indigo-300 hover:text-white hover:border-indigo-400' : '',
        ]"
      >
        {{ option }}
      </p>

      <div
        :class="[explanation ? '' : 'hidden']"
        class="rounded-lg shadow-lg shadow-indigo-300 border-2 border-indigo-200 py-4 px-3 text-xl"
      >
        <span>Объяснение</span>
        <p class="font-medium mt-2">{{ currentQuestion.explanation || 'No explanation available' }}</p>
      </div>
    </div>

    <QuizNavigation
      :currentQuestionIndex="currentQuestionIndex"
      :answered="answered"
      :questions="state.questions"
      @go-to-question="goToQuestion"
      @go-to-next-question="goToNextQuestion"
      @go-to-prev-question="goToPreviousQuestion"
    />
  </div>
</template>
