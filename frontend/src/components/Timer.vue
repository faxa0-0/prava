<script setup>
import router from '@/router';
import { ref, onMounted, watch } from 'vue';

const props = defineProps({
  initialMinutes: {
    type: Number,
    default: 20,
  },
});

const minutes = ref(props.initialMinutes);
const seconds = ref(0);
const isTimeUp = ref(false);
let interval;

const startTimer = () => {
  interval = setInterval(() => {
    if (seconds.value === 0) {
      if (minutes.value > 0) {
        minutes.value--;
        seconds.value = 59;
      } else {
        isTimeUp.value = true;
        clearInterval(interval);
      }
    } else {
      seconds.value--;
    }
  }, 1000);
};

watch(isTimeUp, (newVal) => {
  if (newVal) {
    localStorage.removeItem('questions');
    router.push('/');
  }
});

onMounted(() => {
  startTimer();
});
</script>

<template>
  <div class="flex flex-col items-center justify-center">
    <div class="text-2xl text-gray-700 font-bold">
      <span>{{ String(minutes).padStart(2, '0') }}</span
      >:
      <span>{{ String(seconds).padStart(2, '0') }}</span>
    </div>
    <p v-if="isTimeUp" class="mt-2 text-red-500 font-medium">Time's up!</p>
  </div>
</template>
